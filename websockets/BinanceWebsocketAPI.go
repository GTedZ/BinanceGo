package websockets

import (
	"fmt"
	"time"

	"github.com/GTedZ/binancego/apikeys"
	jsoniter "github.com/json-iterator/go"
)

type BinanceWebsocketAPI struct {
	base *reconnectingPrivateMessageWebsocket

	defaultTimeout_sec int
	rateLimits         []RateLimit

	timestamp_offset int64
	API              apikeys.KeyPair

	OnMessage   func(messageType int, message []byte)
	OnReconnect func()
}

// # Fetches the request weight directly from the response
//
// intervalString => "MINUTE" | "SECOND"
//
// intervalNum    => 1, 2, ...
func (resp *BinanceWebsocketAPI_Response) GetUsedWeight(interval string, intervalNum int) (RateLimit, error) {
	return resp.GetRateLimit("REQUEST_WEIGHT", interval, intervalNum)
}

// # Fetches the orders count ratelimit directly from the response
//
// intervalString => "MINUTE" | "SECOND"
//
// intervalNum    => 1, 2, ...
func (resp *BinanceWebsocketAPI_Response) GetOrderCountLimit(interval string, intervalNum int) (RateLimit, error) {
	return resp.GetRateLimit("ORDERS", interval, intervalNum)
}

// # Fetches the request weight directly from the response
//
// intervalString => "MINUTE" | "SECOND"
//
// intervalNum    => 1, 2, ...
func (resp *BinanceWebsocketAPI_Response) GetRawRequestsLimit(interval string, intervalNum int) (RateLimit, error) {
	return resp.GetRateLimit("RAW_REQUESTS", interval, intervalNum)
}

// # Fetches the ratelimit directly from the response
//
// rateLimitType  => "REQUEST_WEIGHT" | "ORDER" | "RAW_REQUESTS"
//
// intervalString => "MINUTE" | "SECOND"
//
// intervalNum    => 1, 2, ...
func (resp *BinanceWebsocketAPI_Response) GetRateLimit(rateLimitType string, interval string, intervalNum int) (RateLimit, error) {
	for _, rateLimit := range resp.RateLimits {
		if rateLimit.RateLimitType == rateLimitType &&
			rateLimit.Interval == interval &&
			rateLimit.IntervalNum == intervalNum {
			return rateLimit, nil
		}
	}

	return RateLimit{}, fmt.Errorf("ratelimit not found")
}

func (socket *BinanceWebsocketAPI) onMessage(messageType int, msg []byte) {
	Logger.WARN(fmt.Sprintf("Received a non-request message => %d: %s", messageType, msg))
}

func (socket *BinanceWebsocketAPI) onReconnect() {
	if socket.OnReconnect != nil {
		socket.OnReconnect()
	}
}

//// Public methods

func (socket *BinanceWebsocketAPI) SetTimestampOffset(timestamp_offset int64) {
	socket.timestamp_offset = timestamp_offset
}

func (socket *BinanceWebsocketAPI) SendPrivateMessage(message map[string]interface{}) (result []byte, response *BinanceWebsocketAPI_Response, err error) {
	raw_response, _, err := socket.base.SendPrivateMessage(message, socket.defaultTimeout_sec+2)
	if err != nil {
		return nil, nil, err
	}

	err = jsoniter.Unmarshal(raw_response, &response)
	if err != nil {
		fmt.Println("Response", response)
		return nil, response, err
	}

	if response.Error != nil {
		return nil, nil, fmt.Errorf("binance error code: %d msg: %s", response.Error.Code, response.Error.Msg)
	}

	socket.rateLimits = response.RateLimits

	return response.Result, response, nil
}

func (socket *BinanceWebsocketAPI) Send_APIKEY_Request(method string, params map[string]interface{}) (result []byte, response *BinanceWebsocketAPI_Response, err error) {
	params["apiKey"] = socket.API.GetAPIKEY()

	request := make(map[string]interface{})
	request["method"] = method
	request["params"] = params

	return socket.SendPrivateMessage(request)
}

func (socket *BinanceWebsocketAPI) SendSignedRequest(method string, params map[string]interface{}) (result []byte, response *BinanceWebsocketAPI_Response, err error) {
	params["timestamp"] = time.Now().UnixMilli() + socket.timestamp_offset
	params["apiKey"] = socket.API.GetAPIKEY()

	queryString := utils.CreateQueryString(params, true)
	signature, err := socket.API.Sign(queryString)
	if err != nil {
		return nil, nil, err
	}
	params["signature"] = signature

	request := make(map[string]interface{})
	request["method"] = method
	request["params"] = params

	return socket.SendPrivateMessage(request)
}

func (socket *BinanceWebsocketAPI) GetRateLimits() []RateLimit {
	return socket.rateLimits
}
func (socket *BinanceWebsocketAPI) Close() {
	socket.base.Close()
}

////

func CreateBinanceWebsocketAPI(baseURL string, defaultTimeout_sec int, API apikeys.KeyPair) *BinanceWebsocketAPI {
	var socket = &BinanceWebsocketAPI{
		defaultTimeout_sec: defaultTimeout_sec,
		API:                API,
	}

	socket.base = createReconnectingPrivateMessageWebsocket(baseURL, "id")

	socket.base.OnMessage = socket.onMessage
	socket.base.OnReconnect = socket.onReconnect

	return socket
}
