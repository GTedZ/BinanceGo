package spot

// Enum definitions for spot trading

type SymbolStatus string

const (
	// Trading is active for the symbol.
	Trading SymbolStatus = "TRADING"

	// Trading has ended for the current trading day.
	EndOfDay SymbolStatus = "END_OF_DAY"

	// Trading is halted for the symbol.
	Halt SymbolStatus = "HALT"

	// Trading is paused for a scheduled break.
	Break SymbolStatus = "BREAK"
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
	// The order has been accepted by the engine.
	OrderStatusNew OrderStatus = "NEW"

	// The order is pending until the working order of an order list has been fully filled.
	PendingNew OrderStatus = "PENDING_NEW"

	// A part of the order has been filled.
	PartiallyFilled OrderStatus = "PARTIALLY_FILLED"

	// The order has been completely filled.
	Filled OrderStatus = "FILLED"

	// The order has been canceled by the user.
	Canceled OrderStatus = "CANCELED"

	// Currently unused.
	PendingCancel OrderStatus = "PENDING_CANCEL"

	// The order was rejected by the engine and not processed.
	Rejected OrderStatus = "REJECTED"

	// The order expired due to order rules or exchange conditions.
	Expired OrderStatus = "EXPIRED"

	// The order expired due to Self Trade Prevention (STP).
	ExpiredInMatch OrderStatus = "EXPIRED_IN_MATCH"
)

//

type ListStatusType string

const (
	// Returned when the list status responds to a failed action such as order list placement or cancellation.
	ListResponse ListStatusType = "RESPONSE"

	// Indicates the order list has been placed or updated.
	ExecStarted ListStatusType = "EXEC_STARTED"

	// Indicates a clientOrderId in the list has changed.
	Updated ListStatusType = "UPDATED"

	// Indicates the order list has completed execution and is no longer active.
	ListAllDone ListStatusType = "ALL_DONE"
)

//

type ListOrderStatus string

const (
	// The order list is currently executing or has an active update.
	Executing ListOrderStatus = "EXECUTING"

	// The order list has completed execution and is no longer active.
	ListOrderAllDone ListOrderStatus = "ALL_DONE"

	// The order list action failed during placement or cancellation.
	ListOrderReject ListOrderStatus = "REJECT"
)

//

type ContingencyType string

const (
	OCO ContingencyType = "OCO"
	OTO ContingencyType = "OTO"
)

//

type AllocationType string

const (
	SOR AllocationType = "SOR"
)

//

type OrderType string

const (
	Limit           OrderType = "LIMIT"
	Market          OrderType = "MARKET"
	StopLoss        OrderType = "STOP_LOSS"
	StopLossLimit   OrderType = "STOP_LOSS_LIMIT"
	TakeProfit      OrderType = "TAKE_PROFIT"
	TakeProfitLimit OrderType = "TAKE_PROFIT_LIMIT"
	LimitMaker      OrderType = "LIMIT_MAKER"
)

//

type OrderResponseType string

const (
	Ack OrderResponseType = "ACK"

	// Provides the final state of the order after execution.
	Result OrderResponseType = "RESULT"

	// Provides full order information including fills.
	Full OrderResponseType = "FULL"
)

//

type WorkingFloor string

const (
	// Order executed directly on the exchange order book.
	Exchange WorkingFloor = "EXCHANGE"

	// Order executed through Smart Order Routing.
	Sor WorkingFloor = "SOR"
)

//

type OrderSide string

const (
	// Buy order.
	Buy OrderSide = "BUY"

	// Sell order.
	Sell OrderSide = "SELL"
)

//

type TimeInForce string

const (
	// Good Til Canceled — order remains on the book until canceled.
	GTC TimeInForce = "GTC"

	// Immediate Or Cancel — order executes immediately and any unfilled portion is canceled.
	IOC TimeInForce = "IOC"

	// Fill Or Kill — order must be completely filled immediately or canceled.
	FOK TimeInForce = "FOK"
)

//

type RateLimitType string

const (
	RequestWeight RateLimitType = "REQUEST_WEIGHT"
	Orders        RateLimitType = "ORDERS"
	RawRequests   RateLimitType = "RAW_REQUESTS"
)

//

type RateLimitInterval string

const (
	Second RateLimitInterval = "SECOND"
	Minute RateLimitInterval = "MINUTE"
	Day    RateLimitInterval = "DAY"
)

//

type STPMode string

const (
	STPNone      STPMode = "NONE"
	ExpireMaker  STPMode = "EXPIRE_MAKER"
	ExpireTaker  STPMode = "EXPIRE_TAKER"
	ExpireBoth   STPMode = "EXPIRE_BOTH"
	STPDecrement STPMode = "DECREMENT"
	STPTransfer  STPMode = "TRANSFER"
)

//

type ExecutionType string

const (
	// The order has been accepted into the engine.
	ExecNew ExecutionType = "NEW"

	// The order has been canceled by the user.
	ExecCanceled ExecutionType = "CANCELED"

	// The order has been amended.
	ExecReplaced ExecutionType = "REPLACED"

	// The order has been rejected and was not processed (e.g. `Cancel Replace Orders` wherein the new order placement is rejected but the request to cancel request succeeds.)
	ExecRejected ExecutionType = "REJECTED"

	// Part of the order or all of the order's quantity has filled.
	ExecTrade ExecutionType = "TRADE"

	// The order was canceled according to the order type's rules (e.g. `LIMIT FOK` orders with no fill, `LIMIT IOC` or `MARKET` orders that partially fill) or by the exchange, (e.g. orders canceled during liquidation, orders canceled during maintenance).
	ExecExpired ExecutionType = "EXPIRED"

	// The order has expired due to STP.
	ExecTradePrevention ExecutionType = "TRADE_PREVENTION"
)

//

type ExecutionRule string

const (
	PRICE_RANGE ExecutionRule = "PRICE_RANGE"
)

//

type ExchangeFilterType string

const (
	EXCHANGE_MAX_NUM_ORDERS         ExchangeFilterType = "EXCHANGE_MAX_NUM_ORDERS"
	EXCHANGE_MAX_NUM_ALGO_ORDERS    ExchangeFilterType = "EXCHANGE_MAX_NUM_ALGO_ORDERS"
	EXCHANGE_MAX_NUM_ICEBERG_ORDERS ExchangeFilterType = "EXCHANGE_MAX_NUM_ICEBERG_ORDERS"
	EXCHANGE_MAX_NUM_ORDER_LISTS    ExchangeFilterType = "EXCHANGE_MAX_NUM_ORDER_LISTS"
)

type SymbolFilterType string

const (
	PRICE_FILTER           SymbolFilterType = "PRICE_FILTER"
	PERCENT_PRICE          SymbolFilterType = "PERCENT_PRICE"
	PERCENT_PRICE_BY_SIDE  SymbolFilterType = "PERCENT_PRICE_BY_SIDE"
	LOT_SIZE               SymbolFilterType = "LOT_SIZE"
	MIN_NOTIONAL           SymbolFilterType = "MIN_NOTIONAL"
	NOTIONAL               SymbolFilterType = "NOTIONAL"
	ICEBERG_PARTS          SymbolFilterType = "ICEBERG_PARTS"
	MARKET_LOT_SIZE        SymbolFilterType = "MARKET_LOT_SIZE"
	MAX_NUM_ORDERS         SymbolFilterType = "MAX_NUM_ORDERS"
	MAX_NUM_ALGO_ORDERS    SymbolFilterType = "MAX_NUM_ALGO_ORDERS"
	MAX_NUM_ICEBERG_ORDERS SymbolFilterType = "MAX_NUM_ICEBERG_ORDERS"
	MAX_POSITION           SymbolFilterType = "MAX_POSITION"
	TRAILING_DELTA         SymbolFilterType = "TRAILING_DELTA"
	MAX_NUM_ORDER_AMENDS   SymbolFilterType = "MAX_NUM_ORDER_AMENDS"
	MAX_NUM_ORDER_LISTS    SymbolFilterType = "MAX_NUM_ORDER_LISTS"
)

type AssetFilterType string

const (
	MAX_ASSET AssetFilterType = "MAX_ASSET"
)

type KlineInterval string

const (
	// seconds
	Interval1s KlineInterval = "1s"

	// minutes
	Interval1m  KlineInterval = "1m"
	Interval3m  KlineInterval = "3m"
	Interval5m  KlineInterval = "5m"
	Interval15m KlineInterval = "15m"
	Interval30m KlineInterval = "30m"

	// hours
	Interval1h  KlineInterval = "1h"
	Interval2h  KlineInterval = "2h"
	Interval4h  KlineInterval = "4h"
	Interval6h  KlineInterval = "6h"
	Interval8h  KlineInterval = "8h"
	Interval12h KlineInterval = "12h"

	// days
	Interval1d KlineInterval = "1d"
	Interval3d KlineInterval = "3d"

	// weeks
	Interval1w KlineInterval = "1w"

	// months
	Interval1M KlineInterval = "1M"
)

type ReferencePriceCalculationType string

const (
	ArithmeticMean ReferencePriceCalculationType = "ARITHMETIC_MEAN"
	External       ReferencePriceCalculationType = "EXTERNAL"
)
