package websockets

import (
	"fmt"
	"math"
	"sync/atomic"
	"time"
)

type reconnectingPrivateMessageWebsocket struct {
	base                       *privateMessageWebsocket
	privateMessagePropertyName string

	url    string
	ready  atomic.Bool
	closed atomic.Bool

	OnMessage   func(messageType int, msg []byte)
	OnReconnect func()
}

func (socket *reconnectingPrivateMessageWebsocket) onSubsocketClosed() {
	// The reconnecting socket has been terminally closed, so no need to reconnect or do anything
	if socket.closed.Load() {
		return
	}

	go socket.newSubsocket()
}

func (socket *reconnectingPrivateMessageWebsocket) newSubsocket() {
	socket.ready.Store(false)

	var newSocket *privateMessageWebsocket
	var err error
	var retries int = 0
	for {
		Logger.DEBUG(fmt.Sprintf("[%s] Connecting to new subsocket", socket.url))
		newSocket, err = createPrivateMessageWebsocket(socket.url, socket.privateMessagePropertyName)
		if err != nil {
			Logger.ERROR(fmt.Sprintf("[%s] Failed to open subsocket", socket.url), err)

			extraDuration := int(math.Min(float64(retries*250), 2000))
			duration := time.Duration(500 + extraDuration)
			time.Sleep(duration * time.Millisecond)
			retries++
			continue
		}
		break
	}

	socket.base = newSocket

	socket.base.OnClose = socket.onSubsocketClosed
	socket.base.OnMessage = socket.onMessage

	socket.ready.Store(true)

	if socket.OnReconnect != nil {
		socket.OnReconnect()
	}
}

func (socket *reconnectingPrivateMessageWebsocket) onMessage(messageType int, msg []byte) {
	if socket.OnMessage != nil {
		socket.OnMessage(messageType, msg)
	}
}

//// Public methods

func (socket *reconnectingPrivateMessageWebsocket) SetURL(URL string) {
	socket.url = URL
}

func (socket *reconnectingPrivateMessageWebsocket) Reconnect() {
	socket.base.Close()
}

func (socket *reconnectingPrivateMessageWebsocket) Close() {
	socket.closed.Store(true)

	socket.base.Close()
}

func (socket *reconnectingPrivateMessageWebsocket) SendPrivateMessage(message map[string]interface{}, timeout_sec ...int) (response []byte, hasTimedOut bool, err error) {
	for {
		if socket.closed.Load() {
			return nil, false, fmt.Errorf("socket has been closed")
		}
		if !socket.ready.Load() {
			time.Sleep(500 * time.Millisecond)
			continue
		}
		break
	}
	return socket.base.SendPrivateMessage(message, timeout_sec...)
}

////

// WARNING: Since this is a reconnecting socket, it will NEVER error out when trying to connect to the server, it will block until a successful connection is established
func createReconnectingPrivateMessageWebsocket(URL string, privateMessagePropertyName string) *reconnectingPrivateMessageWebsocket {
	var socket = &reconnectingPrivateMessageWebsocket{
		url:                        URL,
		privateMessagePropertyName: privateMessagePropertyName,
	}

	socket.newSubsocket()

	return socket
}
