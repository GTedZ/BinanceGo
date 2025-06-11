package websockets

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
)

type privateMessageWebsocket struct {
	base *baseWebsocket

	privateMessagePropertyName string
	pendingRequests            struct {
		Mu  sync.Mutex
		Map map[string]chan []byte
	}

	OnMessage func(messageType int, msg []byte)
	OnClose   func()
}

func (socket *privateMessageWebsocket) addPendingRequest() (string, chan []byte) {
	socket.pendingRequests.Mu.Lock()
	defer socket.pendingRequests.Mu.Unlock()

	channel := make(chan []byte)
	var requestId string
	for {
		requestId = uuid.New().String()
		_, exists := socket.pendingRequests.Map[requestId]
		if exists {
			continue
		}
		break
	}

	socket.pendingRequests.Map[requestId] = channel
	return requestId, channel
}

func (socket *privateMessageWebsocket) removePendingRequest(id string) {
	channel, exists := socket.getPendingRequest(id)

	socket.pendingRequests.Mu.Lock()
	defer socket.pendingRequests.Mu.Unlock()

	if exists {
		delete(socket.pendingRequests.Map, id)
		close(channel)
	}
}

func (socket *privateMessageWebsocket) getPendingRequest(id string) (channel chan []byte, exists bool) {
	socket.pendingRequests.Mu.Lock()
	defer socket.pendingRequests.Mu.Unlock()

	channel, exists = socket.pendingRequests.Map[id]
	return channel, exists
}

func (socket *privateMessageWebsocket) onClose() {
	if socket.OnClose != nil {
		socket.OnClose()
	}
}

func (socket *privateMessageWebsocket) onMessage(msgType int, msg []byte) {
	id, isPrivate := socket.checkMessageIsPrivate(msg)
	if isPrivate {
		channel, exists := socket.getPendingRequest(id)
		if exists {
			go func() {
				defer func() {
					if r := recover(); r != nil {
						Logger.WARN(fmt.Sprintf("Recovered from panic: 'send on closed channel' trying to respond to Id %s", id))
					}
				}()

				select {
				case channel <- msg:
					socket.removePendingRequest(id)
				case <-time.After(40 * time.Second):
					Logger.WARN("Private message channel send timeout")
					socket.removePendingRequest(id)
				}
			}()

		}
		return
	}

	if socket.OnMessage != nil {
		socket.OnMessage(msgType, msg)
	}
}

func (socket *privateMessageWebsocket) checkMessageIsPrivate(msg []byte) (Id string, isPrivate bool) {
	if len(msg) == 0 {
		return "", false
	}
	if msg[0] == '[' {
		return "", false
	}

	var parsedMessage map[string]interface{}
	err := jsoniter.Unmarshal(msg, &parsedMessage)
	if err != nil {
		Logger.ERROR(fmt.Sprintf("[%s] Failed to unmarshall the following message => %s", socket.base.url, msg), err)
		return "", false
	}

	propertyInterface, exists := parsedMessage[socket.privateMessagePropertyName]
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

// Public Methods

func (socket *privateMessageWebsocket) Close() {
	socket.base.Close()
}

func (socket *privateMessageWebsocket) SendJSON(v interface{}) error {
	Logger.DEBUG(fmt.Sprintf("Sending JSON => %v", v))
	return socket.base.SendJSON(v)
}

func (socket *privateMessageWebsocket) SendPrivateMessage(message map[string]interface{}, timeout_sec ...int) (response []byte, hasTimedOut bool, err error) {
	requestId, respChan := socket.addPendingRequest()

	Logger.INFO(fmt.Sprintf("Sending private request of id %s => %v", requestId, message))

	message[socket.privateMessagePropertyName] = requestId

	err = socket.SendJSON(message)
	if err != nil {
		Logger.ERROR(fmt.Sprintf("[%s] There was an error sending JSON for private message", socket.base.url), err)
		return nil, false, err
	}

	timeout := 4
	if len(timeout_sec) > 0 {
		timeout = timeout_sec[0]
	}

	// Wait for response or timeout
	var timer <-chan time.Time
	if timeout > 0 {
		timer = time.After(time.Duration(timeout) * time.Second)
	}

	select {
	case resp := <-respChan:
		return resp, false, nil
	case <-timer:
		go func() {
			time.Sleep(20 * time.Second)
			socket.removePendingRequest(requestId) // Let's not clean them out in case there is an error, not alot of timed out requests should be expected anyway
		}()
		return nil, true, fmt.Errorf("the request has timed out after %d seconds", timeout)
	}
}

//

func createPrivateMessageWebsocket(URL string, privateMessagePropertyName string) (*privateMessageWebsocket, error) {
	baseSocket, err := createBaseSocket(URL)
	if err != nil {
		return nil, err
	}

	socket := &privateMessageWebsocket{
		base:                       baseSocket,
		privateMessagePropertyName: privateMessagePropertyName,
	}
	socket.pendingRequests.Map = make(map[string]chan []byte)

	socket.base.OnClose = socket.onClose
	socket.base.OnMessage = socket.onMessage

	return socket, nil
}
