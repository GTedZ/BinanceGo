package futures

import (
	"strconv"
	"strings"
)

type wsEndpoints struct {
	client *Client
}

func newWsEndpoints(c *Client) *wsEndpoints {
	return &wsEndpoints{
		client: c,
	}
}

////
// Aggregate Trade
////

type WsAggTrade struct {
	Event      EventType `json:"e"`
	EventTime  int64     `json:"E"`
	Symbol     string    `json:"s"`
	AggTradeID int64     `json:"a"`
	Price      float64   `json:"p,string"`
	Quantity   float64   `json:"q,string"`
	// Normal quantity without the trades involving RPI orders
	NormalQuantity float64 `json:"nq,string"`
	FirstTradeID   int64   `json:"f"`
	LastTradeID    int64   `json:"l"`
	TradeTime      int64   `json:"T"`
	IsMaker        bool    `json:"m"`
	// Symbol type: 1 = USDⓈ-M, 2 = COIN-M
	SymbolType int `json:"st"`
}

type AggTradeWebsocket struct {
	*websocket[*WsAggTrade]
}

func (atws *AggTradeWebsocket) buildStreamNames(symbols ...string) []string {
	streams := make([]string, len(symbols))

	for i, s := range symbols {
		streams[i] = strings.ToLower(s) + "@aggTrade"
	}

	return streams
}

func (atws *AggTradeWebsocket) Subscribe(symbols ...string) Error {
	streams := atws.buildStreamNames(symbols...)
	return atws.subscribe(streams)
}

func (atws *AggTradeWebsocket) Unsubscribe(symbols ...string) Error {
	streams := atws.buildStreamNames(symbols...)
	return atws.unsubscribe(streams)
}

func (ws *wsEndpoints) AggTrade(onMessage func(*WsAggTrade), symbols ...string) (*AggTradeWebsocket, Error) {
	var s AggTradeWebsocket
	streams := s.buildStreamNames(symbols...)

	base, err := newWebsocket(ws.client.wssBaseUrl, WebsocketRouteMarket, streams, onMessage, ws.client.logger)
	if err != nil {
		return nil, err
	}

	s.websocket = base

	return &s, nil
}

////
// Mark Price
////

type WsMarkPrice struct {
	Event     EventType `json:"e"`
	EventTime int64     `json:"E"`
	Symbol    string    `json:"s"`
	MarkPrice float64   `json:"p,string"`
	// TODO: verify. The docs list "ap" as the mark price moving average, but
	// this field is not present on all historical payload versions. It is left
	// here so that callers can access it when present; it defaults to 0 when absent.
	MarkPriceMovingAvg   float64 `json:"ap,string"`
	IndexPrice           float64 `json:"i,string"`
	EstimatedSettlePrice float64 `json:"P,string"`
	FundingRate          float64 `json:"r,string"`
	NextFundingTime      int64   `json:"T"`
	// Symbol type: 1 = UM, 2 = CM
	SymbolType int `json:"st"`
}

// WsMarkPriceInterval controls the update frequency of the mark price stream.
type WsMarkPriceInterval string

const (
	// 1000ms
	WsMarkPriceFast WsMarkPriceInterval = "1s"
	// 3000ms (default)
	WsMarkPriceSlow WsMarkPriceInterval = ""
)

type WsMarkPriceParams struct {
	Symbol string
	Speed  WsMarkPriceInterval
}

type MarkPriceWebsocket struct {
	*websocket[*WsMarkPrice]
}

func (mpws *MarkPriceWebsocket) buildStreamNames(params ...WsMarkPriceParams) []string {
	streams := make([]string, len(params))

	for i, p := range params {
		stream := strings.ToLower(p.Symbol) + "@markPrice"

		if p.Speed != "" {
			stream += "@" + string(p.Speed)
		}

		streams[i] = stream
	}

	return streams
}

func (mpws *MarkPriceWebsocket) Subscribe(params ...WsMarkPriceParams) Error {
	streams := mpws.buildStreamNames(params...)
	return mpws.subscribe(streams)
}

func (mpws *MarkPriceWebsocket) Unsubscribe(params ...WsMarkPriceParams) Error {
	streams := mpws.buildStreamNames(params...)
	return mpws.unsubscribe(streams)
}

func (ws *wsEndpoints) MarkPrice(onMessage func(*WsMarkPrice), params ...WsMarkPriceParams) (*MarkPriceWebsocket, Error) {
	var s MarkPriceWebsocket
	streams := s.buildStreamNames(params...)

	base, err := newWebsocket(ws.client.wssBaseUrl, WebsocketRouteMarket, streams, onMessage, ws.client.logger)
	if err != nil {
		return nil, err
	}

	s.websocket = base

	return &s, nil
}

////
// All Market Mark Price
////

type AllMarkPriceWebsocket struct {
	*websocket[[]*WsMarkPrice]
}

func (ws *wsEndpoints) AllMarkPrice(onMessage func([]*WsMarkPrice), speed ...WsMarkPriceInterval) (*AllMarkPriceWebsocket, Error) {
	stream := "!markPrice@arr"
	if len(speed) > 0 && speed[0] != "" {
		stream += "@" + string(speed[0])
	}

	var s AllMarkPriceWebsocket

	base, err := newWebsocket(ws.client.wssBaseUrl, WebsocketRouteMarket, []string{stream}, onMessage, ws.client.logger)
	if err != nil {
		return nil, err
	}

	s.websocket = base

	return &s, nil
}

////
// Candlesticks
////

type WsKline struct {
	Event     EventType     `json:"e"`
	EventTime int64         `json:"E"`
	Symbol    string        `json:"s"`
	Kline     WsCandlestick `json:"k"`
}

type WsCandlestick struct {
	OpenTime            int64    `json:"t"`
	CloseTime           int64    `json:"T"`
	Symbol              string   `json:"s"`
	Interval            Interval `json:"i"`
	FirstTradeID        int64    `json:"f"`
	LastTradeID         int64    `json:"L"`
	Open                float64  `json:"o,string"`
	Close               float64  `json:"c,string"`
	High                float64  `json:"h,string"`
	Low                 float64  `json:"l,string"`
	Volume              float64  `json:"v,string"`
	TradeCount          int64    `json:"n"`
	IsClosed            bool     `json:"x"`
	QuoteVolume         float64  `json:"q,string"`
	TakerBuyBaseVolume  float64  `json:"V,string"`
	TakerBuyQuoteVolume float64  `json:"Q,string"`
	// ignore
	Unused string `json:"B"`
}

type WsKlineParams struct {
	Symbol   string
	Interval Interval
}

type CandlesticksWebsocket struct {
	*websocket[*WsKline]
}

func (cws *CandlesticksWebsocket) buildStreamNames(params ...WsKlineParams) []string {
	streams := make([]string, len(params))

	for i, p := range params {
		streams[i] = strings.ToLower(p.Symbol) + "@kline_" + string(p.Interval)
	}

	return streams
}

func (cws *CandlesticksWebsocket) Subscribe(params ...WsKlineParams) Error {
	streams := cws.buildStreamNames(params...)
	return cws.subscribe(streams)
}

func (cws *CandlesticksWebsocket) Unsubscribe(params ...WsKlineParams) Error {
	streams := cws.buildStreamNames(params...)
	return cws.unsubscribe(streams)
}

func (ws *wsEndpoints) Candlesticks(onMessage func(*WsKline), params ...WsKlineParams) (*CandlesticksWebsocket, Error) {
	var s CandlesticksWebsocket
	streams := s.buildStreamNames(params...)

	base, err := newWebsocket(ws.client.wssBaseUrl, WebsocketRouteMarket, streams, onMessage, ws.client.logger)
	if err != nil {
		return nil, err
	}

	s.websocket = base

	return &s, nil
}

////
// Continuous Contract Candlesticks
////

type WsContinuousKline struct {
	Event        EventType     `json:"e"`
	EventTime    int64         `json:"E"`
	Pair         string        `json:"ps"`
	ContractType ContractType  `json:"ct"`
	Kline        WsCandlestick `json:"k"`
}

type WsContinuousKlineParams struct {
	Pair         string
	ContractType ContractType
	Interval     Interval
}

type ContinuousCandlesticksWebsocket struct {
	*websocket[*WsContinuousKline]
}

func (ccws *ContinuousCandlesticksWebsocket) buildStreamNames(params ...WsContinuousKlineParams) []string {
	streams := make([]string, len(params))

	for i, p := range params {
		streams[i] = strings.ToLower(p.Pair) + "_" + strings.ToLower(string(p.ContractType)) + "@continuousKline_" + string(p.Interval)
	}

	return streams
}

func (ccws *ContinuousCandlesticksWebsocket) Subscribe(params ...WsContinuousKlineParams) Error {
	streams := ccws.buildStreamNames(params...)
	return ccws.subscribe(streams)
}

func (ccws *ContinuousCandlesticksWebsocket) Unsubscribe(params ...WsContinuousKlineParams) Error {
	streams := ccws.buildStreamNames(params...)
	return ccws.unsubscribe(streams)
}

func (ws *wsEndpoints) ContinuousCandlesticks(onMessage func(*WsContinuousKline), params ...WsContinuousKlineParams) (*ContinuousCandlesticksWebsocket, Error) {
	var s ContinuousCandlesticksWebsocket
	streams := s.buildStreamNames(params...)

	base, err := newWebsocket(ws.client.wssBaseUrl, WebsocketRouteMarket, streams, onMessage, ws.client.logger)
	if err != nil {
		return nil, err
	}

	s.websocket = base

	return &s, nil
}

////
// Mini Ticker
////

type WsMiniTicker struct {
	Event       EventType `json:"e"`
	EventTime   int64     `json:"E"`
	Symbol      string    `json:"s"`
	Close       float64   `json:"c,string"`
	Open        float64   `json:"o,string"`
	High        float64   `json:"h,string"`
	Low         float64   `json:"l,string"`
	Volume      float64   `json:"v,string"`
	QuoteVolume float64   `json:"q,string"`
	Pair        string    `json:"ps"`
	// Symbol type: 1 = UM, 2 = CM
	SymbolType int `json:"st"`
}

type MiniTickerWebsocket struct {
	*websocket[*WsMiniTicker]
}

func (mtws *MiniTickerWebsocket) buildStreamNames(symbols ...string) []string {
	streams := make([]string, len(symbols))

	for i, s := range symbols {
		streams[i] = strings.ToLower(s) + "@miniTicker"
	}

	return streams
}

func (mtws *MiniTickerWebsocket) Subscribe(symbols ...string) Error {
	streams := mtws.buildStreamNames(symbols...)
	return mtws.subscribe(streams)
}

func (mtws *MiniTickerWebsocket) Unsubscribe(symbols ...string) Error {
	streams := mtws.buildStreamNames(symbols...)
	return mtws.unsubscribe(streams)
}

func (ws *wsEndpoints) MiniTicker(onMessage func(*WsMiniTicker), symbols ...string) (*MiniTickerWebsocket, Error) {
	var s MiniTickerWebsocket
	streams := s.buildStreamNames(symbols...)

	base, err := newWebsocket(ws.client.wssBaseUrl, WebsocketRouteMarket, streams, onMessage, ws.client.logger)
	if err != nil {
		return nil, err
	}

	s.websocket = base

	return &s, nil
}

////
// All Market Mini Tickers
////

type AllMiniTickersWebsocket struct {
	*websocket[[]*WsMiniTicker]
}

func (ws *wsEndpoints) AllMiniTickers(onMessage func([]*WsMiniTicker)) (*AllMiniTickersWebsocket, Error) {
	var s AllMiniTickersWebsocket

	base, err := newWebsocket(ws.client.wssBaseUrl, WebsocketRouteMarket, []string{"!miniTicker@arr"}, onMessage, ws.client.logger)
	if err != nil {
		return nil, err
	}

	s.websocket = base

	return &s, nil
}

////
// Ticker
////

type WsTicker struct {
	Event              EventType `json:"e"`
	EventTime          int64     `json:"E"`
	Symbol             string    `json:"s"`
	PriceChange        float64   `json:"p,string"`
	PriceChangePercent float64   `json:"P,string"`
	WeightedAvgPrice   float64   `json:"w,string"`
	LastPrice          float64   `json:"c,string"`
	LastQuantity       float64   `json:"Q,string"`
	Open               float64   `json:"o,string"`
	High               float64   `json:"h,string"`
	Low                float64   `json:"l,string"`
	Volume             float64   `json:"v,string"`
	QuoteVolume        float64   `json:"q,string"`
	OpenTime           int64     `json:"O"`
	CloseTime          int64     `json:"C"`
	FirstTradeID       int64     `json:"F"`
	LastTradeID        int64     `json:"L"`
	TradeCount         int64     `json:"n"`
	Pair               string    `json:"ps"`
	// Symbol type: 1 = UM, 2 = CM
	SymbolType int `json:"st"`
}

type TickerWebsocket struct {
	*websocket[*WsTicker]
}

func (tws *TickerWebsocket) buildStreamNames(symbols ...string) []string {
	streams := make([]string, len(symbols))

	for i, s := range symbols {
		streams[i] = strings.ToLower(s) + "@ticker"
	}

	return streams
}

func (tws *TickerWebsocket) Subscribe(symbols ...string) Error {
	streams := tws.buildStreamNames(symbols...)
	return tws.subscribe(streams)
}

func (tws *TickerWebsocket) Unsubscribe(symbols ...string) Error {
	streams := tws.buildStreamNames(symbols...)
	return tws.unsubscribe(streams)
}

func (ws *wsEndpoints) Ticker(onMessage func(*WsTicker), symbols ...string) (*TickerWebsocket, Error) {
	var s TickerWebsocket
	streams := s.buildStreamNames(symbols...)

	base, err := newWebsocket(ws.client.wssBaseUrl, WebsocketRouteMarket, streams, onMessage, ws.client.logger)
	if err != nil {
		return nil, err
	}

	s.websocket = base

	return &s, nil
}

////
// All Market Tickers
////

type AllTickersWebsocket struct {
	*websocket[[]*WsTicker]
}

func (ws *wsEndpoints) AllTickers(onMessage func([]*WsTicker)) (*AllTickersWebsocket, Error) {
	var s AllTickersWebsocket

	base, err := newWebsocket(ws.client.wssBaseUrl, WebsocketRouteMarket, []string{"!ticker@arr"}, onMessage, ws.client.logger)
	if err != nil {
		return nil, err
	}

	s.websocket = base

	return &s, nil
}

////
// Book Ticker
////

type WsBookTicker struct {
	Event    EventType `json:"e"`
	UpdateID int64     `json:"u"`
	Symbol   string    `json:"s"`
	Pair     string    `json:"ps"`
	// Event time
	EventTime int64 `json:"E"`
	// Transaction time
	TransactTime int64   `json:"T"`
	BestBidPrice float64 `json:"b,string"`
	BestBidQty   float64 `json:"B,string"`
	BestAskPrice float64 `json:"a,string"`
	BestAskQty   float64 `json:"A,string"`
	// Symbol type: 1 = UM, 2 = CM
	SymbolType int `json:"st"`
}

type BookTickerWebsocket struct {
	*websocket[*WsBookTicker]
}

func (btws *BookTickerWebsocket) buildStreamNames(symbols ...string) []string {
	streams := make([]string, len(symbols))

	for i, s := range symbols {
		streams[i] = strings.ToLower(s) + "@bookTicker"
	}

	return streams
}

func (btws *BookTickerWebsocket) Subscribe(symbols ...string) Error {
	streams := btws.buildStreamNames(symbols...)
	return btws.subscribe(streams)
}

func (btws *BookTickerWebsocket) Unsubscribe(symbols ...string) Error {
	streams := btws.buildStreamNames(symbols...)
	return btws.unsubscribe(streams)
}

func (ws *wsEndpoints) BookTicker(onMessage func(*WsBookTicker), symbols ...string) (*BookTickerWebsocket, Error) {
	var s BookTickerWebsocket
	streams := s.buildStreamNames(symbols...)

	base, err := newWebsocket(ws.client.wssBaseUrl, WebsocketRouteMarket, streams, onMessage, ws.client.logger)
	if err != nil {
		return nil, err
	}

	s.websocket = base

	return &s, nil
}

////
// All Book Tickers
////

type AllBookTickersWebsocket struct {
	*websocket[*WsBookTicker]
}

func (ws *wsEndpoints) AllBookTickers(onMessage func(*WsBookTicker)) (*AllBookTickersWebsocket, Error) {
	var s AllBookTickersWebsocket

	base, err := newWebsocket(ws.client.wssBaseUrl, WebsocketRouteMarket, []string{"!bookTicker@arr"}, onMessage, ws.client.logger)
	if err != nil {
		return nil, err
	}

	s.websocket = base

	return &s, nil
}

////
// Liquidation Order (Force Order)
////

type WsForceOrder struct {
	Event     EventType        `json:"e"`
	EventTime int64            `json:"E"`
	Order     WsForceOrderInfo `json:"o"`
}

type WsForceOrderInfo struct {
	Symbol       string      `json:"s"`
	Side         OrderSide   `json:"S"`
	OrderType    OrderType   `json:"o"`
	TimeInForce  TimeInForce `json:"f"`
	OrigQuantity float64     `json:"q,string"`
	Price        float64     `json:"p,string"`
	AvgPrice     float64     `json:"ap,string"`
	OrderStatus  OrderStatus `json:"X"`
	// Order Last Filled Quantity
	LastFilledQty float64 `json:"l,string"`
	// Order Filled Accumulated Quantity
	FilledAccumulatedQty float64 `json:"z,string"`
	TradeTime            int64   `json:"T"`
}

type ForceOrderWebsocket struct {
	*websocket[*WsForceOrder]
}

func (fows *ForceOrderWebsocket) buildStreamNames(symbols ...string) []string {
	streams := make([]string, len(symbols))

	for i, s := range symbols {
		streams[i] = strings.ToLower(s) + "@forceOrder"
	}

	return streams
}

func (fows *ForceOrderWebsocket) Subscribe(symbols ...string) Error {
	streams := fows.buildStreamNames(symbols...)
	return fows.subscribe(streams)
}

func (fows *ForceOrderWebsocket) Unsubscribe(symbols ...string) Error {
	streams := fows.buildStreamNames(symbols...)
	return fows.unsubscribe(streams)
}

func (ws *wsEndpoints) ForceOrder(onMessage func(*WsForceOrder), symbols ...string) (*ForceOrderWebsocket, Error) {
	var s ForceOrderWebsocket
	streams := s.buildStreamNames(symbols...)

	base, err := newWebsocket(ws.client.wssBaseUrl, WebsocketRouteMarket, streams, onMessage, ws.client.logger)
	if err != nil {
		return nil, err
	}

	s.websocket = base

	return &s, nil
}

////
// All Market Liquidation Orders
////

type AllForceOrdersWebsocket struct {
	*websocket[*WsForceOrder]
}

func (ws *wsEndpoints) AllForceOrders(onMessage func(*WsForceOrder)) (*AllForceOrdersWebsocket, Error) {
	var s AllForceOrdersWebsocket

	base, err := newWebsocket(ws.client.wssBaseUrl, WebsocketRouteMarket, []string{"!forceOrder@arr"}, onMessage, ws.client.logger)
	if err != nil {
		return nil, err
	}

	s.websocket = base

	return &s, nil
}

////
// Diff. Book Depth
////

// WsDepth is the payload shared by both the Diff. Book Depth stream and the
// Partial Book Depth stream. On USD-M futures both carry the same shape,
// including the update-ID bookkeeping fields (U, u, pu).
type WsDepth struct {
	Event     EventType `json:"e"`
	EventTime int64     `json:"E"`
	// Transaction time
	TransactTime int64  `json:"T"`
	Symbol       string `json:"s"`
	// First update ID in event
	FirstUpdateID int64 `json:"U"`
	// Final update ID in event
	FinalUpdateID int64 `json:"u"`
	// Final update ID in last stream (used to detect gaps)
	PrevFinalUpdateID int64        `json:"pu"`
	Bids              []PriceLevel `json:"b"`
	Asks              []PriceLevel `json:"a"`
	Pair              string       `json:"ps"`
	// Symbol type: 1 = UM, 2 = CM
	SymbolType int `json:"st"`
}

// WsDepthInterval controls the update frequency of the depth streams.
type WsDepthInterval string

const (
	// 100ms
	WsDepthFast WsDepthInterval = "100ms"
	// 500ms
	WsDepthMedium WsDepthInterval = "500ms"
	// 250ms (default)
	WsDepthSlow WsDepthInterval = ""
)

type WsDiffDepthParams struct {
	Symbol string
	Speed  WsDepthInterval
}

type DiffDepthWebsocket struct {
	*websocket[*WsDepth]
}

func (dws *DiffDepthWebsocket) buildStreamNames(params ...WsDiffDepthParams) []string {
	streams := make([]string, len(params))

	for i, p := range params {
		stream := strings.ToLower(p.Symbol) + "@depth"

		if p.Speed != "" {
			stream += "@" + string(p.Speed)
		}

		streams[i] = stream
	}

	return streams
}

func (dws *DiffDepthWebsocket) Subscribe(params ...WsDiffDepthParams) Error {
	streams := dws.buildStreamNames(params...)
	return dws.subscribe(streams)
}

func (dws *DiffDepthWebsocket) Unsubscribe(params ...WsDiffDepthParams) Error {
	streams := dws.buildStreamNames(params...)
	return dws.unsubscribe(streams)
}

func (ws *wsEndpoints) Depth(onMessage func(*WsDepth), params ...WsDiffDepthParams) (*DiffDepthWebsocket, Error) {
	var s DiffDepthWebsocket
	streams := s.buildStreamNames(params...)

	base, err := newWebsocket(ws.client.wssBaseUrl, WebsocketRoutePublic, streams, onMessage, ws.client.logger)
	if err != nil {
		return nil, err
	}

	s.websocket = base

	return &s, nil
}

////
// Partial Book Depth
////

type WsPartialDepthParams struct {
	Symbol string
	// 5, 10 or 20
	Levels int
	Speed  WsDepthInterval
}

type PartialDepthWebsocket struct {
	*websocket[*WsDepth]
}

func (pdws *PartialDepthWebsocket) buildStreamNames(params ...WsPartialDepthParams) []string {
	streams := make([]string, len(params))

	for i, p := range params {
		stream := strings.ToLower(p.Symbol) + "@depth" + strconv.Itoa(p.Levels)

		if p.Speed != "" {
			stream += "@" + string(p.Speed)
		}

		streams[i] = stream
	}

	return streams
}

func (pdws *PartialDepthWebsocket) Subscribe(params ...WsPartialDepthParams) Error {
	streams := pdws.buildStreamNames(params...)
	return pdws.subscribe(streams)
}

func (pdws *PartialDepthWebsocket) Unsubscribe(params ...WsPartialDepthParams) Error {
	streams := pdws.buildStreamNames(params...)
	return pdws.unsubscribe(streams)
}

func (ws *wsEndpoints) PartialDepth(onMessage func(*WsDepth), params ...WsPartialDepthParams) (*PartialDepthWebsocket, Error) {
	var s PartialDepthWebsocket
	streams := s.buildStreamNames(params...)

	base, err := newWebsocket(ws.client.wssBaseUrl, WebsocketRoutePublic, streams, onMessage, ws.client.logger)
	if err != nil {
		return nil, err
	}

	s.websocket = base

	return &s, nil
}

////
// Composite Index
////

type WsCompositeIndex struct {
	Event     EventType `json:"e"`
	EventTime int64     `json:"E"`
	Symbol    string    `json:"s"`
	Price     float64   `json:"p,string"`
	Component string    `json:"C"`
	// Composition of the index
	Composition []WsCompositeIndexComponent `json:"c"`
}

type WsCompositeIndexComponent struct {
	BaseAsset          string  `json:"b"`
	QuoteAsset         string  `json:"q"`
	WeightInQuantity   float64 `json:"w,string"`
	WeightInPercentage float64 `json:"W,string"`
	IndexPrice         float64 `json:"i,string"`
}

type CompositeIndexWebsocket struct {
	*websocket[*WsCompositeIndex]
}

func (ciws *CompositeIndexWebsocket) buildStreamNames(symbols ...string) []string {
	streams := make([]string, len(symbols))

	for i, s := range symbols {
		streams[i] = strings.ToLower(s) + "@compositeIndex"
	}

	return streams
}

func (ciws *CompositeIndexWebsocket) Subscribe(symbols ...string) Error {
	streams := ciws.buildStreamNames(symbols...)
	return ciws.subscribe(streams)
}

func (ciws *CompositeIndexWebsocket) Unsubscribe(symbols ...string) Error {
	streams := ciws.buildStreamNames(symbols...)
	return ciws.unsubscribe(streams)
}

func (ws *wsEndpoints) CompositeIndex(onMessage func(*WsCompositeIndex), symbols ...string) (*CompositeIndexWebsocket, Error) {
	var s CompositeIndexWebsocket
	streams := s.buildStreamNames(symbols...)

	base, err := newWebsocket(ws.client.wssBaseUrl, WebsocketRouteMarket, streams, onMessage, ws.client.logger)
	if err != nil {
		return nil, err
	}

	s.websocket = base

	return &s, nil
}

////
// Contract Info
////

type WsContractInfo struct {
	Event        EventType      `json:"e"`
	EventTime    int64          `json:"E"`
	Symbol       string         `json:"s"`
	Pair         string         `json:"ps"`
	ContractType ContractType   `json:"ct"`
	DeliveryDate int64          `json:"dt"`
	OnboardDate  int64          `json:"ot"`
	Status       ContractStatus `json:"cs"`
	// Brackets (notional-based leverage brackets)
	Brackets []WsContractBracket `json:"bks"`
}

type WsContractBracket struct {
	// Notional bracket
	Bracket int `json:"bs"`
	// Floor notional of this bracket
	BracketNotionalFloor float64 `json:"bnf"`
	// Cap notional of this bracket
	BracketNotionalCap float64 `json:"bnc"`
	// Maintenance ratio for this bracket
	MaintenanceRatio float64 `json:"mmr"`
	// Auxiliary number for quick calculation
	Auxiliary float64 `json:"cf"`
	// Min leverage for this bracket
	MinLeverage int `json:"mi"`
	// Max leverage for this bracket
	MaxLeverage int `json:"ma"`
}

type ContractInfoWebsocket struct {
	*websocket[*WsContractInfo]
}

// ContractInfo subscribes to the all-market Contract Info stream (!contractInfo),
// which pushes an event whenever a contract's info is updated (leverage bracket,
// funding interval, contract status, etc.).
func (ws *wsEndpoints) ContractInfo(onMessage func(*WsContractInfo)) (*ContractInfoWebsocket, Error) {
	var s ContractInfoWebsocket

	base, err := newWebsocket(ws.client.wssBaseUrl, WebsocketRouteMarket, []string{"!contractInfo"}, onMessage, ws.client.logger)
	if err != nil {
		return nil, err
	}

	s.websocket = base

	return &s, nil
}

////
// Multi-Assets Mode Asset Index
////

type WsAssetIndex struct {
	Event                 EventType `json:"e"`
	EventTime             int64     `json:"E"`
	Symbol                string    `json:"s"`
	Index                 float64   `json:"i,string"`
	BidBuffer             float64   `json:"b,string"`
	AskBuffer             float64   `json:"a,string"`
	BidRate               float64   `json:"B,string"`
	AskRate               float64   `json:"A,string"`
	AutoExchangeBidBuffer float64   `json:"q,string"`
	AutoExchangeAskBuffer float64   `json:"g,string"`
	AutoExchangeBidRate   float64   `json:"Q,string"`
	AutoExchangeAskRate   float64   `json:"G,string"`
}

type AssetIndexWebsocket struct {
	*websocket[*WsAssetIndex]
}

func (aiws *AssetIndexWebsocket) buildStreamNames(assetSymbols ...string) []string {
	streams := make([]string, len(assetSymbols))

	for i, s := range assetSymbols {
		streams[i] = strings.ToLower(s) + "@assetIndex"
	}

	return streams
}

func (aiws *AssetIndexWebsocket) Subscribe(assetSymbols ...string) Error {
	streams := aiws.buildStreamNames(assetSymbols...)
	return aiws.subscribe(streams)
}

func (aiws *AssetIndexWebsocket) Unsubscribe(assetSymbols ...string) Error {
	streams := aiws.buildStreamNames(assetSymbols...)
	return aiws.unsubscribe(streams)
}

func (ws *wsEndpoints) AssetIndex(onMessage func(*WsAssetIndex), assetSymbols ...string) (*AssetIndexWebsocket, Error) {
	var s AssetIndexWebsocket
	streams := s.buildStreamNames(assetSymbols...)

	base, err := newWebsocket(ws.client.wssBaseUrl, WebsocketRouteMarket, streams, onMessage, ws.client.logger)
	if err != nil {
		return nil, err
	}

	s.websocket = base

	return &s, nil
}

////
// All Market Asset Index
////

type AllAssetIndexWebsocket struct {
	*websocket[[]*WsAssetIndex]
}

func (ws *wsEndpoints) AllAssetIndex(onMessage func([]*WsAssetIndex)) (*AllAssetIndexWebsocket, Error) {
	var s AllAssetIndexWebsocket

	base, err := newWebsocket(ws.client.wssBaseUrl, WebsocketRouteMarket, []string{"!assetIndex@arr"}, onMessage, ws.client.logger)
	if err != nil {
		return nil, err
	}

	s.websocket = base

	return &s, nil
}
