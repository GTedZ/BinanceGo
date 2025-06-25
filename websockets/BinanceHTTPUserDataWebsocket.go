package websockets

import (
	"time"
)

type BinanceHTTPUserDataWebsocket struct {
	base *reconnectingPrivateMessageWebsocket

	baseURL string

	listenKey_func func() (string, error)
	keepAlive_func func(string) error

	listenKey_prefix                string
	listenKey_expirationTime_millis int64
	listenKey                       string

	OnMessage   func(messageType int, msg []byte)
	OnReconnect func()
	OnError     func(error)
	OnClose     func()
}

////

func (socket *BinanceHTTPUserDataWebsocket) onMessage(messageType int, msg []byte) {
	if socket.OnMessage != nil {
		socket.OnMessage(messageType, msg)
	}
}

func (socket *BinanceHTTPUserDataWebsocket) onReconnect() {
	if socket.OnReconnect != nil {
		socket.OnReconnect()
	}
}

func (socket *BinanceHTTPUserDataWebsocket) onError(err error) {
	if socket.OnError != nil {
		socket.OnError(err)
	}
}

func (socket *BinanceHTTPUserDataWebsocket) onClose() {
	if socket.OnClose != nil {
		socket.OnClose()
	}
}

func (socket *BinanceHTTPUserDataWebsocket) newListenKey() (err error) {
	socket.listenKey, err = socket.listenKey_func()
	return err
}

func (socket *BinanceHTTPUserDataWebsocket) buildURL() string {
	return socket.baseURL + socket.listenKey_prefix + socket.listenKey
}

func (socket *BinanceHTTPUserDataWebsocket) keepAliveLoop(MAX_TRIES int64) {
	var errCount int64 = 0
	for {
		time.Sleep(time.Duration(socket.listenKey_expirationTime_millis/MAX_TRIES) * time.Millisecond)

		// Check if socket is closed
		if socket.base.closed.Load() {
			return
		}

		err := socket.keepAlive_func(socket.listenKey)
		if err != nil {
			errCount++
			socket.onError(err)

			if errCount == MAX_TRIES {
				socket.Close()
				return
			}

			continue
		} else {
			errCount = 0
		}
	}
}

//// Public Methods

func (socket *BinanceHTTPUserDataWebsocket) SetOnMessage(onMessage func(messageType int, msg []byte)) {
	socket.OnMessage = onMessage
}

func (socket *BinanceHTTPUserDataWebsocket) SetOnReconnect(onReconnect func()) {
	socket.OnReconnect = onReconnect
}

func (socket *BinanceHTTPUserDataWebsocket) SetOnError(onError func(error)) {
	socket.OnError = onError
}

func (socket *BinanceHTTPUserDataWebsocket) SetOnClose(onClose func()) {
	socket.OnClose = onClose
}

func (socket *BinanceHTTPUserDataWebsocket) RestartUserDataStream() error {
	err := socket.newListenKey()
	if err != nil {
		return err
	}

	socket.base.SetURL(socket.buildURL())
	socket.base.Reconnect()

	return nil
}

func (socket *BinanceHTTPUserDataWebsocket) Close() {
	socket.base.Close()

	socket.onClose()
}

////

func CreateHTTPUserDataWebsocket(baseURL string, listenKey_prefix string, listenKey_expirationTime_mins int64, listenKey_func func() (string, error), keepAlive_func func(string) error) (*BinanceHTTPUserDataWebsocket, error) {
	var socket = &BinanceHTTPUserDataWebsocket{
		baseURL: baseURL,

		listenKey_func: listenKey_func,
		keepAlive_func: keepAlive_func,

		listenKey_prefix:                listenKey_prefix,
		listenKey_expirationTime_millis: listenKey_expirationTime_mins * 60 * 1000,
	}

	err := socket.newListenKey()
	if err != nil {
		return nil, err
	}

	baseSocket, err := createReconnectingPrivateMessageWebsocket(socket.buildURL(), "id")
	if err != nil {
		return nil, err
	}
	socket.base = baseSocket

	socket.base.OnMessage = socket.onMessage
	socket.base.OnReconnect = socket.onReconnect

	go socket.keepAliveLoop(10)

	return socket, nil
}
