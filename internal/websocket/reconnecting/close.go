package reconnecting

import "time"

func (s *Websocket) closeHandler() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.closed {
		return
	}

	s.broadcastOnDisconnect()

	s.reconnecting = true

	tries := 0
	for {
		if s.closed {
			return
		}

		err := s.newSubSocket()
		if err == nil {
			break
		}

		d := computeBackoffDuration(tries)
		s.logger.DEBUGf("failed to reconnected to '%s' after try #%d, waiting %d ms before trying again", s.url, tries+1, d.Milliseconds())

		time.Sleep(d)
		tries++

		continue
	}

	s.reconnecting = false
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
	// Closes it so that if closeHandler() loop is still running, it will be able to close early.
	s.closed = true

	s.mu.Lock()
	defer s.mu.Unlock()

	s.base.Close()
}
