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
	return newWebsocketAPI(c.wssApiBaseUrl, c.logger)
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

	return doWsApiRequest[OrderBook](wsapi.client, wsapi, "depth", params, NONE)
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

	result, resp, err := doWsApiRequest[[]*Trade](wsapi.client, wsapi, "trades.recent", params, NONE)
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

	result, resp, err := doWsApiRequest[[]*Trade](wsapi.client, wsapi, "trades.historical", params, NONE)
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

	result, resp, err := doWsApiRequest[[]*AggTrade](wsapi.client, wsapi, "trades.aggregate", params, NONE)
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

	result, resp, err := doWsApiRequest[[]*Candlestick](wsapi.client, wsapi, "klines", params, NONE)
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

	result, resp, err := doWsApiRequest[[]*Candlestick](wsapi.client, wsapi, "uiKlines", params, NONE)
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

	result, resp, err := doWsApiRequest[*AveragePrice](wsapi.client, wsapi, "avgPrice", params, NONE)
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

	result, resp, err := doWsApiRequest[[]*Ticker24h](wsapi.client, wsapi, "ticker.24hr", params, NONE)
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

	return doWsApiRequest[Ticker24h](wsapi.client, wsapi, "ticker.24hr", params, NONE)
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

	result, resp, err := doWsApiRequest[[]*MiniTicker24h](wsapi.client, wsapi, "ticker.24hr", params, NONE)
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

	return doWsApiRequest[MiniTicker24h](wsapi.client, wsapi, "ticker.24hr", params, NONE)
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

	result, resp, err := doWsApiRequest[[]*TradingDayTicker](wsapi.client, wsapi, "ticker.tradingDay", params, NONE)
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

	return doWsApiRequest[TradingDayTicker](wsapi.client, wsapi, "ticker.tradingDay", params, NONE)
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

	result, resp, err := doWsApiRequest[[]*TradingDayMiniTicker](wsapi.client, wsapi, "ticker.tradingDay", params, NONE)
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

	return doWsApiRequest[TradingDayMiniTicker](wsapi.client, wsapi, "ticker.tradingDay", params, NONE)
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

	result, resp, err := doWsApiRequest[[]*RollingWindowTicker](wsapi.client, wsapi, "ticker", params, NONE)
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

	return doWsApiRequest[RollingWindowTicker](wsapi.client, wsapi, "ticker", params, NONE)
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

	result, resp, err := doWsApiRequest[[]*RollingWindowMiniTicker](wsapi.client, wsapi, "ticker", params, NONE)
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

	return doWsApiRequest[RollingWindowMiniTicker](wsapi.client, wsapi, "ticker", params, NONE)
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

	result, resp, err := doWsApiRequest[[]*PriceTicker](wsapi.client, wsapi, "ticker.price", params, NONE)
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

	return doWsApiRequest[PriceTicker](wsapi.client, wsapi, "ticker.price", params, NONE)
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

	result, resp, err := doWsApiRequest[[]*OrderBookTicker](wsapi.client, wsapi, "ticker.book", params, NONE)
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

	return doWsApiRequest[OrderBookTicker](wsapi.client, wsapi, "ticker.book", params, NONE)
}

////
// Reference Price
////

func (wsapi *WebsocketAPI) ReferencePrice(symbol string) (*ReferencePrice, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol

	return doWsApiRequest[ReferencePrice](wsapi.client, wsapi, "referencePrice", params, NONE)
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

	return doWsApiRequest[ReferencePriceCalculation](wsapi.client, wsapi, "referencePrice.calculation", params, NONE)
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

func doWsApiRequest[T any](c *Client, s *WebsocketAPI, method string, params map[string]interface{}, securityType SecurityType) (*T, *WsApiResponse, Error) {
	payload := make(map[string]interface{})

	var apikey, signed bool

	switch securityType {
	case NONE:
		apikey = false
		signed = false
	case TRADE:
		apikey = true
		signed = true
	case USER_DATA:
		apikey = true
		signed = true
	case USER_STREAM:
		apikey = true
		signed = false
	}

	if apikey {
		params["apiKey"] = c.apikey.ApiKey()
	}

	if signed {
		params["timestamp"] = time.Now().UnixMilli()
		paramString := requests.CreateQueryString(params, true)
		signature, err := c.apikey.Sign(paramString)
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

//

func newWebsocketAPI(baseUrl string, logger logging.Logger) (*WebsocketAPI, Error) {
	socket := &WebsocketAPI{
		baseUrl: baseUrl,
		logger:  logger,
	}

	messageHandler := func(data []byte) {
		logger.WARNf("Received message on WebsocketAPI socket when one wasn't expected => %s", string(data))
	}

	base, err := binance.New(baseUrl, messageHandler, nil, nil, logger)
	if err != nil {
		return nil, berror.NewNetworkError(err)
	}

	socket.base = base

	return socket, nil
}
