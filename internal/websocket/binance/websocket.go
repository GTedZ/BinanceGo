package binance

import (
	"sync"

	"github.com/GTedZ/binancego/internal/logging"
	"github.com/GTedZ/binancego/internal/websocket/reconnecting"
)

type Websocket struct {
	base *reconnecting.Websocket
	url  string

	requestPropertyKey string

	logger logging.Logger

	requests struct {
		mu   sync.Mutex
		data map[string]chan []byte
		once map[string]*sync.Once
	}

	////
	// User callbacks
	////

	onMessage func(data []byte)

	// When the current base socket closes
	onDisconnect func()

	// When new socket is successfully connected
	onReconnected func()
}

//

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

//

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

		requestPropertyKey: "id",

		logger: logger,
	}

	socket.requests.data = make(map[string]chan []byte)
	socket.requests.once = make(map[string]*sync.Once)

	// Setting base properties
	socket.SetOnMessage(onMessage)
	socket.SetOnDisconnect(onDisconnect)
	socket.SetOnReconnected(onReconnected)

	base, err := reconnecting.New(URL, socket.messageHandler, socket.disconnectHandler, socket.reconnectedHandler, logger)
	if err != nil {
		return nil, err
	}

	socket.base = base

	return socket, nil
}
