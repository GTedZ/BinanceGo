package base

import ws "github.com/gorilla/websocket"

func (s *Websocket) SendPing(data []byte) error {
	s.writeMu.Lock()
	defer s.writeMu.Unlock()

	return s.conn.WriteMessage(ws.PingMessage, data)
}

func (s *Websocket) SendPong(data []byte) error {
	s.writeMu.Lock()
	defer s.writeMu.Unlock()

	return s.conn.WriteMessage(ws.PongMessage, data)
}

func (s *Websocket) SendJSON(v interface{}) error {
	s.writeMu.Lock()
	defer s.writeMu.Unlock()

	return s.conn.WriteJSON(v)
}

func (s *Websocket) SendMessage(data []byte) error {
	s.writeMu.Lock()
	defer s.writeMu.Unlock()

	return s.conn.WriteMessage(ws.TextMessage, data)
}

func (s *Websocket) SendPreparedMessage(pm *ws.PreparedMessage) error {
	s.writeMu.Lock()
	defer s.writeMu.Unlock()

	return s.conn.WritePreparedMessage(pm)
}
