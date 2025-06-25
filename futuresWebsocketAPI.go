package Binance

import (
	"github.com/GTedZ/binancego/websockets"
)

type Futures_WebsocketAPI struct {
	binance *Binance
}

func (futures_WSAPI *Futures_WebsocketAPI) init(binance *Binance) {
	futures_WSAPI.binance = binance
}

////

type Futures_WSAPI_Socket struct {
	base *websockets.BinanceWebsocketAPI
}

func (socket *Futures_WSAPI_Socket) GetRateLimits() []websockets.RateLimit {
	return socket.base.GetRateLimits()
}

//

type FuturesWSAPI_OrderBook struct {
	// 1027024
	LastUpdateID int `json:"lastUpdateId"`

	// 1589436922972   // Message output time
	EventTime int64 `json:"E"`

	// 1589436922959   // Transaction time
	TransactionTime int64 `json:"T"`

	// [[ "4.00000000", "431.00000000" ]]
	Bids [][2]string `json:"bids"`

	// [[ "4.00000200", "12.00000000" ]]
	Asks [][2]string `json:"asks"`
}

type FuturesWSAPI_OrderBook_Params struct {
	Symbol string `json:"symbol"`
	// optional
	Limit int `json:"limit"`
}

func (socket *Futures_WSAPI_Socket) OrderBook(symbol string, limit ...int) (orderbook *FuturesWSAPI_OrderBook, resp *websockets.BinanceWebsocketAPI_Response, err error) {

	params := make(map[string]interface{})
	params["symbol"] = symbol
	if len(limit) != 0 {
		params["limit"] = limit[0]
	}

	request := make(map[string]interface{})
	request["method"] = "depth"
	request["params"] = params

	data, resp, err := socket.base.SendPrivateMessage(request)
	if err != nil {
		return nil, resp, err
	}

	err = json.Unmarshal(data, &orderbook)
	if err != nil {
		return nil, resp, err
	}

	return orderbook, resp, nil
}

//

type FuturesWSAPI_PriceTicker struct {
	// "BTCUSDT"
	Symbol string `json:"symbol"`

	// "6000.01"
	Price string `json:"price"`

	// 1589437530011   // Transaction time
	Time int64 `json:"time"`
}

func (socket *Futures_WSAPI_Socket) PriceTicker(symbol ...string) (priceTickers []*FuturesWSAPI_PriceTicker, resp *websockets.BinanceWebsocketAPI_Response, err error) {
	request := make(map[string]interface{})
	request["method"] = "ticker.price"

	if len(symbol) != 0 {
		request["params"] = struct {
			Symbol string `json:"symbol"`
		}{
			Symbol: symbol[0],
		}
	}

	data, resp, err := socket.base.SendPrivateMessage(request)
	if err != nil {
		return nil, resp, err
	}

	if len(symbol) == 1 {
		var priceTicker *FuturesWSAPI_PriceTicker
		err = json.Unmarshal(data, &priceTicker)
		priceTickers = []*FuturesWSAPI_PriceTicker{priceTicker}
	} else {
		err = json.Unmarshal(data, &priceTickers)
	}

	if err != nil {
		return nil, resp, err
	}

	return priceTickers, resp, nil
}

//

type FuturesWSAPI_BookTicker struct {
	// 1027024
	LastUpdateID int `json:"lastUpdateId"`

	// "BTCUSDT"
	Symbol string `json:"symbol"`

	// "4.00000000"
	BidPrice string `json:"bidPrice"`

	// "431.00000000"
	BidQty string `json:"bidQty"`

	// "4.00000200"
	AskPrice string `json:"askPrice"`

	// "9.00000000"
	AskQty string `json:"askQty"`

	// 1589437530011   // Transaction time
	Time int64 `json:"time"`
}

func (socket *Futures_WSAPI_Socket) BookTicker(symbol ...string) (bookTickers []*FuturesWSAPI_BookTicker, resp *websockets.BinanceWebsocketAPI_Response, err error) {
	request := make(map[string]interface{})
	request["method"] = "ticker.book"

	if len(symbol) != 0 {
		request["params"] = struct {
			Symbol string `json:"symbol"`
		}{
			Symbol: symbol[0],
		}
	}

	data, resp, err := socket.base.SendPrivateMessage(request)
	if err != nil {
		return nil, resp, err
	}

	if len(symbol) == 1 {
		var priceTicker *FuturesWSAPI_BookTicker
		err = json.Unmarshal(data, &priceTicker)
		bookTickers = []*FuturesWSAPI_BookTicker{priceTicker}
	} else {
		err = json.Unmarshal(data, &bookTickers)
	}

	if err != nil {
		return nil, resp, err
	}

	return bookTickers, resp, nil
}

//

////

func (futures_WSAPI *Futures_WebsocketAPI) NewWebsocketAPI() (*Futures_WSAPI_Socket, error) {
	var socket Futures_WSAPI_Socket
	baseSocket, err := websockets.CreateBinanceWebsocketAPI(FUTURES_Constants.WebsocketAPI.URL, FUTURES_Constants.WebsocketAPI.DefaultRequestTimeout_sec, futures_WSAPI.binance.API)
	if err != nil {
		return nil, err
	}

	socket.base = baseSocket

	return &socket, nil
}
