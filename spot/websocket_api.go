package spot

import (
	"time"

	"github.com/GTedZ/binancego/internal/berror"
	"github.com/GTedZ/binancego/internal/json"
	"github.com/GTedZ/binancego/internal/logging"
	"github.com/GTedZ/binancego/internal/requests"
	"github.com/GTedZ/binancego/internal/validation"
	"github.com/GTedZ/binancego/internal/websocket/binance"
)

type WebsocketAPI struct {
	base    baseWebsocket
	baseUrl string

	client *Client

	logger logging.Logger
}

func (c *Client) WebsocketAPI() (*WebsocketAPI, Error) {
	return newWebsocketAPI(c.wssApiBaseUrl, c)
}

////
// Private Methods
////

func (wsapi *WebsocketAPI) setOnMessage(messageHandler func(date []byte)) {
	wsapi.base.SetOnMessage(messageHandler)
}

////
// Public Methods
////

func (wsapi *WebsocketAPI) Close() {
	wsapi.base.Close()
}

////
// Order Book
////

func (wsapi *WebsocketAPI) OrderBook(symbol string, opts ...OrderBookParams) (*OrderBook, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "limit", opt.Limit)
		validation.SetIfNotZero(params, "symbolStatus", opt.SymbolStatus)
	}

	return doWsApiRequest[OrderBook](wsapi, "depth", params, wsapi.client.apikey, NONE)
}

////
// Recent Trades
////

func (wsapi *WebsocketAPI) RecentTrades(symbol string, opts ...RecentTradesParams) ([]*Trade, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "limit", opt.Limit)
	}

	result, resp, err := doWsApiRequest[[]*Trade](wsapi, "trades.recent", params, wsapi.client.apikey, NONE)
	if err != nil {
		return nil, resp, err
	}

	return *result, resp, err
}

////
// Historical Trades
////

func (wsapi *WebsocketAPI) HistoricalTrades(symbol string, opts ...HistoricalTradesParams) ([]*Trade, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "limit", opt.Limit)
		validation.SetIfNotZero(params, "fromId", opt.FromId)
	}

	result, resp, err := doWsApiRequest[[]*Trade](wsapi, "trades.historical", params, wsapi.client.apikey, NONE)
	if err != nil {
		return nil, resp, err
	}

	return *result, resp, err
}

////
// Aggregate Trades
////

func (wsapi *WebsocketAPI) AggTrades(symbol string, opts ...AggTradeParams) ([]*AggTrade, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "fromId", opt.FromId)
		validation.SetIfNotZero(params, "startTime", opt.StartTime)
		validation.SetIfNotZero(params, "endTime", opt.EndTime)
		validation.SetIfNotZero(params, "limit", opt.Limit)
	}

	result, resp, err := doWsApiRequest[[]*AggTrade](wsapi, "trades.aggregate", params, wsapi.client.apikey, NONE)
	if err != nil {
		return nil, resp, err
	}

	return *result, resp, err
}

////
// Candlesticks
////

func (wsapi *WebsocketAPI) Candlesticks(symbol string, interval Interval, opts ...CandlestickParams) ([]*Candlestick, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol
	params["interval"] = interval

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "startTime", opt.StartTime)
		validation.SetIfNotZero(params, "endTime", opt.EndTime)
		validation.SetIfNotZero(params, "timeZone", opt.TimeZone)
		validation.SetIfNotZero(params, "limit", opt.Limit)
	}

	result, resp, err := doWsApiRequest[[]*Candlestick](wsapi, "klines", params, wsapi.client.apikey, NONE)
	if err != nil {
		return nil, resp, err
	}

	return *result, resp, err
}

////
// UI Klines
////

func (wsapi *WebsocketAPI) UIKlines(symbol string, interval Interval, opts ...CandlestickParams) ([]*Candlestick, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol
	params["interval"] = interval

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "startTime", opt.StartTime)
		validation.SetIfNotZero(params, "endTime", opt.EndTime)
		validation.SetIfNotZero(params, "timeZone", opt.TimeZone)
		validation.SetIfNotZero(params, "limit", opt.Limit)
	}

	result, resp, err := doWsApiRequest[[]*Candlestick](wsapi, "uiKlines", params, wsapi.client.apikey, NONE)
	if err != nil {
		return nil, resp, err
	}

	return *result, resp, err
}

////
// Current Average Price
////

func (wsapi *WebsocketAPI) AveragePrice(symbol string) (*AveragePrice, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol

	result, resp, err := doWsApiRequest[*AveragePrice](wsapi, "avgPrice", params, wsapi.client.apikey, NONE)
	if err != nil {
		return nil, resp, err
	}

	return *result, resp, err
}

////
// 24h Ticker/s
////

func (wsapi *WebsocketAPI) Tickers24h(opts ...Tickers24hParams) ([]*Ticker24h, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	params["type"] = "FULL"

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "symbols", opt.Symbols)
		validation.SetIfNotZero(params, "symbolStatus", opt.SymbolStatus)
	}

	result, resp, err := doWsApiRequest[[]*Ticker24h](wsapi, "ticker.24hr", params, wsapi.client.apikey, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

func (wsapi *WebsocketAPI) Ticker24h(symbol string, opts ...Ticker24hParams) (*Ticker24h, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol
	params["type"] = "FULL"

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "symbolStatus", opt.SymbolStatus)
	}

	return doWsApiRequest[Ticker24h](wsapi, "ticker.24hr", params, wsapi.client.apikey, NONE)
}

////
// 24h MiniTicker/s
////

func (wsapi *WebsocketAPI) MiniTickers24h(opts ...MiniTickers24hParams) ([]*MiniTicker24h, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	params["type"] = "MINI"

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "symbols", opt.Symbols)
		validation.SetIfNotZero(params, "symbolStatus", opt.SymbolStatus)
	}

	result, resp, err := doWsApiRequest[[]*MiniTicker24h](wsapi, "ticker.24hr", params, wsapi.client.apikey, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

func (wsapi *WebsocketAPI) MiniTicker24h(symbol string, opts ...MiniTicker24hParams) (*MiniTicker24h, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol
	params["type"] = "MINI"

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "symbolStatus", opt.SymbolStatus)
	}

	return doWsApiRequest[MiniTicker24h](wsapi, "ticker.24hr", params, wsapi.client.apikey, NONE)
}

////
// Trading Day Ticker
////

func (wsapi *WebsocketAPI) TradingDayTickers(opts ...TradingDayTickersParams) ([]*TradingDayTicker, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	params["type"] = "FULL"

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "symbols", opt.Symbols)
		validation.SetIfNotZero(params, "symbolStatus", opt.SymbolStatus)
		validation.SetIfNotZero(params, "timeZone", opt.TimeZone)
	}

	result, resp, err := doWsApiRequest[[]*TradingDayTicker](wsapi, "ticker.tradingDay", params, wsapi.client.apikey, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

func (wsapi *WebsocketAPI) TradingDayTicker(symbol string, opts ...TradingDayTickerParams) (*TradingDayTicker, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol
	params["type"] = "FULL"

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "symbolStatus", opt.SymbolStatus)
		validation.SetIfNotZero(params, "timeZone", opt.TimeZone)
	}

	return doWsApiRequest[TradingDayTicker](wsapi, "ticker.tradingDay", params, wsapi.client.apikey, NONE)
}

////
// Trading Day MiniTicker
////

func (wsapi *WebsocketAPI) TradingDayMiniTickers(opts ...TradingDayTickersParams) ([]*TradingDayMiniTicker, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	params["type"] = "MINI"

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "symbols", opt.Symbols)
		validation.SetIfNotZero(params, "symbolStatus", opt.SymbolStatus)
		validation.SetIfNotZero(params, "timeZone", opt.TimeZone)
	}

	result, resp, err := doWsApiRequest[[]*TradingDayMiniTicker](wsapi, "ticker.tradingDay", params, wsapi.client.apikey, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

func (wsapi *WebsocketAPI) TradingDayMiniTicker(symbol string, opts ...TradingDayTickerParams) (*TradingDayMiniTicker, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol
	params["type"] = "MINI"

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "symbolStatus", opt.SymbolStatus)
		validation.SetIfNotZero(params, "timeZone", opt.TimeZone)
	}

	return doWsApiRequest[TradingDayMiniTicker](wsapi, "ticker.tradingDay", params, wsapi.client.apikey, NONE)
}

////
// Rolling Window Ticker/s
////

func (wsapi *WebsocketAPI) RollingWindowTickers(opts ...RollingWindowTickersParams) ([]*RollingWindowTicker, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	params["type"] = "FULL"

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "symbols", opt.Symbols)
		validation.SetIfNotZero(params, "symbolStatus", opt.SymbolStatus)
		validation.SetIfNotZero(params, "windowSize", opt.WindowSize)
	}

	result, resp, err := doWsApiRequest[[]*RollingWindowTicker](wsapi, "ticker", params, wsapi.client.apikey, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}
func (wsapi *WebsocketAPI) RollingWindowTicker(symbol string, opts ...RollingWindowTickerParams) (*RollingWindowTicker, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol
	params["type"] = "FULL"

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "symbolStatus", opt.SymbolStatus)
		validation.SetIfNotZero(params, "windowSize", opt.WindowSize)
	}

	return doWsApiRequest[RollingWindowTicker](wsapi, "ticker", params, wsapi.client.apikey, NONE)
}

// //
// Rolling Window MiniTicker/s
// //
func (wsapi *WebsocketAPI) RollingWindowMiniTickers(opts ...RollingWindowMiniTickersParams) ([]*RollingWindowMiniTicker, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	params["type"] = "MINI"

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "symbols", opt.Symbols)
		validation.SetIfNotZero(params, "symbolStatus", opt.SymbolStatus)
		validation.SetIfNotZero(params, "windowSize", opt.WindowSize)
	}

	result, resp, err := doWsApiRequest[[]*RollingWindowMiniTicker](wsapi, "ticker", params, wsapi.client.apikey, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}
func (wsapi *WebsocketAPI) RollingWindowMiniTicker(symbol string, opts ...RollingWindowTickerMiniParams) (*RollingWindowMiniTicker, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol
	params["type"] = "MINI"

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "symbolStatus", opt.SymbolStatus)
		validation.SetIfNotZero(params, "windowSize", opt.WindowSize)
	}

	return doWsApiRequest[RollingWindowMiniTicker](wsapi, "ticker", params, wsapi.client.apikey, NONE)
}

////
// Symbol Price Ticker
////

func (wsapi *WebsocketAPI) PriceTickers(opts ...PriceTickersParams) ([]*PriceTicker, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "symbols", opt.Symbols)
		validation.SetIfNotZero(params, "symbolStatus", opt.SymbolStatus)
	}

	result, resp, err := doWsApiRequest[[]*PriceTicker](wsapi, "ticker.price", params, wsapi.client.apikey, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

func (wsapi *WebsocketAPI) PriceTicker(symbol string, opts ...PriceTickerParams) (*PriceTicker, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "symbolStatus", opt.SymbolStatus)
	}

	return doWsApiRequest[PriceTicker](wsapi, "ticker.price", params, wsapi.client.apikey, NONE)
}

////
// Book Ticker
////

func (wsapi *WebsocketAPI) BookTickers(opts ...OrderBookTickersParams) ([]*OrderBookTicker, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "symbols", opt.Symbols)
		validation.SetIfNotZero(params, "symbolStatus", opt.SymbolStatus)
	}

	result, resp, err := doWsApiRequest[[]*OrderBookTicker](wsapi, "ticker.book", params, wsapi.client.apikey, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

func (wsapi *WebsocketAPI) BookTicker(symbol string, opts ...OrderBookTickerParams) (*OrderBookTicker, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "symbolStatus", opt.SymbolStatus)
	}

	return doWsApiRequest[OrderBookTicker](wsapi, "ticker.book", params, wsapi.client.apikey, NONE)
}

////
// Reference Price
////

func (wsapi *WebsocketAPI) ReferencePrice(symbol string) (*ReferencePrice, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol

	return doWsApiRequest[ReferencePrice](wsapi, "referencePrice", params, wsapi.client.apikey, NONE)
}

////
// Reference Price Calculation
////

func (wsapi *WebsocketAPI) ReferencePriceCalculation(symbol string, opts ...ReferencePriceCalculationParams) (*ReferencePriceCalculation, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol

	if len(opts) > 0 {
		opt := opts[0]
		validation.SetIfNotZero(params, "symbolStatus", opt.SymbolStatus)
	}

	return doWsApiRequest[ReferencePriceCalculation](wsapi, "referencePrice.calculation", params, wsapi.client.apikey, NONE)
}

////////////////////////////////////////////////////////////////////////////////
// Trading requests
////////////////////////////////////////////////////////////////////////////////

////
// Place Order
////

// PlaceOrder sends a new order request over the WebSocket API (method "order.place").
//
// It accepts any orderRequest implementation (e.g. NewLimitBuy, NewMarketSell,
// NewStopLossLimitOrder, etc.) — the exact same request builders used by the
// REST Client.Order — and submits it to the exchange over the WebSocket
// connection.
//
// The response type is controlled by newOrderRespType (ACK, RESULT or FULL),
// exactly as with the REST endpoint.
func (wsapi *WebsocketAPI) PlaceOrder(req orderRequest) (*Order, *WsApiResponse, Error) {
	params := req.build()

	return doWsApiRequest[Order](wsapi, "order.place", params, wsapi.client.apikey, TRADE)
}

////
// Test Order
////

// TestOrder validates a new order without sending it to the matching engine
// (method "order.test").
//
// NOTE:
// If computeCommissionRates is false, Binance returns an empty JSON object `{}`.
// In that case, all fields in OrderTest will remain zero values.
func (wsapi *WebsocketAPI) TestOrder(req orderRequest, computeCommissionRates bool) (*OrderTest, *WsApiResponse, Error) {
	params := req.build()

	params["computeCommissionRates"] = computeCommissionRates

	return doWsApiRequest[OrderTest](wsapi, "order.test", params, wsapi.client.apikey, TRADE)
}

////
// Place SOR Order
////

func (wsapi *WebsocketAPI) PlaceSorOrder(req orderRequest) (*Order, *WsApiResponse, Error) {
	params := req.build()

	return doWsApiRequest[Order](wsapi, "sor.order.place", params, wsapi.client.apikey, TRADE)
}

////
// Test SOR Order
////

func (wsapi *WebsocketAPI) TestSorOrder(req orderRequest, computeCommissionRates bool) (*OrderTest, *WsApiResponse, Error) {
	params := req.build()

	params["computeCommissionRates"] = computeCommissionRates

	return doWsApiRequest[OrderTest](wsapi, "sor.order.test", params, wsapi.client.apikey, TRADE)
}

////
// Cancel Order
////

func (wsapi *WebsocketAPI) CancelOrder(symbol string, orderId int64, opts ...CancelOrderParams) (*Order, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol
	validation.SetIfNotZero(params, "orderId", orderId)

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "origClientOrderId", opt.OrigClientOrderId)
		validation.SetIfNotZero(params, "newClientOrderId", opt.NewClientOrderId)

		validation.SetIfNotZero(params, "cancelRestrictions", opt.CancelRestrictions)

		validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	}

	return doWsApiRequest[Order](wsapi, "order.cancel", params, wsapi.client.apikey, TRADE)
}

////
// Cancel Replace Order
////

// TODO

////
// Cancel Open Orders
////

func (wsapi *WebsocketAPI) CancelOpenOrders(symbol string, opts ...CancelOpenOrdersParams) ([]*Order, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	}

	result, resp, err := doWsApiRequest[[]*Order](wsapi, "openOrders.cancelAll", params, wsapi.client.apikey, TRADE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

////////////////////////////////////////////////////////////////////////////////
// Account requests
////////////////////////////////////////////////////////////////////////////////

////
// Account Information
////

func (wsapi *WebsocketAPI) Account(opts ...AccountParams) (*Account, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "omitZeroBalances", opt.OmitZeroBalances)
		validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	}

	return doWsApiRequest[Account](wsapi, "account.status", params, wsapi.client.apikey, USER_DATA)
}

////
// Query Order
////

func (wsapi *WebsocketAPI) QueryOrder(symbol string, opts ...QueryOrderParams) (*Order, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "orderId", opt.OrderID)
		validation.SetIfNotZero(params, "origClientOrderId", opt.OrigClientOrderID)
		validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	}

	return doWsApiRequest[Order](wsapi, "order.status", params, wsapi.client.apikey, USER_DATA)
}

////
// Open Orders
////

func (wsapi *WebsocketAPI) OpenOrders(opts ...OpenOrdersParams) ([]*Order, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "symbol", opt.Symbol)
		validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	}

	result, resp, err := doWsApiRequest[[]*Order](wsapi, "openOrders.status", params, wsapi.client.apikey, USER_DATA)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

////
// All Orders
////

func (wsapi *WebsocketAPI) AllOrders(symbol string, opts ...AllOrdersParams) ([]*Order, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "orderId", opt.OrderID)
		validation.SetIfNotZero(params, "startTime", opt.StartTime)
		validation.SetIfNotZero(params, "endTime", opt.EndTime)
		validation.SetIfNotZero(params, "limit", opt.Limit)
		validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	}

	result, resp, err := doWsApiRequest[[]*Order](wsapi, "allOrders", params, wsapi.client.apikey, USER_DATA)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

////
// Account Trades
////

func (wsapi *WebsocketAPI) Trades(symbol string, opts ...AccountTradesParam) ([]*Trade, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "orderId", opt.OrderID)
		validation.SetIfNotZero(params, "startTime", opt.StartTime)
		validation.SetIfNotZero(params, "endTime", opt.EndTime)
		validation.SetIfNotZero(params, "fromId", opt.FromID)
		validation.SetIfNotZero(params, "limit", opt.Limit)
		validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	}

	result, resp, err := doWsApiRequest[[]*Trade](wsapi, "myTrades", params, wsapi.client.apikey, USER_DATA)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

////
// Unfilled Order Count
////

func (wsapi *WebsocketAPI) UnfilledOrderCount(opts ...UnfilledOrderCountParams) ([]*UnfilledOrderCount, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	}

	result, resp, err := doWsApiRequest[[]*UnfilledOrderCount](wsapi, "account.rateLimits.orders", params, wsapi.client.apikey, USER_DATA)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

////
// Account Commission
////

func (wsapi *WebsocketAPI) AccountCommission(symbol string, opts ...AccountCommissionParams) (*AccountCommission, *WsApiResponse, Error) {
	params := map[string]interface{}{
		"symbol": symbol,
	}

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	}

	return doWsApiRequest[AccountCommission](wsapi, "account.commission", params, wsapi.client.apikey, USER_DATA)
}

////
// Response
////

// Status codes in the status field are the same as in HTTP.
// Here are some common status codes that you might encounter:
//
// - 200 indicates a successful response.
//
// - 4XX status codes indicate invalid requests; the issue is on your side.
//
// - 400 – your request failed, see error for the reason.
//
// - 403 – you have been blocked by the Web Application Firewall. This can indicate a rate limit violation or a security block.
// See https://www.binance.com/en/support/faq/detail/360004492232 for more details.
//
// - 409 – your request partially failed but also partially succeeded, see error for details.
//
// - 418 – you have been auto-banned for repeated violation of rate limits.
//
// - 429 – you have exceeded API request rate limit, please slow down.
//
// - 5XX status codes indicate internal errors; the issue is on Binance's side.
//
// Important: If a response contains 5xx status code, it does not necessarily mean that your request has failed. Execution status is unknown and the request might have actually succeeded.
// Please use query methods to confirm the status.
// You might also want to establish a new WebSocket connection for that.
type WsApiResponse struct {
	Id         string          `json:"id"`
	Status     int             `json:"status"`
	Error      *WsApiError     `json:"error"`
	Result     json.RawMessage `json:"result"`
	RateLimits []RateLimit     `json:"rateLimits"`
}

type WsApiError struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data json.RawMessage `json:"data"`
}

////
// Request
////

func doWsApiRequest[T any](s *WebsocketAPI, method string, params map[string]interface{}, apikey KeyPair, securityType SecurityType) (*T, *WsApiResponse, Error) {
	payload := make(map[string]interface{})

	var useApikey, isSigned bool

	switch securityType {
	case NONE:
		useApikey = false
		isSigned = false
	case TRADE:
		useApikey = true
		isSigned = true
	case USER_STREAM:
		useApikey = true
		isSigned = false
	case USER_DATA:
		useApikey = true
		isSigned = true
	}

	if useApikey {
		params["apiKey"] = apikey.ApiKey()
	}

	if isSigned {
		params["timestamp"] = time.Now().UnixMilli()
		paramString := requests.CreateQueryString(params, true)
		signature, err := apikey.Sign(paramString)
		if err != nil {
			return nil, nil, berror.NewSignatureError(err)
		}

		params["signature"] = signature
	}

	payload["method"] = method
	payload["params"] = params

	data, err := s.base.SendRequest(payload, websocketRequestTimeout)
	if err != nil {
		return nil, nil, err
	}

	var response *WsApiResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, response, berror.NewParseError(err)
	}

	if response.Status == 200 {
		var result T
		if err := json.Unmarshal(response.Result, &result); err != nil {
			return nil, response, berror.NewParseError(err)
		}

		return &result, response, nil
	}

	return nil, response, berror.NewAPIError(response.Error.Code, response.Error.Msg)
}

////
// Constructor
////

func newWebsocketAPI(baseUrl string, client *Client) (*WebsocketAPI, Error) {
	socket := &WebsocketAPI{
		baseUrl: baseUrl,

		client: client,
		logger: client.logger,
	}

	messageHandler := func(data []byte) {
		client.logger.WARNf("Received message on WebsocketAPI socket when one wasn't expected => %s", string(data))
	}

	base, err := binance.New(baseUrl, messageHandler, nil, nil, client.logger)
	if err != nil {
		return nil, berror.NewNetworkError(err)
	}

	socket.base = base

	return socket, nil
}
