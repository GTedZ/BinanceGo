package spot

import "github.com/GTedZ/binancego/internal/validation"

// Market Order
type OrderBook struct {
	LastUpdateId int64 `json:"lastUpdateId"`
	//"bids": [
	//    [
	//      "4.00000000",     // PRICE
	//      "431.00000000"    // QTY
	//    ]
	//  ]
	Bids [][2]PriceLevel `json:"bids"`
	// 	"asks": [
	//     [
	//       "4.00000200",
	//       "12.00000000"
	//     ]
	//   ]
	Asks [][2]PriceLevel `json:"asks"`
}

type PriceLevel [2]float64

type OrderBookParams struct {
	Limit        int
	SymbolStatus SymbolStatus
}

func (c *Client) OrderBook(symbol string, opts ...OrderBookParams) (*OrderBook, Response, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "limit", opt.Limit)
		validation.SetIfNotZero(params, "symbolStatus", opt.SymbolStatus)
	}

	return doRequest[OrderBook](c, Methods.GET, "/api/v3/depth", params, nil, NONE)
}

// Trade
type Trade struct {
	Id           int64   `json:"id"`
	Price        float64 `json:"price,string"`
	Qty          float64 `json:"qty,string"`
	QuoteQty     float64 `json:"quoteQty,string"`
	Time         int64   `json:"time"`
	IsBuyerMaker bool    `json:"isBuyerMaker"`
	IsBestMatch  bool    `json:"isBestMatch"`
}

// Recent Trades

type RecentTradesParams struct {
	Limit int
}

func (c *Client) RecentTrades(symbol string, opts ...RecentTradesParams) ([]*Trade, Response, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "limit", opt.Limit)
	}

	result, response, err := doRequest[[]*Trade](c, Methods.GET, "/api/v3/trades", params, nil, NONE)
	if result == nil {
		return nil, response, err
	}
	return *result, response, err
}

// Historical Trades

type HistoricalTradesParams struct {
	Limit  int
	FromId int
}

func (c *Client) HistoricalTrades(symbol string, opts ...HistoricalTradesParams) ([]*Trade, Response, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "limit", opt.Limit)
		validation.SetIfNotZero(params, "fromId", opt.FromId)
	}

	result, response, err := doRequest[[]*Trade](c, Methods.GET, "/api/v3/historicalTrades", params, nil, NONE)
	if result == nil {
		return nil, response, err
	}
	return *result, response, err
}

// Aggregate Trades

type AggTrade struct {
	// Aggregate tradeId
	AggTradeId int64 `json:"a"`
	// Price
	Price float64 `json:"p,string"`
	// Quantity
	Quantity float64 `json:"q,string"`
	// First tradeId
	FirstTradeId int64 `json:"f"`
	// Last tradeId
	LastTradeId int64 `json:"l"`
	// Timestamp
	Timestamp int64 `json:"T"`
	// Was the buyer the maker?
	IsMaker bool `json:"m"`
	// Was the trade the best price match?
	IsBestMatch bool `json:"M"`
}

type AggTradeParams struct {
	FromId    int64
	StartTime int64
	EndTime   int64
	Limit     int
}

func (c *Client) AggTrade(symbol string, opts ...AggTradeParams) ([]*AggTrade, Response, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "limit", opt.Limit)
		validation.SetIfNotZero(params, "startTime", opt.StartTime)
		validation.SetIfNotZero(params, "endTime", opt.EndTime)
		validation.SetIfNotZero(params, "fromId", opt.FromId)
	}

	result, resp, err := doRequest[[]*AggTrade](c, Methods.GET, "/api/v3/aggTrades", params, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

// Candlesticks

type Candlestick struct {
	OpenTime int64

	Open  float64
	High  float64
	Low   float64
	Close float64

	Volume float64

	CloseTime int64

	QuoteAssetVolume float64
	TradeCount       int64

	TakerBuyBaseAssetVolume  float64
	TakerBuyQuoteAssetVolume float64

	Unused string
}

type CandlestickParams struct {
	StartTime int64
	EndTime   int64
	TimeZone  string
	Limit     int
}

func (c *Client) Candlesticks(symbol string, interval KlineInterval, opts ...CandlestickParams) ([]*Candlestick, Response, Error) {
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

	result, resp, err := doRequest[[]*Candlestick](c, Methods.GET, "/api/v3/klines", params, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

// UI Klines

func (c *Client) UIKlines(symbol string, interval KlineInterval, opts ...CandlestickParams) ([]*Candlestick, Response, Error) {
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

	result, resp, err := doRequest[[]*Candlestick](c, Methods.GET, "/api/v3/uiKlines", params, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

// Average Price

type AveragePrice struct {
	// Average price interval (in minutes)
	Mins int64 `json:"mins"`
	// Average price
	Price string `json:"price"`
	// Last trade time
	CloseTime int64 `json:"closeTime"`
}

func (c *Client) AveragePrice(symbol string) (*AveragePrice, Response, Error) {
	return doRequest[AveragePrice](c, Methods.GET, "/api/v3/avgPrice", nil, nil, NONE)
}

// Ticker 24h

type Ticker24h struct {
	Symbol string `json:"symbol"`

	PriceChange        float64 `json:"priceChange,string"`
	PriceChangePercent float64 `json:"priceChangePercent,string"`
	WeightedAvgPrice   float64 `json:"weightedAvgPrice,string"`

	PrevClosePrice float64 `json:"prevClosePrice,string"`

	LastPrice float64 `json:"lastPrice,string"`
	LastQty   float64 `json:"lastQty,string"`

	BidPrice float64 `json:"bidPrice,string"`
	BidQty   float64 `json:"bidQty,string"`

	AskPrice float64 `json:"askPrice,string"`
	AskQty   float64 `json:"askQty,string"`

	OpenPrice float64 `json:"openPrice,string"`
	HighPrice float64 `json:"highPrice,string"`
	LowPrice  float64 `json:"lowPrice,string"`

	Volume      float64 `json:"volume,string"`
	QuoteVolume float64 `json:"quoteVolume,string"`

	OpenTime  int64 `json:"openTime"`
	CloseTime int64 `json:"closeTime"`

	FirstID int64 `json:"firstId"`
	LastID  int64 `json:"lastId"`

	Count int64 `json:"count"`
}

type Tickers24hParams struct {
	Symbols      []string
	SymbolStatus SymbolStatus
}

func (c *Client) Tickers24h(opts ...Tickers24hParams) ([]*Ticker24h, Response, Error) {
	params := make(map[string]interface{})

	params["type"] = "FULL"

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "symbols", opt.Symbols)
		validation.SetIfNotZero(params, "symbolStatus", opt.SymbolStatus)
	}

	result, resp, err := doRequest[[]*Ticker24h](c, Methods.GET, "/api/v3/ticker/24hr", params, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

type Ticker24hParams struct {
	SymbolStatus SymbolStatus
}

func (c *Client) Ticker24h(symbol string, opts ...Ticker24hParams) (*Ticker24h, Response, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol
	params["type"] = "FULL"

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "symbolStatus", opt.SymbolStatus)
	}

	return doRequest[Ticker24h](c, Methods.GET, "/api/v3/ticker/24hr", params, nil, NONE)
}

// MiniTicker 24h

type MiniTicker24h struct {
	Symbol string `json:"symbol"`

	OpenPrice float64 `json:"openPrice,string"`
	HighPrice float64 `json:"highPrice,string"`
	LowPrice  float64 `json:"lowPrice,string"`
	LastPrice float64 `json:"lastPrice,string"`

	Volume      float64 `json:"volume,string"`
	QuoteVolume float64 `json:"quoteVolume,string"`

	OpenTime  int64 `json:"openTime"`
	CloseTime int64 `json:"closeTime"`

	FirstID int64 `json:"firstId"`
	LastID  int64 `json:"lastId"`

	Count int64 `json:"count"`
}

type MiniTickers24hParams struct {
	Symbols      []string
	SymbolStatus SymbolStatus
}

func (c *Client) MiniTickers24h(opts ...MiniTickers24hParams) ([]*MiniTicker24h, Response, Error) {
	params := make(map[string]interface{})

	params["type"] = "MINI"

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "symbols", opt.Symbols)
		validation.SetIfNotZero(params, "symbolStatus", opt.SymbolStatus)
	}

	result, resp, err := doRequest[[]*MiniTicker24h](c, Methods.GET, "/api/v3/ticker/24hr", params, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

type MiniTicker24hParams struct {
	SymbolStatus SymbolStatus
}

func (c *Client) MiniTicker24h(symbol string, opts ...MiniTicker24hParams) (*MiniTicker24h, Response, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol
	params["type"] = "MINI"

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "symbolStatus", opt.SymbolStatus)
	}

	return doRequest[MiniTicker24h](c, Methods.GET, "/api/v3/ticker/24hr", params, nil, NONE)
}

// Trading Day Ticker

type TradingDayTicker struct {
	Symbol string `json:"symbol"`

	PriceChange        float64 `json:"priceChange,string"`
	PriceChangePercent float64 `json:"priceChangePercent,string"`
	WeightedAvgPrice   float64 `json:"weightedAvgPrice,string"`

	OpenPrice float64 `json:"openPrice,string"`
	HighPrice float64 `json:"highPrice,string"`
	LowPrice  float64 `json:"lowPrice,string"`
	LastPrice float64 `json:"lastPrice,string"`

	Volume      float64 `json:"volume,string"`
	QuoteVolume float64 `json:"quoteVolume,string"`

	OpenTime  int64 `json:"openTime"`
	CloseTime int64 `json:"closeTime"`

	FirstID int64 `json:"firstId"`
	LastID  int64 `json:"lastId"`

	Count int64 `json:"count"`
}

type TradingDayTickersParams struct {
	Symbols      []string
	SymbolStatus SymbolStatus
	TimeZone     string
}

func (c *Client) TradingDayTickers(opts ...TradingDayTickersParams) ([]*TradingDayTicker, Response, Error) {
	params := make(map[string]interface{})

	params["type"] = "FULL"

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "symbols", opt.Symbols)
		validation.SetIfNotZero(params, "symbolStatus", opt.SymbolStatus)
		validation.SetIfNotZero(params, "timeZone", opt.TimeZone)
	}

	result, resp, err := doRequest[[]*TradingDayTicker](c, Methods.GET, "/api/v3/ticker/tradingDay", params, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

type TradingDayTickerParams struct {
	SymbolStatus SymbolStatus
	TimeZone     string
}

func (c *Client) TradingDayTicker(symbol string, opts ...TradingDayTickerParams) (*TradingDayTicker, Response, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol
	params["type"] = "FULL"

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "symbolStatus", opt.SymbolStatus)
		validation.SetIfNotZero(params, "timeZone", opt.TimeZone)
	}

	return doRequest[TradingDayTicker](c, Methods.GET, "/api/v3/ticker/tradingDay", params, nil, NONE)
}

// Trading Day MiniTicker

type TradingDayMiniTicker struct {
	Symbol string `json:"symbol"`

	OpenPrice float64 `json:"openPrice,string"`
	HighPrice float64 `json:"highPrice,string"`
	LowPrice  float64 `json:"lowPrice,string"`
	LastPrice float64 `json:"lastPrice,string"`

	Volume      float64 `json:"volume,string"`
	QuoteVolume float64 `json:"quoteVolume,string"`

	OpenTime  int64 `json:"openTime"`
	CloseTime int64 `json:"closeTime"`

	FirstID int64 `json:"firstId"`
	LastID  int64 `json:"lastId"`

	Count int64 `json:"count"`
}

type MiniTradingDayTickersParams struct {
	Symbols      []string
	SymbolStatus SymbolStatus
	TimeZone     string
}

func (c *Client) TradingDayMiniTickers(opts ...TradingDayTickersParams) ([]*TradingDayMiniTicker, Response, Error) {
	params := make(map[string]interface{})

	params["type"] = "MINI"

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "symbols", opt.Symbols)
		validation.SetIfNotZero(params, "symbolStatus", opt.SymbolStatus)
		validation.SetIfNotZero(params, "timeZone", opt.TimeZone)
	}

	result, resp, err := doRequest[[]*TradingDayMiniTicker](c, Methods.GET, "/api/v3/ticker/tradingDay", params, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

type MiniTradingDayTickerParams struct {
	SymbolStatus SymbolStatus
	TimeZone     string
}

func (c *Client) TradingDayMiniTicker(symbol string, opts ...TradingDayTickerParams) (*TradingDayMiniTicker, Response, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol
	params["type"] = "MINI"

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "symbolStatus", opt.SymbolStatus)
		validation.SetIfNotZero(params, "timeZone", opt.TimeZone)
	}

	return doRequest[TradingDayMiniTicker](c, Methods.GET, "/api/v3/ticker/tradingDay", params, nil, NONE)
}

// Symbol Price Ticker

type PriceTicker struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price,string"`
}

type PriceTickersParams struct {
	Symbols      []string
	SymbolStatus SymbolStatus
}

func (c *Client) PriceTickers(opts ...PriceTickersParams) ([]*PriceTicker, Response, Error) {
	params := make(map[string]interface{})

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "symbols", opt.Symbols)
		validation.SetIfNotZero(params, "symbolStatus", opt.SymbolStatus)
	}

	result, resp, err := doRequest[[]*PriceTicker](c, Methods.GET, "/api/v3/ticker/price", params, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

type PriceTickerParams struct {
	SymbolStatus SymbolStatus
}

func (c *Client) PriceTicker(symbol string, opts ...PriceTickerParams) (*PriceTicker, Response, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "symbolStatus", opt.SymbolStatus)
	}

	return doRequest[PriceTicker](c, Methods.GET, "/api/v3/ticker/price", params, nil, NONE)
}

// Order Book Ticker

type OrderBookTicker struct {
	Symbol string `json:"symbol"`

	BidPrice float64 `json:"bidPrice,string"`
	BidQty   float64 `json:"bidQty,string"`

	AskPrice float64 `json:"askPrice,string"`
	AskQty   float64 `json:"askQty,string"`
}

type OrderBookTickersParams struct {
	Symbols      []string
	SymbolStatus SymbolStatus
}

func (c *Client) OrderBookTickers(opts ...OrderBookTickersParams) ([]*OrderBookTicker, Response, Error) {
	params := make(map[string]interface{})

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "symbols", opt.Symbols)
		validation.SetIfNotZero(params, "symbolStatus", opt.SymbolStatus)
	}

	result, resp, err := doRequest[[]*OrderBookTicker](c, Methods.GET, "/api/v3/ticker/bookTicker", params, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

type OrderBookTickerParams struct {
	SymbolStatus SymbolStatus
}

func (c *Client) OrderBookTicker(symbol string, opts ...OrderBookTickerParams) (*OrderBookTicker, Response, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "symbolStatus", opt.SymbolStatus)
	}

	return doRequest[OrderBookTicker](c, Methods.GET, "/api/v3/ticker/bookTicker", params, nil, NONE)
}

// Roling Window Ticker

type RollingWindowTicker struct {
	Symbol string `json:"symbol"`

	PriceChange        float64 `json:"priceChange,string"`
	PriceChangePercent float64 `json:"priceChangePercent,string"`
	WeightedAvgPrice   float64 `json:"weightedAvgPrice,string"`

	OpenPrice float64 `json:"openPrice,string"`
	HighPrice float64 `json:"highPrice,string"`
	LowPrice  float64 `json:"lowPrice,string"`
	LastPrice float64 `json:"lastPrice,string"`

	Volume      float64 `json:"volume,string"`
	QuoteVolume float64 `json:"quoteVolume,string"`

	OpenTime  int64 `json:"openTime"`
	CloseTime int64 `json:"closeTime"`

	FirstID int64 `json:"firstId"`
	LastID  int64 `json:"lastId"`

	Count int64 `json:"count"`
}

type RollingWindowTickersParams struct {
	Symbols      []string
	SymbolStatus SymbolStatus
	WindowSize   string
}

func (c *Client) RollingWindowTickers(opts ...RollingWindowTickersParams) ([]*RollingWindowTicker, Response, Error) {
	params := make(map[string]interface{})

	params["type"] = "FULL"

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "symbols", opt.Symbols)
		validation.SetIfNotZero(params, "symbolStatus", opt.SymbolStatus)
		validation.SetIfNotZero(params, "windowSize", opt.WindowSize)
	}

	result, resp, err := doRequest[[]*RollingWindowTicker](c, Methods.GET, "/api/v3/ticker", params, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

type RollingWindowTickerParams struct {
	SymbolStatus SymbolStatus
	WindowSize   string
}

func (c *Client) RollingWindowTicker(symbol string, opts ...RollingWindowTickerParams) (*RollingWindowTicker, Response, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol
	params["type"] = "FULL"

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "symbolStatus", opt.SymbolStatus)
		validation.SetIfNotZero(params, "windowSize", opt.WindowSize)
	}

	return doRequest[RollingWindowTicker](c, Methods.GET, "/api/v3/ticker", params, nil, NONE)
}

// Rolling Window MiniTicker

type RollingWindowMiniTicker struct {
	Symbol string `json:"symbol"`

	OpenPrice float64 `json:"openPrice,string"`
	HighPrice float64 `json:"highPrice,string"`
	LowPrice  float64 `json:"lowPrice,string"`
	LastPrice float64 `json:"lastPrice,string"`

	Volume      float64 `json:"volume,string"`
	QuoteVolume float64 `json:"quoteVolume,string"`

	OpenTime  int64 `json:"openTime"`
	CloseTime int64 `json:"closeTime"`

	FirstID int64 `json:"firstId"`
	LastID  int64 `json:"lastId"`

	Count int64 `json:"count"`
}

type RollingWindowMiniTickersParams struct {
	Symbols      []string
	SymbolStatus SymbolStatus
	WindowSize   string
}

func (c *Client) RollingWindowMiniTickers(opts ...RollingWindowMiniTickersParams) ([]*RollingWindowMiniTicker, Response, Error) {
	params := make(map[string]interface{})

	params["type"] = "MINI"

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "symbols", opt.Symbols)
		validation.SetIfNotZero(params, "symbolStatus", opt.SymbolStatus)
		validation.SetIfNotZero(params, "windowSize", opt.WindowSize)
	}

	result, resp, err := doRequest[[]*RollingWindowMiniTicker](c, Methods.GET, "/api/v3/ticker", params, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

type RollingWindowTickerMiniParams struct {
	SymbolStatus SymbolStatus
	WindowSize   string
}

func (c *Client) RollingWindowMiniTicker(symbol string, opts ...RollingWindowTickerMiniParams) (*RollingWindowMiniTicker, Response, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol
	params["type"] = "MINI"

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "symbolStatus", opt.SymbolStatus)
		validation.SetIfNotZero(params, "windowSize", opt.WindowSize)
	}

	return doRequest[RollingWindowMiniTicker](c, Methods.GET, "/api/v3/ticker", params, nil, NONE)
}

// Reference Price

type ReferencePrice struct {
	Symbol string `json:"symbol"`

	ReferencePrice *float64 `json:"referencePrice,string"`

	Timestamp int64 `json:"timestamp"`
}

func (r ReferencePrice) HasPrice() bool {
	return r.ReferencePrice != nil
}

func (c *Client) ReferencePrice(symbol string) (*ReferencePrice, Response, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol

	return doRequest[ReferencePrice](c, Methods.GET, "/api/v3/referencePrice", params, nil, NONE)
}

// Reference Price Calculations

type ReferencePriceCalculation struct {
	Symbol string `json:"symbol"`

	CalculationType ReferencePriceCalculationType `json:"calculationType"`

	// ARITHMETIC_MEAN
	BucketCount   *int64 `json:"bucketCount,omitempty"`
	BucketWidthMs *int64 `json:"bucketWidthMs,omitempty"`

	// EXTERNAL
	ExternalCalculationID *int64 `json:"externalCalculationId,omitempty"`
}

type ReferencePriceCalculationParams struct {
	SymbolStatus SymbolStatus
}

func (r ReferencePriceCalculation) IsArithmeticMean() bool {
	return r.CalculationType == ArithmeticMean
}

func (r ReferencePriceCalculation) ArithmeticMeanParams() (bucketCount, bucketWidthMs int64, ok bool) {
	if !r.IsArithmeticMean() || r.BucketCount == nil || r.BucketWidthMs == nil {
		return 0, 0, false
	}
	return *r.BucketCount, *r.BucketWidthMs, true
}

func (r ReferencePriceCalculation) IsExternal() bool {
	return r.CalculationType == External
}

func (r ReferencePriceCalculation) ExternalID() (int64, bool) {
	if !r.IsExternal() || r.ExternalCalculationID == nil {
		return 0, false
	}
	return *r.ExternalCalculationID, true
}

func (c *Client) ReferencePriceCalculation(symbol string, opts ...ReferencePriceCalculationParams) (*ReferencePriceCalculation, Response, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol

	if len(opts) > 0 {
		opt := opts[0]
		validation.SetIfNotZero(params, "symbolStatus", opt.SymbolStatus)
	}

	return doRequest[ReferencePriceCalculation](c, Methods.GET, "/api/v3/referencePrice/calculation", params, nil, NONE)
}
