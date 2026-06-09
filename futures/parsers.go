package futures

import (
	"fmt"
	"strconv"

	"github.com/GTedZ/binancego/internal/json"
)

////
// ExchangeInfo
////

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
	var lotSize *SymbolFilterLotSize
	var marketLotSize *SymbolFilterMarketLotSize
	var maxNumOrders *SymbolFilterMaxNumOrders
	var maxNumAlgoOrders *SymbolFilterMaxNumAlgoOrders
	var percentPrice *SymbolFilterPercentPrice
	var minNotional *SymbolFilterMinNotional

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

		case LOT_SIZE:
			var temp SymbolFilterLotSize
			err = json.Unmarshal(rawFilterData, &temp)
			if err != nil {
				return err
			}
			lotSize = &temp

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

		case PERCENT_PRICE:
			var temp SymbolFilterPercentPrice
			err = json.Unmarshal(rawFilterData, &temp)
			if err != nil {
				return err
			}
			percentPrice = &temp

		case MIN_NOTIONAL:
			var temp SymbolFilterMinNotional
			err = json.Unmarshal(rawFilterData, &temp)
			if err != nil {
				return err
			}
			minNotional = &temp

		default:
			fmt.Printf("[BinanceGo] unknown filterType %s found\n", symbolFilter.FilterType)
		}
	}

	*symbolFilters = SymbolFilters{
		PRICE_FILTER:        priceFilter,
		LOT_SIZE:            lotSize,
		MARKET_LOT_SIZE:     marketLotSize,
		MAX_NUM_ORDERS:      maxNumOrders,
		MAX_NUM_ALGO_ORDERS: maxNumAlgoOrders,
		PERCENT_PRICE:       percentPrice,
		MIN_NOTIONAL:        minNotional,
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

// func (u *UserDataEvent) UnmarshalJSON(data []byte) error {
// 	// Step 1: get raw event
// 	var raw struct {
// 		SubscriptionID int             `json:"subscriptionId"`
// 		Event          json.RawMessage `json:"event"`
// 	}
// 	if err := json.Unmarshal(data, &raw); err != nil {
// 		return err
// 	}
// 	u.SubscriptionID = raw.SubscriptionID

// 	// Step 2: detect event type — probe as string first, fall back to numeric
// 	var probe struct {
// 		E     UserDataEventType `json:"e"`
// 		ETime int64             `json:"E"`
// 	}
// 	if err := json.Unmarshal(raw.Event, &probe); err != nil {
// 		return err
// 	}

// 	switch probe.E {
// 	case EventOutboundAccountPosition:
// 		var temp OutboundAccountPositionEvent
// 		err := json.Unmarshal(raw.Event, &temp)
// 		if err != nil {
// 			return err
// 		}
// 		u.EventType = temp.EventType
// 		u.EventTime = temp.EventTime
// 		u.AccountPosition = &temp

// 	case EventBalanceUpdate:
// 		var temp BalanceUpdateEvent
// 		err := json.Unmarshal(raw.Event, &temp)
// 		if err != nil {
// 			return err
// 		}
// 		u.EventType = temp.EventType
// 		u.EventTime = temp.EventTime
// 		u.BalanceUpdate = &temp

// 	case EventExecutionReport:
// 		var temp ExecutionReportEvent
// 		err := json.Unmarshal(raw.Event, &temp)
// 		if err != nil {
// 			return err
// 		}
// 		u.EventType = temp.EventType
// 		u.EventTime = temp.EventTime
// 		u.ExecutionReport = &temp

// 	case EventListStatus:
// 		var temp ListStatusEvent
// 		err := json.Unmarshal(raw.Event, &temp)
// 		if err != nil {
// 			return err
// 		}
// 		u.EventType = temp.EventType
// 		u.EventTime = temp.EventTime
// 		u.ListStatus = &temp

// 	case EventStreamTerminated:
// 		var temp EventStreamTerminatedEvent
// 		err := json.Unmarshal(raw.Event, &temp)
// 		if err != nil {
// 			return err
// 		}
// 		u.EventType = temp.EventType
// 		u.EventTime = temp.EventTime
// 		u.StreamTerminated = &temp

// 	case EventExternalLockUpdate:
// 		var temp ExternalLockUpdateEvent
// 		err := json.Unmarshal(raw.Event, &temp)
// 		if err != nil {
// 			return err
// 		}
// 		u.EventType = temp.EventType
// 		u.EventTime = temp.EventTime
// 		u.ExternalLock = &temp

// 	default:
// 		fmt.Printf("[BinanceGo] unknown user data event type %s found\n", probe.E)
// 	}

// 	return nil
// }
