package base

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/GTedZ/binancego/internal/logging"
	ws "github.com/gorilla/websocket"
)

// This websocket handles simply connecting to binance, reading the messages and successfully disconnecting once the connection goes dead
//
// And handles sending requests via the Id system
type Websocket struct {
	// Host server's URL
	url               string
	lastHeartbeatTime time.Time
	heartbeatMu       sync.Mutex
	closed            atomic.Bool

	conn    *ws.Conn
	writeMu sync.Mutex

	logger logging.Logger

	// User callbacks
	onMessage func(data []byte)
	onClose   func()
}

func (s *Websocket) broadcastOnMessage(data []byte) {
	if s.onMessage != nil {
		s.onMessage(data)
	}
}

func (s *Websocket) broadcastOnClose() {
	if s.onClose != nil {
		s.onClose()
	}
}

////

func New(
	URL string,
	onMessage func([]byte),
	onClose func(),
	logger logging.Logger,

) (*Websocket, error) {
	var socket = &Websocket{
		url:    URL,
		logger: logger,
	}

	socket.recordLastHeartbeat()

	socket.SetOnMessage(onMessage)
	socket.SetOnClose(onClose)

	logger.DEBUG(fmt.Sprintf("Socket connecting to '%v'", URL))

	conn, _, err := ws.DefaultDialer.Dial(URL, nil)
	if err != nil {
		logger.ERROR("There was an error creating websocket", err)
		return nil, err
	}
	socket.conn = conn
	logger.DEBUG(fmt.Sprintf("Socket connected: '%v'", URL))

	////

	socket.conn.SetPingHandler(socket.pingHandler)
	socket.conn.SetPongHandler(socket.pongHandler)
	socket.conn.SetCloseHandler(socket.closeHandler)

	////

	go socket.listen()
	go socket.heartbeatMonitor()

	return socket, nil
}
