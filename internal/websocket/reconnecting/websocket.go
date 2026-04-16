package reconnecting

import (
	"sync"

	"github.com/GTedZ/binancego/internal/logging"
	"github.com/GTedZ/binancego/internal/websocket/base"
)

type Websocket struct {
	base *base.Websocket
	url  string

	mu sync.Mutex

	reconnecting bool
	closed       bool

	logger logging.Logger

	////
	// User callbacks
	////

	onMessage func(data []byte)

	// When the current base socket closes
	onDisconnect func()

	// When new socket is successfully connected
	onReconnected func()
}

////

func (s *Websocket) newSubSocket() error {
	// Connecting socket
	base, err := base.New(s.url, s.messageHandler, s.closeHandler, s.logger)
	if err != nil {
		return err
	}

	// Setting all properties
	s.base = base

	return nil
}

func (s *Websocket) broadcastOnMessage(data []byte) {
	if s.onMessage != nil {
		s.onMessage(data)
	}
}

func (s *Websocket) broadcastOnDisconnect() {
	if s.onDisconnect != nil {
		s.onDisconnect()
	}
}

func (s *Websocket) broadcastOnReconnected() {
	if s.onReconnected != nil {
		s.onReconnected()
	}
}

// Creates a reconnecting websocket connection
// Only errors if initial connection cannot be established
func New(
	URL string,
	onMessage func(data []byte),
	onDisconnect func(),
	onReconnected func(),
	logger logging.Logger,

) (*Websocket, error) {
	socket := &Websocket{
		url: URL,

		logger: logger,
	}

	// Setting base properties
	socket.SetOnMessage(onMessage)
	socket.SetOnDisconnect(onDisconnect)
	socket.SetOnReconnected(onReconnected)

	// Sub socket
	err := socket.newSubSocket()
	if err != nil {
		return nil, err
	}

	return socket, nil
}
