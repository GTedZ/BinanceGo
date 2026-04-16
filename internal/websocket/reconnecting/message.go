package reconnecting

func (s *Websocket) messageHandler(data []byte) {
	s.broadcastOnMessage(data)
}

////
// Public Methods
////

func (s *Websocket) SetOnMessage(cb func(data []byte)) {
	s.onMessage = cb
}
