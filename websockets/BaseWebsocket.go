package websockets

import (
	"fmt"
	"sync/atomic"
	"time"

	ws "github.com/gorilla/websocket"
)

const (
	HEARTBEAT_CHECK_INTERVAL_SEC        = 5
	HEARTBEAT_CLOSE_ON_NO_HEARTBEAT_SEC = 20
)

// type websocket_CombinedStream_Message struct {
// 	Stream string              `json:"stream"`
// 	Data   jsoniter.RawMessage `json:"data"`
// }

// This websocket handles simply connecting to binance, reading the messages and successfully disconnecting once the connection goes dead
//
// And handles sending requests via the Id system
type baseWebsocket struct {
	// Host server's URL
	url                     string
	lastHeartbeat_Timestamp time.Time
	closed                  atomic.Bool

	conn *ws.Conn

	OnMessage func(messageType int, msg []byte)
	OnClose   func()
}

func (socket *baseWebsocket) markAsClosed() {
	// CompareAndSwap returns "true" if the swap was successful
	// Meaning that socket.closed was false, which means that if it were true, we'd want to emit the OnClose()
	wasFalseAndHasBeenSwappedToTrue := socket.closed.CompareAndSwap(false, true)

	if !wasFalseAndHasBeenSwappedToTrue {
		Logger.DEBUG(fmt.Sprintf("[%s] socket already marked as closed", socket.url))
		return
	}

	if socket.OnClose != nil {
		socket.OnClose()
	}
}

func (socket *baseWebsocket) onPing(pingData string) error {
	Logger.DEBUG(fmt.Sprintf("Received a ping: %s", pingData))

	socket.recordLastHeartbeat() // Logs a heartbeat

	err := socket.conn.WriteMessage(ws.PongMessage, []byte(pingData))
	if err != nil {
		Logger.ERROR("Error sending Pong:", err)
		return err
	}

	return nil
}

func (socket *baseWebsocket) onPong(appData string) error {
	socket.recordLastHeartbeat()

	return nil
}

func (socket *baseWebsocket) closeHandler(code int, text string) error {
	Logger.DEBUG(fmt.Sprintf("[*Websocket.CloseHandler()] code: %d, text: %s, isClosed: %v", code, text, socket.closed.Load()))
	socket.markAsClosed()

	return nil
}

func (socket *baseWebsocket) onMessage(messageType int, msg []byte) {
	if socket.closed.Load() {
		return
	}

	if socket.OnMessage != nil {
		socket.OnMessage(messageType, msg)
	}
}

func (socket *baseWebsocket) recordLastHeartbeat() {
	socket.lastHeartbeat_Timestamp = time.Now()
}

func (socket *baseWebsocket) listen() {
	// Goroutine to read messages
	for {
		if socket.closed.Load() {
			return
		}
		msgType, msg, err := socket.conn.ReadMessage()
		if err != nil {
			Logger.ERROR(fmt.Sprintf("[%s] Error reading message", socket.url), err)
			socket.markAsClosed()
			return
		}

		Logger.DEBUG(fmt.Sprintf("Type: %d, message: %s\n", msgType, string(msg)))
		socket.recordLastHeartbeat()

		socket.onMessage(msgType, msg)
	}
}

func (socket *baseWebsocket) checkHeartbeats() {
	ticker := time.NewTicker(HEARTBEAT_CHECK_INTERVAL_SEC * time.Second)
	defer ticker.Stop()

	for {

		<-ticker.C // Wait the appropriate amount of time
		if socket.closed.Load() {
			Logger.DEBUG("[HEARTBEAT] Websocket is closed, skipping check.")

			return
		}

		currentTime := time.Now()

		elapsed := currentTime.Unix() - socket.lastHeartbeat_Timestamp.Unix()

		// Check if the last heartbeat is older than the close interval
		if elapsed >= HEARTBEAT_CLOSE_ON_NO_HEARTBEAT_SEC {
			Logger.DEBUG("[HEARTBEAT] No heartbeat detected, socket will terminate.")
			return
		}

		// Check if the last heartbeat is older than the heartbeat check interval
		if elapsed >= HEARTBEAT_CHECK_INTERVAL_SEC {
			err := socket.conn.WriteMessage(ws.PingMessage, nil)
			if err != nil {
				Logger.ERROR("[HEARTBEAT] Error sending ping", err)

			} else {
				Logger.DEBUG("[HEARTBEAT] Ping sent.")
			}
		}
	}
}

//// Public methods

func (socket *baseWebsocket) Close() {
	socket.conn.Close()

	socket.markAsClosed()
}

func (socket *baseWebsocket) SendJSON(v interface{}) error {
	return socket.conn.WriteJSON(v)
}

////

func createBaseSocket(URL string) (*baseWebsocket, error) {
	var socket = &baseWebsocket{
		url:                     URL,
		lastHeartbeat_Timestamp: time.Now(),
	}

	conn, _, err := ws.DefaultDialer.Dial(URL, nil)
	if err != nil {
		Logger.ERROR("There was an error creating websocket", err)
		return nil, err
	}
	socket.conn = conn
	Logger.DEBUG(fmt.Sprintf("Socket connected: %v", URL))

	////

	socket.conn.SetPingHandler(socket.onPing)
	socket.conn.SetPongHandler(socket.onPong)
	socket.conn.SetCloseHandler(socket.closeHandler)

	////

	go socket.listen()
	go socket.checkHeartbeats()

	return socket, nil
}
