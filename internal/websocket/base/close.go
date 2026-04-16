package base

func (s *Websocket) markClosed(silent bool) {
	// CompareAndSwap returns "true" if the swap was successful
	// Meaning that socket.closed was false, which means that if it were true, we'd want to emit the OnClose()
	wasFalseAndHasBeenSwappedToTrue := s.closed.CompareAndSwap(false, true)

	if !wasFalseAndHasBeenSwappedToTrue {
		s.logger.DEBUGf("[%s] socket already marked as closed", s.url)
		return
	}

	if !silent {
		s.broadcastOnClose()
	}
}

func (s *Websocket) closeHandler(code int, text string) error {
	s.logger.DEBUGf("[*Websocket.CloseHandler()] code: %d, text: %s, isClosed: %v", code, text, s.closed.Load())
	s.markClosed(false)

	return nil
}

////
// Public Method
////

// SetOnClose sets the callback function called upon socket closure
func (s *Websocket) SetOnClose(cb func()) {
	s.onClose = cb
}

func (s *Websocket) Close() {
	s.markClosed(true)

	s.conn.Close()
}
