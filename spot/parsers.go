package spot

import (
	"fmt"
	"strconv"

	"github.com/GTedZ/binancego/internal/json"
)

////
// ExchangeInfo
////

// ExchangeInfoFilters
type exchangeFilterRawType struct {
	FilterType ExchangeFilterType `json:"filterType"`
}

func (exchangeFilters *ExchangeFilters) UnmarshalJSON(data []byte) error {
	var rawExchangeFilters []json.RawMessage
	err := json.Unmarshal(data, &rawExchangeFilters)
	if err != nil {
		return err
	}

	var maxNumOrders *ExchangeFilterMaxNumOrders
	var maxNumAlgoOrders *ExchangeFilterMaxNumAlgoOrders
	var maxNumIcebergOrders *ExchangeFilterMaxNumIcebergOrders
	var maxOrderLists *ExchangeFilterMaxNumOrderLists

	for _, rawFilterData := range rawExchangeFilters {
		var exchangeFilter exchangeFilterRawType
		err = json.Unmarshal(rawFilterData, &exchangeFilter)
		if err != nil {
			return err
		}

		switch exchangeFilter.FilterType {
		case EXCHANGE_MAX_NUM_ORDERS:
			var temp ExchangeFilterMaxNumOrders
			err = json.Unmarshal(rawFilterData, &temp)
			if err != nil {
				return err
			}
			maxNumOrders = &temp

		case EXCHANGE_MAX_NUM_ALGO_ORDERS:
			var temp ExchangeFilterMaxNumAlgoOrders
			err = json.Unmarshal(rawFilterData, &temp)
			if err != nil {
				return err
			}
			maxNumAlgoOrders = &temp

		case EXCHANGE_MAX_NUM_ICEBERG_ORDERS:
			var temp ExchangeFilterMaxNumIcebergOrders
			err = json.Unmarshal(rawFilterData, &temp)
			if err != nil {
				return err
			}
			maxNumIcebergOrders = &temp

		case EXCHANGE_MAX_NUM_ORDER_LISTS:
			var temp ExchangeFilterMaxNumOrderLists
			err = json.Unmarshal(rawFilterData, &temp)
			if err != nil {
				return err
			}
			maxOrderLists = &temp

		default:
			fmt.Printf("[BinanceGo] unknown filterType %s found\n", exchangeFilter.FilterType)
		}
	}

	*exchangeFilters = ExchangeFilters{
		MAX_NUM_ORDERS:         maxNumOrders,
		MAX_NUM_ALGO_ORDERS:    maxNumAlgoOrders,
		MAX_NUM_ICEBERG_ORDERS: maxNumIcebergOrders,
		MAX_NUM_ORDER_LISTS:    maxOrderLists,
	}
	return nil
}

// ExchangeInfo Symbol
type symbolFilterRawType struct {
	FilterType SymbolFilterType `json:"filterType"`
}

func (symbolFilters *SymbolFilters) UnmarshalJSON(data []byte) error {
	var rawSymbolFilters []json.RawMessage
	err := json.Unmarshal(data, &rawSymbolFilters)
	if err != nil {
		return err
	}

	var priceFilter *SymbolFilterPriceFilter
	var percentPrice *SymbolFilterPercentPrice
	var percentPriceBySide *SymbolFilterPercentPriceBySide
	var lotSize *SymbolFilterLotSize
	var minNotional *SymbolFilterMinNotional
	var notional *SymbolFilterNotional
	var icebergParts *SymbolFilterIcebergParts
	var marketLotSize *SymbolFilterMarketLotSize
	var maxNumOrders *SymbolFilterMaxNumOrders
	var maxNumAlgoOrders *SymbolFilterMaxNumAlgoOrders
	var maxNumIcebergOrders *SymbolFilterMaxNumIcebergOrders
	var maxPosition *SymbolFilterMaxPosition
	var trailingDelta *SymbolFilterTrailingDelta
	var maxNumOrderAmends *SymbolFilterMaxNumOrderAmends
	var maxNumOrderLists *SymbolFilterMaxNumOrderLists

	for _, rawFilterData := range rawSymbolFilters {
		var symbolFilter symbolFilterRawType
		err = json.Unmarshal(rawFilterData, &symbolFilter)
		if err != nil {
			return err
		}

		switch symbolFilter.FilterType {
		case PRICE_FILTER:
			var temp SymbolFilterPriceFilter
			err = json.Unmarshal(rawFilterData, &temp)
			if err != nil {
				return err
			}
			priceFilter = &temp

		case PERCENT_PRICE:
			var temp SymbolFilterPercentPrice
			err = json.Unmarshal(rawFilterData, &temp)
			if err != nil {
				return err
			}
			percentPrice = &temp

		case PERCENT_PRICE_BY_SIDE:
			var temp SymbolFilterPercentPriceBySide
			err = json.Unmarshal(rawFilterData, &temp)
			if err != nil {
				return err
			}
			percentPriceBySide = &temp

		case LOT_SIZE:
			var temp SymbolFilterLotSize
			err = json.Unmarshal(rawFilterData, &temp)
			if err != nil {
				return err
			}
			lotSize = &temp

		case MIN_NOTIONAL:
			var temp SymbolFilterMinNotional
			err = json.Unmarshal(rawFilterData, &temp)
			if err != nil {
				return err
			}
			minNotional = &temp

		case NOTIONAL:
			var temp SymbolFilterNotional
			err = json.Unmarshal(rawFilterData, &temp)
			if err != nil {
				return err
			}
			notional = &temp

		case ICEBERG_PARTS:
			var temp SymbolFilterIcebergParts
			err = json.Unmarshal(rawFilterData, &temp)
			if err != nil {
				return err
			}
			icebergParts = &temp

		case MARKET_LOT_SIZE:
			var temp SymbolFilterMarketLotSize
			err = json.Unmarshal(rawFilterData, &temp)
			if err != nil {
				return err
			}
			marketLotSize = &temp

		case MAX_NUM_ORDERS:
			var temp SymbolFilterMaxNumOrders
			err = json.Unmarshal(rawFilterData, &temp)
			if err != nil {
				return err
			}
			maxNumOrders = &temp

		case MAX_NUM_ALGO_ORDERS:
			var temp SymbolFilterMaxNumAlgoOrders
			err = json.Unmarshal(rawFilterData, &temp)
			if err != nil {
				return err
			}
			maxNumAlgoOrders = &temp

		case MAX_NUM_ICEBERG_ORDERS:
			var temp SymbolFilterMaxNumIcebergOrders
			err = json.Unmarshal(rawFilterData, &temp)
			if err != nil {
				return err
			}
			maxNumIcebergOrders = &temp

		case MAX_POSITION:
			var temp SymbolFilterMaxPosition
			err = json.Unmarshal(rawFilterData, &temp)
			if err != nil {
				return err
			}
			maxPosition = &temp

		case TRAILING_DELTA:
			var temp SymbolFilterTrailingDelta
			err = json.Unmarshal(rawFilterData, &temp)
			if err != nil {
				return err
			}
			trailingDelta = &temp

		case MAX_NUM_ORDER_AMENDS:
			var temp SymbolFilterMaxNumOrderAmends
			err = json.Unmarshal(rawFilterData, &temp)
			if err != nil {
				return err
			}
			maxNumOrderAmends = &temp

		case MAX_NUM_ORDER_LISTS:
			var temp SymbolFilterMaxNumOrderLists
			err = json.Unmarshal(rawFilterData, &temp)
			if err != nil {
				return err
			}
			maxNumOrderLists = &temp

		default:
			fmt.Printf("[BinanceGo] unknown filterType %s found\n", symbolFilter.FilterType)
		}
	}

	*symbolFilters = SymbolFilters{
		PRICE_FILTER:           priceFilter,
		PERCENT_PRICE:          percentPrice,
		PERCENT_PRICE_BY_SIDE:  percentPriceBySide,
		LOT_SIZE:               lotSize,
		MIN_NOTIONAL:           minNotional,
		NOTIONAL:               notional,
		ICEBERG_PARTS:          icebergParts,
		MARKET_LOT_SIZE:        marketLotSize,
		MAX_NUM_ORDERS:         maxNumOrders,
		MAX_NUM_ALGO_ORDERS:    maxNumAlgoOrders,
		MAX_NUM_ICEBERG_ORDERS: maxNumIcebergOrders,
		MAX_POSITION:           maxPosition,
		TRAILING_DELTA:         trailingDelta,
		MAX_NUM_ORDER_AMENDS:   maxNumOrderAmends,
		MAX_NUM_ORDER_LISTS:    maxNumOrderLists,
	}
	return nil
}

type symbolRuleRawType struct {
	RuleType ExecutionRule `json:"ruleType"`
}

func (symbolRules *SymbolRules) UnmarshalJSON(data []byte) error {
	var rawSymbolRules []json.RawMessage
	err := json.Unmarshal(data, &rawSymbolRules)
	if err != nil {
		return err
	}

	var priceRange *SymbolRulePriceRange

	for _, rawFilterData := range rawSymbolRules {
		var ruleType symbolRuleRawType
		err = json.Unmarshal(rawFilterData, &ruleType)
		if err != nil {
			return err
		}

		switch ruleType.RuleType {
		case PRICE_RANGE:
			var temp SymbolRulePriceRange
			err = json.Unmarshal(rawFilterData, &temp)
			if err != nil {
				return err
			}
			priceRange = &temp

		default:
			fmt.Printf("[BinanceGo] unknown ruleType %s found\n", ruleType.RuleType)
		}
	}

	*symbolRules = SymbolRules{
		PRICE_RANGE: priceRange,
	}
	return nil
}

////
// Asset Filter
////

type assetFilterRawType struct {
	FilterType AssetFilterType `json:"filterType"`
}

func (assetFilters *AssetFilters) UnmarshalJSON(data []byte) error {
	var rawAssetFilters []json.RawMessage
	err := json.Unmarshal(data, &rawAssetFilters)
	if err != nil {
		return err
	}

	var maxAsset *AssetFilterMaxAsset

	for _, rawFilterData := range rawAssetFilters {
		var assetFilter assetFilterRawType
		err = json.Unmarshal(rawFilterData, &assetFilter)
		if err != nil {
			return err
		}

		switch assetFilter.FilterType {
		case MAX_ASSET:
			var temp AssetFilterMaxAsset
			err = json.Unmarshal(rawFilterData, &temp)
			if err != nil {
				return err
			}
			maxAsset = &temp

		default:
			fmt.Printf("[BinanceGo] unknown filterType %s found\n", assetFilter.FilterType)
		}
	}

	*assetFilters = AssetFilters{
		MAX_ASSET: maxAsset,
	}
	return nil
}

////
// Price Level
////

func (p *PriceLevel) UnmarshalJSON(data []byte) error {
	var raw [2]string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if raw[0] == "" || raw[1] == "" {
		return fmt.Errorf("invalid price level: empty values")
	}

	price, err := strconv.ParseFloat(raw[0], 64)
	if err != nil {
		return err
	}

	qty, err := strconv.ParseFloat(raw[1], 64)
	if err != nil {
		return err
	}

	p.Price = price
	p.Qty = qty
	return nil
}

////
// Candlesticks
////

func (c *Candlestick) UnmarshalJSON(data []byte) error {
	var raw []interface{}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if len(raw) < 12 {
		return fmt.Errorf("invalid candlestick length: %d", len(raw))
	}

	var err error

	c.OpenTime = int64(raw[0].(float64))

	if c.Open, err = strconv.ParseFloat(raw[1].(string), 64); err != nil {
		return err
	}
	if c.High, err = strconv.ParseFloat(raw[2].(string), 64); err != nil {
		return err
	}
	if c.Low, err = strconv.ParseFloat(raw[3].(string), 64); err != nil {
		return err
	}
	if c.Close, err = strconv.ParseFloat(raw[4].(string), 64); err != nil {
		return err
	}
	if c.Volume, err = strconv.ParseFloat(raw[5].(string), 64); err != nil {
		return err
	}

	c.CloseTime = int64(raw[6].(float64))

	if c.QuoteAssetVolume, err = strconv.ParseFloat(raw[7].(string), 64); err != nil {
		return err
	}

	c.TradeCount = int64(raw[8].(float64))

	if c.TakerBuyBaseAssetVolume, err = strconv.ParseFloat(raw[9].(string), 64); err != nil {
		return err
	}
	if c.TakerBuyQuoteAssetVolume, err = strconv.ParseFloat(raw[10].(string), 64); err != nil {
		return err
	}

	c.Unused = raw[11].(string)

	return nil
}

////
// UserData Event
////

func (u *UserDataEvent) UnmarshalJSON(data []byte) error {
	// Step 1: get raw event
	var raw struct {
		SubscriptionID int             `json:"subscriptionId"`
		Event          json.RawMessage `json:"event"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	u.SubscriptionID = raw.SubscriptionID

	// Step 2: detect event type — probe as string first, fall back to numeric
	var probe struct {
		E     UserDataEventType `json:"e"`
		ETime int64             `json:"E"`
	}
	if err := json.Unmarshal(raw.Event, &probe); err != nil {
		return err
	}

	switch probe.E {
	case EventOutboundAccountPosition:
		var temp OutboundAccountPositionEvent
		err := json.Unmarshal(raw.Event, &temp)
		if err != nil {
			return err
		}
		u.EventType = temp.EventType
		u.EventTime = temp.EventTime
		u.AccountPosition = &temp

	case EventBalanceUpdate:
		var temp BalanceUpdateEvent
		err := json.Unmarshal(raw.Event, &temp)
		if err != nil {
			return err
		}
		u.EventType = temp.EventType
		u.EventTime = temp.EventTime
		u.BalanceUpdate = &temp

	case EventExecutionReport:
		var temp ExecutionReportEvent
		err := json.Unmarshal(raw.Event, &temp)
		if err != nil {
			return err
		}
		u.EventType = temp.EventType
		u.EventTime = temp.EventTime
		u.ExecutionReport = &temp

	case EventListStatus:
		var temp ListStatusEvent
		err := json.Unmarshal(raw.Event, &temp)
		if err != nil {
			return err
		}
		u.EventType = temp.EventType
		u.EventTime = temp.EventTime
		u.ListStatus = &temp

	case EventStreamTerminated:
		var temp EventStreamTerminatedEvent
		err := json.Unmarshal(raw.Event, &temp)
		if err != nil {
			return err
		}
		u.EventType = temp.EventType
		u.EventTime = temp.EventTime
		u.StreamTerminated = &temp

	case EventExternalLockUpdate:
		var temp ExternalLockUpdateEvent
		err := json.Unmarshal(raw.Event, &temp)
		if err != nil {
			return err
		}
		u.EventType = temp.EventType
		u.EventTime = temp.EventTime
		u.ExternalLock = &temp

	default:
		fmt.Printf("[BinanceGo] unknown user data event type %s found\n", probe.E)
	}

	return nil
}
