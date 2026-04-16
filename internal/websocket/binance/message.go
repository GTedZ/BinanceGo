package binance

import (
	"time"
)

func (s *Websocket) messageHandler(data []byte) {
	id, isPrivate := s.checkMessageIsPrivate(data)
	if !isPrivate {
		s.broadcastOnMessage(data)
		return
	}

	s.requests.mu.Lock()
	channel, exists := s.requests.data[id]
	s.requests.mu.Unlock()

	if !exists {
		s.logger.DEBUGf("Received response for unknown request ID %s, ignoring", id)
		return
	}

	go func() {
		select {
		case channel <- data:
			s.removePendingRequest(id)

		case <-time.After(SENDER_TIMEOUT_DURATION):
			s.logger.WARN("Private message channel send timeout")
			s.removePendingRequest(id)
		}
	}()
}

////
// Public Methods
////

func (s *Websocket) SetOnMessage(cb func(data []byte)) {
	s.onMessage = cb
}
