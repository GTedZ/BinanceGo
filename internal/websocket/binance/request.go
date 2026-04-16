package binance

import (
	"fmt"
	"sync"
	"time"

	"github.com/GTedZ/binancego/internal/berror"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
)

func (s *Websocket) checkMessageIsPrivate(data []byte) (id string, isPrivate bool) {
	if len(data) == 0 {
		return "", false
	}
	if data[0] == '[' {
		return "", false
	}

	var parsedMessage map[string]interface{}
	err := jsoniter.Unmarshal(data, &parsedMessage)
	if err != nil {
		s.logger.ERRORf("[%s] Failed to unmarshall the following message '%s': '%s'", s.url, data, err.Error())
		return "", false
	}

	propertyInterface, exists := parsedMessage[s.requestPropertyKey]
	if !exists {
		return "", false
	}

	propertyValue, ok := propertyInterface.(string)
	if !ok {
		return "", false
	}

	if propertyValue == "" {
		return "", false
	}

	return propertyValue, true
}

func (s *Websocket) addPendingRequest() (string, chan []byte) {
	s.requests.mu.Lock()
	defer s.requests.mu.Unlock()

	channel := make(chan []byte)
	requestId := uuid.New().String()

	s.requests.data[requestId] = channel
	s.requests.once[requestId] = &sync.Once{}
	return requestId, channel
}

func (s *Websocket) removePendingRequest(id string) {
	s.requests.mu.Lock()
	once, exists := s.requests.once[id]
	s.requests.mu.Unlock()

	if !exists {
		return
	}

	once.Do(func() {
		s.requests.mu.Lock()
		channel, exists := s.requests.data[id]
		if exists {
			delete(s.requests.data, id)
			delete(s.requests.once, id)
			close(channel)
		}
		s.requests.mu.Unlock()
	})
}

////
// Public Methods
////

func (s *Websocket) SendRequest(message map[string]interface{}, timeout time.Duration) (data []byte, err berror.Error) {
	requestId, respChan := s.addPendingRequest()

	s.logger.INFOf("Sending private request of id %s => %v", requestId, message)

	message[s.requestPropertyKey] = requestId

	sendErr := s.SendJSON(message)
	if sendErr != nil {
		s.logger.ERRORf("[%s] There was an error sending JSON for private message: %s", s.url, sendErr.Error())
		return nil, berror.NewNetworkError(sendErr)
	}

	// Wait for response or timeout
	var timer <-chan time.Time
	if timeout > 0 {
		timer = time.After(timeout)
	}

	select {
	case resp := <-respChan:
		return resp, nil
	case <-timer:
		go func() {
			time.Sleep(REQUEST_TIMEOUT_DURATION)
			s.removePendingRequest(requestId)
		}()
		return nil, berror.NewTimeoutError(fmt.Sprintf("request timed out after %f seconds", timeout.Seconds()))
	}
}
