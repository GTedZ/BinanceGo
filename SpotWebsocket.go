package Binance

import (
	"fmt"
	"strconv"
	"strings"

	websockets "github.com/GTedZ/binancego/websockets"
)

type Spot_Websockets struct {
	binance *Binance
}

func (spot_ws *Spot_Websockets) init(binance *Binance) {
	spot_ws.binance = binance
}

type Spot_Websocket struct {
	base *websockets.BinanceWebsocket
}

func (spot_ws *Spot_Websocket) Close() {
	spot_ws.base.Close()
}

func (spot_ws *Spot_Websocket) SetMessageListener(f func(messageType int, msg []byte)) {
	spot_ws.base.OnMessage = f
}

// This is called when the socket has successfully reconnected after a disconnection
func (spot_ws *Spot_Websocket) SetReconnectListener(f func()) {
	spot_ws.base.OnReconnect = f
}

func (spot_ws *Spot_Websocket) ListSubscriptions(timeout_sec ...int) (subscriptions []string, err error) {
	return spot_ws.base.ListSubscriptions(timeout_sec...)
}

func (spot_ws *Spot_Websocket) Subscribe(stream ...string) (hasTimedOut bool, err error) {
	requestObj := make(map[string]interface{})
	requestObj["method"] = "SUBSCRIBE"
	requestObj["params"] = stream
	_, hasTimedOut, err = spot_ws.base.SendPrivateMessage(requestObj)
	if err != nil || hasTimedOut {
		return hasTimedOut, err
	}

	spot_ws.base.SetStreams(append(spot_ws.base.GetStreams(), stream...))
	spot_ws.base.UpdateStreams()

	Logger.INFO(fmt.Sprintf("Successfully Subscribed to %s", stream))

	return false, nil
}

func (spot_ws *Spot_Websocket) Unsubscribe(stream ...string) (hasTimedOut bool, err error) {
	requestObj := make(map[string]interface{})
	requestObj["method"] = "UNSUBSCRIBE"
	requestObj["params"] = stream
	_, timedOut, err := spot_ws.base.SendPrivateMessage(requestObj)
	if err != nil || timedOut {
		return timedOut, err
	}

	// Filter out the streams to remove from spot_ws.Websocket.Streams
	streamMap := make(map[string]bool)
	for _, s := range stream {
		streamMap[s] = true
	}

	var updatedStreams []string
	for _, existingStream := range spot_ws.base.GetStreams() {
		if !streamMap[existingStream] {
			updatedStreams = append(updatedStreams, existingStream)
		}
	}
	spot_ws.base.SetStreams(updatedStreams)

	Logger.INFO(fmt.Sprintf("successfully Unsubscribed from %s", stream))

	return false, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type SpotWS_AggTrade struct {
	Event     string `json:"e"`
	EventTime int64  `json:"E"`
	Symbol    string `json:"s"`
	// Trade ID
	AggTradeId int64  `json:"a"`
	Price      string `json:"p"`
	Quantity   string `json:"q"`
	// First Trade ID
	FirstTradeId int64 `json:"f"`
	// Last Trade ID
	LastTradeId int64 `json:"l"`
	// Trade time
	Timestamp int64 `json:"T"`
	// Is the buyer the market maker?
	IsMaker bool `json:"m"`
	// Ignore
	Ignore bool `json:"M"`
}

type SpotWS_AggTrade_Socket struct {
	Handler *Spot_Websocket
}

func (*SpotWS_AggTrade_Socket) CreateStreamName(symbol string) string {
	return strings.ToLower(symbol) + "@aggTrade"
}

func (socket *SpotWS_AggTrade_Socket) Subscribe(symbol ...string) (hasTimedOut bool, err error) {
	for i := range symbol {
		symbol[i] = socket.CreateStreamName(symbol[i])
	}

	return socket.Handler.Subscribe(symbol...)
}
func (socket *SpotWS_AggTrade_Socket) Unsubscribe(symbol ...string) (hasTimedOut bool, err error) {
	for i := range symbol {
		symbol[i] = socket.CreateStreamName(symbol[i])
	}
	return socket.Handler.Unsubscribe(symbol...)
}

func (spot_ws *Spot_Websockets) AggTrade(publicOnMessage func(aggTrade *SpotWS_AggTrade), symbol ...string) (*SpotWS_AggTrade_Socket, error) {
	var newSocket SpotWS_AggTrade_Socket
	for i := range symbol {
		symbol[i] = newSocket.CreateStreamName(symbol[i])
	}
	socket := spot_ws.CreateSocket(symbol)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var aggTrade SpotWS_AggTrade
		err := json.Unmarshal(msg, &aggTrade)
		if err != nil {
			LocalError(PARSING_ERR, err.Error())
			return
		}
		publicOnMessage(&aggTrade)
	}

	newSocket.Handler = socket
	return &newSocket, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type SpotWS_Trade struct {
	Event     string `json:"e"`
	EventTime int64  `json:"E"`
	Symbol    string `json:"s"`
	// Trade ID
	TradeID  int64  `json:"t"`
	Price    string `json:"p"`
	Quantity string `json:"q"`
	// Trade time
	Timestamp int64 `json:"T"`
	// Is the buyer the market maker?
	IsMaker bool `json:"m"`
	// Ignore
	Ignore bool `json:"M"`
}

type SpotWS_Trade_Socket struct {
	Handler *Spot_Websocket
}

func (*SpotWS_Trade_Socket) CreateStreamName(symbol string) string {
	return strings.ToLower(symbol) + "@trade"
}

func (socket *SpotWS_Trade_Socket) Subscribe(symbol ...string) (hasTimedOut bool, err error) {
	for i := range symbol {
		symbol[i] = socket.CreateStreamName(symbol[i])
	}

	return socket.Handler.Subscribe(symbol...)
}
func (socket *SpotWS_Trade_Socket) Unsubscribe(symbol ...string) (hasTimedOut bool, err error) {
	for i := range symbol {
		symbol[i] = socket.CreateStreamName(symbol[i])
	}
	return socket.Handler.Unsubscribe(symbol...)
}

func (spot_ws *Spot_Websockets) Trade(publicOnMessage func(trade *SpotWS_Trade), symbol ...string) (*SpotWS_Trade_Socket, error) {
	var newSocket SpotWS_Trade_Socket
	for i := range symbol {
		symbol[i] = newSocket.CreateStreamName(symbol[i])
	}
	socket := spot_ws.CreateSocket(symbol)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var trade SpotWS_Trade
		err := json.Unmarshal(msg, &trade)
		if err != nil {
			LocalError(PARSING_ERR, err.Error())
			return
		}
		publicOnMessage(&trade)
	}

	newSocket.Handler = socket
	return &newSocket, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type SpotWS_Candlestick_MSG struct {

	// Event type
	Event string `json:"e"`

	// Event time
	EventTime int64 `json:"E"`

	// Symbol
	Symbol string `json:"s"`

	Candle *SpotWS_Candlestick `json:"k"`
}
type SpotWS_Candlestick struct {

	// Kline start time
	OpenTime int64 `json:"t"`

	// Kline close time
	CloseTime int64 `json:"T"`

	// Symbol
	Symbol string `json:"s"`

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

	// Base asset volume
	BaseAssetVolume string `json:"v"`

	// Number of trades
	TradeCount int64 `json:"n"`

	// Is this kline closed?
	IsClosed bool `json:"x"`

	// Quote asset volume
	QuoteAssetVolume string `json:"q"`

	// Taker buy base asset volume
	TakerBuyBaseAssetVolume string `json:"V"`

	// Taker buy quote asset volume
	TakerBuyQuoteAssetVolume string `json:"Q"`

	// Ignore
	Ignore string `json:"B"`
}

type SpotWS_Candlestick_StreamIdentifier struct {
	Symbol   string
	Interval string
}

type SpotWS_Candlestick_Socket struct {
	Handler *Spot_Websocket
}

func (*SpotWS_Candlestick_Socket) CreateStreamName(symbol string, interval string) string {
	return strings.ToLower(symbol) + "@kline_" + interval
}

func (socket *SpotWS_Candlestick_Socket) Subscribe(identifiers ...SpotWS_Candlestick_StreamIdentifier) (hasTimedOut bool, err error) {
	streams := make([]string, len(identifiers))
	for i, id := range identifiers {
		streams[i] = socket.CreateStreamName(id.Symbol, id.Interval)
	}

	return socket.Handler.Subscribe(streams...)
}
func (socket *SpotWS_Candlestick_Socket) Unsubscribe(identifiers ...SpotWS_Candlestick_StreamIdentifier) (hasTimedOut bool, err error) {
	streams := make([]string, len(identifiers))
	for i, id := range identifiers {
		streams[i] = socket.CreateStreamName(id.Symbol, id.Interval)
	}

	return socket.Handler.Unsubscribe(streams...)
}

func (spot_ws *Spot_Websockets) Candlesticks(publicOnMessage func(candlestick_msg *SpotWS_Candlestick_MSG), identifiers ...SpotWS_Candlestick_StreamIdentifier) (*SpotWS_Candlestick_Socket, error) {
	var newSocket SpotWS_Candlestick_Socket

	streams := make([]string, len(identifiers))
	for i, id := range identifiers {
		streams[i] = newSocket.CreateStreamName(id.Symbol, id.Interval)
	}

	socket := spot_ws.CreateSocket(streams)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var candlestick_msg SpotWS_Candlestick_MSG
		err := json.Unmarshal(msg, &candlestick_msg)
		if err != nil {
			LocalError(PARSING_ERR, err.Error())
			return
		}
		publicOnMessage(&candlestick_msg)
	}

	newSocket.Handler = socket
	return &newSocket, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type SpotWS_Candlestick_TimezoneOffset_Socket struct {
	Handler *Spot_Websocket
}

func (*SpotWS_Candlestick_TimezoneOffset_Socket) CreateStreamName(symbol string, interval string) string {
	return strings.ToLower(symbol) + "@kline_" + interval + "@+08:00"
}

func (socket *SpotWS_Candlestick_TimezoneOffset_Socket) Subscribe(identifiers ...SpotWS_Candlestick_StreamIdentifier) (hasTimedOut bool, err error) {
	streams := make([]string, len(identifiers))
	for i, id := range identifiers {
		streams[i] = socket.CreateStreamName(id.Symbol, id.Interval)
	}

	return socket.Handler.Subscribe(streams...)
}
func (socket *SpotWS_Candlestick_TimezoneOffset_Socket) Unsubscribe(identifiers ...SpotWS_Candlestick_StreamIdentifier) (hasTimedOut bool, err error) {
	streams := make([]string, len(identifiers))
	for i, id := range identifiers {
		streams[i] = socket.CreateStreamName(id.Symbol, id.Interval)
	}

	return socket.Handler.Unsubscribe(streams...)
}

func (spot_ws *Spot_Websockets) Candlestick_WithOffset(publicOnMessage func(candlestick_msg *SpotWS_Candlestick_MSG), identifiers ...SpotWS_Candlestick_StreamIdentifier) (*SpotWS_Candlestick_TimezoneOffset_Socket, error) {
	var newSocket SpotWS_Candlestick_TimezoneOffset_Socket

	streams := make([]string, len(identifiers))
	for i, id := range identifiers {
		streams[i] = newSocket.CreateStreamName(id.Symbol, id.Interval)
	}
	socket := spot_ws.CreateSocket(streams)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var candlestick_msg SpotWS_Candlestick_MSG
		err := json.Unmarshal(msg, &candlestick_msg)
		if err != nil {
			LocalError(PARSING_ERR, err.Error())
			return
		}
		publicOnMessage(&candlestick_msg)
	}

	newSocket.Handler = socket
	return &newSocket, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type SpotWS_MiniTicker struct {

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

type SpotWS_MiniTicker_Socket struct {
	Handler *Spot_Websocket
}

func (*SpotWS_MiniTicker_Socket) CreateStreamName(symbol string) string {
	return strings.ToLower(symbol) + "@miniTicker"
}

func (socket *SpotWS_MiniTicker_Socket) Subscribe(symbol ...string) (hasTimedOut bool, err error) {
	for i := range symbol {
		symbol[i] = socket.CreateStreamName(symbol[i])
	}

	return socket.Handler.Subscribe(symbol...)
}
func (socket *SpotWS_MiniTicker_Socket) Unsubscribe(symbol ...string) (hasTimedOut bool, err error) {
	for i := range symbol {
		symbol[i] = socket.CreateStreamName(symbol[i])
	}

	return socket.Handler.Unsubscribe(symbol...)
}

func (spot_ws *Spot_Websockets) MiniTicker(publicOnMessage func(miniTicker *SpotWS_MiniTicker), symbol ...string) (*SpotWS_MiniTicker_Socket, error) {
	var newSocket SpotWS_MiniTicker_Socket
	for i := range symbol {
		symbol[i] = newSocket.CreateStreamName(symbol[i])
	}
	socket := spot_ws.CreateSocket(symbol)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var miniTicker SpotWS_MiniTicker
		err := json.Unmarshal(msg, &miniTicker)
		if err != nil {
			LocalError(PARSING_ERR, err.Error())
			return
		}
		publicOnMessage(&miniTicker)
	}

	newSocket.Handler = socket
	return &newSocket, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type SpotWS_AllMiniTickers_Socket struct {
	Handler *Spot_Websocket
}

func (*SpotWS_AllMiniTickers_Socket) CreateStreamName() string {
	return "!miniTicker@arr"
}

func (socket *SpotWS_AllMiniTickers_Socket) Subscribe() (hasTimedOut bool, err error) {
	return socket.Handler.Subscribe(socket.CreateStreamName())
}
func (socket *SpotWS_AllMiniTickers_Socket) Unsubscribe() (hasTimedOut bool, err error) {
	return socket.Handler.Unsubscribe(socket.CreateStreamName())
}

func (spot_ws *Spot_Websockets) AllMiniTickers(publicOnMessage func(miniTickers []*SpotWS_MiniTicker)) (*SpotWS_AllMiniTickers_Socket, error) {
	var newSocket SpotWS_AllMiniTickers_Socket
	socket := spot_ws.CreateSocket([]string{newSocket.CreateStreamName()})

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var miniTickers []*SpotWS_MiniTicker
		err := json.Unmarshal(msg, &miniTickers)
		if err != nil {
			LocalError(PARSING_ERR, err.Error())
			return
		}
		publicOnMessage(miniTickers)
	}

	newSocket.Handler = socket
	return &newSocket, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type SpotWS_Ticker struct {

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

	// First trade(F)-1 price (first trade before the 24hr rolling window)
	PreviousFirstTradePrice string `json:"x"`

	// Last price
	LastPrice string `json:"c"`

	// Last quantity
	LastQty string `json:"Q"`

	// Best bid price
	Bid string `json:"b"`

	// Best bid quantity
	BidQty string `json:"B"`

	// Best ask price
	Ask string `json:"a"`

	// Best ask quantity
	AskQty string `json:"A"`

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

type SpotWS_Ticker_Socket struct {
	Handler *Spot_Websocket
}

func (*SpotWS_Ticker_Socket) CreateStreamName(symbol string) string {
	return strings.ToLower(symbol) + "@ticker"
}

func (socket *SpotWS_Ticker_Socket) Subscribe(symbol ...string) (hasTimedOut bool, err error) {
	for i := range symbol {
		symbol[i] = socket.CreateStreamName(symbol[i])
	}

	return socket.Handler.Subscribe(symbol...)
}
func (socket *SpotWS_Ticker_Socket) Unsubscribe(symbol ...string) (hasTimedOut bool, err error) {
	for i := range symbol {
		symbol[i] = socket.CreateStreamName(symbol[i])
	}

	return socket.Handler.Unsubscribe(symbol...)
}

func (spot_ws *Spot_Websockets) Ticker(publicOnMessage func(ticker *SpotWS_Ticker), symbol ...string) (*SpotWS_Ticker_Socket, error) {
	var newSocket SpotWS_Ticker_Socket
	for i := range symbol {
		symbol[i] = newSocket.CreateStreamName(symbol[i])
	}
	socket := spot_ws.CreateSocket(symbol)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var ticker SpotWS_Ticker
		err := json.Unmarshal(msg, &ticker)
		if err != nil {
			LocalError(PARSING_ERR, err.Error())
			return
		}
		publicOnMessage(&ticker)
	}

	newSocket.Handler = socket
	return &newSocket, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type SpotWS_AllTickers_Socket struct {
	Handler *Spot_Websocket
}

func (*SpotWS_AllTickers_Socket) CreateStreamName() string {
	return "!ticker@arr"
}

func (socket *SpotWS_AllTickers_Socket) Subscribe() (hasTimedOut bool, err error) {
	return socket.Handler.Subscribe(socket.CreateStreamName())
}
func (socket *SpotWS_AllTickers_Socket) Unsubscribe() (hasTimedOut bool, err error) {
	return socket.Handler.Unsubscribe(socket.CreateStreamName())
}

func (spot_ws *Spot_Websockets) AllTickers(publicOnMessage func(tickers []*SpotWS_Ticker)) (*SpotWS_AllTickers_Socket, error) {
	var newSocket SpotWS_AllTickers_Socket
	socket := spot_ws.CreateSocket([]string{newSocket.CreateStreamName()})

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var tickers []*SpotWS_Ticker
		err := json.Unmarshal(msg, &tickers)
		if err != nil {
			LocalError(PARSING_ERR, err.Error())
			return
		}
		publicOnMessage(tickers)
	}

	newSocket.Handler = socket
	return &newSocket, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type SpotWS_RollingWindowStatistic struct {

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

	// Open price
	Open string `json:"o"`

	// High price
	High string `json:"h"`

	// Low price
	Low string `json:"l"`

	// Last price
	LastPrice string `json:"c"`

	// Weighted average price
	WeightedAveragePrice string `json:"w"`

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

type SpotWS_RollingWindowStatistics_StreamIdentifier struct {
	Symbol     string
	WindowSize string
}

type SpotWS_RollingWindowStatistics_Socket struct {
	Handler *Spot_Websocket
}

func (*SpotWS_RollingWindowStatistics_Socket) CreateStreamName(symbol string, windowSize string) string {
	return strings.ToLower(symbol) + "@ticker_" + windowSize
}

func (socket *SpotWS_RollingWindowStatistics_Socket) Subscribe(identifiers ...SpotWS_RollingWindowStatistics_StreamIdentifier) (hasTimedOut bool, err error) {
	streams := make([]string, len(identifiers))
	for i, id := range identifiers {
		streams[i] = socket.CreateStreamName(id.Symbol, id.WindowSize)
	}

	return socket.Handler.Subscribe(streams...)
}
func (socket *SpotWS_RollingWindowStatistics_Socket) Unsubscribe(identifiers ...SpotWS_RollingWindowStatistics_StreamIdentifier) (hasTimedOut bool, err error) {
	streams := make([]string, len(identifiers))
	for i, id := range identifiers {
		streams[i] = socket.CreateStreamName(id.Symbol, id.WindowSize)
	}

	return socket.Handler.Unsubscribe(streams...)
}

func (spot_ws *Spot_Websockets) RollingWindowStatistics(publicOnMessage func(rwStat *SpotWS_RollingWindowStatistic), identifiers ...SpotWS_RollingWindowStatistics_StreamIdentifier) (*SpotWS_RollingWindowStatistics_Socket, error) {
	var newSocket SpotWS_RollingWindowStatistics_Socket

	streams := make([]string, len(identifiers))
	for i, id := range identifiers {
		streams[i] = newSocket.CreateStreamName(id.Symbol, id.WindowSize)
	}

	socket := spot_ws.CreateSocket(streams)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var rwStat SpotWS_RollingWindowStatistic
		err := json.Unmarshal(msg, &rwStat)
		if err != nil {
			LocalError(PARSING_ERR, err.Error())
			return
		}
		publicOnMessage(&rwStat)
	}

	newSocket.Handler = socket
	return &newSocket, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type SpotWS_AllRollingWindowStatistics_Socket struct {
	Handler *Spot_Websocket
}

func (*SpotWS_AllRollingWindowStatistics_Socket) CreateStreamName(WindowSize string) string {
	return "!ticker_" + WindowSize + "@arr"
}

func (socket *SpotWS_AllRollingWindowStatistics_Socket) Subscribe(WindowSize ...string) (hasTimedOut bool, err error) {
	for i := range WindowSize {
		WindowSize[i] = socket.CreateStreamName(WindowSize[i])
	}

	return socket.Handler.Subscribe(WindowSize...)
}
func (socket *SpotWS_AllRollingWindowStatistics_Socket) Unsubscribe(WindowSize ...string) (hasTimedOut bool, err error) {
	for i := range WindowSize {
		WindowSize[i] = socket.CreateStreamName(WindowSize[i])
	}

	return socket.Handler.Unsubscribe(WindowSize...)
}

func (spot_ws *Spot_Websockets) AllRollingWindowStatistics(publicOnMessage func(rwStats []*SpotWS_RollingWindowStatistic), WindowSize ...string) (*SpotWS_AllRollingWindowStatistics_Socket, error) {
	var newSocket SpotWS_AllRollingWindowStatistics_Socket

	for i := range WindowSize {
		WindowSize[i] = newSocket.CreateStreamName(WindowSize[i])
	}

	socket := spot_ws.CreateSocket(WindowSize)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var rwStats []*SpotWS_RollingWindowStatistic
		err := json.Unmarshal(msg, &rwStats)
		if err != nil {
			LocalError(PARSING_ERR, err.Error())
			return
		}
		publicOnMessage(rwStats)
	}

	newSocket.Handler = socket
	return &newSocket, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type SpotWS_BookTicker struct {

	// order book updateId
	UpdateId int64 `json:"u"`

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

type SpotWS_BookTicker_Socket struct {
	Handler *Spot_Websocket
}

func (*SpotWS_BookTicker_Socket) CreateStreamName(symbol string) string {
	return strings.ToLower(symbol) + "@bookTicker"
}

func (socket *SpotWS_BookTicker_Socket) Subscribe(symbol ...string) (hasTimedOut bool, err error) {
	for i := range symbol {
		symbol[i] = socket.CreateStreamName(symbol[i])
	}

	return socket.Handler.Subscribe(symbol...)
}
func (socket *SpotWS_BookTicker_Socket) Unsubscribe(symbol ...string) (hasTimedOut bool, err error) {
	for i := range symbol {
		symbol[i] = socket.CreateStreamName(symbol[i])
	}

	return socket.Handler.Unsubscribe(symbol...)
}

func (spot_ws *Spot_Websockets) BookTicker(publicOnMessage func(bookTicker *SpotWS_BookTicker), symbol ...string) (*SpotWS_BookTicker_Socket, error) {
	var newSocket SpotWS_BookTicker_Socket

	for i := range symbol {
		symbol[i] = newSocket.CreateStreamName(symbol[i])
	}

	socket := spot_ws.CreateSocket(symbol)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var bookTicker *SpotWS_BookTicker
		err := json.Unmarshal(msg, &bookTicker)
		if err != nil {
			LocalError(PARSING_ERR, err.Error())
			return
		}
		publicOnMessage(bookTicker)
	}

	newSocket.Handler = socket
	return &newSocket, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type SpotWS_AveragePrice struct {

	// Event type
	Event string `json:"e"`

	// Event time
	EventTime int64 `json:"E"`

	// Symbol
	Symbol string `json:"s"`

	// Average price interval
	Interval string `json:"i"`

	// Average price
	AveragePrice string `json:"w"`

	// Last trade time
	Timestamp int64 `json:"T"`
}

type SpotWS_AveragePrice_Socket struct {
	Handler *Spot_Websocket
}

func (*SpotWS_AveragePrice_Socket) CreateStreamName(symbol string) string {
	return strings.ToLower(symbol) + "@avgPrice"
}

func (socket *SpotWS_AveragePrice_Socket) Subscribe(symbol ...string) (hasTimedOut bool, err error) {
	for i := range symbol {
		symbol[i] = socket.CreateStreamName(symbol[i])
	}

	return socket.Handler.Subscribe(symbol...)
}
func (socket *SpotWS_AveragePrice_Socket) Unsubscribe(symbol ...string) (hasTimedOut bool, err error) {
	for i := range symbol {
		symbol[i] = socket.CreateStreamName(symbol[i])
	}

	return socket.Handler.Unsubscribe(symbol...)
}

func (spot_ws *Spot_Websockets) AveragePrice(publicOnMessage func(averagePrice *SpotWS_AveragePrice), symbol ...string) (*SpotWS_AveragePrice_Socket, error) {
	var newSocket SpotWS_AveragePrice_Socket

	for i := range symbol {
		symbol[i] = newSocket.CreateStreamName(symbol[i])
	}

	socket := spot_ws.CreateSocket(symbol)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var averagePrice *SpotWS_AveragePrice
		err := json.Unmarshal(msg, &averagePrice)
		if err != nil {
			LocalError(PARSING_ERR, err.Error())
			return
		}
		publicOnMessage(averagePrice)
	}

	newSocket.Handler = socket
	return &newSocket, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//	{
//		"LastUpdateId": 160,  // Last update ID
//		"Bids": [             // Bids to be updated
//		  [
//			"0.0024",         // Price level to be updated
//			"10"              // Quantity
//		  ]
//		],
//		"Asks": [             // Asks to be updated
//		  [
//			"0.0026",         // Price level to be updated
//			"100"             // Quantity
//		  ]
//		]
//	}
type SpotWS_PartialBookDepth struct {

	// Last update ID
	LastUpdateId int64 `json:"lastUpdateId"`

	// Bids to be updated
	// [
	//   [
	// 	"0.0024",         // Price level to be updated
	// 	"10"              // Quantity
	//   ]
	// ],
	// ...
	Bids [][2]string `json:"bids"`

	// Asks to be updated
	// [
	//   [
	// 	"0.0026",    // Price level to be updated
	// 	"100"        // Quantity
	//   ],
	//   ...
	// ]
	Asks [][2]string `json:"asks"`
}

type SpotWS_PartialBookDepth_StreamIdentifier struct {
	Symbol string

	// Accepted values: 5, 10 or 20
	Levels int

	// # Stream Push Interval
	//
	// false -> 1000ms updates
	//
	// true  -> 100ms updates
	IsFast bool
}

type SpotWS_PartialBookDepth_Socket struct {
	Handler *Spot_Websocket
}

func (*SpotWS_PartialBookDepth_Socket) CreateStreamName(symbol string, levels int, isFast bool) string {
	streamName := strings.ToLower(symbol) + "@depth" + strconv.FormatInt(int64(levels), 10)
	if isFast {
		streamName += "@100ms"
	}
	return streamName
}

func (socket *SpotWS_PartialBookDepth_Socket) Subscribe(identifiers ...SpotWS_PartialBookDepth_StreamIdentifier) (hasTimedOut bool, err error) {
	streams := make([]string, len(identifiers))
	for i, id := range identifiers {
		streams[i] = socket.CreateStreamName(id.Symbol, id.Levels, id.IsFast)
	}

	return socket.Handler.Subscribe(streams...)
}
func (socket *SpotWS_PartialBookDepth_Socket) Unsubscribe(identifiers ...SpotWS_PartialBookDepth_StreamIdentifier) (hasTimedOut bool, err error) {
	streams := make([]string, len(identifiers))
	for i, id := range identifiers {
		streams[i] = socket.CreateStreamName(id.Symbol, id.Levels, id.IsFast)
	}

	return socket.Handler.Unsubscribe(streams...)
}

func (spot_ws *Spot_Websockets) PartialBookDepth(publicOnMessage func(partialBookDepth *SpotWS_PartialBookDepth), identifiers ...SpotWS_PartialBookDepth_StreamIdentifier) (*SpotWS_PartialBookDepth_Socket, error) {
	var newSocket SpotWS_PartialBookDepth_Socket

	streams := make([]string, len(identifiers))
	for i, id := range identifiers {
		streams[i] = newSocket.CreateStreamName(id.Symbol, id.Levels, id.IsFast)
	}

	socket := spot_ws.CreateSocket(streams)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var partialBookDepth *SpotWS_PartialBookDepth
		err := json.Unmarshal(msg, &partialBookDepth)
		if err != nil {
			LocalError(PARSING_ERR, err.Error())
			return
		}
		publicOnMessage(partialBookDepth)
	}

	newSocket.Handler = socket
	return &newSocket, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//	{
//		"Event": "depthUpdate", // Event type
//		"EventTime": 1672515782136, // Event time
//		"Symbol": "BNBBTC",      // Symbol
//		"FirstUpdateId": 157,           // First update ID in event
//		"FinalUpdateId": 160,           // Final update ID in event
//		"Bids": [
//		  [
//			"0.0024",       // Price level to be updated
//			"10"            // Quantity
//		  ]
//		],
//		"Asks": [
//		  [
//			"0.0026",       // Price level to be updated
//			"100"           // Quantity
//		  ]
//		]
//	  }
type SpotWS_DiffBookDepth struct {

	// Event type
	Event string `json:"e"`

	// Event time
	EventTime int64 `json:"E"`

	// Symbol
	Symbol string `json:"s"`

	// First update ID in event
	FirstUpdateId int64 `json:"U"`

	// Final update ID in event
	FinalUpdateId int64 `json:"u"`

	// Bids to be updated
	Bids [][2]string `json:"b"`

	// Asks to be updated
	Asks [][2]string `json:"a"`
}

type SpotWS_DiffBookDepth_StreamIdentifier struct {
	Symbol string

	// # Stream Push Interval
	//
	// false -> 1000ms updates
	//
	// true  -> 100ms updates
	IsFast bool
}

type SpotWS_DiffBookDepth_Socket struct {
	Handler *Spot_Websocket
}

func (*SpotWS_DiffBookDepth_Socket) CreateStreamName(symbol string, isFast bool) string {
	streamName := strings.ToLower(symbol) + "@depth"
	if isFast {
		streamName += "@100ms"
	}
	return streamName
}

func (socket *SpotWS_DiffBookDepth_Socket) Subscribe(identifiers ...SpotWS_DiffBookDepth_StreamIdentifier) (hasTimedOut bool, err error) {
	streams := make([]string, len(identifiers))
	for i, id := range identifiers {
		streams[i] = socket.CreateStreamName(id.Symbol, id.IsFast)
	}

	return socket.Handler.Subscribe(streams...)
}
func (socket *SpotWS_DiffBookDepth_Socket) Unsubscribe(identifiers ...SpotWS_DiffBookDepth_StreamIdentifier) (hasTimedOut bool, err error) {
	streams := make([]string, len(identifiers))
	for i, id := range identifiers {
		streams[i] = socket.CreateStreamName(id.Symbol, id.IsFast)
	}

	return socket.Handler.Unsubscribe(streams...)
}

func (spot_ws *Spot_Websockets) DiffBookDepth(publicOnMessage func(diffBookDepth *SpotWS_DiffBookDepth), identifiers ...SpotWS_DiffBookDepth_StreamIdentifier) (*SpotWS_DiffBookDepth_Socket, error) {
	var newSocket SpotWS_DiffBookDepth_Socket

	streams := make([]string, len(identifiers))
	for i, id := range identifiers {
		streams[i] = newSocket.CreateStreamName(id.Symbol, id.IsFast)
	}

	socket := spot_ws.CreateSocket(streams)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var diffBookDepth *SpotWS_DiffBookDepth
		err := json.Unmarshal(msg, &diffBookDepth)
		if err != nil {
			LocalError(PARSING_ERR, err.Error())
			return
		}
		publicOnMessage(diffBookDepth)
	}

	newSocket.Handler = socket
	return &newSocket, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (*Spot_Websockets) CreateSocket(streams []string) *Spot_Websocket {
	baseURL := SPOT_Constants.Websocket.URLs[0]

	socket := websockets.CreateBinanceWebsocket(baseURL, streams)

	ws := &Spot_Websocket{
		base: socket,
	}

	return ws
}
