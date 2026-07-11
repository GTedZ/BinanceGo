package futures

// Enum definitions for futures trading

type SymbolType string

const (
	Future SymbolType = "FUTURE"
)

type ContractType string

const (
	Perpetual           ContractType = "PERPETUAL"
	CurrentMonth        ContractType = "CURRENT_MONTH"
	NextMonth           ContractType = "NEXT_MONTH"
	CurrentQuarter      ContractType = "CURRENT_QUARTER"
	NextQuarter         ContractType = "NEXT_QUARTER"
	PerpetualDelivering ContractType = "PERPETUAL_DELIVERING"
)

type ContractStatus string

const (
	PendingTrading ContractStatus = "PENDING_TRADING"
	Trading        ContractStatus = "TRADING"
	PreDelivering  ContractStatus = "PRE_DELIVERING"
	Delivering     ContractStatus = "DELIVERING"
	Delivered      ContractStatus = "DELIVERED"
	PreSettle      ContractStatus = "PRE_SETTLE"
	Settling       ContractStatus = "SETTLING"
	Close          ContractStatus = "CLOSE"
)

//

type Permission string

const (
	SPOT        Permission = "SPOT"
	MARGIN      Permission = "MARGIN"
	LEVERAGED   Permission = "LEVERAGED"
	TRD_GRP_002 Permission = "TRD_GRP_002"
	TRD_GRP_003 Permission = "TRD_GRP_003"
	TRD_GRP_004 Permission = "TRD_GRP_004"
	TRD_GRP_005 Permission = "TRD_GRP_005"
	TRD_GRP_006 Permission = "TRD_GRP_006"
	TRD_GRP_007 Permission = "TRD_GRP_007"
	TRD_GRP_008 Permission = "TRD_GRP_008"
	TRD_GRP_009 Permission = "TRD_GRP_009"
	TRD_GRP_010 Permission = "TRD_GRP_010"
	TRD_GRP_011 Permission = "TRD_GRP_011"
	TRD_GRP_012 Permission = "TRD_GRP_012"
	TRD_GRP_013 Permission = "TRD_GRP_013"
	TRD_GRP_014 Permission = "TRD_GRP_014"
	TRD_GRP_015 Permission = "TRD_GRP_015"
	TRD_GRP_016 Permission = "TRD_GRP_016"
	TRD_GRP_017 Permission = "TRD_GRP_017"
	TRD_GRP_018 Permission = "TRD_GRP_018"
	TRD_GRP_019 Permission = "TRD_GRP_019"
	TRD_GRP_020 Permission = "TRD_GRP_020"
	TRD_GRP_021 Permission = "TRD_GRP_021"
	TRD_GRP_022 Permission = "TRD_GRP_022"
	TRD_GRP_023 Permission = "TRD_GRP_023"
	TRD_GRP_024 Permission = "TRD_GRP_024"
	TRD_GRP_025 Permission = "TRD_GRP_025"
)

//

type OrderStatus string

const (
	StatusNew       OrderStatus = "NEW"
	PartiallyFilled OrderStatus = "PARTIALLY_FILLED"
	Filled          OrderStatus = "FILLED"
	Cancelled       OrderStatus = "CANCELED"
	Rejected        OrderStatus = "REJECTED"
	Expired         OrderStatus = "EXPIRED"
	ExpiredInMatch  OrderStatus = "EXPIRED_IN_MATCH"
)

//

type OrderType string

const (
	Limit              OrderType = "LIMIT"
	Market             OrderType = "MARKET"
	Stop               OrderType = "STOP"
	StopMarket         OrderType = "STOP_MARKET"
	TakeProfit         OrderType = "TAKE_PROFIT"
	TakeProfitMarket   OrderType = "TAKE_PROFIT_MARKET"
	TrailingStopMarket OrderType = "TRAILING_STOP_MARKET"
)

//

type OrderSide string

const (
	Buy  OrderSide = "BUY"
	Sell OrderSide = "SELL"
)

//

type PositionSide string

const (
	BOTH  PositionSide = "BOTH"
	LONG  PositionSide = "LONG"
	SHORT PositionSide = "SHORT"
)

//

type TimeInForce string

const (
	// Good Till Cancel(GTC order valitidy is 1 year from placement)
	GTC TimeInForce = "GTC"

	// Immediate or Cancel
	IOC TimeInForce = "IOC"

	// Fill or Kill
	FOK TimeInForce = "FOK"

	// Good Till Crossing (Post Only)
	GTX TimeInForce = "GTX"

	// Good Till Date
	GTD TimeInForce = "GTD"

	// Retail Price Improvement(RPI order is post only and only be matched with the order from APP or Web)
	RPI TimeInForce = "RPI"
)

//

type WorkingType string

const (
	WorkingTypeMarkPrice     WorkingType = "MARK_PRICE"
	WorkingTypeContractPrice WorkingType = "CONTRACT_PRICE"
)

//

type OrderResponseType string

const (
	Ack    OrderResponseType = "ACK"
	Result OrderResponseType = "RESULT"
)

//

type Interval string

const (
	// seconds
	Interval1s Interval = "1s"

	// minutes
	Interval1m  Interval = "1m"
	Interval3m  Interval = "3m"
	Interval5m  Interval = "5m"
	Interval15m Interval = "15m"
	Interval30m Interval = "30m"

	// hours
	Interval1h  Interval = "1h"
	Interval2h  Interval = "2h"
	Interval4h  Interval = "4h"
	Interval6h  Interval = "6h"
	Interval8h  Interval = "8h"
	Interval12h Interval = "12h"

	// days
	Interval1d Interval = "1d"
	Interval3d Interval = "3d"

	// weeks
	Interval1w Interval = "1w"

	// months
	Interval1M Interval = "1M"
)

//

type STPMode string

const (
	ExpireTaker STPMode = "EXPIRE_TAKER"
	ExpireBoth  STPMode = "EXPIRE_BOTH"
	ExpireMaker STPMode = "EXPIRE_MAKER"
)

//

type PriceMatch string

const (
	// No price match
	None PriceMatch = "NONE"

	// counterparty best price
	Opponent PriceMatch = "OPPONENT"

	// the 5th best price from the counterparty
	Opponent5 PriceMatch = "OPPONENT_5"

	// the 10th best price from the counterparty
	Opponent10 PriceMatch = "OPPONENT_10"

	// the 20th best price from the counterparty
	Opponent20 PriceMatch = "OPPONENT_20"

	// the best price on the same side of the order book
	Queue PriceMatch = "QUEUE"

	// the 5th best price on the same side of the order book
	Queue5 PriceMatch = "QUEUE_5"

	// the 10th best price on the same side of the order book
	Queue10 PriceMatch = "QUEUE_10"

	// the 20th best price on the same side of the order book
	Queue20 PriceMatch = "QUEUE_20"
)

//

type RateLimitType string

const (
	RateLimitRequestWeight RateLimitType = "REQUEST_WEIGHT"

	RateLimitOrders RateLimitType = "ORDERS"
)

//

type RateLimitInterval string

const (
	RateLimitMinute RateLimitInterval = "MINUTE"
)

//

type SymbolFilterType string

const (
	PRICE_FILTER          SymbolFilterType = "PRICE_FILTER"
	LOT_SIZE              SymbolFilterType = "LOT_SIZE"
	MARKET_LOT_SIZE       SymbolFilterType = "MARKET_LOT_SIZE"
	MAX_NUM_ORDERS        SymbolFilterType = "MAX_NUM_ORDERS"
	MAX_NUM_ALGO_ORDERS   SymbolFilterType = "MAX_NUM_ALGO_ORDERS"
	PERCENT_PRICE         SymbolFilterType = "PERCENT_PRICE"
	MIN_NOTIONAL          SymbolFilterType = "MIN_NOTIONAL"
	POSITION_RISK_CONTROL SymbolFilterType = "POSITION_RISK_CONTROL"
)

//

type PositionControl string

const (
	PositionControlNone PositionControl = "NONE"
)

//

type Period string

const (
	Period5m  Period = "5m"
	Period15m Period = "15m"
	Period30m Period = "30m"
	Period1h  Period = "1h"
	Period2h  Period = "2h"
	Period4h  Period = "4h"
	Period6h  Period = "6h"
	Period12h Period = "12h"
	Period1d  Period = "1d"
)

//

type RiskRating string

const (
	HighRisk   RiskRating = "High"
	MediumRisk RiskRating = "Medium"
	LowRisk    RiskRating = "Low"
)

//

type MarketType string

const (
	MarketTypeEquity    MarketType = "EQUITY"
	MarketTypeCommodity MarketType = "COMMODITY"
	MarketTypeKREquity  MarketType = "KR_EQUITY"
)

//

type MarketSessionType string

const (
	// Shared across all markets
	MarketSessionTypeNoTrading MarketSessionType = "NO_TRADING"

	// U.S. equity market
	MarketSessionTypePreMarket   MarketSessionType = "PRE_MARKET"
	MarketSessionTypeRegular     MarketSessionType = "REGULAR"
	MarketSessionTypeAfterMarket MarketSessionType = "AFTER_MARKET"
	MarketSessionTypeOvernight   MarketSessionType = "OVERNIGHT"
)

//

type AutoCloseType string

const (
	AutoCloseTypeLiquidation AutoCloseType = "LIQUIDATION"
	AutoCloseTypeADL         AutoCloseType = "ADL"
)

//

type MarginType string

const (
	Isolated MarginType = "ISOLATED"
	Cross    MarginType = "CROSSED"
)

//

type MarginModificationType int

const (
	IncreaseMargin MarginModificationType = 1
	ReduceMargin   MarginModificationType = 2
)

//

type WebsocketRoute string

const (
	// Public market data streams
	WebsocketRoutePublic WebsocketRoute = "/public"

	WebsocketRouteMarket WebsocketRoute = "/market"

	// Private user data streams
	WebsocketRoutePrivate WebsocketRoute = "/private"
)

//

// EventType is the value of the "e" field present in every market stream payload.
type EventType string

const (
	// Aggregate trade
	EventAggTrade EventType = "aggTrade"

	// Mark price
	EventMarkPriceUpdate EventType = "markPriceUpdate"

	// Kline / Candlestick
	EventKline EventType = "kline"

	// TODO: verify against a live feed. The docs example renders continuous
	// kline events with "e":"kline", but Binance historically emits
	// "continuous_kline" for the continuousKline stream. The concrete parsing
	// does not depend on this constant, it is exposed for callers only.
	EventContinuousKline EventType = "continuous_kline"

	// Mini ticker
	Event24hrMiniTicker EventType = "24hrMiniTicker"

	// Full ticker
	Event24hrTicker EventType = "24hrTicker"

	// Book ticker
	EventBookTicker EventType = "bookTicker"

	// Liquidation order
	EventForceOrder EventType = "forceOrder"

	// Order book depth updates
	EventDepthUpdate EventType = "depthUpdate"

	// Composite index
	EventCompositeIndex EventType = "compositeIndex"

	// Contract info
	EventContractInfo EventType = "contractInfo"

	// Multi-Assets mode asset index
	EventAssetIndexUpdate EventType = "assetIndexUpdate"
)

//

// UserDataEventType is the value of the "e" field present in every user data
// stream payload.
type UserDataEventType string

const (
	EventListenKeyExpired              UserDataEventType = "listenKeyExpired"
	EventMarginCall                    UserDataEventType = "MARGIN_CALL"
	EventAccountUpdate                 UserDataEventType = "ACCOUNT_UPDATE"
	EventOrderTradeUpdate              UserDataEventType = "ORDER_TRADE_UPDATE"
	EventTradeLite                     UserDataEventType = "TRADE_LITE"
	EventAccountConfigUpdate           UserDataEventType = "ACCOUNT_CONFIG_UPDATE"
	EventStrategyUpdate                UserDataEventType = "STRATEGY_UPDATE"
	EventGridUpdate                    UserDataEventType = "GRID_UPDATE"
	EventConditionalOrderTriggerReject UserDataEventType = "CONDITIONAL_ORDER_TRIGGER_REJECT"
)

func (e UserDataEventType) IsValid() bool {
	switch e {
	case EventListenKeyExpired,
		EventMarginCall,
		EventAccountUpdate,
		EventOrderTradeUpdate,
		EventTradeLite,
		EventAccountConfigUpdate,
		EventStrategyUpdate,
		EventGridUpdate,
		EventConditionalOrderTriggerReject:
		return true
	}
	return false
}
