package binance

func (s *Websocket) disconnectHandler() {
	s.broadcastOnDisconnect()
}

func (s *Websocket) reconnectedHandler() {
	s.broadcastOnReconnected()
}

////
// Public Methods
////

func (s *Websocket) SetOnDisconnect(cb func()) {
	s.onDisconnect = cb
}

func (s *Websocket) SetOnReconnected(cb func()) {
	s.onReconnected = cb
}

func (s *Websocket) Close() {
	s.base.Close()
}
