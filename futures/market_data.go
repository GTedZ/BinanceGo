package futures

import "github.com/GTedZ/binancego/internal/validation"

// Market Order
type OrderBook struct {
	LastUpdateId int64 `json:"lastUpdateId"`
	// Message output time
	Timestamp int64 `json:"E"`
	// Transaction time
	TransactTime int64 `json:"T"`
	//"bids": [
	//    [
	//      "4.00000000",     // PRICE
	//      "431.00000000"    // QTY
	//    ]
	//  ]
	Bids []PriceLevel `json:"bids"`
	// 	"asks": [
	//     [
	//       "4.00000200",
	//       "12.00000000"
	//     ]
	//   ]
	Asks []PriceLevel `json:"asks"`
}

type PriceLevel struct {
	Price float64
	Qty   float64
}

type OrderBookParams struct {
	Limit int
}

func (c *Client) OrderBook(symbol string, opts ...OrderBookParams) (*OrderBook, Response, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "limit", opt.Limit)
	}

	return doRequest[OrderBook](c, Methods.GET, "/fapi/v1/depth", params, nil, NONE)
}

func (c *Client) RPIOrderBook(symbol string, opts ...OrderBookParams) (*OrderBook, Response, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "limit", opt.Limit)
	}

	return doRequest[OrderBook](c, Methods.GET, "/fapi/v1/rpiDepth", params, nil, NONE)
}

// Trade
type Trade struct {
	Id           int64   `json:"id"`
	Price        float64 `json:"price,string"`
	Qty          float64 `json:"qty,string"`
	QuoteQty     float64 `json:"quoteQty,string"`
	Time         int64   `json:"time"`
	IsBuyerMaker bool    `json:"isBuyerMaker"`
	IsRPITrade   bool    `json:"isRPITrade"`
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

	result, response, err := doRequest[[]*Trade](c, Methods.GET, "/fapi/v1/trades", params, nil, NONE)
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

	result, response, err := doRequest[[]*Trade](c, Methods.GET, "/fapi/v1/historicalTrades", params, nil, NONE)
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
	// Normal quantity without the trades involving RPI orders
	NormalQuantity float64 `json:"nq,string"`
	// First tradeId
	FirstTradeId int64 `json:"f"`
	// Last tradeId
	LastTradeId int64 `json:"l"`
	// Timestamp
	Timestamp int64 `json:"T"`
	// Was the buyer the maker?
	IsMaker bool `json:"m"`
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

	result, resp, err := doRequest[[]*AggTrade](c, Methods.GET, "/fapi/v1/aggTrades", params, nil, NONE)
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
	Limit     int
}

func (c *Client) Candlesticks(symbol string, interval Interval, opts ...CandlestickParams) ([]*Candlestick, Response, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol
	params["interval"] = interval

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "startTime", opt.StartTime)
		validation.SetIfNotZero(params, "endTime", opt.EndTime)
		validation.SetIfNotZero(params, "limit", opt.Limit)
	}

	result, resp, err := doRequest[[]*Candlestick](c, Methods.GET, "/fapi/v1/klines", params, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

// Continuous Klines

func (c *Client) ContinuousKlines(symbol string, interval Interval, contractType ContractType, opts ...CandlestickParams) ([]*Candlestick, Response, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol
	params["interval"] = interval
	params["contractType"] = contractType

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "startTime", opt.StartTime)
		validation.SetIfNotZero(params, "endTime", opt.EndTime)
		validation.SetIfNotZero(params, "limit", opt.Limit)
	}

	result, resp, err := doRequest[[]*Candlestick](c, Methods.GET, "/fapi/v1/continuousKlines", params, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

// Index Price Klines

func (c *Client) IndexPriceKlines(symbol string, interval Interval, opts ...CandlestickParams) ([]*Candlestick, Response, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol
	params["interval"] = interval

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "startTime", opt.StartTime)
		validation.SetIfNotZero(params, "endTime", opt.EndTime)
		validation.SetIfNotZero(params, "limit", opt.Limit)
	}

	result, resp, err := doRequest[[]*Candlestick](c, Methods.GET, "/fapi/v1/indexPriceKlines", params, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

// Mark Price Klines

func (c *Client) MarkPriceKlines(symbol string, interval Interval, opts ...CandlestickParams) ([]*Candlestick, Response, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol
	params["interval"] = interval

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "startTime", opt.StartTime)
		validation.SetIfNotZero(params, "endTime", opt.EndTime)
		validation.SetIfNotZero(params, "limit", opt.Limit)
	}

	result, resp, err := doRequest[[]*Candlestick](c, Methods.GET, "/fapi/v1/markPriceKlines", params, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

// Premium Index Klines

func (c *Client) PremiumIndexKlines(symbol string, interval Interval, opts ...CandlestickParams) ([]*Candlestick, Response, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol
	params["interval"] = interval

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "startTime", opt.StartTime)
		validation.SetIfNotZero(params, "endTime", opt.EndTime)
		validation.SetIfNotZero(params, "limit", opt.Limit)
	}

	result, resp, err := doRequest[[]*Candlestick](c, Methods.GET, "/fapi/v1/premiumIndexKlines", params, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

// Mark Price

type MarkPrice struct {
	Symbol string `json:"symbol"`
	// mark price
	MarkPrice float64 `json:"markPrice,string"`
	// index price
	IndexPrice float64 `json:"indexPrice,string"`
	// Estimated Settle Price, only useful in the last hour before the settlement starts.
	EstimatedSettlePrice float64 `json:"estimatedSettlePrice,string"`
	// This is the Latest funding rate
	LastFundingRate float64 `json:"lastFundingRate,string"`
	InterestRate    float64 `json:"interestRate,string"`
	NextFundingTime int64   `json:"nextFundingTime"`
	Time            int64   `json:"time"`
}

func (c *Client) AllMarkPrices() ([]*MarkPrice, Response, Error) {
	result, resp, err := doRequest[[]*MarkPrice](c, Methods.GET, "/fapi/v1/premiumIndex", nil, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

func (c *Client) MarkPrice(symbol string) (*MarkPrice, Response, Error) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	return doRequest[MarkPrice](c, Methods.GET, "/fapi/v1/premiumIndex", params, nil, NONE)
}

// Funding Rate

type FundingRate struct {
	Symbol      string  `json:"symbol"`
	FundingRate float64 `json:"fundingRate,string"`
	FundingTime int64   `json:"fundingTime"`
	// mark price associated with a particular funding fee charge
	MarkPrice float64 `json:"markPrice,string"`
}

type FundingRateHistoryParams struct {
	StartTime int64
	EndTime   int64
	Limit     int
}

func (c *Client) AllFundingRateHistory(opts ...FundingRateHistoryParams) ([]*FundingRate, Response, Error) {
	params := make(map[string]interface{})

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "startTime", opt.StartTime)
		validation.SetIfNotZero(params, "endTime", opt.EndTime)
		validation.SetIfNotZero(params, "limit", opt.Limit)
	}

	result, resp, err := doRequest[[]*FundingRate](c, Methods.GET, "/fapi/v1/fundingRate", params, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

func (c *Client) FundingRateHistory(symbol string, opts ...FundingRateHistoryParams) ([]*FundingRate, Response, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "startTime", opt.StartTime)
		validation.SetIfNotZero(params, "endTime", opt.EndTime)
		validation.SetIfNotZero(params, "limit", opt.Limit)
	}
	result, resp, err := doRequest[[]*FundingRate](c, Methods.GET, "/fapi/v1/fundingRate", params, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

// Funding Rate Info

type FundingRateInfo struct {
	Symbol                   string  `json:"symbol"`
	AdjustedFundingRateCap   float64 `json:"adjustedFundingRateCap,string"`
	AdjustedFundingRateFloor float64 `json:"adjustedFundingRateFloor,string"`
	FundingIntervalHours     int     `json:"fundingIntervalHours"`
	// ignore
	Disclaimer bool `json:"disclaimer"`
}

func (c *Client) FundingRateInfo() ([]*FundingRateInfo, Response, Error) {
	result, resp, err := doRequest[[]*FundingRateInfo](c, Methods.GET, "/fapi/v1/fundingInfo", nil, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
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

func (c *Client) Tickers24h() ([]*Ticker24h, Response, Error) {
	result, resp, err := doRequest[[]*Ticker24h](c, Methods.GET, "/fapi/v1/ticker/24hr", nil, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

func (c *Client) Ticker24h(symbol string) (*Ticker24h, Response, Error) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	return doRequest[Ticker24h](c, Methods.GET, "/fapi/v1/ticker/24hr", params, nil, NONE)
}

// Symbol Price Ticker

type PriceTicker struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price,string"`
	Time   int64   `json:"time"`
}

func (c *Client) PriceTickers() ([]*PriceTicker, Response, Error) {
	result, resp, err := doRequest[[]*PriceTicker](c, Methods.GET, "/fapi/v1/ticker/price", nil, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

func (c *Client) PriceTicker(symbol string) (*PriceTicker, Response, Error) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	return doRequest[PriceTicker](c, Methods.GET, "/fapi/v1/ticker/price", params, nil, NONE)
}

// Symbol Price Ticker V2

type PriceTickerV2 struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price,string"`
	Time   int64   `json:"time"`
}

// # PriceTickersV2
//  The field `X-MBX-USED-WEIGHT-1M` in response header is not accurate from this endpoint, please ignore.
func (c *Client) PriceTickersV2() ([]*PriceTickerV2, Response, Error) {
	result, resp, err := doRequest[[]*PriceTickerV2](c, Methods.GET, "/fapi/v2/ticker/price", nil, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

// # PriceTickerV2
//  The field `X-MBX-USED-WEIGHT-1M` in response header is not accurate from this endpoint, please ignore.
func (c *Client) PriceTickerV2(symbol string) (*PriceTickerV2, Response, Error) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	return doRequest[PriceTickerV2](c, Methods.GET, "/fapi/v2/ticker/price", params, nil, NONE)
}

// Order Book Ticker

type OrderBookTicker struct {
	Symbol string `json:"symbol"`

	BidPrice float64 `json:"bidPrice,string"`
	BidQty   float64 `json:"bidQty,string"`

	AskPrice float64 `json:"askPrice,string"`
	AskQty   float64 `json:"askQty,string"`

	Time int64 `json:"time"`
}

func (c *Client) OrderBookTickers() ([]*OrderBookTicker, Response, Error) {
	result, resp, err := doRequest[[]*OrderBookTicker](c, Methods.GET, "/fapi/v1/ticker/bookTicker", nil, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

func (c *Client) OrderBookTicker(symbol string) (*OrderBookTicker, Response, Error) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	return doRequest[OrderBookTicker](c, Methods.GET, "/fapi/v1/ticker/bookTicker", params, nil, NONE)
}

// Delivery Price

type DeliveryPrice struct {
	DeliveryPrice float64 `json:"deliveryPrice,string"`
	DeliveryTime  int64   `json:"deliveryTime"`
}

func (c *Client) DeliveryPrice(pair string) ([]*DeliveryPrice, Response, Error) {
	params := make(map[string]interface{})
	params["pair"] = pair
	result, resp, err := doRequest[[]*DeliveryPrice](c, Methods.GET, "/futures/data/delivery-price", params, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

// Open Interest

type OpenInterest struct {
	Symbol       string  `json:"symbol"`
	OpenInterest float64 `json:"openInterest,string"`
	// Transact Time
	Time int64 `json:"time"`
}

func (c *Client) OpenInterest(symbol string) (*OpenInterest, Response, Error) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	return doRequest[OpenInterest](c, Methods.GET, "/fapi/v1/openInterest", params, nil, NONE)
}

// Open Interest Statistics

type OpenInterestStatistics struct {
	Symbol               string  `json:"symbol"`
	SumOpenInterest      float64 `json:"sumOpenInterest,string"`
	SumOpenInterestValue float64 `json:"sumOpenInterestValue,string"`
	CMCCirculatingSupply float64 `json:"CMCCirculatingSupply,string"`
	Timestamp            int64   `json:"timestamp"`
}

type OpenInterestStatisticsParams struct {
	Limit     int
	StartTime int64
	EndTime   int64
}

func (c *Client) OpenInterestStatistics(symbol string, period Period, opts ...OpenInterestStatisticsParams) ([]*OpenInterestStatistics, Response, Error) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	params["period"] = period

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "limit", opt.Limit)
		validation.SetIfNotZero(params, "startTime", opt.StartTime)
		validation.SetIfNotZero(params, "endTime", opt.EndTime)
	}

	result, resp, err := doRequest[[]*OpenInterestStatistics](c, Methods.GET, "/futures/data/openInterestHist", params, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

// Long/Short Position Ratio

type LongShortPositionRatio struct {
	Symbol string `json:"symbol"`
	// long/short position ratio of top traders
	LongShortRatio float64 `json:"longShortRatio,string"`
	// long positions ratio of top traders
	LongAccount float64 `json:"longAccount,string"`
	// short positions ratio of top traders
	ShortAccount float64 `json:"shortAccount,string"`
	Time         int64   `json:"time"`
}

type TopTraderLongShortPositionRatioParams struct {
	Limit     int
	StartTime int64
	EndTime   int64
}

// TopTraderLongShortPositionRatio
//  The proportion of net long and net short positions relative to the total open positions of the top 20% of users with the highest margin balances.
//  `Long Position %` = `Long positions of top traders / Total open positions of top traders`.
//  `Short Position %` = `Short positions of top traders / Total open positions of top traders`.
//  `Long/Short Ratio (Positions)` = `Long Position % / Short Position %`.
func (c *Client) TopTraderLongShortPositionRatio(symbol string, period Period, opts ...TopTraderLongShortPositionRatioParams) ([]*LongShortPositionRatio, Response, Error) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	params["period"] = period

	if len(opts) > 0 {
		opt := opts[0]
		validation.SetIfNotZero(params, "limit", opt.Limit)
		validation.SetIfNotZero(params, "startTime", opt.StartTime)
		validation.SetIfNotZero(params, "endTime", opt.EndTime)
	}

	result, resp, err := doRequest[[]*LongShortPositionRatio](c, Methods.GET, "/futures/data/topLongShortPositionRatio", params, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

type TopTraderLongShortAccountRatioParams struct {
	Limit     int
	StartTime int64
	EndTime   int64
}

// TopTraderLongShortAccountRatio
//  The proportion of net long and net short accounts relative to the total accounts of the top 20% of users with the highest margin balances. Each account is counted only once.
//  `Long Account %` = `Accounts of top traders with net long positions / Total accounts of top traders with open positions`.
//  `Short Account %` = `Accounts of top traders with net short positions / Total accounts of top traders with open positions`.
//  `Long/Short Ratio (Accounts)` = `Long Account % / Short Account %`.
func (c *Client) TopTraderLongShortAccountRatio(symbol string, period Period, opts ...TopTraderLongShortAccountRatioParams) ([]*LongShortPositionRatio, Response, Error) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	params["period"] = period

	if len(opts) > 0 {
		opt := opts[0]
		validation.SetIfNotZero(params, "limit", opt.Limit)
		validation.SetIfNotZero(params, "startTime", opt.StartTime)
		validation.SetIfNotZero(params, "endTime", opt.EndTime)
	}

	result, resp, err := doRequest[[]*LongShortPositionRatio](c, Methods.GET, "/futures/data/topLongShortAccountRatio", params, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

type SymbolLongShortRatio struct {
	Limit     int
	StartTime int64
	EndTime   int64
}

func (c *Client) SymbolLongShortRatio(symbol string, period Period, opts ...SymbolLongShortRatio) ([]*LongShortPositionRatio, Response, Error) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	params["period"] = period

	if len(opts) > 0 {
		opt := opts[0]
		validation.SetIfNotZero(params, "limit", opt.Limit)
		validation.SetIfNotZero(params, "startTime", opt.StartTime)
		validation.SetIfNotZero(params, "endTime", opt.EndTime)
	}

	result, resp, err := doRequest[[]*LongShortPositionRatio](c, Methods.GET, "/futures/data/globalLongShortAccountRatio", params, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

// Take Buy/Sell Ratio

type TakerBuySellRatio struct {
	BuySellRatio float64 `json:"buySellRatio,string"`
	BuyVolume    float64 `json:"buyVolume,string"`
	SellVolume   float64 `json:"sellVolume,string"`
	Timestamp    int64   `json:"timestamp"`
}

type TakerBuySellRatioParams struct {
	Limit     int
	StartTime int64
	EndTime   int64
}

func (c *Client) TakerBuySellRatio(symbol string, period Period, opts ...TakerBuySellRatioParams) ([]*TakerBuySellRatio, Response, Error) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	params["period"] = period

	if len(opts) > 0 {
		opt := opts[0]
		validation.SetIfNotZero(params, "limit", opt.Limit)
		validation.SetIfNotZero(params, "startTime", opt.StartTime)
		validation.SetIfNotZero(params, "endTime", opt.EndTime)
	}

	result, resp, err := doRequest[[]*TakerBuySellRatio](c, Methods.GET, "/futures/data/takerlongshortRatio", params, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

// Basis

type Basis struct {
	IndexPrice          float64      `json:"indexPrice,string"`
	ContractType        ContractType `json:"contractType"`
	BasisRate           float64      `json:"basisRate,string"`
	FuturesPrice        float64      `json:"futuresPrice,string"`
	AnnualizedBasisRate float64      `json:"annualizedBasisRate,string"`
	Basis               float64      `json:"basis,string"`
	Pair                string       `json:"pair"`
	Timestamp           int64        `json:"timestamp"`
}

type BasisParams struct {
	Limit     int
	StartTime int64
	EndTime   int64
}

func (c *Client) Basis(pair string, contractType ContractType, period Period, opts ...BasisParams) ([]*Basis, Response, Error) {
	params := make(map[string]interface{})
	params["pair"] = pair
	params["contractType"] = contractType
	params["period"] = period

	if len(opts) > 0 {
		opt := opts[0]
		validation.SetIfNotZero(params, "limit", opt.Limit)
		validation.SetIfNotZero(params, "startTime", opt.StartTime)
		validation.SetIfNotZero(params, "endTime", opt.EndTime)
	}
	result, resp, err := doRequest[[]*Basis](c, Methods.GET, "/futures/data/basis", params, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

// CompositeIndexSymbolInfo

type CompositeIndexSymbolInfo struct {
	Symbol        string `json:"symbol"`
	Time          int64  `json:"time"`
	Component     string `json:"component"`
	BaseAssetList []*CompositeIndexSymbolInfoAsset
}

type CompositeIndexSymbolInfoAsset struct {
	BaseAsset          string  `json:"baseAsset"`
	QuoteAsset         string  `json:"quoteAsset"`
	WeightInQuantity   float64 `json:"weightInQuantity,string"`
	WeightInPercentage float64 `json:"weightInPercentage,string"`
}

func (c *Client) AllCompositeIndexSymbolsInfo() ([]*CompositeIndexSymbolInfo, Response, Error) {
	result, resp, err := doRequest[[]*CompositeIndexSymbolInfo](c, Methods.GET, "/fapi/v1/indexInfo", nil, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

func (c *Client) CompositeIndexSymbolInfo(symbol string) ([]*CompositeIndexSymbolInfo, Response, Error) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	result, resp, err := doRequest[[]*CompositeIndexSymbolInfo](c, Methods.GET, "/fapi/v1/indexInfo", params, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

// Multi-Assets Mode Asset Index

type MultiAssetsModeAssetIndex struct {
	Symbol string  `json:"symbol"`
	Time   int64   `json:"time"`
	Index  float64 `json:"index,string"`

	BidBuffer float64 `json:"bidBuffer,string"`
	AskBuffer float64 `json:"askBuffer,string"`

	BidRate float64 `json:"bidRate,string"`
	AskRate float64 `json:"askRate,string"`

	AutoExchangeBidBuffer float64 `json:"autoExchangeBidBuffer,string"`
	AutoExchangeAskBuffer float64 `json:"autoExchangeAskBuffer,string"`

	AutoExchangeBidRate float64 `json:"autoExchangeBidRate,string"`
	AutoExchangeAskRate float64 `json:"autoExchangeAskRate,string"`
}

func (c *Client) AllMultiAssetsModeAssetIndexes() ([]*MultiAssetsModeAssetIndex, Response, Error) {
	result, resp, err := doRequest[[]*MultiAssetsModeAssetIndex](c, Methods.GET, "/fapi/v1/assetIndex", nil, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

func (c *Client) MultiAssetsModeAssetIndex(symbol string) ([]*MultiAssetsModeAssetIndex, Response, Error) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	result, resp, err := doRequest[[]*MultiAssetsModeAssetIndex](c, Methods.GET, "/fapi/v1/assetIndex", params, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

// Index Price Constituents

type IndexPriceConstituent struct {
	Symbol       string `json:"symbol"`
	Time         int64  `json:"time"`
	Constituents []*IndexPriceConstituentSymbol
}

type IndexPriceConstituentSymbol struct {
	Exchange string  `json:"exchange"`
	Symbol   string  `json:"symbol"`
	Price    float64 `json:"price,string"`
	Weight   float64 `json:"weight,string"`
}

func (c *Client) IndexPriceConstituent(symbol string) ([]*IndexPriceConstituent, Response, Error) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	result, resp, err := doRequest[[]*IndexPriceConstituent](c, Methods.GET, "/fapi/v1/constituents", params, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

// Insurance Fund Balance Snapshot

type InsuranceFundBalanceSnapshot struct {
	Symbols []string `json:"symbols"`
	Assets  []*InsuranceFundBalanceSnapshotAsset
}

type InsuranceFundBalanceSnapshotAsset struct {
	Asset         string  `json:"asset"`
	MarginBalance float64 `json:"marginBalance,string"`
	UpdateTime    int64   `json:"updateTime"`
}

func (c *Client) AllInsuranceFundBalancesSnapshot() (*InsuranceFundBalanceSnapshot, Response, Error) {
	result, resp, err := doRequest[InsuranceFundBalanceSnapshot](c, Methods.GET, "/fapi/v1/insuranceBalance", nil, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return result, resp, err
}

func (c *Client) InsuranceFundBalanceSnapshot(symbol string) (*InsuranceFundBalanceSnapshot, Response, Error) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	result, resp, err := doRequest[InsuranceFundBalanceSnapshot](c, Methods.GET, "/fapi/v1/insuranceBalance", params, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return result, resp, err
}

// ADL Risk Rating

type ADLRiskRating struct {
	Symbol string `json:"symbol"`
	// ADL risk rating, valid values are `1`, `2`, `3`, `4`, `5`. The higher the value, the higher the ADL risk.
	ADLRisk    RiskRating `json:"adlRisk"`
	UpdateTime int64      `json:"updateTime"`
}

// ADLRisk returns the symbol-level Auto-Deleveraging (ADL) risk rating.
//
// The ADL risk rating measures the likelihood of ADL occurring during
// liquidation for a specific symbol. The rating is derived from several
// symbol-level factors, including insurance fund balance, position
// concentration, order book depth, price volatility, average leverage,
// unrealized PnL, and margin utilization.
//
// Possible ratings are:
//   - High
//   - Medium
//   - Low
//
// The rating is updated every 30 minutes.
func (c *Client) AllADLRiskRatings() ([]*ADLRiskRating, Response, Error) {
	result, resp, err := doRequest[[]*ADLRiskRating](c, Methods.GET, "/fapi/v1/symbolAdlRisk", nil, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

// ADLRisk returns the symbol-level Auto-Deleveraging (ADL) risk rating.
//
// The ADL risk rating measures the likelihood of ADL occurring during
// liquidation for a specific symbol. The rating is derived from several
// symbol-level factors, including insurance fund balance, position
// concentration, order book depth, price volatility, average leverage,
// unrealized PnL, and margin utilization.
//
// Possible ratings are:
//   - High
//   - Medium
//   - Low
//
// The rating is updated every 30 minutes.
func (c *Client) ADLRiskRating(symbol string) (*ADLRiskRating, Response, Error) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	result, resp, err := doRequest[ADLRiskRating](c, Methods.GET, "/fapi/v1/symbolAdlRisk", params, nil, NONE)
	if result == nil {
		return nil, resp, err
	}
	return result, resp, err
}

// Trading Schedule

type TradingSchedule struct {
	UpdateTime      int64                          `json:"updateTime"`
	MarketSchedules map[MarketType]*MarketSchedule `json:"marketSchedules"`
}

type MarketSchedule struct {
	Sessions []*MarketScheduleSession `json:"sessions"`
}

type MarketScheduleSession struct {
	StartTime int64             `json:"startTime"`
	EndTime   int64             `json:"endTime"`
	Type      MarketSessionType `json:"type"`
}

// Market Schedules
//
// Trading session schedules for the underlying assets of TradFi Perps.
//
// Schedules are provided for a one-week period forward and a one-week
// period backward starting from the day prior to the query time, covering
// the U.S. equity market, Korean equity market, and commodity market.
//
// Session types by market:
//
// U.S. equity market:
//   - PRE_MARKET
//   - REGULAR
//   - AFTER_MARKET
//   - OVERNIGHT
//   - NO_TRADING
//
// Commodity market:
//   - REGULAR
//   - NO_TRADING
//
// Korean equity market:
//   - REGULAR
//   - NO_TRADING
func (c *Client) TradingSchedule() (*TradingSchedule, Response, Error) {
	return doRequest[TradingSchedule](c, Methods.GET, "/fapi/v1/tradingSchedule", nil, nil, NONE)
}
