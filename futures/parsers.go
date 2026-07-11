package futures

import (
	"fmt"
	"strconv"

	"github.com/GTedZ/binancego/internal/json"
)

// //
// ExchangeInfo
// //
func (exchangeFilters *ExchangeFilters) UnmarshalJSON(data []byte) error {
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
	var lotSize *SymbolFilterLotSize
	var marketLotSize *SymbolFilterMarketLotSize
	var maxNumOrders *SymbolFilterMaxNumOrders
	var maxNumAlgoOrders *SymbolFilterMaxNumAlgoOrders
	var percentPrice *SymbolFilterPercentPrice
	var minNotional *SymbolFilterMinNotional
	var positionRiskControl *SymbolFilterPositionRiskControl

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

		case POSITION_RISK_CONTROL:
			var temp SymbolFilterPositionRiskControl
			err = json.Unmarshal(rawFilterData, &temp)
			if err != nil {
				return err
			}
			positionRiskControl = &temp

		default:
			fmt.Printf("[BinanceGo] unknown filterType %s found\n", symbolFilter.FilterType)
		}
	}

	// {"filterType":"POSITION_RISK_CONTROL","positionControlSide":"NONE"}

	*symbolFilters = SymbolFilters{
		PRICE_FILTER:          priceFilter,
		LOT_SIZE:              lotSize,
		MARKET_LOT_SIZE:       marketLotSize,
		MAX_NUM_ORDERS:        maxNumOrders,
		MAX_NUM_ALGO_ORDERS:   maxNumAlgoOrders,
		PERCENT_PRICE:         percentPrice,
		MIN_NOTIONAL:          minNotional,
		POSITION_RISK_CONTROL: positionRiskControl,
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

// UnmarshalJSON dispatches a raw USD-M futures user data event into the correct
// concrete event struct based on its "e" (event type) field.
//
// Unlike the spot user data stream (where every event is wrapped in a
// {subscriptionId, event} envelope), futures pushes raw event objects directly,
// so the payload is decoded in-place here.
func (u *UserDataEvent) UnmarshalJSON(data []byte) error {
	// Step 1: detect event type
	var probe struct {
		E     UserDataEventType `json:"e"`
		ETime int64             `json:"E"`
	}
	if err := json.Unmarshal(data, &probe); err != nil {
		return err
	}

	u.EventType = probe.E
	u.EventTime = probe.ETime

	switch probe.E {
	case EventListenKeyExpired:
		var temp ListenKeyExpiredEvent
		if err := json.Unmarshal(data, &temp); err != nil {
			return err
		}
		u.ListenKeyExpired = &temp

	case EventMarginCall:
		var temp MarginCallEvent
		if err := json.Unmarshal(data, &temp); err != nil {
			return err
		}
		u.MarginCall = &temp

	case EventAccountUpdate:
		var temp AccountUpdateEvent
		if err := json.Unmarshal(data, &temp); err != nil {
			return err
		}
		u.AccountUpdate = &temp

	case EventOrderTradeUpdate:
		var temp OrderTradeUpdateEvent
		if err := json.Unmarshal(data, &temp); err != nil {
			return err
		}
		u.OrderTradeUpdate = &temp

	case EventTradeLite:
		var temp TradeLiteEvent
		if err := json.Unmarshal(data, &temp); err != nil {
			return err
		}
		u.TradeLite = &temp

	case EventAccountConfigUpdate:
		var temp AccountConfigUpdateEvent
		if err := json.Unmarshal(data, &temp); err != nil {
			return err
		}
		u.AccountConfigUpdate = &temp

	case EventStrategyUpdate:
		var temp StrategyUpdateEvent
		if err := json.Unmarshal(data, &temp); err != nil {
			return err
		}
		u.StrategyUpdate = &temp

	case EventGridUpdate:
		var temp GridUpdateEvent
		if err := json.Unmarshal(data, &temp); err != nil {
			return err
		}
		u.GridUpdate = &temp

	case EventConditionalOrderTriggerReject:
		var temp ConditionalOrderTriggerRejectEvent
		if err := json.Unmarshal(data, &temp); err != nil {
			return err
		}
		u.ConditionalOrderTriggerReject = &temp

	default:
		fmt.Printf("[BinanceGo] unknown user data event type %s found\n", probe.E)
	}

	return nil
}
