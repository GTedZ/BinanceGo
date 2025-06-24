package websockets

import jsoniter "github.com/json-iterator/go"

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

type BinanceError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
