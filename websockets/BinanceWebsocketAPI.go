package websockets

import (
	"fmt"
	"time"

	jsoniter "github.com/json-iterator/go"
)

type BinanceWebsocketAPI struct {
	base *reconnectingPrivateMessageWebsocket

	baseURL            string
	defaultTimeout_sec int
	rateLimits         []RateLimit

	timestamp_offset int64
	apikey           string
	apisecret        string

	OnMessage   func(messageType int, message []byte)
	OnReconnect func()
}

type RateLimit struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	IntervalNum   int    `json:"intervalNum"`
	Limit         int    `json:"limit"`
	Count         int    `json:"count"`
}

type BinanceWebsocketAPI_Response struct {
	Id         string              `json:"id"`
	Status     int                 `json:"status"`
	Error      *BinanceError       `json:"error"`
	Result     jsoniter.RawMessage `json:"result"`
	RateLimits []RateLimit         `json:"rateLimits"`
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

type BinanceError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
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

func (socket *BinanceWebsocketAPI) SetAPIKeys(APIKEY string, APISECRET string) {
	socket.apikey = APIKEY
	socket.apisecret = APISECRET
}

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

func (socket *BinanceWebsocketAPI) Send_APIKEY_Request(message map[string]interface{}) (result []byte, response *BinanceWebsocketAPI_Response, err error) {
	message["apiKey"] = socket.apikey

	return socket.SendPrivateMessage(message)
}

func (socket *BinanceWebsocketAPI) SendSignedRequest(message map[string]interface{}) (result []byte, response *BinanceWebsocketAPI_Response, err error) {
	message["timestamp"] = time.Now().UnixMilli() + socket.timestamp_offset
	message["apiKey"] = socket.apikey

	queryString := utils.CreateQueryString(message, true)
	signature, err := utils.CreateHMACSignature(queryString, socket.apisecret)
	if err != nil {
		return nil, nil, err
	}
	message["signature"] = signature

	return socket.SendPrivateMessage(message)
}

func (socket *BinanceWebsocketAPI) GetRateLimits() []RateLimit {
	return socket.rateLimits
}
func (socket *BinanceWebsocketAPI) Close() {
	socket.base.Close()
}

////

func CreateBinanceWebsocketAPI(baseURL string, defaultTimeout_sec int, APIKEY string, APISECRET string) *BinanceWebsocketAPI {
	var socket = &BinanceWebsocketAPI{
		baseURL:            baseURL,
		defaultTimeout_sec: defaultTimeout_sec,
		apikey:             APIKEY,
		apisecret:          APISECRET,
	}

	socket.base = createReconnectingPrivateMessageWebsocket(baseURL, "id")

	socket.base.OnMessage = socket.onMessage
	socket.base.OnReconnect = socket.onReconnect

	return socket
}
