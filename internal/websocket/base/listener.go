package base

func (s *Websocket) listen() {
	// Goroutine to read messages
	for {
		// Stop the loop if the socket is marked as closed
		if s.closed.Load() {
			return
		}

		msgType, msg, err := s.conn.ReadMessage()
		if err != nil {
			if s.closed.Load() {
				return
			}

			s.logger.ERRORf("[%s] Error reading message: %v", s.url, err)
			s.markClosed(false)
			return
		}

		s.logger.DEBUGf("Type: %d, message: %s\n", msgType, string(msg))
		s.recordLastHeartbeat()

		// Send the data over to the message handler
		s.messageHandler(msgType, msg)
	}
}
