package spot

import "github.com/GTedZ/binancego/internal/validation"

// This struct includes all known fields, including those that are conditionally
// returned depending on order type, execution, and advanced features.
//
// NOTE:
// - Binance encodes numeric values as strings; these are parsed into float64.
// - Many fields are optional and will remain zero values if not returned.
// - Presence of fields depends on order type, response type, and exchange features.
type Order struct {

	// --- Core identifiers ---

	// Symbol is the trading pair (e.g., BTCUSDT).
	Symbol string `json:"symbol"`

	// OrderId is the unique identifier assigned by Binance.
	OrderId int64 `json:"orderId"`

	// OrderListId identifies an order list (OCO). -1 if not part of a list.
	OrderListId int64 `json:"orderListId"`

	// ClientOrderId is the client-provided or auto-generated unique ID.
	ClientOrderId string `json:"clientOrderId"`

	// --- Pricing & quantities ---

	// Price is the order price. May be "0.00000000" for MARKET orders.
	Price float64 `json:"price,string"`

	// OrigQty is the original quantity of the order.
	OrigQty float64 `json:"origQty,string"`

	// ExecutedQty is the quantity that has been filled.
	ExecutedQty float64 `json:"executedQty,string"`

	// CummulativeQuoteQty is the total quote asset spent/received.
	CummulativeQuoteQty float64 `json:"cummulativeQuoteQty,string"`

	// OrigQuoteOrderQty is the original quote quantity for MARKET orders using quoteOrderQty.
	OrigQuoteOrderQty float64 `json:"origQuoteOrderQty,string"`

	// --- Order state ---

	// Status represents the current order status (e.g., NEW, FILLED, CANCELED).
	Status OrderStatus `json:"status"`

	// TimeInForce defines how long the order remains active (e.g., GTC, IOC).
	TimeInForce TimeInForce `json:"timeInForce"`

	// Type is the order type (e.g., LIMIT, MARKET, STOP_LOSS).
	Type OrderType `json:"type"`

	// Side indicates BUY or SELL.
	Side OrderSide `json:"side"`

	// --- Stop / iceberg ---

	// StopPrice is the trigger price for STOP_LOSS and TAKE_PROFIT orders.
	// Only present for trigger-based orders.
	StopPrice float64 `json:"stopPrice,string"`

	// IcebergQty is the visible portion of an iceberg order.
	// Only present if icebergQty was used.
	IcebergQty float64 `json:"icebergQty,string"`

	// --- Trailing stop ---

	// TrailingDelta is the price delta required to activate a trailing stop.
	//
	// Only present for trailing stop orders.
	TrailingDelta int64 `json:"trailingDelta"`

	// TrailingTime is the timestamp when the trailing order became active.
	//
	// Only present for trailing stop orders.
	TrailingTime int64 `json:"trailingTime"`

	// --- Strategy metadata ---

	// StrategyId is a user-defined identifier for grouping orders.
	StrategyId int64 `json:"strategyId"`

	// StrategyType is a user-defined strategy type (must be >= 1000000).
	StrategyType int `json:"strategyType"`

	// --- Self Trade Prevention (STP) ---

	// SelfTradePreventionMode indicates the STP mode used for this order.
	SelfTradePreventionMode STPMode `json:"selfTradePreventionMode"`

	// PreventedMatchId identifies a prevented match due to STP.
	// Only present if the order expired due to STP.
	PreventedMatchId int64 `json:"preventedMatchId"`

	// PreventedQuantity is the quantity prevented from matching due to STP.
	PreventedQuantity float64 `json:"preventedQuantity,string"`

	// --- Smart Order Routing (SOR) ---

	// UsedSor indicates whether Smart Order Routing was used.
	UsedSor bool `json:"usedSor"`

	// WorkingFloor indicates where the order is being executed (e.g., SOR or order book).
	WorkingFloor string `json:"workingFloor"`

	// --- Pegged orders ---

	// PegPriceType defines the peg type (PRIMARY_PEG or MARKET_PEG).
	// Only present for pegged orders.
	PegPriceType PegPriceType `json:"pegPriceType"`

	// PegOffsetType defines the peg offset type (currently PRICE_LEVEL).
	//
	// Only present for pegged orders.
	PegOffsetType PegOffsetType `json:"pegOffsetType"`

	// PegOffsetValue defines the offset level from the pegged price.
	PegOffsetValue int `json:"pegOffsetValue"`

	// PeggedPrice is the current dynamically adjusted pegged price.
	PeggedPrice float64 `json:"peggedPrice,string"`

	// --- Expiry ---

	// ExpiryReason indicates why the order expired (e.g., insufficient liquidity).
	ExpiryReason string `json:"expiryReason"`

	// --- Working state ---

	// IsWorking indicates whether the order is currently active on the book.
	IsWorking bool `json:"isWorking"`

	// WorkingTime is the timestamp when the order started working on the book.
	WorkingTime int64 `json:"workingTime"`

	// --- Timestamps ---

	// Time is the order creation timestamp (ms).
	Time int64 `json:"time"`

	// UpdateTime is the last update timestamp (ms).
	UpdateTime int64 `json:"updateTime"`

	// --- FULL response only ---

	// Fills contains trade breakdowns for the order.
	//
	// Only present when newOrderRespType = FULL.
	Fills []Fill `json:"fills"`
}

// Fill represents an individual trade execution within an order.
//
// It is only populated when the order response type is FULL
// (newOrderRespType = FULL). Each entry corresponds to a partial
// or complete fill of the order.
type Fill struct {

	// Price is the execution price of this trade.
	// Returned for all fills.
	Price float64 `json:"price,string"`

	// Qty is the quantity executed in this specific fill.
	// Returned for all fills.
	Qty float64 `json:"qty,string"`

	// Commission is the fee charged for this fill.
	// Returned for all fills.
	Commission float64 `json:"commission,string"`

	// CommissionAsset is the asset in which the commission was paid
	// (e.g., USDT, BNB, BTC).
	// Returned for all fills.
	CommissionAsset string `json:"commissionAsset"`

	// TradeId is the unique identifier of the trade execution.
	// May be -1 for certain execution types (e.g., SOR internal matches).
	// Returned for all fills.
	TradeId int64 `json:"tradeId"`

	// MatchType indicates how the trade was matched.
	// Only returned for SOR (Smart Order Routing) orders.
	// Example: "ONE_PARTY_TRADE_REPORT".
	MatchType string `json:"matchType"`

	// AllocId is an internal allocation identifier used by Binance.
	// Only returned for SOR orders.
	AllocId int64 `json:"allocId"`
}

////
// Global Order Interface
////

type orderRequest interface {
	build() map[string]interface{}
}

////
// Generic Order Handlers
////

func (c *Client) order(params map[string]interface{}) (*Order, Response, Error) {
	return doRequest[Order](c, Methods.POST, "/api/v3/order", params, nil, TRADE)
}

func (c *Client) orderTest(params map[string]interface{}) (*OrderTest, Response, Error) {
	return doRequest[OrderTest](c, Methods.POST, "/api/v3/order/test", params, nil, TRADE)
}

func (c *Client) orderSor(params map[string]interface{}) (*Order, Response, Error) {
	return doRequest[Order](c, Methods.POST, "/api/v3/sor/order", params, nil, TRADE)
}

func (c *Client) orderSorTest(params map[string]interface{}) (*OrderTest, Response, Error) {
	return doRequest[OrderTest](c, Methods.POST, "/api/v3/sor/order/test", params, nil, TRADE)
}

////
// Order Handler
////

// Order sends a new order request to Binance (/api/v3/order).
//
// This method accepts any orderRequest implementation (e.g. NewLimitBuy,
// NewMarketSell, NewStopLossLimitOrder, etc.) and submits it to the exchange.
//
// Response type:
// - Controlled by newOrderRespType:
//   - ACK (minimal)
//   - RESULT (execution info)
//   - FULL (includes fills)
//
// This method is a thin wrapper over Binance and does not enforce validation.
// Invalid parameter combinations will be rejected by the API.
func (c *Client) Order(req orderRequest) (*Order, Response, Error) {
	return c.order(req.build())
}

////
// Order Test Handler
////

type OrderTest struct {
	StandardCommissionForOrder CommissionRates `json:"standardCommissionForOrder"`
	SpecialCommissionForOrder  CommissionRates `json:"specialCommissionForOrder"`
	TaxCommissionForOrder      CommissionRates `json:"taxCommissionForOrder"`

	Discount CommissionDiscount `json:"discount"`
}

// OrderTest sends a new test order request to Binance.
//
// NOTE:
// If computeCommissionRates is false, Binance returns an empty JSON object `{}`.
// In that case, all fields in this struct will remain zero values.
// Users must explicitly enable computeCommissionRates to populate these fields.
func (c *Client) OrderTest(req orderRequest, computeCommissionRates bool) (*OrderTest, Response, Error) {
	params := req.build()

	params["computeCommissionRates"] = computeCommissionRates

	return c.orderTest(params)
}

////
// Order Sor Handler
////

func (c *Client) OrderSor(req orderRequest) (*Order, Response, Error) {
	return c.orderSor(req.build())
}

////
// Order Sor Test Handler
////

func (c *Client) OrderSorTest(req orderRequest, computeCommissionRates bool) (*OrderTest, Response, Error) {
	params := req.build()

	params["computeCommissionRates"] = computeCommissionRates

	return c.orderSorTest(params)
}

////
// Order Request Structs
////

type LimitOrderRequest struct {
	Symbol      string
	Side        OrderSide
	TimeInForce TimeInForce
	Quantity    float64
	Price       float64
	Params      LimitOrderParams
}

func (o LimitOrderRequest) build() map[string]interface{} {
	params := make(map[string]interface{})

	params["symbol"] = o.Symbol
	params["side"] = o.Side
	params["type"] = Limit

	validation.SetIfNotZero(params, "timeInForce", o.TimeInForce)
	validation.SetIfNotZero(params, "quantity", o.Quantity)
	validation.SetIfNotZero(params, "price", o.Price)

	validation.SetIfNotZero(params, "newClientOrderId", o.Params.NewClientOrderId)
	validation.SetIfNotZero(params, "strategyId", o.Params.StrategyId)
	validation.SetIfNotZero(params, "strategyType", o.Params.StrategyType)
	validation.SetIfNotZero(params, "icebergQty", o.Params.IcebergQty)
	validation.SetIfNotZero(params, "newOrderRespType", o.Params.NewOrderRespType)
	validation.SetIfNotZero(params, "selfTradePreventionMode", o.Params.SelfTradePreventionMode)
	validation.SetIfNotZero(params, "pegPriceType", o.Params.PegPriceType)
	validation.SetIfNotZero(params, "pegOffsetValue", o.Params.PegOffsetValue)
	validation.SetIfNotZero(params, "pegOffsetType", o.Params.PegOffsetType)
	validation.SetIfNotZero(params, "recvWindow", o.Params.RecvWindow)

	return params
}

type LimitMakerOrderRequest struct {
	Symbol   string
	Side     OrderSide
	Quantity float64
	Price    float64
	Params   LimitMakerOrderParams
}

func (o LimitMakerOrderRequest) build() map[string]interface{} {
	params := make(map[string]interface{})

	params["symbol"] = o.Symbol
	params["side"] = o.Side
	params["type"] = LimitMaker

	validation.SetIfNotZero(params, "quantity", o.Quantity)
	validation.SetIfNotZero(params, "price", o.Price)

	validation.SetIfNotZero(params, "newClientOrderId", o.Params.NewClientOrderId)
	validation.SetIfNotZero(params, "strategyId", o.Params.StrategyId)
	validation.SetIfNotZero(params, "strategyType", o.Params.StrategyType)
	validation.SetIfNotZero(params, "icebergQty", o.Params.IcebergQty)
	validation.SetIfNotZero(params, "newOrderRespType", o.Params.NewOrderRespType)
	validation.SetIfNotZero(params, "selfTradePreventionMode", o.Params.SelfTradePreventionMode)
	validation.SetIfNotZero(params, "pegPriceType", o.Params.PegPriceType)
	validation.SetIfNotZero(params, "pegOffsetValue", o.Params.PegOffsetValue)
	validation.SetIfNotZero(params, "pegOffsetType", o.Params.PegOffsetType)
	validation.SetIfNotZero(params, "recvWindow", o.Params.RecvWindow)

	return params
}

type MarketOrderRequest struct {
	Symbol   string
	Side     OrderSide
	Quantity float64
	Params   MarketOrderParams
}

func (o MarketOrderRequest) build() map[string]interface{} {
	params := make(map[string]interface{})

	params["symbol"] = o.Symbol
	params["side"] = o.Side
	params["type"] = Market

	usedQuantity := o.Quantity
	if validation.IsNotDefault(o.Params.QuoteOrderQty) {
		usedQuantity = 0
	}

	validation.SetIfNotZero(params, "quantity", usedQuantity)
	validation.SetIfNotZero(params, "quoteOrderQty", o.Params.QuoteOrderQty)

	validation.SetIfNotZero(params, "newClientOrderId", o.Params.NewClientOrderId)
	validation.SetIfNotZero(params, "strategyId", o.Params.StrategyId)
	validation.SetIfNotZero(params, "strategyType", o.Params.StrategyType)
	validation.SetIfNotZero(params, "newOrderRespType", o.Params.NewOrderRespType)
	validation.SetIfNotZero(params, "selfTradePreventionMode", o.Params.SelfTradePreventionMode)
	validation.SetIfNotZero(params, "recvWindow", o.Params.RecvWindow)

	return params
}

type StopOrderRequest struct {
	Symbol        string
	Side          OrderSide
	Quantity      float64
	StopPrice     float64
	TrailingDelta int64
	Type          OrderType // StopLoss or TakeProfit
	Params        StopOrderParams
}

func (o StopOrderRequest) build() map[string]interface{} {
	params := make(map[string]interface{})

	params["symbol"] = o.Symbol
	params["side"] = o.Side
	params["type"] = o.Type

	validation.SetIfNotZero(params, "quantity", o.Quantity)
	validation.SetIfNotZero(params, "stopPrice", o.StopPrice)
	validation.SetIfNotZero(params, "trailingDelta", o.TrailingDelta)

	validation.SetIfNotZero(params, "newClientOrderId", o.Params.NewClientOrderId)
	validation.SetIfNotZero(params, "strategyId", o.Params.StrategyId)
	validation.SetIfNotZero(params, "strategyType", o.Params.StrategyType)
	validation.SetIfNotZero(params, "newOrderRespType", o.Params.NewOrderRespType)
	validation.SetIfNotZero(params, "selfTradePreventionMode", o.Params.SelfTradePreventionMode)
	validation.SetIfNotZero(params, "recvWindow", o.Params.RecvWindow)

	return params
}

type StopLimitOrderRequest struct {
	Symbol        string
	Side          OrderSide
	Price         float64
	Quantity      float64
	TimeInForce   TimeInForce
	StopPrice     float64
	TrailingDelta int64
	IcebergQty    float64
	Type          OrderType // StopLossLimit or TakeProfitLimit
	Params        StopLimitOrderParams
}

func (o StopLimitOrderRequest) build() map[string]interface{} {
	params := make(map[string]interface{})

	params["symbol"] = o.Symbol
	params["side"] = o.Side
	params["type"] = o.Type

	validation.SetIfNotZero(params, "timeInForce", o.TimeInForce)
	validation.SetIfNotZero(params, "quantity", o.Quantity)
	validation.SetIfNotZero(params, "price", o.Price)
	validation.SetIfNotZero(params, "stopPrice", o.StopPrice)
	validation.SetIfNotZero(params, "trailingDelta", o.TrailingDelta)
	validation.SetIfNotZero(params, "icebergQty", o.IcebergQty)

	validation.SetIfNotZero(params, "newClientOrderId", o.Params.NewClientOrderId)
	validation.SetIfNotZero(params, "strategyId", o.Params.StrategyId)
	validation.SetIfNotZero(params, "strategyType", o.Params.StrategyType)
	validation.SetIfNotZero(params, "newOrderRespType", o.Params.NewOrderRespType)
	validation.SetIfNotZero(params, "selfTradePreventionMode", o.Params.SelfTradePreventionMode)
	validation.SetIfNotZero(params, "pegPriceType", o.Params.PegPriceType)
	validation.SetIfNotZero(params, "pegOffsetValue", o.Params.PegOffsetValue)
	validation.SetIfNotZero(params, "pegOffsetType", o.Params.PegOffsetType)
	validation.SetIfNotZero(params, "recvWindow", o.Params.RecvWindow)

	return params
}

////
// LIMIT Handler
////

type LimitOrderParams struct {
	NewClientOrderId string

	StrategyId   int64
	StrategyType int

	IcebergQty float64

	NewOrderRespType OrderResponseType

	SelfTradePreventionMode STPMode

	PegPriceType   PegPriceType
	PegOffsetValue int
	PegOffsetType  PegOffsetType

	RecvWindow int
}

func NewLimitOrder(symbol string, side OrderSide, price float64, quantity float64, timeInForce TimeInForce, opts ...LimitOrderParams) orderRequest {
	var opt LimitOrderParams
	if len(opts) > 0 {
		opt = opts[0]
	} else {
		opt = LimitOrderParams{}
	}

	return LimitOrderRequest{
		Symbol:      symbol,
		Side:        side,
		TimeInForce: timeInForce,
		Quantity:    quantity,
		Price:       price,
		Params:      opt,
	}
}

func NewLimitBuy(symbol string, price float64, quantity float64, timeInForce TimeInForce, opts ...LimitOrderParams) orderRequest {
	return NewLimitOrder(symbol, Buy, price, quantity, timeInForce, opts...)
}

func NewLimitSell(symbol string, price float64, quantity float64, timeInForce TimeInForce, opts ...LimitOrderParams) orderRequest {
	return NewLimitOrder(symbol, Sell, price, quantity, timeInForce, opts...)
}

////
// LIMIT MAKER Orders
////

type LimitMakerOrderParams struct {
	NewClientOrderId string

	StrategyId   int64
	StrategyType int

	IcebergQty float64

	NewOrderRespType OrderResponseType

	SelfTradePreventionMode STPMode

	PegPriceType   PegPriceType
	PegOffsetValue int
	PegOffsetType  PegOffsetType

	RecvWindow int
}

func NewLimitMakerOrder(symbol string, side OrderSide, price float64, quantity float64, opts ...LimitMakerOrderParams) orderRequest {
	var opt LimitMakerOrderParams
	if len(opts) > 0 {
		opt = opts[0]
	} else {
		opt = LimitMakerOrderParams{}
	}

	return LimitMakerOrderRequest{
		Symbol:   symbol,
		Side:     side,
		Quantity: quantity,
		Price:    price,
		Params:   opt,
	}
}

func NewLimitMakerBuy(symbol string, price float64, quantity float64, opts ...LimitMakerOrderParams) orderRequest {
	return NewLimitMakerOrder(symbol, Buy, price, quantity, opts...)
}

func NewLimitMakerSell(symbol string, price float64, quantity float64, opts ...LimitMakerOrderParams) orderRequest {
	return NewLimitMakerOrder(symbol, Sell, price, quantity, opts...)
}

////
// MARKET Orders
////

type MarketOrderParams struct {
	TimeInForce TimeInForce

	// If set, `quantity` will be ignored
	QuoteOrderQty float64

	NewClientOrderId string

	StrategyId   int64
	StrategyType int

	NewOrderRespType OrderResponseType

	SelfTradePreventionMode STPMode

	RecvWindow int
}

func NewMarketOrder(symbol string, side OrderSide, quantity float64, opts ...MarketOrderParams) orderRequest {
	var opt MarketOrderParams
	if len(opts) > 0 {
		opt = opts[0]
	} else {
		opt = MarketOrderParams{}
	}

	return MarketOrderRequest{
		Symbol:   symbol,
		Side:     side,
		Quantity: quantity,
		Params:   opt,
	}
}

func NewMarketBuy(symbol string, quantity float64, opts ...MarketOrderParams) orderRequest {
	return NewMarketOrder(symbol, Buy, quantity, opts...)
}

func NewMarketSell(symbol string, quantity float64, opts ...MarketOrderParams) orderRequest {
	return NewMarketOrder(symbol, Sell, quantity, opts...)
}

////
// STOP_LOSS Order
////

type StopOrderParams struct {
	NewClientOrderId string

	StrategyId   int64
	StrategyType int

	StopPrice     float64
	TrailingDelta int64

	NewOrderRespType        OrderResponseType
	SelfTradePreventionMode STPMode

	RecvWindow int
}

func NewStopLossOrder(symbol string, side OrderSide, quantity float64, opts ...StopOrderParams) orderRequest {
	var opt StopOrderParams
	if len(opts) > 0 {
		opt = opts[0]
	}

	return StopOrderRequest{
		Symbol:        symbol,
		Side:          side,
		Quantity:      quantity,
		StopPrice:     opt.StopPrice,
		TrailingDelta: opt.TrailingDelta,
		Type:          StopLoss,
		Params:        opt,
	}
}

func NewStopLossBuy(symbol string, quantity float64, opts ...StopOrderParams) orderRequest {
	return NewStopLossOrder(symbol, Buy, quantity, opts...)
}

func NewStopLossSell(symbol string, quantity float64, opts ...StopOrderParams) orderRequest {
	return NewStopLossOrder(symbol, Sell, quantity, opts...)
}

func NewTakeProfitOrder(symbol string, side OrderSide, quantity float64, opts ...StopOrderParams) orderRequest {
	var opt StopOrderParams
	if len(opts) > 0 {
		opt = opts[0]
	}

	return StopOrderRequest{
		Symbol:        symbol,
		Side:          side,
		Quantity:      quantity,
		StopPrice:     opt.StopPrice,
		TrailingDelta: opt.TrailingDelta,
		Type:          TakeProfit,
		Params:        opt,
	}
}

func NewTakeProfitBuy(symbol string, quantity float64, opts ...StopOrderParams) orderRequest {
	return NewTakeProfitOrder(symbol, Buy, quantity, opts...)
}

func NewTakeProfitSell(symbol string, quantity float64, opts ...StopOrderParams) orderRequest {
	return NewTakeProfitOrder(symbol, Sell, quantity, opts...)
}

////
// STOP_LOSS_LIMIT Order
////

type StopLimitOrderParams struct {
	NewClientOrderId string

	StrategyId   int64
	StrategyType int

	StopPrice     float64
	TrailingDelta int64
	IcebergQty    float64

	NewOrderRespType        OrderResponseType
	SelfTradePreventionMode STPMode

	PegPriceType   PegPriceType
	PegOffsetValue int
	PegOffsetType  PegOffsetType

	RecvWindow int
}

func NewStopLossLimitOrder(symbol string, side OrderSide, price float64, quantity float64, timeInForce TimeInForce, opts ...StopLimitOrderParams) orderRequest {
	var opt StopLimitOrderParams
	if len(opts) > 0 {
		opt = opts[0]
	}

	return StopLimitOrderRequest{
		Symbol:        symbol,
		Side:          side,
		Price:         price,
		Quantity:      quantity,
		TimeInForce:   timeInForce,
		StopPrice:     opt.StopPrice,
		TrailingDelta: opt.TrailingDelta,
		IcebergQty:    opt.IcebergQty,
		Type:          StopLossLimit,
		Params:        opt,
	}
}

func NewStopLossLimitBuy(symbol string, price float64, quantity float64, timeInForce TimeInForce, opts ...StopLimitOrderParams) orderRequest {
	return NewStopLossLimitOrder(symbol, Buy, price, quantity, timeInForce, opts...)
}

func NewStopLossLimitSell(symbol string, price float64, quantity float64, timeInForce TimeInForce, opts ...StopLimitOrderParams) orderRequest {
	return NewStopLossLimitOrder(symbol, Sell, price, quantity, timeInForce, opts...)
}

func NewTakeProfitLimitOrder(symbol string, side OrderSide, price float64, quantity float64, timeInForce TimeInForce, opts ...StopLimitOrderParams) orderRequest {
	var opt StopLimitOrderParams
	if len(opts) > 0 {
		opt = opts[0]
	}

	return StopLimitOrderRequest{
		Symbol:        symbol,
		Side:          side,
		Price:         price,
		Quantity:      quantity,
		TimeInForce:   timeInForce,
		StopPrice:     opt.StopPrice,
		TrailingDelta: opt.TrailingDelta,
		IcebergQty:    opt.IcebergQty,
		Type:          TakeProfitLimit,
		Params:        opt,
	}
}

func NewTakeProfitLimitBuy(symbol string, price float64, quantity float64, timeInForce TimeInForce, opts ...StopLimitOrderParams) orderRequest {
	return NewTakeProfitLimitOrder(symbol, Buy, price, quantity, timeInForce, opts...)
}

func NewTakeProfitLimitSell(symbol string, price float64, quantity float64, timeInForce TimeInForce, opts ...StopLimitOrderParams) orderRequest {
	return NewTakeProfitLimitOrder(symbol, Sell, price, quantity, timeInForce, opts...)
}

////
// CANCEL Order
////

type CancelOrderParams struct {
	OrigClientOrderId string
	NewClientOrderId  string

	CancelRestrictions CancelRestrictions

	RecvWindow int
}

func (c *Client) CancelOrder(symbol string, orderId int64, opts ...CancelOrderParams) (*Order, Response, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol
	validation.SetIfNotZero(params, "orderId", orderId)

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "origClientOrderId", opt.OrigClientOrderId)
		validation.SetIfNotZero(params, "newClientOrderId", opt.NewClientOrderId)

		validation.SetIfNotZero(params, "cancelRestrictions", opt.CancelRestrictions)

		validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	}

	return doRequest[Order](c, Methods.DELETE, "/api/v3/order", params, nil, TRADE)
}

////
// Cancel Orders
////

type CancelOpenOrdersParams struct {
	RecvWindow int
}

func (c *Client) CancelOpenOrders(symbol string, opts ...CancelOpenOrdersParams) ([]*Order, Response, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	}

	result, resp, err := doRequest[[]*Order](c, Methods.DELETE, "/api/v3/openOrders", params, nil, TRADE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

////
// Cancel Replace Order
////

// TODO
