package base

import (
	"time"
)

func (s *Websocket) recordLastHeartbeat() {
	s.heartbeatMu.Lock()
	s.lastHeartbeatTime = time.Now()
	s.heartbeatMu.Unlock()
}

func (s *Websocket) getLastHeartbeatTime() time.Time {
	s.heartbeatMu.Lock()
	defer s.heartbeatMu.Unlock()
	return s.lastHeartbeatTime
}

func (s *Websocket) heartbeatMonitor() {
	ticker := time.NewTicker(HEARTBEAT_CHECK_INTERVAL_SEC * time.Second)
	defer ticker.Stop()

	for {

		<-ticker.C // Wait the appropriate amount of time
		if s.closed.Load() {
			return
		}

		currentTime := time.Now()

		elapsed := currentTime.Unix() - s.getLastHeartbeatTime().Unix()

		// Check if the last heartbeat is older than the close interval
		if elapsed >= HEARTBEAT_CLOSE_ON_NO_HEARTBEAT_SEC {
			s.logger.DEBUG("[HEARTBEAT] No heartbeat detected, socket will terminate.")
			s.markClosed(false)
			s.conn.Close()

			return
		}

		// Check if the last heartbeat is older than the heartbeat check interval
		if elapsed >= HEARTBEAT_CHECK_INTERVAL_SEC {
			err := s.SendPing(nil)
			if err != nil {
				s.logger.ERROR("[HEARTBEAT] Error sending ping", err)

			} else {
				s.logger.DEBUG("[HEARTBEAT] Ping sent.")
			}
		}
	}
}
