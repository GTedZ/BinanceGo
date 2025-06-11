package Binance

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"slices"

	websockets "github.com/GTedZ/binancego/websockets"
)

type Futures_Websockets struct {
	binance *Binance
}

func (futures_ws *Futures_Websockets) init(binance *Binance) {
	futures_ws.binance = binance
}

type Futures_Websocket struct {
	base *websockets.BinanceWebsocket
}

func (futures_ws *Futures_Websocket) Close() {
	futures_ws.base.Close()
}

func (futures_ws *Futures_Websocket) SetMessageListener(f func(messageType int, msg []byte)) {
	futures_ws.base.OnMessage = f
}

func (futures_ws *Futures_Websocket) ListSubscriptions(timeout_sec ...int) (subscriptions []string, err error) {
	return futures_ws.base.ListSubscriptions(timeout_sec...)
}

func (futures_ws *Futures_Websocket) Subscribe(stream ...string) (hasTimedOut bool, err error) {
	requestObj := make(map[string]interface{})
	requestObj["method"] = "SUBSCRIBE"
	requestObj["params"] = stream
	_, timedOut, err := futures_ws.base.SendPrivateMessage(requestObj)
	if err != nil || timedOut {
		return timedOut, err
	}

	futures_ws.base.SetStreams(append(futures_ws.base.GetStreams(), stream...))
	futures_ws.base.UpdateStreams()
	Logger.INFO(fmt.Sprint("Successfully Subscribed to", stream))

	return false, nil
}

func (futures_ws *Futures_Websocket) Unsubscribe(stream ...string) (hasTimedOut bool, err error) {
	requestObj := make(map[string]interface{})
	requestObj["method"] = "UNSUBSCRIBE"
	requestObj["params"] = stream
	_, timedOut, err := futures_ws.base.SendPrivateMessage(requestObj)
	if err != nil || timedOut {
		return timedOut, err
	}

	// Filter out the streams to remove from futures_ws.Websocket.Streams
	streamMap := make(map[string]bool)
	for _, s := range stream {
		streamMap[s] = true
	}

	var updatedStreams []string
	for _, existingStream := range futures_ws.base.GetStreams() {
		if !streamMap[existingStream] {
			updatedStreams = append(updatedStreams, existingStream)
		}
	}
	futures_ws.base.SetStreams(updatedStreams)

	Logger.INFO(fmt.Sprint("Successfully Unsubscribed from", stream))

	return false, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type FuturesWS_AggTrade struct {

	// Event type
	Event string `json:"e"`

	// Event time
	EventTime int64 `json:"E"`

	// Symbol
	Symbol string `json:"s"`

	// Aggregate trade ID
	AggTradeId int64 `json:"a"`

	// Price
	Price string `json:"p"`

	// Quantity
	Quantity string `json:"q"`

	// First trade ID
	FirstTradeId int64 `json:"f"`

	// Last trade ID
	LastTradeId int64 `json:"l"`

	// Trade time
	Timestamp int64 `json:"T"`

	// Is the buyer the market maker?
	IsMaker bool `json:"m"`
}

type FuturesWS_AggTrade_Socket struct {
	Handler *Futures_Websocket
}

func (*FuturesWS_AggTrade_Socket) CreateStreamName(symbol ...string) []string {
	streamNames := make([]string, len(symbol))
	for i := range symbol {
		streamNames[i] = strings.ToLower(symbol[i]) + "@aggTrade"
	}
	return streamNames
}

func (socket *FuturesWS_AggTrade_Socket) Subscribe(symbol ...string) (hasTimedOut bool, err error) {
	streamNames := socket.CreateStreamName(symbol...)
	return socket.Handler.Subscribe(streamNames...)
}

func (socket *FuturesWS_AggTrade_Socket) Unsubscribe(symbol ...string) (hasTimedOut bool, err error) {
	streamNames := socket.CreateStreamName(symbol...)
	return socket.Handler.Unsubscribe(streamNames...)
}

func (futures_ws *Futures_Websockets) AggTrade(publicOnMessage func(aggTrade *FuturesWS_AggTrade), symbol ...string) *FuturesWS_AggTrade_Socket {
	var newSocket FuturesWS_AggTrade_Socket

	streamNames := newSocket.CreateStreamName(symbol...)

	socket := futures_ws.CreateSocket(streamNames, false)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var aggTrade FuturesWS_AggTrade
		err := json.Unmarshal(msg, &aggTrade)
		if err != nil {
			LocalError(PARSING_ERR, err.Error())
			return
		}
		publicOnMessage(&aggTrade)
	}

	newSocket.Handler = socket
	return &newSocket
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type FuturesWS_MarkPrice struct {

	// Event type
	Event string `json:"e"`

	// Event time
	EventTime int64 `json:"E"`

	// Symbol
	Symbol string `json:"s"`

	// Mark price
	MarkPrice string `json:"p"`

	// Index price
	IndexPrice string `json:"i"`

	// Estimated Settle Price, only useful in the last hour before the settlement starts
	EstimatedSettlePrice string `json:"P"`

	// Funding rate
	FundingRate string `json:"r"`

	// Next funding time
	NextFundingTime int64 `json:"T"`
}

type FuturesWS_MarkPrice_Params struct {
	Symbol string
	IsFast bool
}

type FuturesWS_MarkPrice_Socket struct {
	Handler *Futures_Websocket
}

func (*FuturesWS_MarkPrice_Socket) CreateStreamName(params ...FuturesWS_MarkPrice_Params) []string {
	streamNames := make([]string, len(params))
	for i := range params {
		streamNames[i] = strings.ToLower(params[i].Symbol) + "@markPrice"
		if params[i].IsFast {
			streamNames[i] += "@1s"
		}
	}
	return streamNames
}

func (socket *FuturesWS_MarkPrice_Socket) Subscribe(params ...FuturesWS_MarkPrice_Params) (hasTimedOut bool, err error) {
	streams := socket.CreateStreamName(params...)
	return socket.Handler.Subscribe(streams...)
}

func (socket *FuturesWS_MarkPrice_Socket) Unsubscribe(params ...FuturesWS_MarkPrice_Params) (hasTimedOut bool, err error) {
	streams := socket.CreateStreamName(params...)
	return socket.Handler.Unsubscribe(streams...)
}

func (futures_ws *Futures_Websockets) MarkPrice(publicOnMessage func(markPrice *FuturesWS_MarkPrice), params ...FuturesWS_MarkPrice_Params) *FuturesWS_MarkPrice_Socket {
	var newSocket FuturesWS_MarkPrice_Socket

	streams := newSocket.CreateStreamName(params...)

	socket := futures_ws.CreateSocket(streams, false)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var markPrice FuturesWS_MarkPrice
		err := json.Unmarshal(msg, &markPrice)
		if err != nil {
			LocalError(PARSING_ERR, err.Error())
			return
		}
		publicOnMessage(&markPrice)
	}

	newSocket.Handler = socket
	return &newSocket
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type FuturesWS_AllMarkPrices_Socket struct {
	Handler *Futures_Websocket
}

func (*FuturesWS_AllMarkPrices_Socket) CreateStreamName(isFast bool) string {
	streamName := "!markPrice@arr"
	if isFast {
		streamName += "@1s"
	}
	return streamName
}

func (socket *FuturesWS_AllMarkPrices_Socket) Subscribe(isFast ...bool) (hasTimedOut bool, err error) {
	streams := make([]string, len(isFast))
	for i := range isFast {
		streams[i] = socket.CreateStreamName(isFast[i])
	}

	return socket.Handler.Subscribe(streams...)
}

func (socket *FuturesWS_AllMarkPrices_Socket) Unsubscribe(isFast ...bool) (hasTimedOut bool, err error) {
	streams := make([]string, len(isFast))
	for i := range isFast {
		streams[i] = socket.CreateStreamName(isFast[i])
	}

	return socket.Handler.Unsubscribe(streams...)
}

func (futures_ws *Futures_Websockets) AllMarkPrices(publicOnMessage func(markPrices []*FuturesWS_MarkPrice), isFast ...bool) *FuturesWS_AllMarkPrices_Socket {
	var newSocket FuturesWS_AllMarkPrices_Socket

	streams := make([]string, len(isFast))
	for i := range isFast {
		streams[i] = newSocket.CreateStreamName(isFast[i])
	}
	socket := futures_ws.CreateSocket(streams, false)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var markPrices []*FuturesWS_MarkPrice
		err := json.Unmarshal(msg, &markPrices)
		if err != nil {
			LocalError(PARSING_ERR, err.Error())
			return
		}
		publicOnMessage(markPrices)
	}

	newSocket.Handler = socket
	return &newSocket
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type FuturesWS_Candlestick struct {

	// Event type
	Event string `json:"e"`

	// Event time
	EventTime int64 `json:"E"`

	// Symbol
	Symbol string `json:"s"`

	Kline *FuturesWS_Candlestick_Kline `json:"k"`
}
type FuturesWS_Candlestick_Kline struct {

	// Symbol
	Symbol string `json:"s"`

	// Kline start time
	OpenTime int64 `json:"t"`

	// Kline close time
	CloseTime int64 `json:"T"`

	// Is this kline closed?
	IsClosed bool `json:"x"`

	// Interval
	Interval string `json:"i"`

	// First trade ID
	FirstTradeId int64 `json:"f"`

	// Last trade ID
	LastTradeId int64 `json:"L"`

	// Open price
	Open string `json:"o"`

	// Close price
	Close string `json:"c"`

	// High price
	High string `json:"h"`

	// Low price
	Low string `json:"l"`

	// Number of trades
	TradeCount int64 `json:"n"`

	// Base asset volume
	BaseAssetVolume string `json:"v"`

	// Quote asset volume
	QuoteAssetVolume string `json:"q"`

	// Taker buy base asset volume
	TakerBuyBaseAssetVolume string `json:"V"`

	// Taker buy quote asset volume
	TakerBuyQuoteAssetVolume string `json:"Q"`

	// Ignore
	Ignore string `json:"B"`
}

type FuturesWS_Candlesticks_Socket struct {
	Handler *Futures_Websocket
}

type FuturesWS_Candlestick_Params struct {
	Symbol   string
	Interval string
}

func (*FuturesWS_Candlesticks_Socket) CreateStreamName(params ...FuturesWS_Candlestick_Params) []string {
	streamNames := make([]string, len(params))
	for i := range params {
		streamNames[i] = strings.ToLower(params[i].Symbol) + "@kline_" + params[i].Interval
	}
	return streamNames
}

func (socket *FuturesWS_Candlesticks_Socket) Subscribe(params ...FuturesWS_Candlestick_Params) (hasTimedOut bool, err error) {
	streamNames := socket.CreateStreamName(params...)
	return socket.Handler.Subscribe(streamNames...)
}

func (socket *FuturesWS_Candlesticks_Socket) Unsubscribe(params ...FuturesWS_Candlestick_Params) (hasTimedOut bool, err error) {
	streamNames := socket.CreateStreamName(params...)
	return socket.Handler.Unsubscribe(streamNames...)
}

func (futures_ws *Futures_Websockets) Candlesticks(publicOnMessage func(candlestick *FuturesWS_Candlestick), params ...FuturesWS_Candlestick_Params) *FuturesWS_Candlesticks_Socket {
	var newSocket FuturesWS_Candlesticks_Socket

	streamNames := newSocket.CreateStreamName(params...)

	socket := futures_ws.CreateSocket(streamNames, false)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var kline *FuturesWS_Candlestick
		err := json.Unmarshal(msg, &kline)
		if err != nil {
			LocalError(PARSING_ERR, err.Error())
			return
		}
		publicOnMessage(kline)
	}

	newSocket.Handler = socket
	return &newSocket
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type FuturesWS_ContinuousCandlestick struct {

	// Event type
	Event string `json:"e"`

	// Event time
	EventTime int64 `json:"E"`

	// Pair
	Pair string `json:"ps"`

	// Contract type
	ContractType string `json:"ct"`

	Kline *FuturesWS_ContinuousCandlestick_Kline `json:"k"`
}
type FuturesWS_ContinuousCandlestick_Kline struct {

	// Kline start time
	OpenTime int64 `json:"t"`

	// Kline close time
	CloseTime int64 `json:"T"`

	// Interval
	Interval string `json:"i"`

	// First updateId
	FirstUpdateId int64 `json:"f"`

	// Last updateId
	LastUpdateId int64 `json:"L"`

	// Open price
	Open string `json:"o"`

	// Close price
	Close string `json:"c"`

	// High price
	High string `json:"h"`

	// Low price
	Low string `json:"l"`

	// volume
	Volume string `json:"v"`

	// Number of trades
	TradeCount int64 `json:"n"`

	// Is this kline closed?
	IsClosed bool `json:"x"`

	// Quote asset volume
	QuoteAssetVolume string `json:"q"`

	// Taker buy volume
	TakerBuyVolume string `json:"V"`

	// Taker buy quote asset volume
	TakerBuyQuoteAssetVolume string `json:"Q"`

	// Ignore
	Ignore string `json:"B"`
}

type FuturesWS_ContinuousCandlestick_Socket struct {
	Handler *Futures_Websocket
}

type FuturesWS_ContinuousCandlestick_Params struct {
	Symbol       string
	ContractType string
	Interval     string
}

func (*FuturesWS_ContinuousCandlestick_Socket) CreateStreamName(params ...FuturesWS_ContinuousCandlestick_Params) []string {
	streamNames := make([]string, len(params))
	for i := range params {
		streamNames[i] = strings.ToLower(params[i].Symbol) + "_" + strings.ToLower(params[i].ContractType) + "@continuousKline_" + params[i].Interval
	}
	return streamNames
}

func (socket *FuturesWS_ContinuousCandlestick_Socket) Subscribe(params ...FuturesWS_ContinuousCandlestick_Params) (hasTimedOut bool, err error) {
	streamNames := socket.CreateStreamName(params...)
	return socket.Handler.Subscribe(streamNames...)
}

func (socket *FuturesWS_ContinuousCandlestick_Socket) Unsubscribe(params ...FuturesWS_ContinuousCandlestick_Params) (hasTimedOut bool, err error) {
	streamNames := socket.CreateStreamName(params...)
	return socket.Handler.Unsubscribe(streamNames...)
}

// This is the only endpoint where binance unofficially supports the "1s" interval
//
// So using it should be okay.
func (futures_ws *Futures_Websockets) ContinuousCandlesticks(publicOnMessage func(candlestick *FuturesWS_ContinuousCandlestick), params ...FuturesWS_ContinuousCandlestick_Params) *FuturesWS_ContinuousCandlestick_Socket {
	var newSocket FuturesWS_ContinuousCandlestick_Socket

	streamNames := newSocket.CreateStreamName(params...)

	socket := futures_ws.CreateSocket(streamNames, false)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var kline *FuturesWS_ContinuousCandlestick
		err := json.Unmarshal(msg, &kline)
		if err != nil {
			LocalError(PARSING_ERR, err.Error())
			return
		}
		publicOnMessage(kline)
	}

	newSocket.Handler = socket
	return &newSocket
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type FuturesWS_MiniTicker struct {

	// Event type
	Event string `json:"e"`

	// Event time
	EventTime int64 `json:"E"`

	// Symbol
	Symbol string `json:"s"`

	// Close price
	Close string `json:"c"`

	// Open price
	Open string `json:"o"`

	// High price
	High string `json:"h"`

	// Low price
	Low string `json:"l"`

	// Total traded base asset volume
	BaseAssetVolume string `json:"v"`

	// Total traded quote asset volume
	QuoteAssetVolume string `json:"q"`
}

type FuturesWS_MiniTicker_Socket struct {
	Handler *Futures_Websocket
}

func (*FuturesWS_MiniTicker_Socket) CreateStreamName(symbol ...string) []string {
	streamNames := make([]string, len(symbol))
	for i := range symbol {
		streamNames[i] = strings.ToLower(symbol[i]) + "@miniTicker"
	}
	return streamNames
}

func (socket *FuturesWS_MiniTicker_Socket) Subscribe(symbol ...string) (hasTimedOut bool, err error) {
	streamNames := socket.CreateStreamName(symbol...)
	return socket.Handler.Subscribe(streamNames...)
}

func (socket *FuturesWS_MiniTicker_Socket) Unsubscribe(symbol ...string) (hasTimedOut bool, err error) {
	streamNames := socket.CreateStreamName(symbol...)
	return socket.Handler.Unsubscribe(streamNames...)
}

func (futures_ws *Futures_Websockets) MiniTicker(publicOnMessage func(miniTicker *FuturesWS_MiniTicker), symbol ...string) *FuturesWS_MiniTicker_Socket {
	var newSocket FuturesWS_MiniTicker_Socket

	streamNames := newSocket.CreateStreamName(symbol...)

	socket := futures_ws.CreateSocket(streamNames, false)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var miniTicker *FuturesWS_MiniTicker
		err := json.Unmarshal(msg, &miniTicker)
		if err != nil {
			LocalError(PARSING_ERR, err.Error())
			return
		}
		publicOnMessage(miniTicker)
	}

	newSocket.Handler = socket
	return &newSocket
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type FuturesWS_AllMiniTickers_Socket struct {
	Handler *Futures_Websocket
}

func (*FuturesWS_AllMiniTickers_Socket) CreateStreamName() string {
	return "!miniTicker@arr"
}

func (socket *FuturesWS_AllMiniTickers_Socket) Subscribe() (hasTimedOut bool, err error) {
	return socket.Handler.Subscribe(socket.CreateStreamName())
}

func (socket *FuturesWS_AllMiniTickers_Socket) Unsubscribe() (hasTimedOut bool, err error) {
	return socket.Handler.Unsubscribe(socket.CreateStreamName())
}

func (futures_ws *Futures_Websockets) AllMiniTickers(publicOnMessage func(miniTickers []*FuturesWS_MiniTicker)) *FuturesWS_AllMiniTickers_Socket {
	var newSocket FuturesWS_AllMiniTickers_Socket

	streamName := newSocket.CreateStreamName()

	socket := futures_ws.CreateSocket([]string{streamName}, false)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var miniTickers []*FuturesWS_MiniTicker
		err := json.Unmarshal(msg, &miniTickers)
		if err != nil {
			LocalError(PARSING_ERR, err.Error())
			return
		}
		publicOnMessage(miniTickers)
	}

	newSocket.Handler = socket
	return &newSocket
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type FuturesWS_Ticker struct {

	// Event type
	Event string `json:"e"`

	// Event time
	EventTime int64 `json:"E"`

	// Symbol
	Symbol string `json:"s"`

	// Price change
	PriceChange string `json:"p"`

	// Price change percent
	PriceChangePercent string `json:"P"`

	// Weighted average price
	WeightedAveragePrice string `json:"w"`

	// Last price
	LastPrice string `json:"c"`

	// Last quantity
	LastQty string `json:"Q"`

	// Open price
	Open string `json:"o"`

	// High price
	High string `json:"h"`

	// Low price
	Low string `json:"l"`

	// Total traded base asset volume
	BaseAssetVolume string `json:"v"`

	// Total traded quote asset volume
	QuoteAssetVolume string `json:"q"`

	// Statistics open time
	OpenTime int64 `json:"O"`

	// Statistics close time
	CloseTime int64 `json:"C"`

	// First trade ID
	FirstTradeId int64 `json:"F"`

	// Last trade Id
	LastTradeId int64 `json:"L"`

	// Total number of trades
	TradeCount int64 `json:"n"`
}

type FuturesWS_Ticker_Socket struct {
	Handler *Futures_Websocket
}

func (*FuturesWS_Ticker_Socket) CreateStreamName(symbol ...string) []string {
	streamNames := make([]string, len(symbol))
	for i := range symbol {
		streamNames[i] = strings.ToLower(symbol[i]) + "@ticker"
	}
	return streamNames
}

func (socket *FuturesWS_Ticker_Socket) Subscribe(symbol ...string) (hasTimedOut bool, err error) {
	streamNames := socket.CreateStreamName(symbol...)
	return socket.Handler.Subscribe(streamNames...)
}

func (socket *FuturesWS_Ticker_Socket) Unsubscribe(symbol ...string) (hasTimedOut bool, err error) {
	streamNames := socket.CreateStreamName(symbol...)
	return socket.Handler.Unsubscribe(streamNames...)
}

func (futures_ws *Futures_Websockets) Ticker(publicOnMessage func(ticker *FuturesWS_Ticker), symbol ...string) *FuturesWS_Ticker_Socket {
	var newSocket FuturesWS_Ticker_Socket

	streamNames := newSocket.CreateStreamName(symbol...)

	socket := futures_ws.CreateSocket(streamNames, false)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var ticker *FuturesWS_Ticker
		err := json.Unmarshal(msg, &ticker)
		if err != nil {
			LocalError(PARSING_ERR, err.Error())
			return
		}
		publicOnMessage(ticker)
	}

	newSocket.Handler = socket
	return &newSocket
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type FuturesWS_AllTickers_Socket struct {
	Handler *Futures_Websocket
}

func (*FuturesWS_AllTickers_Socket) CreateStreamName() string {
	return "!ticker@arr"
}

func (socket *FuturesWS_AllTickers_Socket) Subscribe() (hasTimedOut bool, err error) {
	return socket.Handler.Subscribe(socket.CreateStreamName())
}

func (socket *FuturesWS_AllTickers_Socket) Unsubscribe() (hasTimedOut bool, err error) {
	return socket.Handler.Unsubscribe(socket.CreateStreamName())
}

func (futures_ws *Futures_Websockets) AllTickers(publicOnMessage func(tickers []*FuturesWS_Ticker)) *FuturesWS_AllTickers_Socket {
	var newSocket FuturesWS_AllTickers_Socket

	streamName := newSocket.CreateStreamName()

	socket := futures_ws.CreateSocket([]string{streamName}, false)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var tickers []*FuturesWS_Ticker
		err := json.Unmarshal(msg, &tickers)
		if err != nil {
			LocalError(PARSING_ERR, err.Error())
			return
		}
		publicOnMessage(tickers)
	}

	newSocket.Handler = socket
	return &newSocket
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type FuturesWS_BookTicker struct {

	// event type
	Event string `json:"e"`

	// order book updateId
	UpdateId int64 `json:"u"`

	// event time
	EventTime int64 `json:"E"`

	// transaction time
	TransactTime int64 `json:"T"`

	// symbol
	Symbol string `json:"s"`

	// best bid price
	Bid string `json:"b"`

	// best bid qty
	BidQty string `json:"B"`

	// best ask price
	Ask string `json:"a"`

	// best ask qty
	AskQty string `json:"A"`
}

type FuturesWS_BookTicker_Socket struct {
	Handler *Futures_Websocket
}

func (*FuturesWS_BookTicker_Socket) CreateStreamName(symbol ...string) []string {
	for i := range symbol {
		symbol[i] = strings.ToLower(symbol[i]) + "@bookTicker"
	}
	return symbol
}

func (socket *FuturesWS_BookTicker_Socket) Subscribe(symbol ...string) (hasTimedOut bool, err error) {
	symbol = socket.CreateStreamName(symbol...)
	return socket.Handler.Subscribe(symbol...)
}

func (socket *FuturesWS_BookTicker_Socket) Unsubscribe(symbol ...string) (hasTimedOut bool, err error) {
	symbol = socket.CreateStreamName(symbol...)
	return socket.Handler.Unsubscribe(symbol...)
}

func (futures_ws *Futures_Websockets) BookTicker(publicOnMessage func(bookTicker *FuturesWS_BookTicker), symbol ...string) *FuturesWS_BookTicker_Socket {
	var newSocket FuturesWS_BookTicker_Socket

	symbol = newSocket.CreateStreamName(symbol...)
	socket := futures_ws.CreateSocket(symbol, false)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var bookTicker FuturesWS_BookTicker
		err := json.Unmarshal(msg, &bookTicker)
		if err != nil {
			LocalError(PARSING_ERR, err.Error())
			return
		}
		publicOnMessage(&bookTicker)
	}

	newSocket.Handler = socket
	return &newSocket
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type FuturesWS_AllBookTickers_Socket struct {
	Handler *Futures_Websocket
}

func (*FuturesWS_AllBookTickers_Socket) CreateStreamName() string {
	return "!bookTicker"
}

func (socket *FuturesWS_AllBookTickers_Socket) Subscribe() (hasTimedOut bool, err error) {
	streamName := socket.CreateStreamName()
	return socket.Handler.Subscribe(streamName)
}

func (socket *FuturesWS_AllBookTickers_Socket) Unsubscribe() (hasTimedOut bool, err error) {
	streamName := socket.CreateStreamName()
	return socket.Handler.Unsubscribe(streamName)
}

func (futures_ws *Futures_Websockets) AllBookTickers(publicOnMessage func(bookTickers []*FuturesWS_BookTicker)) *FuturesWS_AllBookTickers_Socket {
	var newSocket FuturesWS_AllBookTickers_Socket

	streamName := newSocket.CreateStreamName()

	socket := futures_ws.CreateSocket([]string{streamName}, false)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var bookTickers []*FuturesWS_BookTicker
		err := json.Unmarshal(msg, &bookTickers)
		if err != nil {
			LocalError(PARSING_ERR, err.Error())
			return
		}
		publicOnMessage(bookTickers)
	}

	newSocket.Handler = socket
	return &newSocket
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type FuturesWS_LiquidationOrder struct {

	// Event Type
	Event string `json:"e"`

	// Event Time
	EventTime int64 `json:"E"`

	Type *FuturesWS_LiquidationOrder_Order `json:"o"`
}
type FuturesWS_LiquidationOrder_Order struct {

	// Symbol
	Symbol string `json:"s"`

	// Side
	Side string `json:"S"`

	// Order Type
	Type string `json:"o"`

	// Time in Force
	TimeInForce string `json:"f"`

	// Original Quantity
	OrigQty string `json:"q"`

	// Price
	Price string `json:"p"`

	// Average Price
	AvgPrice string `json:"ap"`

	// Order Status
	Status string `json:"X"`

	// Order Last Filled Quantity
	LastFilledQty string `json:"l"`

	// Order Filled Accumulated Quantity
	CumQty string `json:"z"`

	// Order Trade Time
	TradeTime int64 `json:"T"`
}

type FuturesWS_LiquidationOrder_Socket struct {
	Handler *Futures_Websocket
}

func (*FuturesWS_LiquidationOrder_Socket) CreateStreamName(symbol ...string) []string {
	streams := make([]string, len(symbol))
	for i := range streams {
		streams[i] = strings.ToLower(symbol[i])
	}

	return streams
}

func (socket *FuturesWS_LiquidationOrder_Socket) Subscribe(symbol ...string) (hasTimedOut bool, err error) {
	streamNames := socket.CreateStreamName(symbol...)
	return socket.Handler.Subscribe(streamNames...)
}

func (socket *FuturesWS_LiquidationOrder_Socket) Unsubscribe(symbol ...string) (hasTimedOut bool, err error) {
	streamNames := socket.CreateStreamName(symbol...)
	return socket.Handler.Unsubscribe(streamNames...)
}

func (futures_ws *Futures_Websockets) LiquidationOrders(publicOnMessage func(liquidationOrder *FuturesWS_LiquidationOrder), symbol ...string) *FuturesWS_LiquidationOrder_Socket {
	var newSocket FuturesWS_LiquidationOrder_Socket

	streamNames := newSocket.CreateStreamName(symbol...)

	socket := futures_ws.CreateSocket(streamNames, false)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var liquidationOrder *FuturesWS_LiquidationOrder
		err := json.Unmarshal(msg, &liquidationOrder)
		if err != nil {
			LocalError(PARSING_ERR, err.Error())
			return
		}
		publicOnMessage(liquidationOrder)
	}

	newSocket.Handler = socket
	return &newSocket
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type FuturesWS_AllLiquidationOrders_Socket struct {
	Handler *Futures_Websocket
}

func (*FuturesWS_AllLiquidationOrders_Socket) CreateStreamName() string {
	return "!forceOrder@arr"
}

func (socket *FuturesWS_AllLiquidationOrders_Socket) Subscribe() (hasTimedOut bool, err error) {
	return socket.Handler.Subscribe(socket.CreateStreamName())
}

func (socket *FuturesWS_AllLiquidationOrders_Socket) Unsubscribe() (hasTimedOut bool, err error) {
	return socket.Handler.Unsubscribe(socket.CreateStreamName())
}

func (futures_ws *Futures_Websockets) AllLiquidationOrders(publicOnMessage func(liquidationOrder *FuturesWS_LiquidationOrder)) *FuturesWS_AllLiquidationOrders_Socket {
	var newSocket FuturesWS_AllLiquidationOrders_Socket

	streamName := newSocket.CreateStreamName()

	socket := futures_ws.CreateSocket([]string{streamName}, false)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var liquidationOrder *FuturesWS_LiquidationOrder
		err := json.Unmarshal(msg, &liquidationOrder)
		if err != nil {
			LocalError(PARSING_ERR, err.Error())
			return
		}
		publicOnMessage(liquidationOrder)
	}

	newSocket.Handler = socket
	return &newSocket
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type FuturesWS_PartialBookDepth struct {

	// Event Type
	Event string `json:"e"`

	// Event Time
	EventTime int64 `json:"E"`

	TransactTime int64 `json:"T"`

	Symbol string `json:"s"`

	FirstUpdateId int64 `json:"U"`

	LastUpdateId string `json:"u"`

	Previous_LastUpdateId string `json:"pu"`

	// Bids to be updated
	//
	// [
	//     [
	//       "7405.96",      // Price level to be
	//       "3.340"         // Quantity
	//     ],
	// ]
	Bids [][2]string `json:"b"`

	// Asks to be updated
	//
	// [
	//     [
	//       "7405.96",      // Price level to be
	//       "3.340"         // Quantity
	//     ],
	// ]
	Asks [][2]string `json:"a"`
}

type FuturesWS_PartialBookDepth_Socket struct {
	Handler *Futures_Websocket
}

type FuturesWS_PartialBookDepth_Params struct {
	Symbol string

	// Possible values: 5, 10 or 20
	Levels int

	// Possible values: "500ms", "250ms" or "100ms"
	//
	// Default: "250ms"
	UpdateSpeed string
}

func (*FuturesWS_PartialBookDepth_Socket) CreateStreamName(params ...FuturesWS_PartialBookDepth_Params) []string {
	streamNames := make([]string, len(params))
	for i := range params {
		streamNames[i] = strings.ToLower(params[i].Symbol) + "@depth" + fmt.Sprint(params[i].Levels)
		if params[i].UpdateSpeed != "" && params[i].UpdateSpeed != "250ms" {
			streamNames[i] += "@" + params[i].UpdateSpeed
		}
	}

	return streamNames
}

func (socket *FuturesWS_PartialBookDepth_Socket) Subscribe(params ...FuturesWS_PartialBookDepth_Params) (hasTimedOut bool, err error) {
	streamNames := socket.CreateStreamName(params...)
	return socket.Handler.Subscribe(streamNames...)
}

func (socket *FuturesWS_PartialBookDepth_Socket) Unsubscribe(params ...FuturesWS_PartialBookDepth_Params) (hasTimedOut bool, err error) {
	streamNames := socket.CreateStreamName(params...)
	return socket.Handler.Unsubscribe(streamNames...)
}

func (futures_ws *Futures_Websockets) PartialBookDepth(publicOnMessage func(partialBookDepth *FuturesWS_PartialBookDepth), params ...FuturesWS_PartialBookDepth_Params) *FuturesWS_PartialBookDepth_Socket {
	var newSocket FuturesWS_PartialBookDepth_Socket

	streamNames := newSocket.CreateStreamName(params...)

	socket := futures_ws.CreateSocket(streamNames, false)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var partialBookDepth *FuturesWS_PartialBookDepth
		err := json.Unmarshal(msg, &partialBookDepth)
		if err != nil {
			LocalError(PARSING_ERR, err.Error())
			return
		}
		publicOnMessage(partialBookDepth)
	}

	newSocket.Handler = socket
	return &newSocket
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type FuturesWS_DiffBookDepth struct {

	// Event Type
	Event string `json:"e"`

	// Event Time
	EventTime int64 `json:"E"`

	TransactTime int64 `json:"T"`

	Symbol string `json:"s"`

	FirstUpdateId int64 `json:"U"`

	LastUpdateId int64 `json:"u"`

	Previous_LastUpdateId int64 `json:"pu"`

	// Bids to be updated
	//
	// [
	//     [
	//       "7405.96",      // Price level to be
	//       "3.340"         // Quantity
	//     ],
	// ]
	Bids [][2]string `json:"b"`

	// Asks to be updated
	//
	// [
	//     [
	//       "7405.96",      // Price level to be
	//       "3.340"         // Quantity
	//     ],
	// ]
	Asks [][2]string `json:"a"`
}

type FuturesWS_DiffBookDepth_Socket struct {
	Handler *Futures_Websocket
}

type FuturesWS_DiffBookDepth_Params struct {
	Symbol string

	// Possible values: "500ms", "250ms" or "100ms"
	//
	// Default: "250ms"
	UpdateSpeed string
}

func (*FuturesWS_DiffBookDepth_Socket) CreateStreamName(params ...FuturesWS_DiffBookDepth_Params) []string {
	streamNames := make([]string, len(params))
	for i := range params {
		streamNames[i] = strings.ToLower(params[i].Symbol) + "@depth"
		if params[i].UpdateSpeed != "" && params[i].UpdateSpeed != "250ms" {
			streamNames[i] += "@" + params[i].UpdateSpeed
		}
	}

	return streamNames
}

func (socket *FuturesWS_DiffBookDepth_Socket) Subscribe(params ...FuturesWS_DiffBookDepth_Params) (hasTimedOut bool, err error) {
	streamNames := socket.CreateStreamName(params...)
	return socket.Handler.Subscribe(streamNames...)
}

func (socket *FuturesWS_DiffBookDepth_Socket) Unsubscribe(params ...FuturesWS_DiffBookDepth_Params) (hasTimedOut bool, err error) {
	streamNames := socket.CreateStreamName(params...)
	return socket.Handler.Unsubscribe(streamNames...)
}

func (futures_ws *Futures_Websockets) DiffBookDepth(publicOnMessage func(diffBookDepth *FuturesWS_DiffBookDepth), params ...FuturesWS_DiffBookDepth_Params) *FuturesWS_DiffBookDepth_Socket {
	var newSocket FuturesWS_DiffBookDepth_Socket

	streamNames := newSocket.CreateStreamName(params...)

	socket := futures_ws.CreateSocket(streamNames, false)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var diffBookDepth *FuturesWS_DiffBookDepth
		err := json.Unmarshal(msg, &diffBookDepth)
		if err != nil {
			LocalError(PARSING_ERR, err.Error())
			return
		}
		publicOnMessage(diffBookDepth)
	}

	newSocket.Handler = socket
	return &newSocket
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type FuturesWS_CompositeIndexSymbolInfo struct {

	// Event Type
	Event string `json:"e"`

	// Event Time
	EventTime int64 `json:"E"`

	Symbol string `json:"s"`

	Price int64 `json:"p"`

	BaseAsset string `json:"C"`

	Composition []*FuturesWS_CompositeIndexSymbolInfo_Composition
}

type FuturesWS_CompositeIndexSymbolInfo_Composition struct {

	// Base asset
	BaseAsset string `json:"b"`

	// Quote asset
	QuoteAsset string `json:"q"`

	// Weight in quantity
	Weight string `json:"w"`

	// Weight in percentage
	WeightPercent string `json:"W"`

	// Index Price
	IndexPrice string `json:"i"`
}

type FuturesWS_CompositeIndexSymbolInfo_Socket struct {
	Handler *Futures_Websocket
}

func (*FuturesWS_CompositeIndexSymbolInfo_Socket) CreateStreamName(symbol ...string) []string {
	streamNames := make([]string, len(symbol))
	for i := range symbol {
		streamNames[i] = strings.ToLower(symbol[i]) + "@compositeIndex"
	}

	return streamNames
}

func (socket *FuturesWS_CompositeIndexSymbolInfo_Socket) Subscribe(symbol ...string) (hasTimedOut bool, err error) {
	streamNames := socket.CreateStreamName(symbol...)
	return socket.Handler.Subscribe(streamNames...)
}

func (socket *FuturesWS_CompositeIndexSymbolInfo_Socket) Unsubscribe(symbol ...string) (hasTimedOut bool, err error) {
	streamNames := socket.CreateStreamName(symbol...)
	return socket.Handler.Unsubscribe(streamNames...)
}

func (futures_ws *Futures_Websockets) CompositeIndexSymbolInfo(publicOnMessage func(compositeIndexSymbolInfo *FuturesWS_CompositeIndexSymbolInfo), symbol ...string) *FuturesWS_CompositeIndexSymbolInfo_Socket {
	var newSocket FuturesWS_CompositeIndexSymbolInfo_Socket

	streamNames := newSocket.CreateStreamName(symbol...)

	socket := futures_ws.CreateSocket(streamNames, false)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var compositeIndexSymbolInfo *FuturesWS_CompositeIndexSymbolInfo
		err := json.Unmarshal(msg, &compositeIndexSymbolInfo)
		if err != nil {
			LocalError(PARSING_ERR, err.Error())
			return
		}
		publicOnMessage(compositeIndexSymbolInfo)
	}

	newSocket.Handler = socket
	return &newSocket
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type FuturesWS_ContractInfo struct {

	// Event Type
	Event string `json:"e"`

	// Event Time
	EventTime int64 `json:"E"`

	// Symbol
	Symbol string `json:"s"`

	// Pair
	Pair string `json:"ps"`

	// Contract type
	ContractType string `json:"ct"`

	// Delivery date time
	DeliveryDate int64 `json:"dt"`

	// onboard date time
	OnboardDateTime int64 `json:"ot"`

	// Contract status
	ContractStatus string `json:"cs"`

	Bks []*FuturesWS_ContractInfo_Bracket `json:"bks"`
}
type FuturesWS_ContractInfo_Bracket struct {

	// Notional bracket
	NotionalBracket int64 `json:"bs"`

	// Floor notional of this bracket
	FloorNotional int64 `json:"bnf"`

	// Cap notional of this bracket
	MaxNotional int64 `json:"bnc"`

	// Maintenance ratio for this bracket
	MaintenanceRatio float64 `json:"mmr"`

	// Auxiliary number for quick calculation
	Auxiliary int64 `json:"cf"`

	// Min leverage for this bracket
	MinLeverage int64 `json:"mi"`

	// Max leverage for this bracket
	MaxLeverage int64 `json:"ma"`
}

type FuturesWS_ContractInfo_Socket struct {
	Handler *Futures_Websocket
}

func (*FuturesWS_ContractInfo_Socket) CreateStreamName() string {
	return "!contractInfo"
}

func (socket *FuturesWS_ContractInfo_Socket) Subscribe() (hasTimedOut bool, err error) {
	streamName := socket.CreateStreamName()
	return socket.Handler.Subscribe(streamName)
}

func (socket *FuturesWS_ContractInfo_Socket) Unsubscribe() (hasTimedOut bool, err error) {
	streamName := socket.CreateStreamName()
	return socket.Handler.Unsubscribe(streamName)
}

func (futures_ws *Futures_Websockets) ContractInfo(publicOnMessage func(contractInfo *FuturesWS_ContractInfo)) *FuturesWS_ContractInfo_Socket {
	var newSocket FuturesWS_ContractInfo_Socket
	streamName := newSocket.CreateStreamName()
	socket := futures_ws.CreateSocket([]string{streamName}, false)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var aggTrade FuturesWS_ContractInfo
		err := json.Unmarshal(msg, &aggTrade)
		if err != nil {
			LocalError(PARSING_ERR, err.Error())
			return
		}
		publicOnMessage(&aggTrade)
	}

	newSocket.Handler = socket
	return &newSocket
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type FuturesWS_MultiAssetsModeAssetIndex struct {
	Event string `json:"e"`

	EventTime int64 `json:"E"`

	// asset index symbol
	Symbol string `json:"s"`

	// index price
	IndexPrice string `json:"i"`

	// bid buffer
	BidBuffer string `json:"b"`

	// ask buffer
	AskBuffer string `json:"a"`

	// bid rate
	BidRate string `json:"B"`

	// ask rate
	AskRate string `json:"A"`

	// auto exchange bid buffer
	AutoExchange_BidBuffer string `json:"q"`

	// auto exchange ask buffer
	AutoExchange_AskBuffer string `json:"g"`

	// auto exchange bid rate
	AutoExchange_BidRate string `json:"Q"`

	// auto exchange ask rate
	AutoExchange_AskRate string `json:"G"`
}

type FuturesWS_MultiAssetsModeAssetIndex_Socket struct {
	Handler *Futures_Websocket
}

func (*FuturesWS_MultiAssetsModeAssetIndex_Socket) CreateStreamName(assetSymbol ...string) []string {
	streamNames := make([]string, len(assetSymbol))
	for i := range assetSymbol {
		streamNames[i] = strings.ToLower(assetSymbol[i]) + "@assetIndex"
	}

	return streamNames
}

func (socket *FuturesWS_MultiAssetsModeAssetIndex_Socket) Subscribe(assetSymbol ...string) (hasTimedOut bool, err error) {
	streamNames := socket.CreateStreamName(assetSymbol...)
	return socket.Handler.Subscribe(streamNames...)
}

func (socket *FuturesWS_MultiAssetsModeAssetIndex_Socket) Unsubscribe(assetSymbol ...string) (hasTimedOut bool, err error) {
	streamNames := socket.CreateStreamName(assetSymbol...)
	return socket.Handler.Unsubscribe(streamNames...)
}

func (futures_ws *Futures_Websockets) MultiAssetsModeAssetIndex(publicOnMessage func(assetIndexes []*FuturesWS_MultiAssetsModeAssetIndex), assetSymbol ...string) *FuturesWS_MultiAssetsModeAssetIndex_Socket {
	var newSocket FuturesWS_MultiAssetsModeAssetIndex_Socket

	streamNames := newSocket.CreateStreamName(assetSymbol...)

	socket := futures_ws.CreateSocket(streamNames, false)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var assetIndexes []*FuturesWS_MultiAssetsModeAssetIndex
		err := json.Unmarshal(msg, &assetIndexes)
		if err != nil {
			LocalError(PARSING_ERR, err.Error())
			return
		}
		publicOnMessage(assetIndexes)
	}

	newSocket.Handler = socket
	return &newSocket
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type FuturesWS_AllMultiAssetsModeAssetIndexes_Socket struct {
	Handler *Futures_Websocket
}

func (*FuturesWS_AllMultiAssetsModeAssetIndexes_Socket) CreateStreamName() string {
	return "!assetIndex@arr"
}

func (socket *FuturesWS_AllMultiAssetsModeAssetIndexes_Socket) Subscribe() (hasTimedOut bool, err error) {
	return socket.Handler.Subscribe(socket.CreateStreamName())
}

func (socket *FuturesWS_AllMultiAssetsModeAssetIndexes_Socket) Unsubscribe() (hasTimedOut bool, err error) {
	return socket.Handler.Unsubscribe(socket.CreateStreamName())
}

func (futures_ws *Futures_Websockets) AllMultiAssetsModeAssetIndexes(publicOnMessage func(assetIndexes []*FuturesWS_MultiAssetsModeAssetIndex), assetSymbol ...string) *FuturesWS_AllMultiAssetsModeAssetIndexes_Socket {
	var newSocket FuturesWS_AllMultiAssetsModeAssetIndexes_Socket

	streamName := newSocket.CreateStreamName()

	socket := futures_ws.CreateSocket([]string{streamName}, false)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var assetIndexes []*FuturesWS_MultiAssetsModeAssetIndex
		err := json.Unmarshal(msg, &assetIndexes)
		if err != nil {
			LocalError(PARSING_ERR, err.Error())
			return
		}
		publicOnMessage(assetIndexes)
	}

	newSocket.Handler = socket
	return &newSocket
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Futures_ManagedOrderbook struct {
	Symbol       string
	LastUpdateId int64
	Bids         [][2]float64
	Asks         [][2]float64

	isReadyToUpdate bool
	isFetching      bool

	previousEvent *FuturesWS_DiffBookDepth
}

func (managedOrderBook *Futures_ManagedOrderbook) addEvent(event *FuturesWS_DiffBookDepth) {
	var new_bidAsk_placeholder [2]float64

	for _, newBid_str := range event.Bids {
		new_bidAsk_placeholder[0], _ = Utils.ParseFloat(newBid_str[0])
		new_bidAsk_placeholder[1], _ = Utils.ParseFloat(newBid_str[1])

		found := false
		for i, bid := range managedOrderBook.Bids {
			if bid[0] == new_bidAsk_placeholder[0] {
				found = true

				if new_bidAsk_placeholder[1] == 0 {
					managedOrderBook.Bids = slices.Delete(managedOrderBook.Bids, i, i+1)
				} else {
					managedOrderBook.Bids[i][1] = new_bidAsk_placeholder[1]
				}
			}
		}

		if !found {
			if new_bidAsk_placeholder[1] != 0 {
				inserted := false
				for i := range managedOrderBook.Bids {
					// Compare the price of the new bid with the existing bid
					if managedOrderBook.Bids[i][0] < new_bidAsk_placeholder[0] {
						// Insert the new bid before this index
						managedOrderBook.Bids = append(managedOrderBook.Bids[:i], append([][2]float64{new_bidAsk_placeholder}, managedOrderBook.Bids[i:]...)...)
						inserted = true
						break
					}
				}

				// If the new bid is lower than all existing bids, append it to the end
				if !inserted {
					managedOrderBook.Bids = append(managedOrderBook.Bids, new_bidAsk_placeholder)
				}
			}
		}
	}

	for _, newAsk_str := range event.Asks {
		new_bidAsk_placeholder[0], _ = Utils.ParseFloat(newAsk_str[0])
		new_bidAsk_placeholder[1], _ = Utils.ParseFloat(newAsk_str[1])

		found := false
		for i, ask := range managedOrderBook.Asks {
			if ask[0] == new_bidAsk_placeholder[0] {

				if new_bidAsk_placeholder[1] == 0 {
					managedOrderBook.Asks = slices.Delete(managedOrderBook.Asks, i, i+1)
				} else {
					managedOrderBook.Asks[i][1] = new_bidAsk_placeholder[1]
				}
				found = true
			}
		}

		if !found {
			if new_bidAsk_placeholder[1] != 0 {
				inserted := false
				for i := range managedOrderBook.Asks {
					// Compare the price of the new ask with the existing ask
					if managedOrderBook.Asks[i][0] > new_bidAsk_placeholder[0] {
						// Insert the new ask before this index
						managedOrderBook.Asks = append(managedOrderBook.Asks[:i], append([][2]float64{new_bidAsk_placeholder}, managedOrderBook.Asks[i:]...)...)
						inserted = true
						break
					}
				}

				// If the new ask is higher than all existing asks, append it to the end
				if !inserted {
					managedOrderBook.Asks = append(managedOrderBook.Asks, new_bidAsk_placeholder)
				}
			}
		}

	}

	// start := time.Now().UnixMicro()
	// sort.SliceStable(managedOrderBook.Asks, func(i, j int) bool {
	// 	return managedOrderBook.Asks[i][0] < managedOrderBook.Asks[j][0]
	// })

	// sort.SliceStable(managedOrderBook.Bids, func(i, j int) bool {
	// 	return managedOrderBook.Bids[i][0] > managedOrderBook.Bids[j][0]
	// })
	// fmt.Println(time.Now().UnixMicro()-start, "micro")

	managedOrderBook.previousEvent = event
}

type FuturesWS_ManagedOrderBook_Handler struct {
	DiffBookDepth_Socket *FuturesWS_DiffBookDepth_Socket

	Orderbooks struct {
		Mu      sync.Mutex
		Symbols map[string]struct {
			bufferedEvents []*FuturesWS_DiffBookDepth
			Orderbook      *Futures_ManagedOrderbook
		}
	}
}

func (handler *FuturesWS_ManagedOrderBook_Handler) handleWSMessage(futures_ws *Futures_Websockets, diffBookDepth *FuturesWS_DiffBookDepth) (shouldPushEvent bool, managedOrderBook *Futures_ManagedOrderbook) {
	handler.Orderbooks.Mu.Lock()
	defer handler.Orderbooks.Mu.Unlock()
	Orderbook_symbol, exists := handler.Orderbooks.Symbols[diffBookDepth.Symbol]
	if !exists {
		Logger.DEBUG(fmt.Sprintf("Orderbook data for %s weren't found in handleWSMessage", diffBookDepth.Symbol))
		return false, nil
	}

	Orderbook_symbol.bufferedEvents = append(Orderbook_symbol.bufferedEvents, diffBookDepth)

	if !Orderbook_symbol.Orderbook.isReadyToUpdate && Orderbook_symbol.Orderbook.isFetching {
		Logger.DEBUG(fmt.Sprintf("[MANAGEDORDERBOOK] '%s' isNotReadyToUpdate and isFetching", diffBookDepth.Symbol))
		return false, nil
	}

	if !Orderbook_symbol.Orderbook.isReadyToUpdate {
		Logger.DEBUG(fmt.Sprintf("[MANAGEDORDERBOOK] Fetching '%s' snapshot", diffBookDepth.Symbol))
		Orderbook_symbol.Orderbook.isFetching = true

		newOrderBook, _, err := futures_ws.binance.Futures.OrderBook(diffBookDepth.Symbol, 1000)
		Orderbook_symbol.Orderbook.isFetching = false
		if err != nil {
			Logger.DEBUG(fmt.Sprintf("Failed to refetch orderbook data for %s: %s", diffBookDepth.Symbol, err.Error()))
			return false, nil
		}

		Orderbook_symbol.Orderbook.LastUpdateId = newOrderBook.LastUpdateId
		Orderbook_symbol.Orderbook.Asks = make([][2]float64, len(newOrderBook.Asks))
		Orderbook_symbol.Orderbook.Bids = make([][2]float64, len(newOrderBook.Bids))

		for i, ask := range newOrderBook.Asks {
			priceLvl, err1 := Utils.ParseFloat(ask[0])
			priceQty, err2 := Utils.ParseFloat(ask[1])
			if err1 != nil {
				Logger.DEBUG(fmt.Sprintf("Failed to parse bid priceLevel '%s'", ask[0]))
				// continue
			}
			if err2 != nil {
				Logger.DEBUG(fmt.Sprintf("Failed to parse bid priceQty '%s'", ask[1]))
				// continue
			}
			Orderbook_symbol.Orderbook.Asks[i][0] = priceLvl
			Orderbook_symbol.Orderbook.Asks[i][1] = priceQty
		}

		for i, bid := range newOrderBook.Bids {
			priceLvl, err1 := Utils.ParseFloat(bid[0])
			priceQty, err2 := Utils.ParseFloat(bid[1])
			if err1 != nil {
				Logger.DEBUG(fmt.Sprintf("Failed to parse ask priceLevel '%s'", bid[0]))
				// continue
			}
			if err2 != nil {
				Logger.DEBUG(fmt.Sprintf("Failed to parse ask priceQty '%s'", bid[1]))
				// continue
			}
			Orderbook_symbol.Orderbook.Bids[i][0] = priceLvl
			Orderbook_symbol.Orderbook.Bids[i][1] = priceQty
		}

		Orderbook_symbol.Orderbook.isReadyToUpdate = true
	}

	if len(Orderbook_symbol.bufferedEvents) == 0 {
		Logger.DEBUG(fmt.Sprintf("[MANAGEDORDERBOOK] '%s' bufferedEvents empty", diffBookDepth.Symbol))
		return false, nil
	}

	for len(Orderbook_symbol.bufferedEvents) > 0 {

		event := Orderbook_symbol.bufferedEvents[0]
		Orderbook_symbol.bufferedEvents = Orderbook_symbol.bufferedEvents[1:]

		if event.LastUpdateId < Orderbook_symbol.Orderbook.LastUpdateId {
			Logger.DEBUG(fmt.Sprintf("[MANAGEDORDERBOOK] '%s' ignored since 'u' < lastUpdateId: %d < %d", diffBookDepth.Symbol, event.LastUpdateId, Orderbook_symbol.Orderbook.LastUpdateId))
			continue
		}

		if Orderbook_symbol.Orderbook.previousEvent == nil {
			if event.FirstUpdateId <= Orderbook_symbol.Orderbook.LastUpdateId && Orderbook_symbol.Orderbook.LastUpdateId <= event.LastUpdateId {
				Logger.DEBUG(fmt.Sprintf("[MANAGEDORDERBOOK] '%s' first valid event found", diffBookDepth.Symbol))
				Orderbook_symbol.Orderbook.addEvent(event)
			} else {
				Logger.DEBUG(fmt.Sprintf("[MANAGEDORDERBOOK] '%s' INVALID FIRST EVENT FOUND, U: %d, lastUpdateId: %d, u: %d", diffBookDepth.Symbol, event.FirstUpdateId, Orderbook_symbol.Orderbook.LastUpdateId, event.LastUpdateId))
			}
		} else if event.Previous_LastUpdateId == Orderbook_symbol.Orderbook.previousEvent.LastUpdateId {
			Logger.DEBUG(fmt.Sprintf("[MANAGEDORDERBOOK] '%s' good event found, pu == [lastEvent].u, %d == %d", diffBookDepth.Symbol, event.Previous_LastUpdateId, Orderbook_symbol.Orderbook.previousEvent.LastUpdateId))
			Orderbook_symbol.Orderbook.addEvent(event)
		} else {
			Orderbook_symbol.Orderbook.isReadyToUpdate = false
			Orderbook_symbol.Orderbook.isFetching = false
			Orderbook_symbol.Orderbook.previousEvent = nil

			return false, nil
		}
	}

	return true, Orderbook_symbol.Orderbook
}

func (handler *FuturesWS_ManagedOrderBook_Handler) openNewSymbols(params ...FuturesWS_DiffBookDepth_Params) {
	handler.Orderbooks.Mu.Lock()
	defer handler.Orderbooks.Mu.Unlock()

	for _, param := range params {
		symbol := param.Symbol

		_, exists := handler.Orderbooks.Symbols[symbol]
		if exists {
			continue
		}

		handler.Orderbooks.Symbols[symbol] = struct {
			bufferedEvents []*FuturesWS_DiffBookDepth
			Orderbook      *Futures_ManagedOrderbook
		}{
			bufferedEvents: make([]*FuturesWS_DiffBookDepth, 0),
			Orderbook: &Futures_ManagedOrderbook{
				Symbol: symbol,

				isFetching:      false,
				isReadyToUpdate: false,
			},
		}
	}
}

func (handler *FuturesWS_ManagedOrderBook_Handler) removeSymbols(params ...FuturesWS_DiffBookDepth_Params) {
	handler.Orderbooks.Mu.Lock()
	defer handler.Orderbooks.Mu.Unlock()

	for _, param := range params {
		symbol := param.Symbol

		_, exists := handler.Orderbooks.Symbols[symbol]
		if !exists {
			continue
		}

		delete(handler.Orderbooks.Symbols, symbol)
	}
}

func (handler *FuturesWS_ManagedOrderBook_Handler) Unsubscribe(params ...FuturesWS_DiffBookDepth_Params) (hasTimedOut bool, err error) {
	hasTimedOut, err = handler.DiffBookDepth_Socket.Unsubscribe(params...)
	if err != nil {
		return hasTimedOut, err
	}

	handler.removeSymbols(params...)

	return false, nil
}

func (handler *FuturesWS_ManagedOrderBook_Handler) Subscribe(params ...FuturesWS_DiffBookDepth_Params) (hasTimedOut bool, err error) {
	hasTimedOut, err = handler.DiffBookDepth_Socket.Subscribe(params...)
	if err != nil {
		return hasTimedOut, err
	}

	handler.openNewSymbols(params...)

	return false, nil
}

func (futures_ws *Futures_Websockets) Managed_OrderBook(publicOnMessage func(orderBook *Futures_ManagedOrderbook), params ...FuturesWS_DiffBookDepth_Params) *FuturesWS_ManagedOrderBook_Handler {
	handler := &FuturesWS_ManagedOrderBook_Handler{}

	var newSocket FuturesWS_DiffBookDepth_Socket
	streamNames := newSocket.CreateStreamName(params...)
	socket := futures_ws.CreateSocket(streamNames, false)
	newSocket.Handler = socket

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var diffBookDepth *FuturesWS_DiffBookDepth
		err := json.Unmarshal(msg, &diffBookDepth)
		if err != nil {
			LocalError(PARSING_ERR, err.Error())
			return
		}

		shouldPushEvent, managedOrderbook := handler.handleWSMessage(futures_ws, diffBookDepth)
		if shouldPushEvent {
			publicOnMessage(managedOrderbook)
		}
	}

	handler.DiffBookDepth_Socket = &newSocket
	handler.Orderbooks.Symbols = make(map[string]struct {
		bufferedEvents []*FuturesWS_DiffBookDepth
		Orderbook      *Futures_ManagedOrderbook
	})
	handler.openNewSymbols(params...)

	return handler
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type FuturesWS_ManagedCandlesticks_Handler struct {
	futures_ws *Futures_Websockets

	Candlesticks_Socket *FuturesWS_Candlesticks_Socket
	AggTrades_Socket    *FuturesWS_AggTrade_Socket

	intervals []*Binance_Interval

	disable_aggTrades_stream    bool
	disable_candlesticks_stream bool

	Candlesticks struct {
		Mu      sync.Mutex
		Symbols map[string]*FuturesWS_ManagedCandlesticks_Symbol
	}
}

type FuturesWS_ManagedCandlesticks_Symbol struct {
	parent    *FuturesWS_ManagedCandlesticks_Handler
	Symbol    string
	AggTrades []*FuturesWS_ManagedCandlesticks_AggTrade
	Intervals struct {
		Mu  sync.Mutex
		Map map[string]*FuturesWS_ManagedCandlesticks_Interval
	}
	OnChange *Event[*FuturesWS_ManagedCandlesticks_Symbol]
}

type FuturesWS_ManagedCandlesticks_Interval struct {
	symbol       *FuturesWS_ManagedCandlesticks_Symbol
	Interval     *Binance_Interval
	Candlesticks []*FuturesWS_ManagedCandlestick
}

type FuturesWS_Candlestick_Float64 struct {
	OpenTime  int64
	CloseTime int64

	Open  float64
	High  float64
	Low   float64
	Close float64

	Volume                   float64
	QuoteAssetVolume         float64
	TakerBuyBaseAssetVolume  float64
	TakerBuyQuoteAssetVolume float64
	TradeCount               int64
}

type FuturesWS_ManagedCandlestick struct {
	OpenTime  int64
	CloseTime int64

	Open  float64
	High  float64
	Low   float64
	Close float64

	Volume                   float64
	QuoteAssetVolume         float64
	TakerBuyBaseAssetVolume  float64
	TakerBuyQuoteAssetVolume float64
	TradeCount               int64

	// Calculated from the incoming aggTrades
	//
	// If being used, always use math.Max(Calculated_Volume, Volume)
	Calculated_Volume float64

	// Calculated from the incoming aggTrades
	//
	// If being used, always use math.Max(Calculated_QuoteAssetVolume, QuoteAssetVolume)
	Calculated_QuoteAssetVolume float64

	// # Not sure if calculated correctly
	//
	// Currently calculated via adding to it ONLY if the aggTrade has 'IsMaker' as false
	//
	// If being used, always use math.Max(Calculated_TakerBuyBaseAssetVolume, TakerBuyBaseAssetVolume)
	Calculated_TakerBuyBaseAssetVolume float64

	// # Not sure if calculated correctly
	//
	// Currently calculated via adding to it ONLY if the aggTrade has 'IsMaker' as false
	//
	// If being used, always use math.Max(Calculated_TakerBuyQuoteAssetVolume, TakerBuyQuoteAssetVolume)
	Calculated_TakerBuyQuoteAssetVolume float64
	Calculated_TradeCount               int64

	AggTrades []*FuturesWS_ManagedCandlesticks_AggTrade
}

type FuturesWS_ManagedCandlesticks_AggTrade struct {
	Timestamp    int64
	IsMaker      bool
	AggTradeId   int64
	FirstTradeId int64
	LastTradeId  int64

	Price float64
	Qty   float64
}

////

func parseFloat_FuturesWS_Candlestick(candlestick *FuturesWS_Candlestick) (*FuturesWS_Candlestick_Float64, error) {
	kline := candlestick.Kline
	open, err := Utils.ParseFloat(kline.Open)
	if err != nil {
		errStr := fmt.Sprintf("There was an error parsing float '%s' of 'kline.Open' for '%s', interval '%s' in parseFloat_FuturesWS_Candlestick: %s", kline.Open, candlestick.Symbol, kline.Interval, err.Error())
		Logger.ERROR(errStr)
		return nil, err
	}
	high, err := Utils.ParseFloat(kline.High)
	if err != nil {
		errStr := fmt.Sprintf("There was an error parsing float '%s' of 'kline.High' for '%s', interval '%s' in parseFloat_FuturesWS_Candlestick: %s", kline.High, candlestick.Symbol, kline.Interval, err.Error())
		Logger.ERROR(errStr)
		return nil, err
	}
	low, err := Utils.ParseFloat(kline.Low)
	if err != nil {
		errStr := fmt.Sprintf("There was an error parsing float '%s' of 'kline.Low' for '%s', interval '%s' in parseFloat_FuturesWS_Candlestick: %s", kline.Low, candlestick.Symbol, kline.Interval, err.Error())
		Logger.ERROR(errStr)
		return nil, err
	}
	close, err := Utils.ParseFloat(kline.Close)
	if err != nil {
		errStr := fmt.Sprintf("There was an error parsing float '%s' of 'kline.Close' for '%s', interval '%s' in parseFloat_FuturesWS_Candlestick: %s", kline.Close, candlestick.Symbol, kline.Interval, err.Error())
		Logger.ERROR(errStr)
		return nil, err
	}

	baseAssetVolune, err := Utils.ParseFloat(kline.BaseAssetVolume)
	if err != nil {
		errStr := fmt.Sprintf("There was an error parsing float '%s' of 'kline.BaseAssetVolume' for '%s', interval '%s' in parseFloat_FuturesWS_Candlestick: %s", kline.BaseAssetVolume, candlestick.Symbol, kline.Interval, err.Error())
		Logger.ERROR(errStr)
		return nil, err
	}
	quoteAssetVolume, err := Utils.ParseFloat(kline.QuoteAssetVolume)
	if err != nil {
		errStr := fmt.Sprintf("There was an error parsing float '%s' of 'kline.QuoteAssetVolume' for '%s', interval '%s' in parseFloat_FuturesWS_Candlestick: %s", kline.QuoteAssetVolume, candlestick.Symbol, kline.Interval, err.Error())
		Logger.ERROR(errStr)
		return nil, err
	}

	takerBuyBaseAssetVolume, err := Utils.ParseFloat(kline.TakerBuyBaseAssetVolume)
	if err != nil {
		errStr := fmt.Sprintf("There was an error parsing float '%s' of 'kline.TakerBuyBaseAssetVolume' for '%s', interval '%s' in parseFloat_FuturesWS_Candlestick: %s", kline.TakerBuyBaseAssetVolume, candlestick.Symbol, kline.Interval, err.Error())
		Logger.ERROR(errStr)
		return nil, err
	}
	takerBuyQuoteAssetVolume, err := Utils.ParseFloat(kline.TakerBuyQuoteAssetVolume)
	if err != nil {
		errStr := fmt.Sprintf("There was an error parsing float '%s' of 'kline.TakerBuyQuoteAssetVolume' for '%s', interval '%s' in parseFloat_FuturesWS_Candlestick: %s", kline.TakerBuyQuoteAssetVolume, candlestick.Symbol, kline.Interval, err.Error())
		Logger.ERROR(errStr)
		return nil, err
	}

	return &FuturesWS_Candlestick_Float64{
		OpenTime:  kline.OpenTime,
		CloseTime: kline.CloseTime,

		Open:  open,
		High:  high,
		Low:   low,
		Close: close,

		Volume:                   baseAssetVolune,
		QuoteAssetVolume:         quoteAssetVolume,
		TakerBuyBaseAssetVolume:  takerBuyBaseAssetVolume,
		TakerBuyQuoteAssetVolume: takerBuyQuoteAssetVolume,
		TradeCount:               kline.TradeCount,
	}, nil
}

func (interval *FuturesWS_ManagedCandlesticks_Interval) Fetch_Newest_Candlesticks() error {
	if interval.Interval.Custom {
		return LocalError(DATA_NOT_FOUND_ERR, fmt.Sprintf("Unable to fetch candlesticks for custom interval '%s'", interval.Interval.Name))
	}

	newRawCandlesticks, _, err := interval.symbol.parent.futures_ws.binance.Futures.Candlesticks(
		interval.symbol.Symbol,
		interval.Interval.Name,
		Futures_Candlesticks_Params{
			Limit: 1500,
		},
	)
	if err != nil {
		return err
	}

	for i := len(newRawCandlesticks) - 1; i >= 0; i-- {
		newCandlestick, err := parseFloat_Futures_Candlestick(newRawCandlesticks[i])
		if err != nil {
			continue
		}
		interval.handleCandlestick(newCandlestick)
	}

	return nil
}

func (interval *FuturesWS_ManagedCandlesticks_Interval) Fetch_Older_Candlesticks() error {
	if interval.Interval.Custom {
		return LocalError(DATA_NOT_FOUND_ERR, fmt.Sprintf("Unable to fetch candlesticks for custom interval '%s'", interval.Interval.Name))
	}
	// oldest_candlestick
	var endTime int64 = 0
	if len(interval.Candlesticks) != 0 {
		endTime = interval.Candlesticks[0].OpenTime
	}

	newRawCandlesticks, _, err := interval.symbol.parent.futures_ws.binance.Futures.Candlesticks(
		interval.symbol.Symbol,
		interval.Interval.Name,
		Futures_Candlesticks_Params{
			EndTime: endTime,
			Limit:   1500,
		},
	)
	if err != nil {
		return err
	}

	for i := len(newRawCandlesticks) - 1; i >= 0; i-- {
		newCandlestick, err := parseFloat_Futures_Candlestick(newRawCandlesticks[i])
		if err != nil {
			continue
		}
		interval.handleCandlestick(newCandlestick)
	}

	return nil
}

func (managedCandlestick_symbol *FuturesWS_ManagedCandlesticks_Symbol) handleCandlestick(candlestick *FuturesWS_Candlestick) {
	managedCandlestick_symbol.Intervals.Mu.Lock()
	defer managedCandlestick_symbol.Intervals.Mu.Unlock()

	interval, _, parseErr := Utils.GetIntervalFromString(candlestick.Kline.Interval)
	if parseErr != nil {
		return
	}

	newCandlestick, err := parseFloat_FuturesWS_Candlestick(candlestick)
	if err == nil {
		return
	}

	for _, storedInterval := range managedCandlestick_symbol.Intervals.Map {
		// # The following makes sure of two things
		// 1- A candlestick of a higher interval doesn't update candlesticks of a lower interval (like receiving a 3m candle and updating a 1m candle, not ideal for tradecounts and such)
		// 2- A candlestick of the same interval rune (like 'm' for minutes) can only update similar candlesticks with the same rune IF it divides the interval (a candle of 1m can update 3m candles, but a 3m interval cannot update 5m candles)
		//
		// I am aware that since we are doing this, NONE of the Volumes will be updated accurately, lots of overwriting of bad data, but they will be accurate with the aggTrades received, thus the use of math.Max(<candleVolume>, <calculated_candleVolume>) is required
		// But the point of this is to have as much accuracy as possible, with as little data consumption as possible
		if storedInterval.Interval.Value < interval.Value { //  || interval.Interval.Value%intervalValue != 0
			continue
		}

		storedInterval.handleCandlestick(newCandlestick)
	}

}

func (managedCandlestick_symbol *FuturesWS_ManagedCandlesticks_Symbol) handleAggTrade(aggTrade *FuturesWS_AggTrade) {
	managedCandlestick_symbol.Intervals.Mu.Lock()
	defer managedCandlestick_symbol.Intervals.Mu.Unlock()

	price, err := Utils.ParseFloat(aggTrade.Price)
	if err != nil {
		Logger.ERROR(fmt.Sprintf("Failed to parse aggTrade.price '%s' in AddAggTrade for '%s': %s", aggTrade.Price, aggTrade.Symbol, err.Error()))
		return
	}
	quantity, err := Utils.ParseFloat(aggTrade.Quantity)
	if err != nil {
		Logger.ERROR(fmt.Sprintf("Failed to parse aggTrade.quantity '%s' in AddAggTrade for '%s': %s", aggTrade.Price, aggTrade.Symbol, err.Error()))
		return
	}

	new_ManagedAggTrade := &FuturesWS_ManagedCandlesticks_AggTrade{
		Timestamp:    aggTrade.Timestamp,
		IsMaker:      aggTrade.IsMaker,
		AggTradeId:   aggTrade.AggTradeId,
		FirstTradeId: aggTrade.FirstTradeId,
		LastTradeId:  aggTrade.LastTradeId,

		Price: price,
		Qty:   quantity,
	}

	managedCandlestick_symbol.AggTrades = append(managedCandlestick_symbol.AggTrades, new_ManagedAggTrade)

	for _, interval := range managedCandlestick_symbol.Intervals.Map {
		interval.handleAggTrade(new_ManagedAggTrade)
	}
}

////

func (interval *FuturesWS_ManagedCandlesticks_Interval) handleCandlestick(newCandlestick *FuturesWS_Candlestick_Float64) {
	supposed_openTime, supposed_closeTime, err := Utils.GetOpenCloseTimes(newCandlestick.OpenTime, interval.Interval.Name)
	if err != nil {
		return
	}

	if len(interval.Candlesticks) == 0 {
		new_managedCandlestick := &FuturesWS_ManagedCandlestick{
			OpenTime:  supposed_openTime,
			CloseTime: supposed_closeTime,

			Open:  newCandlestick.Open,
			High:  newCandlestick.High,
			Low:   newCandlestick.Low,
			Close: newCandlestick.Close,
		}

		new_managedCandlestick.update(newCandlestick)
		interval.Candlesticks = append(interval.Candlesticks, new_managedCandlestick)
		return
	}

	if interval.Candlesticks[len(interval.Candlesticks)-1].OpenTime < supposed_openTime {
		for {
			openTime, closeTime, err := Utils.GetOpenCloseTimes(interval.Candlesticks[len(interval.Candlesticks)-1].CloseTime+1, interval.Interval.Name)
			if err != nil {
				return
			}

			new_managedCandlestick := &FuturesWS_ManagedCandlestick{
				OpenTime:  openTime,
				CloseTime: closeTime,
			}
			new_managedCandlestick.Open = interval.Candlesticks[len(interval.Candlesticks)-1].Close
			new_managedCandlestick.High = interval.Candlesticks[len(interval.Candlesticks)-1].Close
			new_managedCandlestick.Low = interval.Candlesticks[len(interval.Candlesticks)-1].Close
			new_managedCandlestick.Close = interval.Candlesticks[len(interval.Candlesticks)-1].Close

			interval.Candlesticks = append(interval.Candlesticks, new_managedCandlestick)

			if new_managedCandlestick.OpenTime == supposed_openTime {
				new_managedCandlestick.update(newCandlestick)
				break
			}
		}
		return
	}

	for i := len(interval.Candlesticks) - 1; i >= 0; i-- {
		if interval.Candlesticks[i].OpenTime == supposed_openTime {
			interval.Candlesticks[i].update(newCandlestick)
			return
		}
	}

	// inserting older candle
	for {
		firstCandle := interval.Candlesticks[0]
		openTime, closeTime, err := Utils.GetOpenCloseTimes(firstCandle.OpenTime-1, interval.Interval.Name)
		if err != nil {
			continue
		}

		new_managedCandlestick := &FuturesWS_ManagedCandlestick{
			OpenTime:  openTime,
			CloseTime: closeTime,

			Open:  newCandlestick.Open,
			High:  newCandlestick.Open,
			Low:   newCandlestick.Open,
			Close: newCandlestick.Open,
		}

		interval.Candlesticks = append([]*FuturesWS_ManagedCandlestick{new_managedCandlestick}, interval.Candlesticks...)

		if new_managedCandlestick.OpenTime == supposed_openTime {
			new_managedCandlestick.update(newCandlestick)
			break
		}
	}
}

func (candle *FuturesWS_ManagedCandlestick) update(newCandlestick *FuturesWS_Candlestick_Float64) {
	candle.Open = newCandlestick.Open
	candle.High = math.Max(candle.High, newCandlestick.High)
	candle.Low = math.Min(candle.Low, newCandlestick.Low)
	candle.Close = newCandlestick.Close

	candle.Volume = newCandlestick.Volume
	candle.QuoteAssetVolume = newCandlestick.QuoteAssetVolume
	candle.TakerBuyBaseAssetVolume = newCandlestick.TakerBuyBaseAssetVolume
	candle.TakerBuyQuoteAssetVolume = newCandlestick.TakerBuyQuoteAssetVolume
	candle.TradeCount = newCandlestick.TradeCount
}

//

func (interval *FuturesWS_ManagedCandlesticks_Interval) handleAggTrade(aggTrade *FuturesWS_ManagedCandlesticks_AggTrade) {
	supposed_openTime, supposed_closeTime, err := Utils.GetOpenCloseTimes(aggTrade.Timestamp, interval.Interval.Name)
	if err != nil {
		return
	}

	if len(interval.Candlesticks) == 0 {
		new_managedCandlestick := &FuturesWS_ManagedCandlestick{
			OpenTime:  supposed_openTime,
			CloseTime: supposed_closeTime,

			Open:  aggTrade.Price,
			High:  aggTrade.Price,
			Low:   aggTrade.Price,
			Close: aggTrade.Price,
		}

		new_managedCandlestick.insertAggTrade(aggTrade)
		interval.Candlesticks = append(interval.Candlesticks, new_managedCandlestick)
		return
	}

	if interval.Candlesticks[len(interval.Candlesticks)-1].OpenTime < supposed_openTime {
		for {
			openTime, closeTime, err := Utils.GetOpenCloseTimes(interval.Candlesticks[len(interval.Candlesticks)-1].CloseTime+1, interval.Interval.Name)
			if err != nil {
				return
			}

			new_managedCandlestick := &FuturesWS_ManagedCandlestick{
				OpenTime:  openTime,
				CloseTime: closeTime,
			}

			new_managedCandlestick.Open = interval.Candlesticks[len(interval.Candlesticks)-1].Close
			new_managedCandlestick.High = interval.Candlesticks[len(interval.Candlesticks)-1].Close
			new_managedCandlestick.Low = interval.Candlesticks[len(interval.Candlesticks)-1].Close
			new_managedCandlestick.Close = interval.Candlesticks[len(interval.Candlesticks)-1].Close

			interval.Candlesticks = append(interval.Candlesticks, new_managedCandlestick)

			if new_managedCandlestick.OpenTime == supposed_openTime {
				new_managedCandlestick.insertAggTrade(aggTrade)
				break
			}
		}
		return
	}

	for i := len(interval.Candlesticks) - 1; i >= 0; i-- {
		if interval.Candlesticks[i].OpenTime == supposed_openTime {
			interval.Candlesticks[i].insertAggTrade(aggTrade)
			return
		}
	}

	// inserting older candle
	for {
		firstCandle := interval.Candlesticks[0]
		openTime, closeTime, err := Utils.GetOpenCloseTimes(firstCandle.OpenTime-1, interval.Interval.Name)
		if err != nil {
			return
		}

		new_managedCandlestick := &FuturesWS_ManagedCandlestick{
			OpenTime:  openTime,
			CloseTime: closeTime,

			Open:  firstCandle.Open,
			High:  firstCandle.Open,
			Low:   firstCandle.Open,
			Close: firstCandle.Open,
		}

		if new_managedCandlestick.OpenTime == supposed_openTime {
			new_managedCandlestick.insertAggTrade(aggTrade)
			break
		}

		interval.Candlesticks = append([]*FuturesWS_ManagedCandlestick{new_managedCandlestick}, interval.Candlesticks...)
	}
}

func (candle *FuturesWS_ManagedCandlestick) insertAggTrade(managedAggTrade *FuturesWS_ManagedCandlesticks_AggTrade) {
	candle.High = math.Max(candle.High, managedAggTrade.Price)
	candle.Low = math.Min(candle.Low, managedAggTrade.Price)
	candle.Close = managedAggTrade.Price

	quoteAsset_size := managedAggTrade.Qty * managedAggTrade.Price

	candle.Calculated_Volume += managedAggTrade.Qty
	candle.Calculated_QuoteAssetVolume += quoteAsset_size

	if !managedAggTrade.IsMaker {
		candle.Calculated_TakerBuyBaseAssetVolume += managedAggTrade.Qty
		candle.Calculated_TakerBuyQuoteAssetVolume += quoteAsset_size
	}

	candle.Calculated_TradeCount += managedAggTrade.LastTradeId - managedAggTrade.FirstTradeId

	candle.AggTrades = append(candle.AggTrades, managedAggTrade)
}

////

func (handler *FuturesWS_ManagedCandlesticks_Handler) Subscribe(symbols ...string) (hasTimedOut bool, err error) {
	if !handler.disable_aggTrades_stream {
		hasTimedOut, err = handler.AggTrades_Socket.Subscribe(symbols...)
		if err != nil {
			return hasTimedOut, err
		}
	}

	if !handler.disable_candlesticks_stream {
		params := make([]FuturesWS_Candlestick_Params, len(symbols))
		for i := range params {
			params[i].Symbol = symbols[i]
			params[i].Interval = "1m" // the smallest interval on futures
		}
		hasTimedOut, err = handler.Candlesticks_Socket.Subscribe(params...)
		if err != nil {
			return hasTimedOut, err
		}
	}

	handler.addSymbols(symbols...)

	return false, nil
}

func (handler *FuturesWS_ManagedCandlesticks_Handler) Unsubscribe(symbols ...string) (hasTimedOut bool, err error) {
	if !handler.disable_aggTrades_stream {
		hasTimedOut, err = handler.AggTrades_Socket.Unsubscribe(symbols...)
		if err != nil {
			return hasTimedOut, err
		}
	}

	if !handler.disable_candlesticks_stream {
		params := make([]FuturesWS_Candlestick_Params, len(symbols))
		for i := range params {
			params[i].Symbol = symbols[i]
			params[i].Interval = "1m" // the smallest interval on futures
		}
		hasTimedOut, err = handler.Candlesticks_Socket.Unsubscribe(params...)
		if err != nil {
			return hasTimedOut, err
		}
	}

	handler.removeSymbols(symbols...)

	return false, nil
}

////

func (handler *FuturesWS_ManagedCandlesticks_Handler) onCandlestick(candlestick *FuturesWS_Candlestick) *FuturesWS_ManagedCandlesticks_Symbol {
	symbol, exists := handler.Candlesticks.Symbols[candlestick.Symbol]
	if !exists {
		return nil
	}

	symbol.handleCandlestick(candlestick)

	return symbol
}

func (handler *FuturesWS_ManagedCandlesticks_Handler) onAggTrade(aggTrade *FuturesWS_AggTrade) *FuturesWS_ManagedCandlesticks_Symbol {
	symbol, exists := handler.Candlesticks.Symbols[aggTrade.Symbol]
	if !exists {
		return nil
	}

	symbol.handleAggTrade(aggTrade)

	return symbol
}

func (handler *FuturesWS_ManagedCandlesticks_Handler) addSymbols(symbols ...string) {
	handler.Candlesticks.Mu.Lock()
	defer handler.Candlesticks.Mu.Unlock()

	for _, symbol := range symbols {
		_, exists := handler.Candlesticks.Symbols[symbol]
		if exists {
			continue
		}

		managedCandlestick_symbol := &FuturesWS_ManagedCandlesticks_Symbol{
			parent:   handler,
			OnChange: New[*FuturesWS_ManagedCandlesticks_Symbol](),
		}
		managedCandlestick_symbol.Symbol = symbol
		managedCandlestick_symbol.AggTrades = make([]*FuturesWS_ManagedCandlesticks_AggTrade, 0)

		managedCandlestick_symbol.Intervals.Map = make(map[string]*FuturesWS_ManagedCandlesticks_Interval)

		for _, interval := range handler.intervals {
			managedCandlestick_symbol.Intervals.Map[interval.Name] = &FuturesWS_ManagedCandlesticks_Interval{
				symbol:       managedCandlestick_symbol,
				Interval:     interval,
				Candlesticks: make([]*FuturesWS_ManagedCandlestick, 0),
			}
		}

		handler.Candlesticks.Symbols[symbol] = managedCandlestick_symbol
	}
}

func (handler *FuturesWS_ManagedCandlesticks_Handler) removeSymbols(symbols ...string) {
	for _, symbol := range symbols {
		delete(handler.Candlesticks.Symbols, symbol)
	}
}

// These parameters are optional, not necessary to be filled out
type Managed_CustomCandlesticks_Params struct {
	Disable_aggTrades_stream    bool
	Disable_candlesticks_stream bool
	CustomIntervals             []string
}

// # WARNING
//
// This function is custom made for SPECIFIC purposes!
//
// This handler allowed for full candlestick intervals support (along with custom intervals as well), along with the aggTrades that come with it.
//
// That is done by opening a candlestick stream to the smallest possible interval (here its '1m') and an aggTrade stream to fetch the much faster updates, and update all the local candlestick intervals with both streams' data.
//
// When using the candlestick data (and specifically the volumes and all of their varieties, make sure to use `math.Max(<propertyName>, Calculated_<propertyName>)` to get the accurate binance data)
//
// NOTE: the first element (element '0') might be inaccurate since the streams might've started before fetching the full data (although rare), so make sure to wait a bit before using the first candlestick element.
// This happens due to the fact that the aggTrade and candlestick stream might've began in the middle of the candlestick's interval; using FuturesWS_ManageCandlesticks_Symbol.Fetch_Older_Candlesticks() fixes this.
//
// The library also 'creates' a new candlestick if the aggTrade stream pushed an aggTrade that falls within a newer candlestick than locally present, the OpenTime (although tested to be correct) might be incorrectly calculated, so help double checking here might be helpful.
func (futures_ws *Futures_Websockets) Managed_CustomCandlesticks(publicOnMessage func(symbol *FuturesWS_ManagedCandlesticks_Symbol), opt_params *Managed_CustomCandlesticks_Params, symbols ...string) (*FuturesWS_ManagedCandlesticks_Handler, error) {
	handler := &FuturesWS_ManagedCandlesticks_Handler{}
	handler.futures_ws = futures_ws

	handler.disable_aggTrades_stream = false
	handler.disable_candlesticks_stream = false
	if opt_params != nil {
		handler.disable_aggTrades_stream = opt_params.Disable_aggTrades_stream
		handler.disable_candlesticks_stream = opt_params.Disable_candlesticks_stream
	}

	customIntervals := []string{}
	if opt_params != nil && opt_params.CustomIntervals != nil {
		customIntervals = opt_params.CustomIntervals
	}

	intervals := []string{
		FUTURES_Constants.ChartIntervals.MIN,
		FUTURES_Constants.ChartIntervals.MINS_3,
		FUTURES_Constants.ChartIntervals.MINS_5,
		FUTURES_Constants.ChartIntervals.MINS_15,
		FUTURES_Constants.ChartIntervals.MINS_30,

		FUTURES_Constants.ChartIntervals.HOUR,
		FUTURES_Constants.ChartIntervals.HOURS_2,
		FUTURES_Constants.ChartIntervals.HOURS_4,
		FUTURES_Constants.ChartIntervals.HOURS_6,
		FUTURES_Constants.ChartIntervals.HOURS_8,
		FUTURES_Constants.ChartIntervals.HOURS_12,

		FUTURES_Constants.ChartIntervals.DAY,
		FUTURES_Constants.ChartIntervals.DAYS_3,
		FUTURES_Constants.ChartIntervals.WEEK,
		FUTURES_Constants.ChartIntervals.MONTH,
	}

	for _, interval := range intervals {
		found := false
		for _, presentInterval := range handler.intervals {
			if presentInterval.Name == interval {
				found = true
				break
			}
		}
		if found {
			continue
		}
		interval, _, parseErr := Utils.GetIntervalFromString(interval)
		if parseErr != nil {
			return handler, LocalError(PARSING_ERR, fmt.Sprintf("Failed to parse custom interval '%s': %s", interval.Name, parseErr.Error()))
		}
		interval.Custom = false
		handler.intervals = append(handler.intervals, interval)
	}

	forcedCustomIntervals := []string{"1s", "15s", "30s", "1Y"}
	for _, forcedCustomInterval := range forcedCustomIntervals {
		found := false
		for _, presentInterval := range handler.intervals {
			if presentInterval.Name == forcedCustomInterval {
				found = true
				break
			}
		}
		if found {
			continue
		}
		interval, _, parseErr := Utils.GetIntervalFromString(forcedCustomInterval)
		if parseErr != nil {
			return handler, LocalError(PARSING_ERR, fmt.Sprintf("Failed to parse custom interval '%s': %s", forcedCustomInterval, parseErr.Error()))
		}
		interval.Custom = true
		handler.intervals = append(handler.intervals, interval)
	}

	for _, customInterval := range customIntervals {
		found := false
		for _, interval := range handler.intervals {
			if interval.Name == customInterval {
				found = true
				break
			}
		}
		if found {
			continue
		}
		interval, _, parseErr := Utils.GetIntervalFromString(customInterval)
		if parseErr != nil {
			return handler, LocalError(PARSING_ERR, fmt.Sprintf("Failed to parse custom interval '%s': %s", customInterval, parseErr.Error()))
		}
		interval.Custom = true
		handler.intervals = append(handler.intervals, interval)
	}

	handler.Candlesticks.Symbols = make(map[string]*FuturesWS_ManagedCandlesticks_Symbol)

	handler.addSymbols(symbols...)

	//

	params := make([]FuturesWS_Candlestick_Params, len(symbols))
	for i := range symbols {
		params[i].Symbol = symbols[i]
		params[i].Interval = "1m" // smallest interval in futures
	}

	var candlesticks_socket *FuturesWS_Candlesticks_Socket

	if !handler.disable_candlesticks_stream {
		candlesticks_socket = futures_ws.Candlesticks(
			func(candlestick *FuturesWS_Candlestick) {
				handler.Candlesticks.Mu.Lock()
				defer handler.Candlesticks.Mu.Unlock()

				symbol := handler.onCandlestick(candlestick)
				if symbol != nil {
					symbol.OnChange.Emit(symbol)
					publicOnMessage(symbol)
				}
			},
			params...,
		)
		handler.Candlesticks_Socket = candlesticks_socket
	}

	if !handler.disable_aggTrades_stream {
		aggTrade_socket := futures_ws.AggTrade(
			func(aggTrade *FuturesWS_AggTrade) {
				handler.Candlesticks.Mu.Lock()
				defer handler.Candlesticks.Mu.Unlock()

				symbol := handler.onAggTrade(aggTrade)
				if symbol != nil {
					symbol.OnChange.Emit(symbol)
					publicOnMessage(symbol)
				}
			},
			symbols...,
		)
		handler.AggTrades_Socket = aggTrade_socket
	}

	return handler, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (*Futures_Websockets) CreateSocket(streams []string, isCombined bool) *Futures_Websocket {
	baseURL := FUTURES_Constants.Websocket.URLs[0]

	socket := websockets.CreateBinanceWebsocket(baseURL, streams)

	ws := &Futures_Websocket{
		base: socket,
	}

	return ws
}
