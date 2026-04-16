package reconnecting

import (
	"fmt"

	ws "github.com/gorilla/websocket"
)

func (s *Websocket) SendPing(data []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.closed {
		return fmt.Errorf("socket is closed")
	}

	return s.base.SendPing(data)
}

func (s *Websocket) SendPong(data []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.closed {
		return fmt.Errorf("socket is closed")
	}

	return s.base.SendPong(data)
}

func (s *Websocket) SendJSON(v interface{}) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.closed {
		return fmt.Errorf("socket is closed")
	}

	return s.base.SendJSON(v)
}

func (s *Websocket) SendMessage(data []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.closed {
		return fmt.Errorf("socket is closed")
	}

	return s.base.SendMessage(data)
}

func (s *Websocket) SendPreparedMessage(pm *ws.PreparedMessage) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.closed {
		return fmt.Errorf("socket is closed")
	}

	return s.base.SendPreparedMessage(pm)
}
