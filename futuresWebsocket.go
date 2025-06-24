package Binance

import (
	"fmt"
	"strings"
	"sync"

	"slices"

	"github.com/GTedZ/binancego/lib"
	websockets "github.com/GTedZ/binancego/websockets"
)

type futures_ws struct {
	binance *Binance
}

func (futures_ws *futures_ws) init(binance *Binance) {
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

	futures_ws.base.AddStreams(stream)
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

	futures_ws.base.RemoveStreams(stream)
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

func (futures_ws *futures_ws) AggTrade(publicOnMessage func(aggTrade *FuturesWS_AggTrade), symbol ...string) *FuturesWS_AggTrade_Socket {
	var newSocket FuturesWS_AggTrade_Socket

	streamNames := newSocket.CreateStreamName(symbol...)

	socket := futures_ws.CreateSocket(streamNames, false)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var aggTrade FuturesWS_AggTrade
		err := json.Unmarshal(msg, &aggTrade)
		if err != nil {
			lib.LocalError(Errors.LibraryCodes.PARSE_ERR, err.Error())
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

func (futures_ws *futures_ws) MarkPrice(publicOnMessage func(markPrice *FuturesWS_MarkPrice), params ...FuturesWS_MarkPrice_Params) *FuturesWS_MarkPrice_Socket {
	var newSocket FuturesWS_MarkPrice_Socket

	streams := newSocket.CreateStreamName(params...)

	socket := futures_ws.CreateSocket(streams, false)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var markPrice FuturesWS_MarkPrice
		err := json.Unmarshal(msg, &markPrice)
		if err != nil {
			lib.LocalError(Errors.LibraryCodes.PARSE_ERR, err.Error())
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

func (futures_ws *futures_ws) AllMarkPrices(publicOnMessage func(markPrices []*FuturesWS_MarkPrice), isFast ...bool) *FuturesWS_AllMarkPrices_Socket {
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
			lib.LocalError(Errors.LibraryCodes.PARSE_ERR, err.Error())
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

func (kline *FuturesWS_Candlestick_Kline) ParseFloat() (*FuturesWS_Candlestick_Kline_f64, error) {
	open, err := Utils.ParseFloat(kline.Open)
	if err != nil {
		errStr := fmt.Sprintf("error parsing float '%s' of 'kline.Open': %s", kline.Open, err.Error())
		Logger.ERROR(errStr)
		return nil, err
	}
	high, err := Utils.ParseFloat(kline.High)
	if err != nil {
		errStr := fmt.Sprintf("error parsing float '%s' of 'kline.High': %s", kline.High, err.Error())
		Logger.ERROR(errStr)
		return nil, err
	}
	low, err := Utils.ParseFloat(kline.Low)
	if err != nil {
		errStr := fmt.Sprintf("error parsing float '%s' of 'kline.Low': %s", kline.Low, err.Error())
		Logger.ERROR(errStr)
		return nil, err
	}
	close, err := Utils.ParseFloat(kline.Close)
	if err != nil {
		errStr := fmt.Sprintf("error parsing float '%s' of 'kline.Close': %s", kline.Close, err.Error())
		Logger.ERROR(errStr)
		return nil, err
	}

	baseAssetVolume, err := Utils.ParseFloat(kline.BaseAssetVolume)
	if err != nil {
		errStr := fmt.Sprintf("error parsing float '%s' of 'kline.Volume': %s", kline.BaseAssetVolume, err.Error())
		Logger.ERROR(errStr)
		return nil, err
	}
	quoteAssetVolume, err := Utils.ParseFloat(kline.QuoteAssetVolume)
	if err != nil {
		errStr := fmt.Sprintf("error parsing float '%s' of 'kline.QuoteAssetVolume': %s", kline.QuoteAssetVolume, err.Error())
		Logger.ERROR(errStr)
		return nil, err
	}

	takerBuyBaseAssetVolume, err := Utils.ParseFloat(kline.TakerBuyBaseAssetVolume)
	if err != nil {
		errStr := fmt.Sprintf("error parsing float '%s' of 'kline.TakerBuyBaseAssetVolume': %s", kline.TakerBuyBaseAssetVolume, err.Error())
		Logger.ERROR(errStr)
		return nil, err
	}
	takerBuyQuoteAssetVolume, err := Utils.ParseFloat(kline.TakerBuyQuoteAssetVolume)
	if err != nil {
		errStr := fmt.Sprintf("error parsing float '%s' of 'kline.TakerBuyQuoteAssetVolume': %s", kline.TakerBuyQuoteAssetVolume, err.Error())
		Logger.ERROR(errStr)
		return nil, err
	}

	return &FuturesWS_Candlestick_Kline_f64{
		OpenTime:  kline.OpenTime,
		CloseTime: kline.CloseTime,

		Open:  open,
		High:  high,
		Low:   low,
		Close: close,

		BaseAssetVolume:          baseAssetVolume,
		QuoteAssetVolume:         quoteAssetVolume,
		TakerBuyBaseAssetVolume:  takerBuyBaseAssetVolume,
		TakerBuyQuoteAssetVolume: takerBuyQuoteAssetVolume,
		TradeCount:               kline.TradeCount,
	}, nil
}

type FuturesWS_Candlestick_Kline_f64 struct {

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
	Open float64 `json:"o"`

	// Close price
	Close float64 `json:"c"`

	// High price
	High float64 `json:"h"`

	// Low price
	Low float64 `json:"l"`

	// Number of trades
	TradeCount int64 `json:"n"`

	// Base asset volume
	BaseAssetVolume float64 `json:"v"`

	// Quote asset volume
	QuoteAssetVolume float64 `json:"q"`

	// Taker buy base asset volume
	TakerBuyBaseAssetVolume float64 `json:"V"`

	// Taker buy quote asset volume
	TakerBuyQuoteAssetVolume float64 `json:"Q"`

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

func (futures_ws *futures_ws) Candlesticks(publicOnMessage func(candlestick *FuturesWS_Candlestick), params ...FuturesWS_Candlestick_Params) *FuturesWS_Candlesticks_Socket {
	var newSocket FuturesWS_Candlesticks_Socket

	streamNames := newSocket.CreateStreamName(params...)

	socket := futures_ws.CreateSocket(streamNames, false)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var kline *FuturesWS_Candlestick
		err := json.Unmarshal(msg, &kline)
		if err != nil {
			lib.LocalError(Errors.LibraryCodes.PARSE_ERR, err.Error())
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
func (futures_ws *futures_ws) ContinuousCandlesticks(publicOnMessage func(candlestick *FuturesWS_ContinuousCandlestick), params ...FuturesWS_ContinuousCandlestick_Params) *FuturesWS_ContinuousCandlestick_Socket {
	var newSocket FuturesWS_ContinuousCandlestick_Socket

	streamNames := newSocket.CreateStreamName(params...)

	socket := futures_ws.CreateSocket(streamNames, false)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var kline *FuturesWS_ContinuousCandlestick
		err := json.Unmarshal(msg, &kline)
		if err != nil {
			lib.LocalError(Errors.LibraryCodes.PARSE_ERR, err.Error())
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

func (futures_ws *futures_ws) MiniTicker(publicOnMessage func(miniTicker *FuturesWS_MiniTicker), symbol ...string) *FuturesWS_MiniTicker_Socket {
	var newSocket FuturesWS_MiniTicker_Socket

	streamNames := newSocket.CreateStreamName(symbol...)

	socket := futures_ws.CreateSocket(streamNames, false)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var miniTicker *FuturesWS_MiniTicker
		err := json.Unmarshal(msg, &miniTicker)
		if err != nil {
			lib.LocalError(Errors.LibraryCodes.PARSE_ERR, err.Error())
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

func (futures_ws *futures_ws) AllMiniTickers(publicOnMessage func(miniTickers []*FuturesWS_MiniTicker)) *FuturesWS_AllMiniTickers_Socket {
	var newSocket FuturesWS_AllMiniTickers_Socket

	streamName := newSocket.CreateStreamName()

	socket := futures_ws.CreateSocket([]string{streamName}, false)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var miniTickers []*FuturesWS_MiniTicker
		err := json.Unmarshal(msg, &miniTickers)
		if err != nil {
			lib.LocalError(Errors.LibraryCodes.PARSE_ERR, err.Error())
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

func (futures_ws *futures_ws) Ticker(publicOnMessage func(ticker *FuturesWS_Ticker), symbol ...string) *FuturesWS_Ticker_Socket {
	var newSocket FuturesWS_Ticker_Socket

	streamNames := newSocket.CreateStreamName(symbol...)

	socket := futures_ws.CreateSocket(streamNames, false)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var ticker *FuturesWS_Ticker
		err := json.Unmarshal(msg, &ticker)
		if err != nil {
			lib.LocalError(Errors.LibraryCodes.PARSE_ERR, err.Error())
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

func (futures_ws *futures_ws) AllTickers(publicOnMessage func(tickers []*FuturesWS_Ticker)) *FuturesWS_AllTickers_Socket {
	var newSocket FuturesWS_AllTickers_Socket

	streamName := newSocket.CreateStreamName()

	socket := futures_ws.CreateSocket([]string{streamName}, false)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var tickers []*FuturesWS_Ticker
		err := json.Unmarshal(msg, &tickers)
		if err != nil {
			lib.LocalError(Errors.LibraryCodes.PARSE_ERR, err.Error())
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

func (futures_ws *futures_ws) BookTicker(publicOnMessage func(bookTicker *FuturesWS_BookTicker), symbol ...string) *FuturesWS_BookTicker_Socket {
	var newSocket FuturesWS_BookTicker_Socket

	symbol = newSocket.CreateStreamName(symbol...)
	socket := futures_ws.CreateSocket(symbol, false)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var bookTicker FuturesWS_BookTicker
		err := json.Unmarshal(msg, &bookTicker)
		if err != nil {
			lib.LocalError(Errors.LibraryCodes.PARSE_ERR, err.Error())
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

func (futures_ws *futures_ws) AllBookTickers(publicOnMessage func(bookTickers []*FuturesWS_BookTicker)) *FuturesWS_AllBookTickers_Socket {
	var newSocket FuturesWS_AllBookTickers_Socket

	streamName := newSocket.CreateStreamName()

	socket := futures_ws.CreateSocket([]string{streamName}, false)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var bookTickers []*FuturesWS_BookTicker
		err := json.Unmarshal(msg, &bookTickers)
		if err != nil {
			lib.LocalError(Errors.LibraryCodes.PARSE_ERR, err.Error())
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

func (futures_ws *futures_ws) LiquidationOrders(publicOnMessage func(liquidationOrder *FuturesWS_LiquidationOrder), symbol ...string) *FuturesWS_LiquidationOrder_Socket {
	var newSocket FuturesWS_LiquidationOrder_Socket

	streamNames := newSocket.CreateStreamName(symbol...)

	socket := futures_ws.CreateSocket(streamNames, false)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var liquidationOrder *FuturesWS_LiquidationOrder
		err := json.Unmarshal(msg, &liquidationOrder)
		if err != nil {
			lib.LocalError(Errors.LibraryCodes.PARSE_ERR, err.Error())
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

func (futures_ws *futures_ws) AllLiquidationOrders(publicOnMessage func(liquidationOrder *FuturesWS_LiquidationOrder)) *FuturesWS_AllLiquidationOrders_Socket {
	var newSocket FuturesWS_AllLiquidationOrders_Socket

	streamName := newSocket.CreateStreamName()

	socket := futures_ws.CreateSocket([]string{streamName}, false)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var liquidationOrder *FuturesWS_LiquidationOrder
		err := json.Unmarshal(msg, &liquidationOrder)
		if err != nil {
			lib.LocalError(Errors.LibraryCodes.PARSE_ERR, err.Error())
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

func (futures_ws *futures_ws) PartialBookDepth(publicOnMessage func(partialBookDepth *FuturesWS_PartialBookDepth), params ...FuturesWS_PartialBookDepth_Params) *FuturesWS_PartialBookDepth_Socket {
	var newSocket FuturesWS_PartialBookDepth_Socket

	streamNames := newSocket.CreateStreamName(params...)

	socket := futures_ws.CreateSocket(streamNames, false)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var partialBookDepth *FuturesWS_PartialBookDepth
		err := json.Unmarshal(msg, &partialBookDepth)
		if err != nil {
			lib.LocalError(Errors.LibraryCodes.PARSE_ERR, err.Error())
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

func (futures_ws *futures_ws) DiffBookDepth(publicOnMessage func(diffBookDepth *FuturesWS_DiffBookDepth), params ...FuturesWS_DiffBookDepth_Params) *FuturesWS_DiffBookDepth_Socket {
	var newSocket FuturesWS_DiffBookDepth_Socket

	streamNames := newSocket.CreateStreamName(params...)

	socket := futures_ws.CreateSocket(streamNames, false)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var diffBookDepth *FuturesWS_DiffBookDepth
		err := json.Unmarshal(msg, &diffBookDepth)
		if err != nil {
			lib.LocalError(Errors.LibraryCodes.PARSE_ERR, err.Error())
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

func (futures_ws *futures_ws) CompositeIndexSymbolInfo(publicOnMessage func(compositeIndexSymbolInfo *FuturesWS_CompositeIndexSymbolInfo), symbol ...string) *FuturesWS_CompositeIndexSymbolInfo_Socket {
	var newSocket FuturesWS_CompositeIndexSymbolInfo_Socket

	streamNames := newSocket.CreateStreamName(symbol...)

	socket := futures_ws.CreateSocket(streamNames, false)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var compositeIndexSymbolInfo *FuturesWS_CompositeIndexSymbolInfo
		err := json.Unmarshal(msg, &compositeIndexSymbolInfo)
		if err != nil {
			lib.LocalError(Errors.LibraryCodes.PARSE_ERR, err.Error())
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

func (futures_ws *futures_ws) ContractInfo(publicOnMessage func(contractInfo *FuturesWS_ContractInfo)) *FuturesWS_ContractInfo_Socket {
	var newSocket FuturesWS_ContractInfo_Socket
	streamName := newSocket.CreateStreamName()
	socket := futures_ws.CreateSocket([]string{streamName}, false)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var aggTrade FuturesWS_ContractInfo
		err := json.Unmarshal(msg, &aggTrade)
		if err != nil {
			lib.LocalError(Errors.LibraryCodes.PARSE_ERR, err.Error())
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

func (futures_ws *futures_ws) MultiAssetsModeAssetIndex(publicOnMessage func(assetIndexes []*FuturesWS_MultiAssetsModeAssetIndex), assetSymbol ...string) *FuturesWS_MultiAssetsModeAssetIndex_Socket {
	var newSocket FuturesWS_MultiAssetsModeAssetIndex_Socket

	streamNames := newSocket.CreateStreamName(assetSymbol...)

	socket := futures_ws.CreateSocket(streamNames, false)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var assetIndexes []*FuturesWS_MultiAssetsModeAssetIndex
		err := json.Unmarshal(msg, &assetIndexes)
		if err != nil {
			lib.LocalError(Errors.LibraryCodes.PARSE_ERR, err.Error())
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

func (futures_ws *futures_ws) AllMultiAssetsModeAssetIndexes(publicOnMessage func(assetIndexes []*FuturesWS_MultiAssetsModeAssetIndex), assetSymbol ...string) *FuturesWS_AllMultiAssetsModeAssetIndexes_Socket {
	var newSocket FuturesWS_AllMultiAssetsModeAssetIndexes_Socket

	streamName := newSocket.CreateStreamName()

	socket := futures_ws.CreateSocket([]string{streamName}, false)

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var assetIndexes []*FuturesWS_MultiAssetsModeAssetIndex
		err := json.Unmarshal(msg, &assetIndexes)
		if err != nil {
			lib.LocalError(Errors.LibraryCodes.PARSE_ERR, err.Error())
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

type FuturesWS_ManagedOrderbook struct {
	Symbol       string
	LastUpdateId int64
	Bids         [][2]float64
	Asks         [][2]float64

	isReadyToUpdate bool
	isFetching      bool

	previousEvent *FuturesWS_DiffBookDepth
}

func (managedOrderBook *FuturesWS_ManagedOrderbook) addEvent(event *FuturesWS_DiffBookDepth) {
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
			Orderbook      *FuturesWS_ManagedOrderbook
		}
	}
}

func (handler *FuturesWS_ManagedOrderBook_Handler) handleWSMessage(futures_ws *futures_ws, diffBookDepth *FuturesWS_DiffBookDepth) (shouldPushEvent bool, managedOrderBook *FuturesWS_ManagedOrderbook) {
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
			Orderbook      *FuturesWS_ManagedOrderbook
		}{
			bufferedEvents: make([]*FuturesWS_DiffBookDepth, 0),
			Orderbook: &FuturesWS_ManagedOrderbook{
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

func (futures_ws *futures_ws) Managed_OrderBook(publicOnMessage func(orderBook *FuturesWS_ManagedOrderbook), params ...FuturesWS_DiffBookDepth_Params) *FuturesWS_ManagedOrderBook_Handler {
	handler := &FuturesWS_ManagedOrderBook_Handler{}

	var newSocket FuturesWS_DiffBookDepth_Socket
	streamNames := newSocket.CreateStreamName(params...)
	socket := futures_ws.CreateSocket(streamNames, false)
	newSocket.Handler = socket

	socket.base.OnMessage = func(messageType int, msg []byte) {
		var diffBookDepth *FuturesWS_DiffBookDepth
		err := json.Unmarshal(msg, &diffBookDepth)
		if err != nil {
			lib.LocalError(Errors.LibraryCodes.PARSE_ERR, err.Error())
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
		Orderbook      *FuturesWS_ManagedOrderbook
	})
	handler.openNewSymbols(params...)

	return handler
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//

////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (*futures_ws) CreateSocket(streams []string, isCombined bool) *Futures_Websocket {
	baseURL := FUTURES_Constants.Websocket.URLs[0]

	socket := websockets.CreateBinanceWebsocket(baseURL, streams)

	ws := &Futures_Websocket{
		base: socket,
	}

	return ws
}
