package Binance

import (
	"fmt"
	"time"

	"github.com/GTedZ/binancego/websockets"
)

type Spot_WebsocketAPI struct {
	binance *Binance
}

func (spot_WSAPI *Spot_WebsocketAPI) init(binance *Binance) {
	spot_WSAPI.binance = binance
}

////

type spot_WSAPI_Socket struct {
	base *websockets.BinanceWebsocketAPI
}

func (socket *spot_WSAPI_Socket) GetRateLimits() []websockets.RateLimit {
	return socket.base.GetRateLimits()
}

//

type SpotWSAPI_Ping struct {
	// roundtrip time of the request in milliseconds
	Latency int64
}

func (socket *spot_WSAPI_Socket) Ping() (*SpotWSAPI_Ping, *websockets.BinanceWebsocketAPI_Response, error) {
	request := make(map[string]interface{})
	request["method"] = "ping"

	startTime := time.Now()
	_, resp, err := socket.base.SendPrivateMessage(request)
	ping := SpotWSAPI_Ping{
		Latency: time.Now().UnixMilli() - startTime.UnixMilli(),
	}

	return &ping, resp, err
}

//

type SpotWSAPI_ServerTime struct {
	ServerTime int64
	// roundtrip time of the request in milliseconds
	Latency int64
}

func (socket *spot_WSAPI_Socket) ServerTime() (*SpotWSAPI_ServerTime, *websockets.BinanceWebsocketAPI_Response, error) {
	request := make(map[string]interface{})
	request["method"] = "time"

	startTime := time.Now()
	data, resp, err := socket.base.SendPrivateMessage(request)
	if err != nil {
		return nil, resp, err
	}
	var serverTime SpotWSAPI_ServerTime

	err = json.Unmarshal(data, &serverTime)
	if err != nil {
		return nil, resp, err
	}
	serverTime.Latency = time.Now().UnixMilli() - startTime.UnixMilli()

	return &serverTime, resp, err
}

//

func (socket *spot_WSAPI_Socket) ExchangeInfo(opt_params ...Spot_ExchangeInfo_Params) (*Spot_ExchangeInfo, *websockets.BinanceWebsocketAPI_Response, error) {
	request := make(map[string]interface{})
	request["method"] = "exchangeInfo"

	if len(opt_params) != 0 {
		params := make(map[string]interface{})

		requestParams := opt_params[0]

		if len(requestParams.Symbols) != 0 {
			params["symbols"] = requestParams.Symbols
		} else if isDifferentFromDefault(requestParams.Symbol) {
			params["symbol"] = requestParams.Symbol
		} else {
			if len(requestParams.Permissions) != 0 {
				params["permissions"] = requestParams.Permissions
			}
			if isDifferentFromDefault(requestParams.SymbolStatus) {
				params["symbolStatus"] = requestParams.SymbolStatus
			}
		}
		if requestParams.DontShowPermissionSets {
			params["showPermissionSets"] = !requestParams.DontShowPermissionSets
		}

		request["params"] = params
	}

	data, resp, err := socket.base.SendPrivateMessage(request)
	if err != nil {
		return nil, resp, err
	}

	exchangeInfo, exchParseErr := parseSpotExchangeInfo(data)
	if exchParseErr != nil {
		return nil, resp, fmt.Errorf("failed to parse spot exchangeInfo")
	}

	return exchangeInfo, resp, err
}

////

func (spot_WSAPI *Spot_WebsocketAPI) NewWebsocketAPI() (*spot_WSAPI_Socket, error) {
	var socket spot_WSAPI_Socket
	baseSocket, err := websockets.CreateBinanceWebsocketAPI(SPOT_Constants.WebsocketAPI.URLs[0], SPOT_Constants.WebsocketAPI.DefaultRequestTimeout_sec, spot_WSAPI.binance.API)
	if err != nil {
		return nil, err
	}

	socket.base = baseSocket

	return &socket, nil
}
