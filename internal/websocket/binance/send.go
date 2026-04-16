package binance

import (
	ws "github.com/gorilla/websocket"
)

func (s *Websocket) SendPing(data []byte) error {
	return s.base.SendPing(data)
}

func (s *Websocket) SendPong(data []byte) error {
	return s.base.SendPong(data)
}

func (s *Websocket) SendJSON(v interface{}) error {
	return s.base.SendJSON(v)
}

func (s *Websocket) SendMessage(data []byte) error {
	return s.base.SendMessage(data)
}

func (s *Websocket) SendPreparedMessage(pm *ws.PreparedMessage) error {
	return s.base.SendPreparedMessage(pm)
}
