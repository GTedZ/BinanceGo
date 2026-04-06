package spot

import (
	"sync"
	"time"

	"github.com/GTedZ/binancego/internal/validation"
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
	httpResp, err := c.MakeRequest(Methods.GET, "/api/v3/ping", nil, nil, NONE)
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
	st, httpResp, err := doRequest[serverTime](c, Methods.GET, "/api/v3/time", nil, nil, NONE)
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
	Timezone        string       `json:"timezone"`
	ServerTime      int64        `json:"serverTime"`
	RateLimits      []*RateLimit `json:"rateLimits"`
	ExchangeFilters ExchangeFilters
	Symbols_arr     []*Symbol `json:"symbols"`
	Symbols         struct {
		Mu  sync.Mutex
		Map map[string]*Symbol
	}
	Sors []*SORS `json:"sors"`
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

func (ei *ExchangeInfo) IsOrdersLimitHit(value int) bool {
	if ei.ExchangeFilters.MAX_NUM_ORDERS == nil {
		return false
	}
	return value >= ei.ExchangeFilters.MAX_NUM_ORDERS.MaxNumOrders
}

func (ei *ExchangeInfo) IsAlgoOrdersLimitHit(value int) bool {
	if ei.ExchangeFilters.MAX_NUM_ALGO_ORDERS == nil {
		return false
	}
	return value >= ei.ExchangeFilters.MAX_NUM_ALGO_ORDERS.MaxNumAlgoOrders
}

func (ei *ExchangeInfo) IsIcebergOrdersLimitHit(value int) bool {
	if ei.ExchangeFilters.MAX_NUM_ICEBERG_ORDERS == nil {
		return false
	}
	return value >= ei.ExchangeFilters.MAX_NUM_ICEBERG_ORDERS.MaxNumIcebergOrders
}

func (ei *ExchangeInfo) IsMaxNumOrderLists(value int) bool {
	if ei.ExchangeFilters.MAX_NUM_ORDER_LISTS == nil {
		return false
	}
	return value >= ei.ExchangeFilters.MAX_NUM_ORDER_LISTS.MaxNumOrderLists
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
	MAX_NUM_ORDERS         *ExchangeFilterMaxNumOrders
	MAX_NUM_ALGO_ORDERS    *ExchangeFilterMaxNumAlgoOrders
	MAX_NUM_ICEBERG_ORDERS *ExchangeFilterMaxNumIcebergOrders
	MAX_NUM_ORDER_LISTS    *ExchangeFilterMaxNumOrderLists
}

type ExchangeFilterMaxNumOrders struct {
	FilterType   ExchangeFilterType `json:"filterType"`
	MaxNumOrders int                `json:"maxNumOrders"`
}

type ExchangeFilterMaxNumAlgoOrders struct {
	FilterType       ExchangeFilterType `json:"filterType"`
	MaxNumAlgoOrders int                `json:"maxNumAlgoOrders"`
}

type ExchangeFilterMaxNumIcebergOrders struct {
	FilterType          ExchangeFilterType `json:"filterType"`
	MaxNumIcebergOrders int                `json:"maxNumIcebergOrders"`
}

type ExchangeFilterMaxNumOrderLists struct {
	FilterType       ExchangeFilterType `json:"filterType"`
	MaxNumOrderLists int                `json:"maxNumOrderLists"`
}

// Symbol

type Symbol struct {
	Symbol                          string       `json:"symbol"`
	Status                          SymbolStatus `json:"status"`
	BaseAsset                       string       `json:"baseAsset"`
	BaseAssetPrecision              int          `json:"baseAssetPrecision"`
	QuoteAsset                      string       `json:"quoteAsset"`
	QuotePrecision                  int          `json:"quotePrecision"`
	BaseCommissionPrecision         int          `json:"baseCommissionPrecision"`
	QuoteCommissionPrecision        int          `json:"quoteCommissionPrecision"`
	OrderTypes                      []OrderType  `json:"orderTypes"`
	IcebergAllowed                  bool         `json:"icebergAllowed"`
	OcoAllowed                      bool         `json:"ocoAllowed"`
	OtoAllowed                      bool         `json:"otoAllowed"`
	QuoteOrderQtyMarketAllowed      bool         `json:"quoteOrderQtyMarketAllowed"`
	AllowTrailingStop               bool         `json:"allowTrailingStop"`
	CancelReplaceAllowedbool        bool         `json:"cancelReplaceAllowedbool"`
	IsSpotTradingAllowed            bool         `json:"isSpotTradingAllowed"`
	IsMarginTradingAllowed          bool         `json:"isMarginTradingAllowed"`
	Filters                         SymbolFilters
	Permissions                     []Permission   `json:"permissions"`
	PermissionSets                  [][]Permission `json:"permissionSets"`
	DefaultSelfTradePreventionMode  string         `json:"defaultSelfTradePreventionMode"`
	AllowedSelfTradePreventionModes []string       `json:"allowedSelfTradePreventionModes"`
}

type SymbolFilters struct {
	PRICE_FILTER           *SymbolFilterPriceFilter
	PERCENT_PRICE          *SymbolFilterPercentPrice
	PERCENT_PRICE_BY_SIDE  *SymbolFilterPercentPriceBySide
	LOT_SIZE               *SymbolFilterLotSize
	MIN_NOTIONAL           *SymbolFilterMinNotional
	NOTIONAL               *SymbolFilterNotional
	ICEBERG_PARTS          *SymbolFilterIcebergParts
	MARKET_LOT_SIZE        *SymbolFilterMarketLotSize
	MAX_NUM_ORDERS         *SymbolFilterMaxNumOrders
	MAX_NUM_ALGO_ORDERS    *SymbolFilterMaxNumAlgoOrders
	MAX_NUM_ICEBERG_ORDERS *SymbolFilterMaxNumIcebergOrders
	MAX_POSITION           *SymbolFilterMaxPosition
	TRAILING_DELTA         *SymbolFilterTrailingDelta
	MAX_NUM_ORDER_AMENDS   *SymbolFilterMaxNumOrderAmends
	MAX_NUM_ORDER_LISTS    *SymbolFilterMaxNumOrderLists
}

type SymbolFilterPriceFilter struct {
	FilterType SymbolFilterType `json:"filterType"`
	MinPrice   float64          `json:"minPrice,string"`
	MaxPrice   float64          `json:"maxPrice,string"`
	TickSize   float64          `json:"tickSize,string"`
}

type SymbolFilterPercentPrice struct {
	FilterType     SymbolFilterType `json:"filterType"`
	MultiplierUp   float64          `json:"multiplierUp,string"`
	MultiplierDown float64          `json:"multiplierDown,string"`
	AvgPriceMins   int              `json:"avgPriceMins"`
}

type SymbolFilterPercentPriceBySide struct {
	FilterType        SymbolFilterType `json:"filterType"`
	BidMultiplierUp   float64          `json:"bidMultiplierUp,string"`
	BidMultiplierDown float64          `json:"bidMultiplierDown,string"`
	AskMultiplierUp   float64          `json:"askMultiplierUp,string"`
	AskMultiplierDown float64          `json:"askMultiplierDown,string"`
	AvgPriceMins      int              `json:"avgPriceMins"`
}

type SymbolFilterLotSize struct {
	FilterType SymbolFilterType `json:"filterType"`
	MinQty     float64          `json:"minQty,string"`
	MaxQty     float64          `json:"maxQty,string"`
	StepSize   float64          `json:"stepSize,string"`
}

type SymbolFilterMinNotional struct {
	FilterType    SymbolFilterType `json:"filterType"`
	MinNotional   float64          `json:"minNotional"`
	ApplyToMarket bool             `json:"applyToMarket"`
	AvgPriceMins  int              `json:"avgPriceMins"`
}

type SymbolFilterNotional struct {
	FilterType       SymbolFilterType `json:"filterType"`
	MinNotional      float64          `json:"minNotional,string"`
	ApplyMinToMarket bool             `json:"applyMinToMarket"`
	MaxNotional      float64          `json:"maxNotional,string"`
	ApplyMaxToMarket bool             `json:"applyMaxToMarket"`
	AvgPriceMins     int              `json:"avgPriceMins"`
}

type SymbolFilterIcebergParts struct {
	FilterType SymbolFilterType `json:"filterType"`
	Limit      int              `json:"limit"`
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

type SymbolFilterMaxNumIcebergOrders struct {
	FilterType          SymbolFilterType `json:"filterType"`
	MaxNumIcebergOrders int              `json:"maxNumIcebergOrders"`
}

type SymbolFilterMaxPosition struct {
	FilterType  SymbolFilterType `json:"filterType"`
	MaxPosition float64          `json:"maxPosition,string"`
}

type SymbolFilterTrailingDelta struct {
	FilterType            SymbolFilterType `json:"filterType"`
	MinTrailingAboveDelta int              `json:"minTrailingAboveDelta"`
	MaxTrailingAboveDelta int              `json:"maxTrailingAboveDelta"`
	MinTrailingBelowDelta int              `json:"minTrailingBelowDelta"`
	MaxTrailingBelowDelta int              `json:"maxTrailingBelowDelta"`
}

type SymbolFilterMaxNumOrderAmends struct {
	FilterType        SymbolFilterType `json:"filterType"`
	MaxNumOrderAmends int              `json:"maxNumOrderAmends"`
}

type SymbolFilterMaxNumOrderLists struct {
	FilterType       SymbolFilterType `json:"filterType"`
	MaxNumOrderLists int              `json:"maxNumOrderLists"`
}

type SORS struct {
	BaseAsset string   `json:"baseAsset"`
	Symbols   []string `json:"symbols"`
}

func (sors *SORS) HasSymbol(symbol string) bool {
	for _, s := range sors.Symbols {
		if s == symbol {
			return true
		}
	}

	return false
}

// # Current exchange trading rules and symbol information
//
// Weight: 20
func (c *Client) ExchangeInfo(opts ...ExchangeInfoOption) (*ExchangeInfo, Response, Error) {
	params, err := buildExchangeInfoParams(opts...)
	if err != nil {
		return nil, nil, err
	}

	exchangeInfo, httpResp, err := doRequest[ExchangeInfo](c, Methods.GET, "/api/v3/exchangeInfo", params, nil, NONE)
	if err != nil {
		return nil, httpResp, err
	}

	// Setting up the symbols map
	exchangeInfo.Symbols.Map = make(map[string]*Symbol)
	for _, symbol := range exchangeInfo.Symbols_arr {
		exchangeInfo.Symbols.Map[symbol.Symbol] = symbol
	}

	return exchangeInfo, httpResp, err
}

/////////////////////////////////////////////////////////////////////////////////
// Query Execution Rules
/////////////////////////////////////////////////////////////////////////////////

type QueryExecutionRules struct {
	SymbolRules []SymbolRule
}

type SymbolRule struct {
	Symbol string      `json:"symbol"`
	Rules  SymbolRules `json:"rules"`
}

type SymbolRules struct {
	PRICE_RANGE *SymbolRulePriceRange
}

type SymbolRulePriceRange struct {
	RuleType         ExecutionRule `json:"ruleType"`
	BidLimitMultUp   float64       `json:"bidLimitMultUp,string"`
	BidLimitMultDown float64       `json:"bidLimitMultDown,string"`
	AskLimitMultUp   float64       `json:"askLimitMultUp,string"`
	AskLimitMultDown float64       `json:"askLimitMultDown,string"`
}

type QueryExecutionRulesParams struct {
	Symbols      []string
	SymbolStatus SymbolStatus
}

func (c *Client) QueryExecutionRules(opts ...QueryExecutionRulesParams) (*QueryExecutionRules, Response, Error) {
	params := make(map[string]interface{})

	if len(opts) != 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "symbols", opt.Symbols)
		validation.SetIfNotZero(params, "symbolStatus", opt.Symbols)
	}

	return doRequest[QueryExecutionRules](c, Methods.GET, "/api/v3/exchangeInfo", params, nil, NONE)
}
