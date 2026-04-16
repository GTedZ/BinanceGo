package base

func (s *Websocket) pongHandler(pongData string) error {
	s.logger.DEBUGf("Received a pong: %s", pongData)
	s.recordLastHeartbeat()

	return nil
}
