package Binance

type Spot_WebsocketAPI struct {
	binance *Binance
}

// func (spot_websocketAPI *Spot_WebsocketAPI) init(binance *Binance) {
// 	spot_websocketAPI.binance = binance
// }

// type Spot_WebsocketAPI_Socket struct {
// 	Websocket *Websocket
// 	Conn      *ws.Conn
// 	// Host server's URL
// 	BaseURL string
// }

// func (spot_ws *Spot_WebsocketAPI_Socket) Close() error {
// 	return spot_ws.Websocket.Close()
// }

// // Forcefully reconnects the socket
// // Also makes it a reconnecting socket if it weren't before
// // Useless, but there nonetheless...
// func (spot_ws *Spot_WebsocketAPI_Socket) Reconnect() {
// 	spot_ws.Websocket.Reconnect()
// }

// func (spot_ws *Spot_WebsocketAPI_Socket) SetMessageListener(f func(messageType int, msg []byte)) {
// 	spot_ws.Websocket.OnMessage = f
// }

// func (spot_ws *Spot_WebsocketAPI_Socket) SetPingListener(f func(appData string)) {
// 	spot_ws.Websocket.OnPing = f
// }

// func (spot_ws *Spot_WebsocketAPI_Socket) SetPongListener(f func(appData string)) {
// 	spot_ws.Websocket.OnPong = f
// }

// // This is called when socket has been disconnected
// // Called when the detected a disconnection and wants to reconnect afterwards
// // Usually called right before the 'ReconnectingListener'
// func (spot_ws *Spot_WebsocketAPI_Socket) SetDisconnectListener(f func(code int, text string)) {
// 	spot_ws.Websocket.OnDisconnect = f
// }

// // This is called when socket began reconnecting
// func (spot_ws *Spot_WebsocketAPI_Socket) SetReconnectingListener(f func()) {
// 	spot_ws.Websocket.OnReconnecting = f
// }

// // This is called when the socket has successfully reconnected after a disconnection
// func (spot_ws *Spot_WebsocketAPI_Socket) SetReconnectListener(f func()) {
// 	spot_ws.Websocket.OnReconnect = f
// }

// // This is called when the websocket closes indefinitely
// // Meaning when you invoke the 'Close()' method
// // Or any other way a websocket is set to never reconnect on a disconnection
// func (spot_ws *Spot_WebsocketAPI_Socket) SetCloseListener(f func(code int, text string)) {
// 	spot_ws.Websocket.OnClose = f
// }

// ////

// type SpotWebsocketAPI_Request struct {
// 	Method string
// 	Params map[string]interface{}
// }

// func (socket *Spot_WebsocketAPI_Socket) createRequestObject(request SpotWebsocketAPI_Request) map[string]interface{} {
// 	requestObj := make(map[string]interface{})

// 	requestObj["id"] = uuid.New().String()
// 	requestObj["method"] = request.Method
// 	requestObj["params"] = request.Params

// 	return requestObj
// }

// ////

// func (websocketAPI *Spot_WebsocketAPI) CreateSocket() (*Spot_WebsocketAPI_Socket, *Error) {
// 	URL := SPOT_Constants.WebsocketAPI.URL
// 	socket, err := CreateSocket(URL, nil, false)
// 	if err != nil {
// 		return nil, err
// 	}

// 	socket.privateMessageValidator = func(msg []byte) (isPrivate bool, Id string) {
// 		if len(msg) > 0 && msg[0] == '[' {
// 			return false, ""
// 		}

// 		var privateMessage SpotWS_PrivateMessage
// 		err := json.Unmarshal(msg, &privateMessage)
// 		if err != nil {
// 			Logger.ERROR(fmt.Sprint("[PRIVATEMESSAGEVALIDATOR ERR] WS Message is the following:", string(msg)))
// 			LocalError(PARSING_ERR, err.Error())
// 			return false, ""
// 		}

// 		if privateMessage.Id == "" {
// 			return false, ""
// 		}

// 		return true, privateMessage.Id
// 	}

// 	ws := &Spot_WebsocketAPI_Socket{
// 		Websocket: socket,
// 		Conn:      socket.Conn,
// 		BaseURL:   URL,
// 	}

// 	return ws, nil
// }

// ////
