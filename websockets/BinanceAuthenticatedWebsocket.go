package websockets

// type RateLimit struct {
// 	RateLimitType string `json:"rateLimitType"`
// 	Interval      string `json:"interval"`
// 	IntervalNum   int    `json:"intervalNum"`
// 	Limit         int    `json:"limit"`
// 	Count         int    `json:"count"`
// }

// type BinanceAuthenticatedWebsocket struct {
// 	base *reconnectingPrivateMessageWebsocket

// 	apikey    string
// 	apisecret string

// 	Ratelimits []*RateLimit
// }

// type BinanceAuthenticatedWebsocket_Status struct {
// 	ApiKey           string `json:"apiKey"`
// 	AuthorizedSince  int64  `json:"authorizedSince"`
// 	ConnectedSince   int64  `json:"connectedSince"`
// 	ReturnRateLimits bool   `json:"returnRateLimits"`
// 	ServerTime       int64  `json:"serverTime"`
// }

// ////

// func (socket *BinanceAuthenticatedWebsocket) logon() error {
// 	request := make(map[string]interface{})
// 	request["method"] = "session.logon"
// 	request["params"] = struct {
// 		ApiKey    string `json:"apiKey"`
// 		Signature string `json:"signature"`
// 		Timestamp int64  `json:"timestamp"`
// 	}{
// 		ApiKey: socket.apikey,

// 	}

// 	_, _, err := socket.base.SendPrivateMessage(request)
// 	if err != nil {
// 		Logger.ERROR("Failed to logon for authenticated websocket")
// 		return err
// 	}

// 	return nil
// }

// func (socket *BinanceAuthenticatedWebsocket) GetStatus()

// // This function WILL block until a successful connection can be established
// func CreateBinanceWebsocketAPI(baseURL string, APIKEY string, APISECRET string) (*BinanceWebsocket, error) {
// 	var socket = &BinanceAuthenticatedWebsocket{
// 		apikey: APIKEY,
// 	}

// 	socket.base = createReconnectingPrivateMessageWebsocket(baseURL, "id")

// 	err := socket.logon()

// 	socket.base.OnMessage = socket.onMessage
// 	socket.base.OnReconnect = socket.onReconnect

// 	return socket
// }
