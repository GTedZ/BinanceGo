package Binance

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/GTedZ/binancego/lib"
	websockets "github.com/GTedZ/binancego/websockets"
)

type spot_ws struct {
	binance *Binance
}

func (spot_ws *spot_ws) init(binance *Binance) {
	spot_ws.binance = binance
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Spot_Websocket struct {
	base *websockets.BinanceWebsocket

	// To be used by the library
	onMessage func(messageType int, msg []byte)
	// To be used by the library
	onReconnect func()

	// You can use this to register your own callback
	OnMessage func(messageType int, msg []byte)
	// You can use this to register your own callback
	OnReconnect func()
}

func (spot_ws *Spot_Websocket) onMsg(messageType int, msg []byte) {
	if spot_ws.onMessage != nil {
		spot_ws.onMessage(messageType, msg)
	}

	if spot_ws.OnMessage != nil {
		spot_ws.OnMessage(messageType, msg)
	}
}

func (spot_ws *Spot_Websocket) onReconn() {
	if spot_ws.onReconnect != nil {
		spot_ws.onReconnect()
	}

	if spot_ws.OnReconnect != nil {
		spot_ws.OnReconnect()
	}
}

////

func (spot_ws *Spot_Websocket) Close() {
	spot_ws.base.Close()
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

	spot_ws.base.RemoveStreams(stream)

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
	Socket *Spot_Websocket
}

func (*SpotWS_AggTrade_Socket) CreateStreamName(symbol string) string {
	return strings.ToLower(symbol) + "@aggTrade"
}

func (socket *SpotWS_AggTrade_Socket) Subscribe(symbol ...string) (hasTimedOut bool, err error) {
	for i := range symbol {
		symbol[i] = socket.CreateStreamName(symbol[i])
	}

	return socket.Socket.Subscribe(symbol...)
}
func (socket *SpotWS_AggTrade_Socket) Unsubscribe(symbol ...string) (hasTimedOut bool, err error) {
	for i := range symbol {
		symbol[i] = socket.CreateStreamName(symbol[i])
	}
	return socket.Socket.Unsubscribe(symbol...)
}

func (spot_ws *spot_ws) AggTrade(publicOnMessage func(aggTrade *SpotWS_AggTrade), symbol ...string) (*SpotWS_AggTrade_Socket, error) {
	var newSocket SpotWS_AggTrade_Socket
	for i := range symbol {
		symbol[i] = newSocket.CreateStreamName(symbol[i])
	}
	socket, err := spot_ws.CreateSocket(symbol)
	if err != nil {
		return nil, err
	}

	socket.onMessage = func(messageType int, msg []byte) {
		var aggTrade SpotWS_AggTrade
		err := json.Unmarshal(msg, &aggTrade)
		if err != nil {
			lib.LocalError(LibraryErrorCodes.PARSE_ERR, err.Error())
			return
		}
		publicOnMessage(&aggTrade)
	}

	newSocket.Socket = socket
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
	Socket *Spot_Websocket
}

func (*SpotWS_Trade_Socket) CreateStreamName(symbol string) string {
	return strings.ToLower(symbol) + "@trade"
}

func (socket *SpotWS_Trade_Socket) Subscribe(symbol ...string) (hasTimedOut bool, err error) {
	for i := range symbol {
		symbol[i] = socket.CreateStreamName(symbol[i])
	}

	return socket.Socket.Subscribe(symbol...)
}
func (socket *SpotWS_Trade_Socket) Unsubscribe(symbol ...string) (hasTimedOut bool, err error) {
	for i := range symbol {
		symbol[i] = socket.CreateStreamName(symbol[i])
	}
	return socket.Socket.Unsubscribe(symbol...)
}

func (spot_ws *spot_ws) Trade(publicOnMessage func(trade *SpotWS_Trade), symbol ...string) (*SpotWS_Trade_Socket, error) {
	var newSocket SpotWS_Trade_Socket
	for i := range symbol {
		symbol[i] = newSocket.CreateStreamName(symbol[i])
	}

	socket, err := spot_ws.CreateSocket(symbol)
	if err != nil {
		return nil, err
	}

	socket.onMessage = func(messageType int, msg []byte) {
		var trade SpotWS_Trade
		err := json.Unmarshal(msg, &trade)
		if err != nil {
			lib.LocalError(LibraryErrorCodes.PARSE_ERR, err.Error())
			return
		}
		publicOnMessage(&trade)
	}

	newSocket.Socket = socket
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
	Socket *Spot_Websocket
}

func (*SpotWS_Candlestick_Socket) CreateStreamName(symbol string, interval string) string {
	return strings.ToLower(symbol) + "@kline_" + interval
}

func (socket *SpotWS_Candlestick_Socket) Subscribe(identifiers ...SpotWS_Candlestick_StreamIdentifier) (hasTimedOut bool, err error) {
	streams := make([]string, len(identifiers))
	for i, id := range identifiers {
		streams[i] = socket.CreateStreamName(id.Symbol, id.Interval)
	}

	return socket.Socket.Subscribe(streams...)
}
func (socket *SpotWS_Candlestick_Socket) Unsubscribe(identifiers ...SpotWS_Candlestick_StreamIdentifier) (hasTimedOut bool, err error) {
	streams := make([]string, len(identifiers))
	for i, id := range identifiers {
		streams[i] = socket.CreateStreamName(id.Symbol, id.Interval)
	}

	return socket.Socket.Unsubscribe(streams...)
}

func (spot_ws *spot_ws) Candlesticks(publicOnMessage func(candlestick_msg *SpotWS_Candlestick_MSG), identifiers ...SpotWS_Candlestick_StreamIdentifier) (*SpotWS_Candlestick_Socket, error) {
	var newSocket SpotWS_Candlestick_Socket

	streams := make([]string, len(identifiers))
	for i, id := range identifiers {
		streams[i] = newSocket.CreateStreamName(id.Symbol, id.Interval)
	}

	socket, err := spot_ws.CreateSocket(streams)
	if err != nil {
		return nil, err
	}

	socket.onMessage = func(messageType int, msg []byte) {
		var candlestick_msg SpotWS_Candlestick_MSG
		err := json.Unmarshal(msg, &candlestick_msg)
		if err != nil {
			lib.LocalError(LibraryErrorCodes.PARSE_ERR, err.Error())
			return
		}
		publicOnMessage(&candlestick_msg)
	}

	newSocket.Socket = socket
	return &newSocket, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type SpotWS_Candlestick_TimezoneOffset_Socket struct {
	Socket *Spot_Websocket
}

func (*SpotWS_Candlestick_TimezoneOffset_Socket) CreateStreamName(symbol string, interval string) string {
	return strings.ToLower(symbol) + "@kline_" + interval + "@+08:00"
}

func (socket *SpotWS_Candlestick_TimezoneOffset_Socket) Subscribe(identifiers ...SpotWS_Candlestick_StreamIdentifier) (hasTimedOut bool, err error) {
	streams := make([]string, len(identifiers))
	for i, id := range identifiers {
		streams[i] = socket.CreateStreamName(id.Symbol, id.Interval)
	}

	return socket.Socket.Subscribe(streams...)
}
func (socket *SpotWS_Candlestick_TimezoneOffset_Socket) Unsubscribe(identifiers ...SpotWS_Candlestick_StreamIdentifier) (hasTimedOut bool, err error) {
	streams := make([]string, len(identifiers))
	for i, id := range identifiers {
		streams[i] = socket.CreateStreamName(id.Symbol, id.Interval)
	}

	return socket.Socket.Unsubscribe(streams...)
}

func (spot_ws *spot_ws) Candlestick_WithOffset(publicOnMessage func(candlestick_msg *SpotWS_Candlestick_MSG), identifiers ...SpotWS_Candlestick_StreamIdentifier) (*SpotWS_Candlestick_TimezoneOffset_Socket, error) {
	var newSocket SpotWS_Candlestick_TimezoneOffset_Socket

	streams := make([]string, len(identifiers))
	for i, id := range identifiers {
		streams[i] = newSocket.CreateStreamName(id.Symbol, id.Interval)
	}
	socket, err := spot_ws.CreateSocket(streams)
	if err != nil {
		return nil, err
	}

	socket.onMessage = func(messageType int, msg []byte) {
		var candlestick_msg SpotWS_Candlestick_MSG
		err := json.Unmarshal(msg, &candlestick_msg)
		if err != nil {
			lib.LocalError(LibraryErrorCodes.PARSE_ERR, err.Error())
			return
		}
		publicOnMessage(&candlestick_msg)
	}

	newSocket.Socket = socket
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
	Socket *Spot_Websocket
}

func (*SpotWS_MiniTicker_Socket) CreateStreamName(symbol string) string {
	return strings.ToLower(symbol) + "@miniTicker"
}

func (socket *SpotWS_MiniTicker_Socket) Subscribe(symbol ...string) (hasTimedOut bool, err error) {
	for i := range symbol {
		symbol[i] = socket.CreateStreamName(symbol[i])
	}

	return socket.Socket.Subscribe(symbol...)
}
func (socket *SpotWS_MiniTicker_Socket) Unsubscribe(symbol ...string) (hasTimedOut bool, err error) {
	for i := range symbol {
		symbol[i] = socket.CreateStreamName(symbol[i])
	}

	return socket.Socket.Unsubscribe(symbol...)
}

func (spot_ws *spot_ws) MiniTicker(publicOnMessage func(miniTicker *SpotWS_MiniTicker), symbol ...string) (*SpotWS_MiniTicker_Socket, error) {
	var newSocket SpotWS_MiniTicker_Socket
	for i := range symbol {
		symbol[i] = newSocket.CreateStreamName(symbol[i])
	}
	socket, err := spot_ws.CreateSocket(symbol)
	if err != nil {
		return nil, err
	}

	socket.onMessage = func(messageType int, msg []byte) {
		var miniTicker SpotWS_MiniTicker
		err := json.Unmarshal(msg, &miniTicker)
		if err != nil {
			lib.LocalError(LibraryErrorCodes.PARSE_ERR, err.Error())
			return
		}
		publicOnMessage(&miniTicker)
	}

	newSocket.Socket = socket
	return &newSocket, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type SpotWS_AllMiniTickers_Socket struct {
	Socket *Spot_Websocket
}

func (*SpotWS_AllMiniTickers_Socket) CreateStreamName() string {
	return "!miniTicker@arr"
}

func (socket *SpotWS_AllMiniTickers_Socket) Subscribe() (hasTimedOut bool, err error) {
	return socket.Socket.Subscribe(socket.CreateStreamName())
}
func (socket *SpotWS_AllMiniTickers_Socket) Unsubscribe() (hasTimedOut bool, err error) {
	return socket.Socket.Unsubscribe(socket.CreateStreamName())
}

func (spot_ws *spot_ws) AllMiniTickers(publicOnMessage func(miniTickers []*SpotWS_MiniTicker)) (*SpotWS_AllMiniTickers_Socket, error) {
	var newSocket SpotWS_AllMiniTickers_Socket
	socket, err := spot_ws.CreateSocket([]string{newSocket.CreateStreamName()})
	if err != nil {
		return nil, err
	}

	socket.onMessage = func(messageType int, msg []byte) {
		var miniTickers []*SpotWS_MiniTicker
		err := json.Unmarshal(msg, &miniTickers)
		if err != nil {
			lib.LocalError(LibraryErrorCodes.PARSE_ERR, err.Error())
			return
		}
		publicOnMessage(miniTickers)
	}

	newSocket.Socket = socket
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
	Socket *Spot_Websocket
}

func (*SpotWS_Ticker_Socket) CreateStreamName(symbol string) string {
	return strings.ToLower(symbol) + "@ticker"
}

func (socket *SpotWS_Ticker_Socket) Subscribe(symbol ...string) (hasTimedOut bool, err error) {
	for i := range symbol {
		symbol[i] = socket.CreateStreamName(symbol[i])
	}

	return socket.Socket.Subscribe(symbol...)
}
func (socket *SpotWS_Ticker_Socket) Unsubscribe(symbol ...string) (hasTimedOut bool, err error) {
	for i := range symbol {
		symbol[i] = socket.CreateStreamName(symbol[i])
	}

	return socket.Socket.Unsubscribe(symbol...)
}

func (spot_ws *spot_ws) Ticker(publicOnMessage func(ticker *SpotWS_Ticker), symbol ...string) (*SpotWS_Ticker_Socket, error) {
	var newSocket SpotWS_Ticker_Socket
	for i := range symbol {
		symbol[i] = newSocket.CreateStreamName(symbol[i])
	}
	socket, err := spot_ws.CreateSocket(symbol)
	if err != nil {
		return nil, err
	}

	socket.onMessage = func(messageType int, msg []byte) {
		var ticker SpotWS_Ticker
		err := json.Unmarshal(msg, &ticker)
		if err != nil {
			lib.LocalError(LibraryErrorCodes.PARSE_ERR, err.Error())
			return
		}
		publicOnMessage(&ticker)
	}

	newSocket.Socket = socket
	return &newSocket, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type SpotWS_AllTickers_Socket struct {
	Socket *Spot_Websocket
}

func (*SpotWS_AllTickers_Socket) CreateStreamName() string {
	return "!ticker@arr"
}

func (socket *SpotWS_AllTickers_Socket) Subscribe() (hasTimedOut bool, err error) {
	return socket.Socket.Subscribe(socket.CreateStreamName())
}
func (socket *SpotWS_AllTickers_Socket) Unsubscribe() (hasTimedOut bool, err error) {
	return socket.Socket.Unsubscribe(socket.CreateStreamName())
}

func (spot_ws *spot_ws) AllTickers(publicOnMessage func(tickers []*SpotWS_Ticker)) (*SpotWS_AllTickers_Socket, error) {
	var newSocket SpotWS_AllTickers_Socket
	socket, err := spot_ws.CreateSocket([]string{newSocket.CreateStreamName()})
	if err != nil {
		return nil, err
	}

	socket.onMessage = func(messageType int, msg []byte) {
		var tickers []*SpotWS_Ticker
		err := json.Unmarshal(msg, &tickers)
		if err != nil {
			lib.LocalError(LibraryErrorCodes.PARSE_ERR, err.Error())
			return
		}
		publicOnMessage(tickers)
	}

	newSocket.Socket = socket
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
	Socket *Spot_Websocket
}

func (*SpotWS_RollingWindowStatistics_Socket) CreateStreamName(symbol string, windowSize string) string {
	return strings.ToLower(symbol) + "@ticker_" + windowSize
}

func (socket *SpotWS_RollingWindowStatistics_Socket) Subscribe(identifiers ...SpotWS_RollingWindowStatistics_StreamIdentifier) (hasTimedOut bool, err error) {
	streams := make([]string, len(identifiers))
	for i, id := range identifiers {
		streams[i] = socket.CreateStreamName(id.Symbol, id.WindowSize)
	}

	return socket.Socket.Subscribe(streams...)
}
func (socket *SpotWS_RollingWindowStatistics_Socket) Unsubscribe(identifiers ...SpotWS_RollingWindowStatistics_StreamIdentifier) (hasTimedOut bool, err error) {
	streams := make([]string, len(identifiers))
	for i, id := range identifiers {
		streams[i] = socket.CreateStreamName(id.Symbol, id.WindowSize)
	}

	return socket.Socket.Unsubscribe(streams...)
}

func (spot_ws *spot_ws) RollingWindowStatistics(publicOnMessage func(rwStat *SpotWS_RollingWindowStatistic), identifiers ...SpotWS_RollingWindowStatistics_StreamIdentifier) (*SpotWS_RollingWindowStatistics_Socket, error) {
	var newSocket SpotWS_RollingWindowStatistics_Socket

	streams := make([]string, len(identifiers))
	for i, id := range identifiers {
		streams[i] = newSocket.CreateStreamName(id.Symbol, id.WindowSize)
	}

	socket, err := spot_ws.CreateSocket(streams)
	if err != nil {
		return nil, err
	}

	socket.onMessage = func(messageType int, msg []byte) {
		var rwStat SpotWS_RollingWindowStatistic
		err := json.Unmarshal(msg, &rwStat)
		if err != nil {
			lib.LocalError(LibraryErrorCodes.PARSE_ERR, err.Error())
			return
		}
		publicOnMessage(&rwStat)
	}

	newSocket.Socket = socket
	return &newSocket, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type SpotWS_AllRollingWindowStatistics_Socket struct {
	Socket *Spot_Websocket
}

func (*SpotWS_AllRollingWindowStatistics_Socket) CreateStreamName(WindowSize string) string {
	return "!ticker_" + WindowSize + "@arr"
}

func (socket *SpotWS_AllRollingWindowStatistics_Socket) Subscribe(WindowSize ...string) (hasTimedOut bool, err error) {
	for i := range WindowSize {
		WindowSize[i] = socket.CreateStreamName(WindowSize[i])
	}

	return socket.Socket.Subscribe(WindowSize...)
}
func (socket *SpotWS_AllRollingWindowStatistics_Socket) Unsubscribe(WindowSize ...string) (hasTimedOut bool, err error) {
	for i := range WindowSize {
		WindowSize[i] = socket.CreateStreamName(WindowSize[i])
	}

	return socket.Socket.Unsubscribe(WindowSize...)
}

func (spot_ws *spot_ws) AllRollingWindowStatistics(publicOnMessage func(rwStats []*SpotWS_RollingWindowStatistic), WindowSize ...string) (*SpotWS_AllRollingWindowStatistics_Socket, error) {
	var newSocket SpotWS_AllRollingWindowStatistics_Socket

	for i := range WindowSize {
		WindowSize[i] = newSocket.CreateStreamName(WindowSize[i])
	}

	socket, err := spot_ws.CreateSocket(WindowSize)
	if err != nil {
		return nil, err
	}

	socket.onMessage = func(messageType int, msg []byte) {
		var rwStats []*SpotWS_RollingWindowStatistic
		err := json.Unmarshal(msg, &rwStats)
		if err != nil {
			lib.LocalError(LibraryErrorCodes.PARSE_ERR, err.Error())
			return
		}
		publicOnMessage(rwStats)
	}

	newSocket.Socket = socket
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
	Socket *Spot_Websocket
}

func (*SpotWS_BookTicker_Socket) CreateStreamName(symbol string) string {
	return strings.ToLower(symbol) + "@bookTicker"
}

func (socket *SpotWS_BookTicker_Socket) Subscribe(symbol ...string) (hasTimedOut bool, err error) {
	for i := range symbol {
		symbol[i] = socket.CreateStreamName(symbol[i])
	}

	return socket.Socket.Subscribe(symbol...)
}
func (socket *SpotWS_BookTicker_Socket) Unsubscribe(symbol ...string) (hasTimedOut bool, err error) {
	for i := range symbol {
		symbol[i] = socket.CreateStreamName(symbol[i])
	}

	return socket.Socket.Unsubscribe(symbol...)
}

func (spot_ws *spot_ws) BookTicker(publicOnMessage func(bookTicker *SpotWS_BookTicker), symbol ...string) (*SpotWS_BookTicker_Socket, error) {
	var newSocket SpotWS_BookTicker_Socket

	for i := range symbol {
		symbol[i] = newSocket.CreateStreamName(symbol[i])
	}

	socket, err := spot_ws.CreateSocket(symbol)
	if err != nil {
		return nil, err
	}

	socket.onMessage = func(messageType int, msg []byte) {
		var bookTicker *SpotWS_BookTicker
		err := json.Unmarshal(msg, &bookTicker)
		if err != nil {
			lib.LocalError(LibraryErrorCodes.PARSE_ERR, err.Error())
			return
		}
		publicOnMessage(bookTicker)
	}

	newSocket.Socket = socket
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
	Socket *Spot_Websocket
}

func (*SpotWS_AveragePrice_Socket) CreateStreamName(symbol string) string {
	return strings.ToLower(symbol) + "@avgPrice"
}

func (socket *SpotWS_AveragePrice_Socket) Subscribe(symbol ...string) (hasTimedOut bool, err error) {
	for i := range symbol {
		symbol[i] = socket.CreateStreamName(symbol[i])
	}

	return socket.Socket.Subscribe(symbol...)
}
func (socket *SpotWS_AveragePrice_Socket) Unsubscribe(symbol ...string) (hasTimedOut bool, err error) {
	for i := range symbol {
		symbol[i] = socket.CreateStreamName(symbol[i])
	}

	return socket.Socket.Unsubscribe(symbol...)
}

func (spot_ws *spot_ws) AveragePrice(publicOnMessage func(averagePrice *SpotWS_AveragePrice), symbol ...string) (*SpotWS_AveragePrice_Socket, error) {
	var newSocket SpotWS_AveragePrice_Socket

	for i := range symbol {
		symbol[i] = newSocket.CreateStreamName(symbol[i])
	}

	socket, err := spot_ws.CreateSocket(symbol)
	if err != nil {
		return nil, err
	}

	socket.onMessage = func(messageType int, msg []byte) {
		var averagePrice *SpotWS_AveragePrice
		err := json.Unmarshal(msg, &averagePrice)
		if err != nil {
			lib.LocalError(LibraryErrorCodes.PARSE_ERR, err.Error())
			return
		}
		publicOnMessage(averagePrice)
	}

	newSocket.Socket = socket
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
	Socket *Spot_Websocket
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

	return socket.Socket.Subscribe(streams...)
}
func (socket *SpotWS_PartialBookDepth_Socket) Unsubscribe(identifiers ...SpotWS_PartialBookDepth_StreamIdentifier) (hasTimedOut bool, err error) {
	streams := make([]string, len(identifiers))
	for i, id := range identifiers {
		streams[i] = socket.CreateStreamName(id.Symbol, id.Levels, id.IsFast)
	}

	return socket.Socket.Unsubscribe(streams...)
}

func (spot_ws *spot_ws) PartialBookDepth(publicOnMessage func(partialBookDepth *SpotWS_PartialBookDepth), identifiers ...SpotWS_PartialBookDepth_StreamIdentifier) (*SpotWS_PartialBookDepth_Socket, error) {
	var newSocket SpotWS_PartialBookDepth_Socket

	streams := make([]string, len(identifiers))
	for i, id := range identifiers {
		streams[i] = newSocket.CreateStreamName(id.Symbol, id.Levels, id.IsFast)
	}

	socket, err := spot_ws.CreateSocket(streams)
	if err != nil {
		return nil, err
	}

	socket.onMessage = func(messageType int, msg []byte) {
		var partialBookDepth *SpotWS_PartialBookDepth
		err := json.Unmarshal(msg, &partialBookDepth)
		if err != nil {
			lib.LocalError(LibraryErrorCodes.PARSE_ERR, err.Error())
			return
		}
		publicOnMessage(partialBookDepth)
	}

	newSocket.Socket = socket
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
	Socket *Spot_Websocket
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

	return socket.Socket.Subscribe(streams...)
}
func (socket *SpotWS_DiffBookDepth_Socket) Unsubscribe(identifiers ...SpotWS_DiffBookDepth_StreamIdentifier) (hasTimedOut bool, err error) {
	streams := make([]string, len(identifiers))
	for i, id := range identifiers {
		streams[i] = socket.CreateStreamName(id.Symbol, id.IsFast)
	}

	return socket.Socket.Unsubscribe(streams...)
}

func (spot_ws *spot_ws) DiffBookDepth(publicOnMessage func(diffBookDepth *SpotWS_DiffBookDepth), identifiers ...SpotWS_DiffBookDepth_StreamIdentifier) (*SpotWS_DiffBookDepth_Socket, error) {
	var newSocket SpotWS_DiffBookDepth_Socket

	streams := make([]string, len(identifiers))
	for i, id := range identifiers {
		streams[i] = newSocket.CreateStreamName(id.Symbol, id.IsFast)
	}

	socket, err := spot_ws.CreateSocket(streams)
	if err != nil {
		return nil, err
	}

	socket.onMessage = func(messageType int, msg []byte) {
		var diffBookDepth *SpotWS_DiffBookDepth
		err := json.Unmarshal(msg, &diffBookDepth)
		if err != nil {
			lib.LocalError(LibraryErrorCodes.PARSE_ERR, err.Error())
			return
		}
		publicOnMessage(diffBookDepth)
	}

	newSocket.Socket = socket

	return &newSocket, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type SpotWS_UserData_Socket struct {
	base websockets.BinanceUserDataWebsocket

	onMessage   func(messageType int, msg []byte)
	onReconnect func()
	onError     func(error)
	onClose     func()

	OnOutboundAccountPosition func(*SpotWS_OutboundAccountPosition)
	OnBalanceUpdate           func(*SpotWS_BalanceUpdate)
	OnExecutionReport         func(*SpotWS_ExecutionReport)
	OnListStatus              func(*SpotWS_ListStatus)
	// Ignore since it shouldn't ever be called
	OnListenKeyExpired func(*SpotWS_ListenKeyExpired)
	// Ignore since it shouldn't ever be called
	OnEventStreamterminated func(*SpotWS_EventStreamTerminated)
	OnExternalLockUpdate    func(*SpotWS_ExternalLockUpdate)

	// Only called if there is a new event that is not yet supported by the library
	OnUnknownEvent func(event string, msg []byte)

	OnMessage   func(messageType int, msg []byte)
	OnReconnect func()
	OnError     func(error)
	OnClose     func()
}

func (socket *SpotWS_UserData_Socket) onMsg(messageType int, msg []byte) {
	if socket.onMessage != nil {
		socket.onMessage(messageType, msg)
	}
	if socket.OnMessage != nil {
		socket.OnMessage(messageType, msg)
	}
}

func (socket *SpotWS_UserData_Socket) onReconn() {
	if socket.onReconnect != nil {
		socket.onReconnect()
	}
	if socket.OnReconnect != nil {
		socket.OnReconnect()
	}
}

func (socket *SpotWS_UserData_Socket) onErr(err error) {
	if socket.onError != nil {
		socket.onError(err)
	}
	if socket.OnError != nil {
		socket.OnError(err)
	}
}

func (socket *SpotWS_UserData_Socket) onCls() {
	if socket.onClose != nil {
		socket.onClose()
	}
	if socket.OnClose != nil {
		socket.OnClose()
	}
}

//// Public Methods

func (socket *SpotWS_UserData_Socket) Close() {
	socket.base.Close()
}

func (socket *SpotWS_UserData_Socket) RestartUserDataStream() error {
	return socket.base.RestartUserDataStream()
}

//

var spotWS_UserData_Events_Registry = map[string]func() interface{}{
	"outboundAccountPosition": func() interface{} { return &SpotWS_OutboundAccountPosition{} },
	"balanceUpdate":           func() interface{} { return &SpotWS_BalanceUpdate{} },
	"executionReport":         func() interface{} { return &SpotWS_ExecutionReport{} },
	"listStatus":              func() interface{} { return &SpotWS_ListStatus{} },
	"listenKeyExpired":        func() interface{} { return &SpotWS_ListenKeyExpired{} },
	"eventStreamTerminated":   func() interface{} { return &SpotWS_EventStreamTerminated{} },
	"externalLockUpdate":      func() interface{} { return &SpotWS_ExternalLockUpdate{} },
}

func (spot_ws *spot_ws) UserData() (*SpotWS_UserData_Socket, error) {
	var newSocket SpotWS_UserData_Socket

	var startUserData_func = func() (string, error) {
		listenKey, _, err := spot_ws.binance.Spot.StartUserDataStream()
		if err != nil {
			return "", err
		}
		return listenKey, nil
	}

	var keepAliveUserData_func = func(listenKey string) error {
		_, err := spot_ws.binance.Spot.KeepAlive_UserData_ListenKey(listenKey)
		if err != nil {
			return err
		}
		return nil
	}

	socket, err := websockets.CreateHTTPUserDataWebsocket(SPOT_Constants.Websocket.URLs[0], "/ws/", 60, startUserData_func, keepAliveUserData_func)
	if err != nil {
		return nil, err
	}
	newSocket.base = socket

	newSocket.base.SetOnMessage(newSocket.onMsg)
	newSocket.base.SetOnReconnect(newSocket.onReconn)
	newSocket.base.SetOnError(newSocket.onErr)
	newSocket.base.SetOnClose(newSocket.onCls)

	newSocket.onMessage = func(messageType int, msg []byte) {
		var event struct {
			Event string `json:"e"`
		}

		err := json.Unmarshal(msg, &event)
		if err != nil {
			fmt.Println("PRINT ERROR IS FROM HERE")
			fmt.Printf("[LIBRARY] Failed to unmarshal userData stream message: %s\n\tmsg => %s\n", err.Error(), msg)
			return
		}

		eventStr := event.Event
		if eventStr == "" {
			var nested struct {
				Event struct {
					E string `json:"e"`
				} `json:"event"`
			}
			if err := json.Unmarshal(msg, &nested); err != nil {
				fmt.Printf("[LIBRARY] Failed to unmarshal nested event: %s\n\tmsg => %s\n", err, msg)
				return
			}
			eventStr = nested.Event.E
			if eventStr == "" {
				fmt.Printf("[LIBRARY] No event type found in message: %s\n", msg)
				return
			}
		}

		// Look up factory function
		structFactory, ok := spotWS_UserData_Events_Registry[eventStr]
		if !ok {
			if newSocket.OnUnknownEvent != nil {
				newSocket.OnUnknownEvent(eventStr, msg)
			}
			return
		}

		target := structFactory()
		err = json.Unmarshal(msg, target)
		if err != nil {
			fmt.Printf("[LIBRARY] Failed to unmarshal '%s' event: %s\n", eventStr, err)
			return
		}

		// Dispatch by concrete type
		switch evt := target.(type) {
		case *SpotWS_OutboundAccountPosition:
			if newSocket.OnOutboundAccountPosition != nil {
				newSocket.OnOutboundAccountPosition(evt)
			}

		case *SpotWS_BalanceUpdate:
			if newSocket.OnBalanceUpdate != nil {
				newSocket.OnBalanceUpdate(evt)
			}

		case *SpotWS_ExecutionReport:
			if newSocket.OnExecutionReport != nil {
				newSocket.OnExecutionReport(evt)
			}

		case *SpotWS_ListStatus:
			if newSocket.OnListStatus != nil {
				newSocket.OnListStatus(evt)
			}

		case *SpotWS_ListenKeyExpired:
			if newSocket.OnListenKeyExpired != nil {
				newSocket.OnListenKeyExpired(evt)
			}

		case *SpotWS_EventStreamTerminated:
			if newSocket.OnEventStreamterminated != nil {
				newSocket.OnEventStreamterminated(evt)
			}

		case *SpotWS_ExternalLockUpdate:
			if newSocket.OnExternalLockUpdate != nil {
				newSocket.OnExternalLockUpdate(evt)
			}

		default:
			fmt.Printf("[LIBRARY] No handler for event type '%s'\n", eventStr)
		}
	}

	return &newSocket, nil
}

//// Types

type SpotWS_OutboundAccountPosition struct {
	// "outboundAccountPosition"   // Event type
	Event string `json:"e"`

	// 1564034571105               // Event Time
	EventTime int64 `json:"E"`

	// 1564034571073               // Time of last account update
	LastUpdateTime int64 `json:"u"`

	// [ { "a": "ETH", "f": "10000.000000", "l": "0.000000" } ]   // Balances Array
	Balances []SpotWS_OutboundAccountPosition_Balance `json:"B"`
}
type SpotWS_OutboundAccountPosition_Balance struct {
	// "ETH"       // Asset
	Asset string `json:"a"`

	// "10000.000000"   // Free
	Free string `json:"f"`

	// "0.000000"       // Locked
	Locked string `json:"l"`
}

//

type SpotWS_BalanceUpdate struct {
	// "balanceUpdate"           // Event Type
	Event string `json:"e"`

	// 1573200697110             // Event Time
	EventTime int64 `json:"E"`

	// "BTC"                     // Asset
	Asset string `json:"a"`

	// "100.00000000"            // Balance Delta
	Delta string `json:"d"`

	// 1573200697068             // Clear Time
	ClearTime int64 `json:"T"`
}

//

type SpotWS_ExecutionReport struct {
	// "executionReport"          // Event type
	Event string `json:"e"`

	// 1499405658658              // Event time
	EventTime int64 `json:"E"`

	// "ETHBTC"                   // Symbol
	Symbol string `json:"s"`

	// "mUvoqJxFIILMdfAW5iGSOW"   // Client order ID
	ClientOrderID string `json:"c"`

	// "BUY"                      // Side
	Side string `json:"S"`

	// "LIMIT"                    // Order type
	OrderType string `json:"o"`

	// "GTC"                      // Time in force
	TimeInForce string `json:"f"`

	// "1.00000000"               // Order quantity
	Quantity string `json:"q"`

	// "0.10264410"               // Order price
	Price string `json:"p"`

	// "0.00000000"               // Stop price
	StopPrice string `json:"P"`

	// "0.00000000"               // Iceberg quantity
	IcebergQty string `json:"F"`

	// -1                         // OrderListId
	OrderListID int64 `json:"g"`

	// ""                         // Original client order ID
	OrigClientOrderID string `json:"C"`

	// "NEW"                      // Current execution type
	ExecutionType string `json:"x"`

	// "NEW"                      // Current order status
	OrderStatus string `json:"X"`

	// "NONE"                     // Order reject reason
	RejectReason string `json:"r"`

	// 4293153                    // Order ID
	OrderID int64 `json:"i"`

	// "0.00000000"               // Last executed quantity
	LastExecutedQty string `json:"l"`

	// "0.00000000"               // Cumulative filled quantity
	CumulativeFilledQty string `json:"z"`

	// "0.00000000"               // Last executed price
	LastExecutedPrice string `json:"L"`

	// "0"                        // Commission amount
	Commission string `json:"n"`

	// null                       // Commission asset
	CommissionAsset *string `json:"N"`

	// 1499405658657              // Transaction time
	TransactionTime int64 `json:"T"`

	// -1                         // Trade ID
	TradeID int64 `json:"t"`

	// 3                          // Prevented Match Id
	PreventedMatchID int64 `json:"v"`

	// 8641984                    // Execution Id
	ExecutionID int64 `json:"I"`

	// true                       // Is the order on the book?
	IsOnBook bool `json:"w"`

	// false                      // Is this trade the maker side?
	IsMaker bool `json:"m"`

	// false                      // Ignore
	Ignore bool `json:"M"`

	// 1499405658657              // Order creation time
	OrderCreationTime int64 `json:"O"`

	// "0.00000000"               // Cumulative quote asset transacted quantity
	CumulativeQuoteQty string `json:"Z"`

	// "0.00000000"               // Last quote asset transacted quantity
	LastQuoteQty string `json:"Y"`

	// "0.00000000"               // Quote Order Quantity
	QuoteOrderQty string `json:"Q"`

	// 1499405658657              // Working Time
	WorkingTime int64 `json:"W"`

	// "NONE"                     // SelfTradePreventionMode
	SelfTradePreventionMode string `json:"V"`
}

//

type SpotWS_ListStatus struct {
	// "listStatus"                // Event Type
	Event string `json:"e"`

	// 1564035303637              // Event Time
	EventTime int64 `json:"E"`

	// "ETHBTC"                   // Symbol
	Symbol string `json:"s"`

	// 2                          // OrderListId
	OrderListID int64 `json:"g"`

	// "OCO"                      // Contingency Type
	ContingencyType string `json:"c"`

	// "EXEC_STARTED"             // List Status Type
	ListStatusType string `json:"l"`

	// "EXECUTING"                // List Order Status
	ListOrderStatus string `json:"L"`

	// "NONE"                     // List Reject Reason
	ListRejectReason string `json:"r"`

	// "F4QN4G8DlFATFlIUQ0cjdD"   // List Client Order ID
	ListClientOrderID string `json:"C"`

	// 1564035303625              // Transaction Time
	TransactionTime int64 `json:"T"`

	// [ { ... }, { ... } ]       // An array of objects
	Orders []SpotWS_ListStatus_Order `json:"O"`
}

type SpotWS_ListStatus_Order struct {
	// "ETHBTC"                   // Symbol
	Symbol string `json:"s"`

	// 17                         // OrderId
	OrderID int64 `json:"i"`

	// "AJYsMjErWJesZvqlJCTUgL"   // ClientOrderId
	ClientOrderID string `json:"c"`
}

//

type SpotWS_ListenKeyExpired struct {
	// "listenKeyExpired"      // Event type
	Event string `json:"e"`

	// 1699596037418           // Event time
	EventTime int64 `json:"E"`

	// "OfYGbUzi3PraNagEkdKuFwUHn48brFsItTdsuiIXrucEvD0rhRXZ7I6URWfE8YE8"
	ListenKey string `json:"listenKey"`
}

//

type SpotWS_EventStreamTerminated struct {
	// { "e": "eventStreamTerminated", "E": 1728973001334 }
	Event SpotWS_EventStreamTerminated_Event `json:"event"`
}

type SpotWS_EventStreamTerminated_Event struct {
	// "eventStreamTerminated"   // Event Type
	Event string `json:"e"`

	// 1728973001334              // Event Time
	EventTime int64 `json:"E"`
}

//

type SpotWS_ExternalLockUpdate struct {
	// "externalLockUpdate"     // Event Type
	Event string `json:"e"`

	// 1581557507324            // Event Time
	EventTime int64 `json:"E"`

	// "NEO"                    // Asset
	Asset string `json:"a"`

	// "10.00000000"            // Delta
	Delta string `json:"d"`

	// 1581557507268            // Transaction Time
	TransactionTime int64 `json:"T"`
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (*spot_ws) CreateSocket(streams []string) (*Spot_Websocket, error) {
	baseURL := SPOT_Constants.Websocket.URLs[0]

	socket, err := websockets.CreateBinanceWebsocket(baseURL, streams)
	if err != nil {
		return nil, err
	}

	ws := &Spot_Websocket{
		base: socket,
	}

	ws.onMessage = ws.onMsg
	ws.onReconnect = ws.onReconn

	return ws, nil
}
