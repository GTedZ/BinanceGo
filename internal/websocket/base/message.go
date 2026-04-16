package base

import (
	ws "github.com/gorilla/websocket"
)

func (s *Websocket) messageHandler(messageType int, data []byte) {
	if s.closed.Load() {
		return
	}

	switch messageType {
	case ws.TextMessage:
		s.broadcastOnMessage(data)

	case ws.BinaryMessage:
		s.logger.DEBUGf("Received binary message data: %s", string(data))
	}
}

////
// Public Methods
////

func (s *Websocket) SetOnMessage(cb func(data []byte)) {
	s.onMessage = cb
}
