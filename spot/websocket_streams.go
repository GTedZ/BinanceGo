package spot

import (
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
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
// Reference Price
////

type WsReferencePrice struct {
	Event          EventType `json:"e"`
	Symbol         string    `json:"s"`
	ReferencePrice float64   `json:"r,string"`
	Timestamp      int64     `json:"t"`
}

type ReferencePriceWebsocket struct {
	*websocket[*WsReferencePrice]
}

func (rpws *ReferencePriceWebsocket) buildStreamNames(symbols ...string) []string {
	streams := make([]string, len(symbols))

	for i, s := range symbols {
		streams[i] = strings.ToLower(s) + "@referencePrice"
	}

	return streams
}

func (rpws *ReferencePriceWebsocket) Subscribe(symbols ...string) Error {
	streams := rpws.buildStreamNames(symbols...)
	return rpws.subscribe(streams)
}

func (rpws *ReferencePriceWebsocket) Unsubscribe(symbols ...string) Error {
	streams := rpws.buildStreamNames(symbols...)
	return rpws.unsubscribe(streams)
}

func (ws *wsEndpoints) ReferencePrice(onMessage func(*WsReferencePrice), symbols ...string) (*ReferencePriceWebsocket, Error) {
	var s ReferencePriceWebsocket
	streams := s.buildStreamNames(symbols...)

	base, err := newWebsocket(ws.client.wssBaseUrlMarketData, streams, onMessage, ws.client.logger)
	if err != nil {
		return nil, err
	}

	s.websocket = base

	return &s, nil
}

////
// AggTrade
////

type WsAggTrade struct {
	Event        EventType `json:"e"`
	EventTime    int64     `json:"E"`
	Symbol       string    `json:"s"`
	AggTradeID   int64     `json:"a"`
	Price        float64   `json:"p,string"`
	Quantity     float64   `json:"q,string"`
	FirstTradeID int64     `json:"f"`
	LastTradeID  int64     `json:"l"`
	TradeTime    int64     `json:"T"`
	IsMaker      bool      `json:"m"`
}

type AggTradeWebsocket struct {
	*websocket[*WsAggTrade]
}

func (rpws *AggTradeWebsocket) buildStreamNames(symbols ...string) []string {
	streams := make([]string, len(symbols))

	for i, s := range symbols {
		streams[i] = strings.ToLower(s) + "@aggTrade"
	}

	return streams
}

func (rpws *AggTradeWebsocket) Subscribe(symbols ...string) Error {
	streams := rpws.buildStreamNames(symbols...)
	return rpws.subscribe(streams)
}

func (rpws *AggTradeWebsocket) Unsubscribe(symbols ...string) Error {
	streams := rpws.buildStreamNames(symbols...)
	return rpws.unsubscribe(streams)
}

func (ws *wsEndpoints) AggTrade(onMessage func(*WsAggTrade), symbols ...string) (*AggTradeWebsocket, Error) {
	var s AggTradeWebsocket
	streams := s.buildStreamNames(symbols...)

	base, err := newWebsocket(ws.client.wssBaseUrlMarketData, streams, onMessage, ws.client.logger)
	if err != nil {
		return nil, err
	}

	s.websocket = base

	return &s, nil
}

////
// Trade
////

type WsTrade struct {
	Event     EventType `json:"e"`
	EventTime int64     `json:"E"`
	Symbol    string    `json:"s"`
	TradeID   int64     `json:"t"`
	Price     float64   `json:"p,string"`
	Quantity  float64   `json:"q,string"`
	TradeTime int64     `json:"T"`
	IsMaker   bool      `json:"m"`
}

type TradeWebsocket struct {
	*websocket[*WsTrade]
}

func (rpws *TradeWebsocket) buildStreamNames(symbols ...string) []string {
	streams := make([]string, len(symbols))

	for i, s := range symbols {
		streams[i] = strings.ToLower(s) + "@trade"
	}

	return streams
}

func (rpws *TradeWebsocket) Subscribe(symbols ...string) Error {
	streams := rpws.buildStreamNames(symbols...)
	return rpws.subscribe(streams)
}

func (rpws *TradeWebsocket) Unsubscribe(symbols ...string) Error {
	streams := rpws.buildStreamNames(symbols...)
	return rpws.unsubscribe(streams)
}

func (ws *wsEndpoints) Trade(onMessage func(*WsTrade), symbols ...string) (*TradeWebsocket, Error) {
	var s TradeWebsocket
	streams := s.buildStreamNames(symbols...)

	base, err := newWebsocket(ws.client.wssBaseUrlMarketData, streams, onMessage, ws.client.logger)
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
}

type CandlesticksWebsocket struct {
	*websocket[*WsKline]
}

type WsCandlesticksParams struct {
	Symbol   string
	Interval WsCandlesticksInterval
}

type WsCandlesticksInterval string

const (
	WsKlineFast WsCandlesticksInterval = "1000ms"
	WsKlineSlow WsCandlesticksInterval = "2000ms"
)

func (rpws *CandlesticksWebsocket) buildStreamNames(params ...WsCandlesticksParams) []string {
	streams := make([]string, len(params))

	for i, p := range params {
		streams[i] = strings.ToLower(p.Symbol) + "@kline_" + string(p.Interval)
	}

	return streams
}

func (rpws *CandlesticksWebsocket) Subscribe(params ...WsCandlesticksParams) Error {
	streams := rpws.buildStreamNames(params...)
	return rpws.subscribe(streams)
}

func (rpws *CandlesticksWebsocket) Unsubscribe(params ...WsCandlesticksParams) Error {
	streams := rpws.buildStreamNames(params...)
	return rpws.unsubscribe(streams)
}

func (ws *wsEndpoints) Candlesticks(onMessage func(*WsKline), params ...WsCandlesticksParams) (*CandlesticksWebsocket, Error) {
	var s CandlesticksWebsocket
	streams := s.buildStreamNames(params...)

	base, err := newWebsocket(ws.client.wssBaseUrlMarketData, streams, onMessage, ws.client.logger)
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
	FirstPrice         float64   `json:"x,string"`
	LastPrice          float64   `json:"c,string"`
	LastQuantity       float64   `json:"Q,string"`
	BestBidPrice       float64   `json:"b,string"`
	BestBidQty         float64   `json:"B,string"`
	BestAskPrice       float64   `json:"a,string"`
	BestAskQty         float64   `json:"A,string"`
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

	base, err := newWebsocket(ws.client.wssBaseUrlMarketData, streams, onMessage, ws.client.logger)
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
	Open        float64   `json:"o,string"`
	High        float64   `json:"h,string"`
	Low         float64   `json:"l,string"`
	Close       float64   `json:"c,string"`
	Volume      float64   `json:"v,string"`
	QuoteVolume float64   `json:"q,string"`
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

	base, err := newWebsocket(ws.client.wssBaseUrlMarketData, streams, onMessage, ws.client.logger)
	if err != nil {
		return nil, err
	}

	s.websocket = base

	return &s, nil
}

////
// Rolling Window Ticker
////

type WsRollingWindowTicker struct {
	Event              EventType `json:"e"`
	EventTime          int64     `json:"E"`
	Symbol             string    `json:"s"`
	PriceChange        float64   `json:"p,string"`
	PriceChangePercent float64   `json:"P,string"`
	Open               float64   `json:"o,string"`
	HighPrice          float64   `json:"h,string"`
	LowPrice           float64   `json:"l,string"`
	LastPrice          float64   `json:"c,string"`
	WeightedAvgPrice   float64   `json:"w,string"`
	Volume             float64   `json:"v,string"`
	QuoteVolume        float64   `json:"q,string"`
	OpenTime           int64     `json:"O"`
	CloseTime          int64     `json:"C"`
	FirstTradeID       int64     `json:"F"`
	LastTradeID        int64     `json:"L"`
	TradeCount         int64     `json:"n"`
}

type RollingWindowTickerWebsocket struct {
	*websocket[*WsRollingWindowTicker]
}

type WsRollingWindowParams struct {
	Symbol string
	Window string // "1h", "4h", "1d"
}

func (rtws *RollingWindowTickerWebsocket) buildStreamNames(params ...WsRollingWindowParams) []string {
	streams := make([]string, len(params))

	for i, p := range params {
		streams[i] = strings.ToLower(p.Symbol) + "@ticker_" + p.Window
	}

	return streams
}

func (rtws *RollingWindowTickerWebsocket) Subscribe(params ...WsRollingWindowParams) Error {
	streams := rtws.buildStreamNames(params...)
	return rtws.subscribe(streams)
}

func (rtws *RollingWindowTickerWebsocket) Unsubscribe(params ...WsRollingWindowParams) Error {
	streams := rtws.buildStreamNames(params...)
	return rtws.unsubscribe(streams)
}

func (ws *wsEndpoints) RollingWindowTicker(
	onMessage func(*WsRollingWindowTicker),
	params ...WsRollingWindowParams,
) (*RollingWindowTickerWebsocket, Error) {
	var s RollingWindowTickerWebsocket
	streams := s.buildStreamNames(params...)

	base, err := newWebsocket(ws.client.wssBaseUrlMarketData, streams, onMessage, ws.client.logger)
	if err != nil {
		return nil, err
	}

	s.websocket = base

	return &s, nil
}

////
// Average Price
////

type WsAveragePrice struct {
	Event     EventType `json:"e"`
	EventTime int64     `json:"E"`
	Symbol    string    `json:"s"`
	Interval  Interval  `json:"i"`
	AvgPrice  float64   `json:"w,string"`
	LastTime  int64     `json:"T"`
}

type AveragePriceWebsocket struct {
	*websocket[*WsAveragePrice]
}

func (apws *AveragePriceWebsocket) buildStreamNames(symbols ...string) []string {
	streams := make([]string, len(symbols))

	for i, s := range symbols {
		streams[i] = strings.ToLower(s) + "@avgPrice"
	}

	return streams
}

func (apws *AveragePriceWebsocket) Subscribe(symbols ...string) Error {
	streams := apws.buildStreamNames(symbols...)
	return apws.subscribe(streams)
}

func (apws *AveragePriceWebsocket) Unsubscribe(symbols ...string) Error {
	streams := apws.buildStreamNames(symbols...)
	return apws.unsubscribe(streams)
}

func (ws *wsEndpoints) AveragePrice(onMessage func(*WsAveragePrice), symbols ...string) (*AveragePriceWebsocket, Error) {
	var s AveragePriceWebsocket
	streams := s.buildStreamNames(symbols...)

	base, err := newWebsocket(ws.client.wssBaseUrlMarketData, streams, onMessage, ws.client.logger)
	if err != nil {
		return nil, err
	}

	s.websocket = base

	return &s, nil
}

////
// Order Book
////

type WsOrderBook struct {
	LastUpdateID int64        `json:"lastUpdateId"`
	Bids         []PriceLevel `json:"bids"`
	Asks         []PriceLevel `json:"asks"`
}

type OrderBookWebsocket struct {
	*websocket[*WsOrderBook]
}

type WsOrderBookParams struct {
	Symbol string
	Speed  WsDiffDepthInterval
	// 5, 10 or 20
	Levels int
}

type WsDiffDepthInterval string

const (
	WsDepthFast WsDiffDepthInterval = "100ms"
	WsDepthSlow WsDiffDepthInterval = ""
)

func (ows *OrderBookWebsocket) buildStreamNames(params ...WsOrderBookParams) []string {
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

func (ows *OrderBookWebsocket) Subscribe(params ...WsOrderBookParams) Error {
	streams := ows.buildStreamNames(params...)
	return ows.subscribe(streams)
}

func (ows *OrderBookWebsocket) Unsubscribe(params ...WsOrderBookParams) Error {
	streams := ows.buildStreamNames(params...)
	return ows.unsubscribe(streams)
}

func (ws *wsEndpoints) OrderBook(onMessage func(*WsOrderBook), params ...WsOrderBookParams) (*OrderBookWebsocket, Error) {
	var s OrderBookWebsocket
	streams := s.buildStreamNames(params...)

	base, err := newWebsocket(ws.client.wssBaseUrlMarketData, streams, onMessage, ws.client.logger)
	if err != nil {
		return nil, err
	}

	s.websocket = base

	return &s, nil
}

////
// Depth Update
////

type WsDiffDepth struct {
	Event     EventType    `json:"e"`
	EventTime int64        `json:"E"`
	Symbol    string       `json:"s"`
	FirstID   int64        `json:"U"`
	FinalID   int64        `json:"u"`
	Bids      []PriceLevel `json:"b"`
	Asks      []PriceLevel `json:"a"`
}

type DiffDepthWebsocket struct {
	*websocket[*WsDiffDepth]
}

type WsDiffDepthParams struct {
	Symbol string
	Speed  WsDiffDepthInterval
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

func (ws *wsEndpoints) Depth(onMessage func(*WsDiffDepth), params ...WsDiffDepthParams) (*DiffDepthWebsocket, Error) {
	var s DiffDepthWebsocket
	streams := s.buildStreamNames(params...)

	base, err := newWebsocket(ws.client.wssBaseUrlMarketData, streams, onMessage, ws.client.logger)
	if err != nil {
		return nil, err
	}

	s.websocket = base

	return &s, nil
}

////
// Managed Order Book
////

type ManagedOrderBookHandler struct {
	client *Client

	book *managedOrderBook

	buffer []*WsDiffDepth
	mu     sync.RWMutex

	ready bool

	ws *DiffDepthWebsocket
}

func (h *ManagedOrderBookHandler) RLock() {
	h.mu.RLock()
}

func (h *ManagedOrderBookHandler) RUnlock() {
	h.mu.RUnlock()
}

func (h *ManagedOrderBookHandler) buildBidsUnsafe() []PriceLevel {
	bids := make([]PriceLevel, 0, len(h.book.Bids))
	for p, q := range h.book.Bids {
		bids = append(bids, PriceLevel{
			Price: p,
			Qty:   q,
		})
	}

	return bids
}

// Bids return copies of the internal state.
// They are safe to use without additional locking.
//
// If you need multiple fields atomically (e.g. bids + asks + updateID),
// use RLock()/RUnlock() manually.
func (h *ManagedOrderBookHandler) Bids() []PriceLevel {
	h.mu.RLock()
	defer h.mu.RUnlock()

	return h.buildBidsUnsafe()
}

func (h *ManagedOrderBookHandler) BidsSorted(ascending bool) []PriceLevel {
	bids := h.Bids()
	h.SortPriceLevels(bids, ascending)

	return bids
}

func (h *ManagedOrderBookHandler) buildAsksUnsafe() []PriceLevel {
	asks := make([]PriceLevel, 0, len(h.book.Asks))
	for p, q := range h.book.Asks {
		asks = append(asks, PriceLevel{
			Price: p,
			Qty:   q,
		})
	}

	return asks
}

func (h *ManagedOrderBookHandler) Asks() []PriceLevel {
	h.mu.RLock()
	defer h.mu.RUnlock()

	return h.buildAsksUnsafe()
}

func (h *ManagedOrderBookHandler) AsksSorted(ascending bool) []PriceLevel {
	asks := h.Asks()
	h.SortPriceLevels(asks, ascending)

	return asks
}

// SortPriceLevels sorts the array in place.
func (h *ManagedOrderBookHandler) SortPriceLevels(levels []PriceLevel, ascending bool) {
	if ascending {
		// Ascending
		sort.Slice(levels, func(i, j int) bool {
			return levels[i].Price < levels[j].Price
		})
	} else {
		// Descending
		sort.Slice(levels, func(i, j int) bool {
			return levels[i].Price > levels[j].Price
		})
	}
}

func (h *ManagedOrderBookHandler) LastUpdateID() int64 {
	h.mu.RLock()
	defer h.mu.RUnlock()

	return h.book.LastUpdateID
}

func (h *ManagedOrderBookHandler) Symbol() string {
	return h.book.Symbol
}

func (h *ManagedOrderBookHandler) BestBid() (PriceLevel, bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if len(h.book.Bids) == 0 {
		return PriceLevel{}, false
	}

	var bestPrice float64
	var bestQty float64
	first := true

	for p, q := range h.book.Bids {
		if first || p > bestPrice {
			bestPrice = p
			bestQty = q
			first = false
		}
	}

	return PriceLevel{
		Price: bestPrice,
		Qty:   bestQty,
	}, true
}

func (h *ManagedOrderBookHandler) BestAsk() (PriceLevel, bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if len(h.book.Asks) == 0 {
		return PriceLevel{}, false
	}

	var bestPrice float64
	var bestQty float64
	first := true

	for p, q := range h.book.Asks {
		if first || p < bestPrice {
			bestPrice = p
			bestQty = q
			first = false
		}
	}

	return PriceLevel{
		Price: bestPrice,
		Qty:   bestQty,
	}, true
}

// Snapshot returns a copy of the ManagedOrderBook
//
// If sorted is `true`.
// Bids are sorted in descending order
// Asks are sorted in ascending order
func (h *ManagedOrderBookHandler) Snapshot(sorted bool) ManagedOrderBook {
	h.mu.RLock()
	defer h.mu.RUnlock()

	bids := h.buildBidsUnsafe()
	asks := h.buildAsksUnsafe()

	if sorted {
		h.SortPriceLevels(bids, false)
		h.SortPriceLevels(asks, true)
	}

	return ManagedOrderBook{
		Symbol:       h.book.Symbol,
		LastUpdateID: h.book.LastUpdateID,
		Bids:         bids,
		Asks:         asks,
	}
}

type managedOrderBook struct {
	Symbol       string
	LastUpdateID int64

	Bids map[float64]float64
	Asks map[float64]float64
}

type ManagedOrderBook struct {
	Symbol       string
	LastUpdateID int64
	Bids         []PriceLevel
	Asks         []PriceLevel
}

// ManagedOrderBook
//
// WARNING:
// The ManagedOrderBook passed to the callback is mutable.
// It is reused internally and updated in-place.
//
// If you need to keep a snapshot, you must copy it.
func (ws *wsEndpoints) ManagedOrderBook(symbol string, onUpdate func()) (*ManagedOrderBookHandler, Error) {

	handler := &ManagedOrderBookHandler{
		client: ws.client,
		book: &managedOrderBook{
			Symbol: symbol,
			Bids:   make(map[float64]float64),
			Asks:   make(map[float64]float64),
		},
	}

	// Step 1: start WS (buffering begins immediately)
	socket, err := ws.Depth(
		func(event *WsDiffDepth) {
			handler.handleEvent(event, onUpdate)
		},
		WsDiffDepthParams{
			Symbol: symbol,
			Speed:  WsDepthFast,
		},
	)
	if err != nil {
		return nil, err
	}

	handler.ws = socket

	// Step 2: block until initialized
	if err := handler.initializeSync(); err != nil {
		return nil, err
	}

	return handler, nil
}

func (h *ManagedOrderBookHandler) initializeSync() Error {
	for {
		// --- Phase 1: wait for buffer (under lock)
		h.mu.Lock()

		if len(h.buffer) == 0 {
			h.mu.Unlock()
			time.Sleep(50 * time.Millisecond)
			continue
		}

		h.mu.Unlock()

		// --- Phase 2: fetch snapshot (no lock)
		snapshot, _, err := h.client.OrderBook(h.book.Symbol)
		if err != nil {
			continue
		}

		// --- Phase 3: re-lock and validate
		h.mu.Lock()

		// buffer may have changed
		if len(h.buffer) == 0 {
			h.mu.Unlock()
			continue
		}

		firstEvent := h.buffer[0]

		// snapshot too old → retry
		if snapshot.LastUpdateId < firstEvent.FirstID {
			h.mu.Unlock()
			continue
		}

		// --- Apply snapshot
		h.book.LastUpdateID = snapshot.LastUpdateId
		h.book.Bids = make(map[float64]float64)
		h.book.Asks = make(map[float64]float64)

		for _, b := range snapshot.Bids {
			h.book.Bids[b.Price] = b.Qty
		}
		for _, a := range snapshot.Asks {
			h.book.Asks[a.Price] = a.Qty
		}

		// --- Drop outdated buffered events
		var newBuffer []*WsDiffDepth
		for _, e := range h.buffer {
			if e.FinalID > h.book.LastUpdateID {
				newBuffer = append(newBuffer, e)
			}
		}
		h.buffer = newBuffer

		// --- Apply buffered events
		for _, e := range h.buffer {
			if !h.applyEvent(e) {
				h.mu.Unlock()
				return h.initializeSync()
			}
		}

		h.buffer = nil
		h.ready = true

		h.mu.Unlock()
		return nil
	}
}

func (h *ManagedOrderBookHandler) handleEvent(event *WsDiffDepth, callback func()) {
	h.mu.Lock()

	// Buffer until ready
	if !h.ready {
		h.buffer = append(h.buffer, event)
		h.mu.Unlock()
		return
	}

	// Apply normally
	if !h.applyEvent(event) {
		h.ready = false
		h.buffer = nil
		h.mu.Unlock()

		go h.initializeSync()
		return
	}

	h.mu.Unlock()

	callback()
}

func (h *ManagedOrderBookHandler) applyEvent(e *WsDiffDepth) bool {
	// Ignore old
	if e.FinalID < h.book.LastUpdateID {
		return true
	}

	// Gap detected → resync required
	if e.FirstID > h.book.LastUpdateID+1 {
		return false
	}

	// Apply bids
	for _, b := range e.Bids {
		if b.Qty == 0 {
			delete(h.book.Bids, b.Price)
		} else {
			h.book.Bids[b.Price] = b.Qty
		}
	}

	// Apply asks
	for _, a := range e.Asks {
		if a.Qty == 0 {
			delete(h.book.Asks, a.Price)
		} else {
			h.book.Asks[a.Price] = a.Qty
		}
	}

	h.book.LastUpdateID = e.FinalID
	return true
}
