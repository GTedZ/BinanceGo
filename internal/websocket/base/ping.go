package base

func (s *Websocket) pingHandler(pingData string) error {
	s.logger.DEBUGf("Received a ping: %s", pingData)

	s.recordLastHeartbeat() // Logs a heartbeat

	err := s.SendPong([]byte(pingData))
	if err != nil {
		s.logger.ERROR("Error sending Pong:", err)
		return err
	}

	return nil
}
