package futures

import (
	"sync"
	"time"
)

/////////////////////////////////////////////////////////////////////////////////

type Ping struct {
	//
	Latency  time.Duration
	Duration time.Duration
}

// # Test connectivity to the Rest API.
//
// Weight: 1
//
// Data Source: Memory
func (c *Client) Ping() (*Ping, Response, Error) {
	httpResp, err := c.MakeRequest(Methods.GET, "/fapi/v1/ping", nil, nil, NONE)
	if err != nil {
		return nil, httpResp, err
	}

	ping := &Ping{
		Latency:  httpResp.Latency(),
		Duration: httpResp.Duration(),
	}

	return ping, httpResp, nil
}

/////////////////////////////////////////////////////////////////////////////////

type serverTime struct {
	ServerTime int64 `json:"serverTime"`
}

type ServerTime struct {
	ServerTime time.Time `json:"serverTime"`
	Latency    time.Duration
	Duration   time.Duration
}

// # Check server time
//
// Test connectivity to the Rest API and get the current server time.
//
// Weight: 1
//
// Data Source: Memory
func (c *Client) ServerTime() (*ServerTime, Response, Error) {
	st, httpResp, err := doRequest[serverTime](c, Methods.GET, "/fapi/v1/time", nil, nil, NONE)
	if err != nil {
		return nil, httpResp, err
	}

	serverTime := &ServerTime{
		ServerTime: time.UnixMilli(st.ServerTime),
		Latency:    httpResp.Latency(),
		Duration:   httpResp.Duration(),
	}

	return serverTime, httpResp, nil
}

/////////////////////////////////////////////////////////////////////////////////
// EXCHANGE INFO AND EXCHANGE SYMBOLS
/////////////////////////////////////////////////////////////////////////////////

type ExchangeInfo struct {
	Timezone   string       `json:"timezone"`
	RateLimits []*RateLimit `json:"rateLimits"`
	ServerTime int64        `json:"serverTime"`

	ExchangeFilters ExchangeFilters `json:"exchangeFilters"`

	Assets_arr []*Asset `json:"assets"`
	Assets     struct {
		Mu  sync.Mutex
		Map map[string]*Asset
	}

	Symbols_arr []*Symbol `json:"symbols"`
	Symbols     struct {
		Mu  sync.Mutex
		Map map[string]*Symbol
	}
}

// IsRateLimitHit reports whether the specified rate limit rule has been reached.
//
// The method searches ExchangeInfo.RateLimits for a RateLimit rule matching
// the provided parameters (rateLimitType, interval, and intervalNum). If a
// matching rule is found, its Limit is compared against the provided value.
// If value is greater than or equal to the rule's limit, the method returns true.
//
// Parameters:
//   - value: the current usage value to compare against the rate limit
//   - rateLimitType: the type of rate limit to check (e.g., REQUEST_WEIGHT)
//   - interval: the time interval unit for the rule (e.g., SECOND, MINUTE)
//   - intervalNum: the number of interval units defining the rate window
//
// Returns true if the rate limit has been reached or exceeded, otherwise false.
func (ei *ExchangeInfo) IsRateLimitHit(value int, rateLimitType RateLimitType, interval RateLimitInterval, intervalNum int) bool {
	for _, rl := range ei.RateLimits {
		if rl.RateLimitType != rateLimitType {
			continue
		}

		if rl.Interval != interval {
			continue
		}

		if rl.IntervalNum != intervalNum {
			continue
		}

		return value >= rl.Limit
	}

	return false
}

// RateLimits

type RateLimit struct {
	RateLimitType RateLimitType     `json:"rateLimitType"`
	Interval      RateLimitInterval `json:"interval"`
	IntervalNum   int               `json:"intervalNum"`
	Limit         int               `json:"limit"`
}

// ExchangeFilters

type ExchangeFilters struct {
}

// Asset
type Asset struct {
	Asset string `json:"asset"`
	// whether the asset can be used as margin in Multi-Assets mode
	MarginAvailable bool `json:"marginAvailable"`
	// auto-exchange threshold in Multi-Assets margin mode
	AutoAssetExchange float64 `json:"autoAssetExchange,string"`
}

// Symbol

type Symbol struct {
	Symbol       string         `json:"symbol"`
	Pair         string         `json:"pair"`
	ContractType ContractType   `json:"contractType"`
	DeliveryDate int64          `json:"deliveryDate"`
	OnboardDate  int64          `json:"onboardDate"`
	Status       ContractStatus `json:"status"`

	// ignore
	MaintMarginPercent float64 `json:"maintMarginPercent,string"`
	// ignore
	RequiredMarginPercent float64 `json:"requiredMarginPercent,string"`

	BaseAsset   string `json:"baseAsset"`
	QuoteAsset  string `json:"quoteAsset"`
	MarginAsset string `json:"marginAsset"`

	// please do not use it as tickSize
	PricePrecision int `json:"pricePrecision"`
	// please do not use it as stepSize
	QuantityPrecision  int `json:"quantityPrecision"`
	BaseAssetPrecision int `json:"baseAssetPrecision"`
	QuotePrecision     int `json:"quotePrecision"`

	UnderlyingCoin    string   `json:"underlyingCoin"`
	UnderlyingType    string   `json:"underlyingType"`
	UnderlyingSubType []string `json:"underlyingSubType"`

	SettlePlan int `json:"settlePlan"`

	// threshold for algo order with "priceProtect"
	TriggerProtect float64 `json:"triggerProtect,string"`

	Filters         SymbolFilters
	OrderTypes      []OrderType   `json:"orderTypes"`
	TimeInForce     []TimeInForce `json:"timeInForce"`
	LiquidationFee  float64       `json:"liquidationFee,string"`
	MarketTakeBound float64       `json:"marketTakeBound,string"`
}

type SymbolFilters struct {
	PRICE_FILTER        *SymbolFilterPriceFilter
	LOT_SIZE            *SymbolFilterLotSize
	MARKET_LOT_SIZE     *SymbolFilterMarketLotSize
	MAX_NUM_ORDERS      *SymbolFilterMaxNumOrders
	MAX_NUM_ALGO_ORDERS *SymbolFilterMaxNumAlgoOrders
	PERCENT_PRICE       *SymbolFilterPercentPrice
	MIN_NOTIONAL        *SymbolFilterMinNotional
}

type SymbolFilterPriceFilter struct {
	FilterType SymbolFilterType `json:"filterType"`
	MinPrice   float64          `json:"minPrice,string"`
	MaxPrice   float64          `json:"maxPrice,string"`
	TickSize   float64          `json:"tickSize,string"`
}

type SymbolFilterLotSize struct {
	FilterType SymbolFilterType `json:"filterType"`
	MinQty     float64          `json:"minQty,string"`
	MaxQty     float64          `json:"maxQty,string"`
	StepSize   float64          `json:"stepSize,string"`
}

type SymbolFilterMarketLotSize struct {
	FilterType SymbolFilterType `json:"filterType"`
	MinQty     float64          `json:"minQty,string"`
	MaxQty     float64          `json:"maxQty,string"`
	StepSize   float64          `json:"stepSize,string"`
}

type SymbolFilterMaxNumOrders struct {
	FilterType   SymbolFilterType `json:"filterType"`
	MaxNumOrders int              `json:"maxNumOrders"`
}

type SymbolFilterMaxNumAlgoOrders struct {
	FilterType       SymbolFilterType `json:"filterType"`
	MaxNumAlgoOrders int              `json:"maxNumAlgoOrders"`
}

type SymbolFilterPercentPrice struct {
	FilterType     SymbolFilterType `json:"filterType"`
	MultiplierUp   float64          `json:"multiplierUp,string"`
	MultiplierDown float64          `json:"multiplierDown,string"`
	AvgPriceMins   int              `json:"avgPriceMins"`
}

type SymbolFilterMinNotional struct {
	FilterType    SymbolFilterType `json:"filterType"`
	MinNotional   float64          `json:"minNotional"`
	ApplyToMarket bool             `json:"applyToMarket"`
	AvgPriceMins  int              `json:"avgPriceMins"`
}

// # Current exchange trading rules and symbol information
//
// Weight: 20
func (c *Client) ExchangeInfo() (*ExchangeInfo, Response, Error) {
	exchangeInfo, httpResp, err := doRequest[ExchangeInfo](c, Methods.GET, "/fapi/v1/exchangeInfo", nil, nil, NONE)
	if err != nil {
		return nil, httpResp, err
	}

	// Setting up the assets map
	exchangeInfo.Assets.Map = make(map[string]*Asset)
	for _, asset := range exchangeInfo.Assets_arr {
		exchangeInfo.Assets.Map[asset.Asset] = asset
	}

	// Setting up the symbols map
	exchangeInfo.Symbols.Map = make(map[string]*Symbol)
	for _, symbol := range exchangeInfo.Symbols_arr {
		exchangeInfo.Symbols.Map[symbol.Symbol] = symbol
	}

	return exchangeInfo, httpResp, err
}
