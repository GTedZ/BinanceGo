package futures

import (
	"encoding/json"

	"github.com/GTedZ/binancego/internal/validation"
)

////
// Futures Order Response
////

type Order struct {
	ClientOrderId string `json:"clientOrderId"`

	CumQty      float64 `json:"cumQty,string"`
	CumQuote    float64 `json:"cumQuote,string"`
	ExecutedQty float64 `json:"executedQty,string"`

	OrderId int64 `json:"orderId"`

	AvgPrice float64 `json:"avgPrice,string"`
	OrigQty  float64 `json:"origQty,string"`
	Price    float64 `json:"price,string"`

	ReduceOnly bool `json:"reduceOnly"`

	Side         OrderSide    `json:"side"`
	PositionSide PositionSide `json:"positionSide"`

	Status OrderStatus `json:"status"`

	StopPrice     float64 `json:"stopPrice,string"`
	ClosePosition bool    `json:"closePosition"`

	Symbol string `json:"symbol"`

	TimeInForce TimeInForce `json:"timeInForce"`
	Type        OrderType   `json:"type"`
	OrigType    OrderType   `json:"origType"`

	UpdateTime int64 `json:"updateTime"`

	WorkingType  string `json:"workingType"`
	PriceProtect bool   `json:"priceProtect"`

	PriceMatch              PriceMatch `json:"priceMatch"`
	SelfTradePreventionMode STPMode    `json:"selfTradePreventionMode"`
	GoodTillDate            int64      `json:"goodTillDate"`
}

////
// Batch Order Result
////

// BatchOrderResult supports Binance's mixed batch response:
// - successful order object
// - error object: {"code": -2022, "msg": "..."}
type BatchOrderResult struct {
	Order

	Code int    `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

func (r BatchOrderResult) IsError() bool {
	return r.Code != 0 || r.Msg != ""
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
	return doRequest[Order](c, Methods.POST, "/fapi/v1/order", params, nil, TRADE)
}

func (c *Client) batchOrders(params map[string]interface{}) ([]*BatchOrderResult, Response, Error) {
	result, resp, err := doRequest[[]*BatchOrderResult](c, Methods.POST, "/fapi/v1/batchOrders", params, nil, TRADE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

func (c *Client) modifyOrder(params map[string]interface{}) (*Order, Response, Error) {
	return doRequest[Order](c, Methods.PUT, "/fapi/v1/order", params, nil, TRADE)
}

func (c *Client) batchModifyOrders(params map[string]interface{}) ([]*BatchOrderResult, Response, Error) {
	result, resp, err := doRequest[[]*BatchOrderResult](c, Methods.PUT, "/fapi/v1/batchOrders", params, nil, TRADE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

func (c *Client) cancelOrder(params map[string]interface{}) (*Order, Response, Error) {
	return doRequest[Order](c, Methods.DELETE, "/fapi/v1/order", params, nil, TRADE)
}

func (c *Client) batchCancelOrders(params map[string]interface{}) ([]*BatchOrderResult, Response, Error) {
	result, resp, err := doRequest[[]*BatchOrderResult](c, Methods.DELETE, "/fapi/v1/batchOrders", params, nil, TRADE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

////
// Public Order Handlers
////

// Order sends a new futures order request to Binance.
//
// Endpoint:
// POST /fapi/v1/order
//
// Supported here:
// - LIMIT
// - MARKET
//
// Timestamp is expected to be handled by doRequest for TRADE requests,
// same as the spot implementation.
func (c *Client) Order(req orderRequest) (*Order, Response, Error) {
	return c.order(req.build())
}

// BatchOrders places multiple futures orders.
//
// Endpoint:
// POST /fapi/v1/batchOrders
//
// Binance allows max 5 orders.
// This wrapper keeps validation thin and lets Binance reject invalid combinations.
func (c *Client) BatchOrders(reqs ...orderRequest) ([]*BatchOrderResult, Response, Error) {
	return c.BatchOrdersWithParams(0, reqs...)
}

func (c *Client) BatchOrdersWithParams(recvWindow int64, reqs ...orderRequest) ([]*BatchOrderResult, Response, Error) {
	orders := make([]map[string]interface{}, 0, len(reqs))

	for _, req := range reqs {
		orders = append(orders, req.build())
	}

	batchOrdersJson, _ := json.Marshal(orders)

	params := make(map[string]interface{})
	params["batchOrders"] = string(batchOrdersJson)

	validation.SetIfNotZero(params, "recvWindow", recvWindow)

	return c.batchOrders(params)
}

////
// LIMIT Order
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

	validation.SetIfNotZero(params, "positionSide", o.Params.PositionSide)
	validation.SetIfNotZero(params, "timeInForce", o.TimeInForce)
	validation.SetIfNotZero(params, "quantity", o.Quantity)

	// priceMatch cannot be sent together with price.
	if validation.IsNotDefault(o.Params.PriceMatch) {
		validation.SetIfNotZero(params, "priceMatch", o.Params.PriceMatch)
	} else {
		validation.SetIfNotZero(params, "price", o.Price)
	}

	if o.Params.ReduceOnly {
		params["reduceOnly"] = true
	}

	validation.SetIfNotZero(params, "newClientOrderId", o.Params.NewClientOrderId)
	validation.SetIfNotZero(params, "newOrderRespType", o.Params.NewOrderRespType)
	validation.SetIfNotZero(params, "selfTradePreventionMode", o.Params.SelfTradePreventionMode)
	validation.SetIfNotZero(params, "goodTillDate", o.Params.GoodTillDate)
	validation.SetIfNotZero(params, "recvWindow", o.Params.RecvWindow)

	return params
}

type LimitOrderParams struct {
	PositionSide PositionSide

	// Cannot be sent in Hedge Mode.
	ReduceOnly bool

	NewClientOrderId string
	NewOrderRespType OrderResponseType

	// Cannot be sent together with Price.
	PriceMatch PriceMatch

	SelfTradePreventionMode STPMode

	// Mandatory when TimeInForce is GTD.
	GoodTillDate int64

	RecvWindow int64
}

func NewLimitOrder(
	symbol string,
	side OrderSide,
	price float64,
	quantity float64,
	timeInForce TimeInForce,
	opts ...LimitOrderParams,
) orderRequest {
	var opt LimitOrderParams
	if len(opts) > 0 {
		opt = opts[0]
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
// MARKET Order
////

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

	validation.SetIfNotZero(params, "positionSide", o.Params.PositionSide)
	validation.SetIfNotZero(params, "quantity", o.Quantity)

	if o.Params.ReduceOnly {
		params["reduceOnly"] = true
	}

	validation.SetIfNotZero(params, "newClientOrderId", o.Params.NewClientOrderId)
	validation.SetIfNotZero(params, "newOrderRespType", o.Params.NewOrderRespType)
	validation.SetIfNotZero(params, "selfTradePreventionMode", o.Params.SelfTradePreventionMode)
	validation.SetIfNotZero(params, "recvWindow", o.Params.RecvWindow)

	return params
}

type MarketOrderParams struct {
	PositionSide PositionSide

	// Cannot be sent in Hedge Mode.
	ReduceOnly bool

	NewClientOrderId string
	NewOrderRespType OrderResponseType

	SelfTradePreventionMode STPMode

	RecvWindow int64
}

func NewMarketOrder(symbol string, side OrderSide, quantity float64, opts ...MarketOrderParams) orderRequest {
	var opt MarketOrderParams
	if len(opts) > 0 {
		opt = opts[0]
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
// Modify Order
////

type modifyOrderRequest interface {
	build() map[string]interface{}
}

type ModifyOrderRequest struct {
	OrderId      int64
	OrderRequest orderRequest
}

func (r ModifyOrderRequest) build() map[string]interface{} {
	params := r.OrderRequest.build()
	params["orderId"] = r.OrderId
	return params
}

func (c *Client) ModifyOrder(modifyOrderRequest modifyOrderRequest) (*Order, Response, Error) {
	return c.modifyOrder(modifyOrderRequest.build())
}

func (c *Client) BatchModifyOrders(modifyOrderRequests ...modifyOrderRequest) ([]*BatchOrderResult, Response, Error) {
	// Building orders array
	orders := make([]map[string]interface{}, len(modifyOrderRequests))
	for i, req := range modifyOrderRequests {
		orders[i] = req.build()
	}

	params := make(map[string]interface{})
	params["batchOrders"] = orders
	return c.batchModifyOrders(params)
}

////
// Order Amendment History
////

type OrderAmendmentEntry struct {
	AmendmentID   int64          `json:"amendmentId"`
	Symbol        string         `json:"symbol"`
	Pair          string         `json:"pair"`
	OrderID       int64          `json:"orderId"`
	ClientOrderID string         `json:"clientOrderId"`
	Time          int64          `json:"time"`
	Amendment     OrderAmendment `json:"amendment"`
}

type OrderAmendment struct {
	Price   Change `json:"price"`
	OrigQty Change `json:"origQty"`
	Count   int    `json:"count"`
}

type Change struct {
	Before string `json:"before"`
	After  string `json:"after"`
}

type OrderAmendmentParams struct {
	OrderId           int64  `json:"orderId"`
	OrigClientOrderId string `json:"origClientOrderId"`
	StartTime         int64  `json:"startTime"`
	EndTime           int64  `json:"endTime"`
	Limit             int    `json:"limit"`
	RecvWindow        int64  `json:"recvWindow"`
}

func OrderModifyHistory(symbol string, opts ...OrderAmendmentParams) ([]*OrderAmendmentEntry, Response, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol

	if len(opts) > 0 {
		opt := opts[0]
		validation.SetIfNotZero(params, "orderId", opt.OrderId)
		validation.SetIfNotZero(params, "origClientOrderId", opt.OrigClientOrderId)
		validation.SetIfNotZero(params, "startTime", opt.StartTime)
		validation.SetIfNotZero(params, "endTime", opt.EndTime)
		validation.SetIfNotZero(params, "limit", opt.Limit)
		validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	}

	result, resp, err := doRequest[[]*OrderAmendmentEntry](nil, Methods.GET, "/fapi/v1/orderAmendment", params, nil, TRADE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

////
// Cancel Order
////

type CancelOrderRequest struct {
	OrderId       int64
	ClientOrderId string
	RecvWindow    int64
}

func (c *Client) CancelOrder(req CancelOrderRequest) (*Order, Response, Error) {
	params := make(map[string]interface{})
	validation.SetIfNotZero(params, "orderId", req.OrderId)
	validation.SetIfNotZero(params, "origClientOrderId", req.ClientOrderId)
	validation.SetIfNotZero(params, "recvWindow", req.RecvWindow)
	return c.cancelOrder(params)
}

func (c *Client) BatchCancelOrders(req []CancelOrderRequest) ([]*BatchOrderResult, Response, Error) {
	params := make(map[string]interface{})
	cancelOrders := make([]map[string]interface{}, len(req))
	for i, r := range req {
		cancelOrders[i] = make(map[string]interface{})
		validation.SetIfNotZero(cancelOrders[i], "orderId", r.OrderId)
		validation.SetIfNotZero(cancelOrders[i], "origClientOrderId", r.ClientOrderId)
		validation.SetIfNotZero(cancelOrders[i], "recvWindow", r.RecvWindow)
	}
	params["batchOrders"] = cancelOrders
	return c.batchCancelOrders(params)
}

func (c *Client) CancelAllOpenOrders(symbol string) ([]*BatchOrderResult, Response, Error) {
	params := make(map[string]interface{})
	params["symbol"] = symbol

	result, resp, err := doRequest[[]*BatchOrderResult](c, Methods.DELETE, "/fapi/v1/allOpenOrders", params, nil, TRADE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

////
// Auto-Cancel All Open Orders
////

type AutoCancelAllOpenOrders struct {
	Symbol        string
	CountdownTime int64
}

type AutoCancelAllOpenOrdersParams struct {
	Symbol        string
	CountdownTime int64
	RecvWindow    int64
}

// CountdownCancelAllOrders schedules the cancellation of all open orders for the
// specified symbol after the given countdown period.
//
// This endpoint is designed to be called periodically as a heartbeat. Each call
// replaces any existing countdown with the new countdownTime value.
//
// Example:
//   - Call this endpoint every 30 seconds with countdownTime=120000 (120s).
//   - If no heartbeat is received within 120 seconds, all open orders for the
//     symbol are automatically canceled.
//   - Set countdownTime=0 to stop and clear the active countdown timer.
//
// The exchange checks countdown timers approximately every 10ms. Because timer
// execution is not guaranteed to be exact, applications should include sufficient
// timing redundancy and avoid relying on highly precise or very short countdowns.
func (c *Client) AutoCancelAllOpenOrders(symbol string, recvWindow int64) (*AutoCancelAllOpenOrders, Response, Error) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	validation.SetIfNotZero(params, "recvWindow", recvWindow)

	return doRequest[AutoCancelAllOpenOrders](c, Methods.POST, "/fapi/v1/countdownCancelAll", params, nil, TRADE)
}

////
// Query Order
////

type OrderParams struct {
	OrderId           int64
	OrigClientOrderId string
	RecvWindow        int64
}

// GetOrder checks the status of an order.
//
// Note that some historical orders are not queryable:
//
//   - CANCELED or EXPIRED orders with no fills are removed 3 days after
//     creation.
//   - All orders are removed 90 days after creation.
func (c *Client) GetOrder(symbol string, opt OrderParams) (*Order, Response, Error) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	validation.SetIfNotZero(params, "orderId", opt.OrderId)
	validation.SetIfNotZero(params, "origClientOrderId", opt.OrigClientOrderId)
	validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)

	return doRequest[Order](c, Methods.GET, "/fapi/v1/order", params, nil, TRADE)
}

////
// All Orders
////

type AllOrdersParams struct {
	OrderId    int64
	StartTime  int64
	EndTime    int64
	Limit      int
	RecvWindow int64
}

func (c *Client) QueryAllOrders(symbol string, opt AllOrdersParams) ([]*Order, Response, Error) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	validation.SetIfNotZero(params, "orderId", opt.OrderId)
	validation.SetIfNotZero(params, "startTime", opt.StartTime)
	validation.SetIfNotZero(params, "endTime", opt.EndTime)
	validation.SetIfNotZero(params, "limit", opt.Limit)
	validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)

	result, resp, err := doRequest[[]*Order](c, Methods.GET, "/fapi/v1/allOrders", params, nil, TRADE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

////
// Query All Current Open Orders
////

type CurrentOpenOrdersParams struct {
	RecvWindow int64
}

func (c *Client) SymbolOpenOrders(opt CurrentOpenOrdersParams) ([]*Order, Response, Error) {
	params := make(map[string]interface{})
	validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	result, resp, err := doRequest[[]*Order](c, Methods.GET, "/fapi/v1/openOrders", params, nil, TRADE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

func (c *Client) AllOpenOrders(opt CurrentOpenOrdersParams) ([]*Order, Response, Error) {
	params := make(map[string]interface{})
	validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	result, resp, err := doRequest[[]*Order](c, Methods.GET, "/fapi/v1/openOrders", params, nil, TRADE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

////
// Current Open Order
////

type CurrentOpenOrderParams struct {
	OrderId           int64
	OrigClientOrderId string
	RecvWindow        int64
}

func (c *Client) CurrentOpenOrder(symbol string, opt CurrentOpenOrderParams) (*Order, Response, Error) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	validation.SetIfNotZero(params, "orderId", opt.OrderId)
	validation.SetIfNotZero(params, "origClientOrderId", opt.OrigClientOrderId)
	validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)

	return doRequest[Order](c, Methods.GET, "/fapi/v1/openOrder", params, nil, TRADE)
}

////
// Users Force Orders
////

type ForceOrdersParams struct {
	AutoCloseType AutoCloseType
	StartTime     int64
	EndTime       int64
	Limit         int
	RecvWindow    int64
}

func (c *Client) ForceOrders(symbol string, opt ForceOrdersParams) ([]*Order, Response, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol
	validation.SetIfNotZero(params, "autoCloseType", opt.AutoCloseType)
	validation.SetIfNotZero(params, "startTime", opt.StartTime)
	validation.SetIfNotZero(params, "endTime", opt.EndTime)
	validation.SetIfNotZero(params, "limit", opt.Limit)
	validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)

	result, resp, err := doRequest[[]*Order](c, Methods.GET, "/fapi/v1/forceOrders", params, nil, TRADE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

func (c *Client) AllForceOrders(opt ForceOrdersParams) ([]*Order, Response, Error) {
	params := make(map[string]interface{})

	validation.SetIfNotZero(params, "autoCloseType", opt.AutoCloseType)
	validation.SetIfNotZero(params, "startTime", opt.StartTime)
	validation.SetIfNotZero(params, "endTime", opt.EndTime)
	validation.SetIfNotZero(params, "limit", opt.Limit)
	validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)

	result, resp, err := doRequest[[]*Order](c, Methods.GET, "/fapi/v1/forceOrders", params, nil, TRADE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

////
// Account Trade List
////

type AccountTradeListParams struct {
	OrderId    int64
	StartTime  int64
	EndTime    int64
	FromId     int64
	Limit      int
	RecvWindow int64
}

func (c *Client) AccountTradeList(symbol string, opt AccountTradeListParams) ([]*Trade, Response, Error) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	validation.SetIfNotZero(params, "orderId", opt.OrderId)
	validation.SetIfNotZero(params, "startTime", opt.StartTime)
	validation.SetIfNotZero(params, "endTime", opt.EndTime)
	validation.SetIfNotZero(params, "fromId", opt.FromId)
	validation.SetIfNotZero(params, "limit", opt.Limit)
	validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)

	result, resp, err := doRequest[[]*Trade](c, Methods.GET, "/fapi/v1/userTrades", params, nil, TRADE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

////
// Change Margin Type
////

type ChangeMarginType struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type ChangeMarginTypeParams struct {
	RecvWindow int64
}

func (c *Client) ChangeMarginType(symbol string, marginType MarginType, opts ...ChangeMarginTypeParams) (*ChangeMarginType, Response, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol
	params["marginType"] = marginType

	if len(opts) > 0 {
		opt := opts[0]
		validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	}

	return doRequest[ChangeMarginType](c, Methods.POST, "/fapi/v1/marginType", params, nil, TRADE)
}

////
// Change Position Mode
////

type ChangePositionMode struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type ChangePositionModeParams struct {
	RecvWindow int64
}

// ChangePositionMode switches between One-way Mode and Hedge Mode for the user's futures account.
//
//   - `true`: Hedge Mode
//   - `false`: One-way Mode
//
// In One-way Mode, a symbol has a single position that can be long or short.
//
// In Hedge Mode, a symbol can have two separate positions: one long and one short.
func (c *Client) ChangePositionMode(dualSidePosition bool, opts ...ChangePositionModeParams) (*ChangePositionMode, Response, Error) {
	params := make(map[string]interface{})
	params["dualSidePosition"] = dualSidePosition

	if len(opts) > 0 {
		opt := opts[0]
		validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	}

	return doRequest[ChangePositionMode](c, Methods.POST, "/fapi/v1/positionSide/dual", params, nil, TRADE)
}

////
// Change Initial Leverage
////

type ChangeInitialLeverage struct {
	Leverage         int     `json:"leverage"`
	MaxNotionalValue float64 `json:"maxNotionalValue,string"`
	Symbol           string  `json:"symbol"`
}

type ChangeInitialLeverageParams struct {
	RecvWindow int64
}

func (c *Client) ChangeInitialLeverage(symbol string, leverage int, opts ...ChangeInitialLeverageParams) (*ChangeInitialLeverage, Response, Error) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	params["leverage"] = leverage

	if len(opts) > 0 {
		opt := opts[0]
		validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	}

	return doRequest[ChangeInitialLeverage](c, Methods.POST, "/fapi/v1/leverage", params, nil, TRADE)
}

////
// Change Multi-Assets Mode
////

type ChangeMultiAssetsMode struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type ChangeMultiAssetsModeParams struct {
	RecvWindow int64
}

// ChangeMultiAssetsMode enables or disables Multi-Assets Mode for the user's futures account.
//
//   - `true`: Enable Multi-Assets Mode
//   - `false`: Disable Multi-Assets Mode
func (c *Client) ChangeMultiAssetsMode(multiAssetsMargin bool, opts ...ChangeMultiAssetsModeParams) (*ChangeMultiAssetsMode, Response, Error) {
	params := make(map[string]interface{})
	params["multiAssetsMargin"] = multiAssetsMargin

	if len(opts) > 0 {
		opt := opts[0]
		validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	}

	return doRequest[ChangeMultiAssetsMode](c, Methods.POST, "/fapi/v1/multiAssetsMargin", params, nil, TRADE)
}

////
// Modify Isolated Position Margin
////

type ModifyIsolatedPositionMargin struct {
	Amount float64                `json:"amount,string"`
	Code   int                    `json:"code"`
	Msg    string                 `json:"msg"`
	Type   MarginModificationType `json:"type"`
}

type ModifyIsolatedPositionMarginParams struct {
	RecvWindow int64
}

// ModifyIsolatedPositionMargin adjusts the margin of an isolated position.
//
// Only for isolated symbols. Cannot be used in Hedge Mode.
func (c *Client) ModifyIsolatedPositionMargin(symbol string, positionSide PositionSide, amount float64, modificationType MarginModificationType, opts ...ModifyIsolatedPositionMarginParams) (*ModifyIsolatedPositionMargin, Response, Error) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	params["positionSide"] = positionSide
	params["amount"] = amount
	params["type"] = modificationType

	if len(opts) > 0 {
		opt := opts[0]
		validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	}

	return doRequest[ModifyIsolatedPositionMargin](c, Methods.POST, "/fapi/v1/positionMargin", params, nil, TRADE)
}

////
// Position Information V2
////

type PositionInformationV2 struct {
	EntryPrice       float64      `json:"entryPrice,string"`
	BreakEvenPrice   float64      `json:"breakEvenPrice,string"`
	MarginType       MarginType   `json:"marginType"`
	IsAutoAddMargin  bool         `json:"isAutoAddMargin,string"`
	IsolatedMargin   float64      `json:"isolatedMargin,string"`
	Leverage         int          `json:"leverage,string"`
	LiquidationPrice float64      `json:"liquidationPrice,string"`
	MarkPrice        float64      `json:"markPrice,string"`
	MaxNotionalValue float64      `json:"maxNotionalValue,string"`
	PositionAmt      float64      `json:"positionAmt,string"`
	Notional         float64      `json:"notional,string"`
	IsolatedWallet   float64      `json:"isolatedWallet,string"`
	Symbol           string       `json:"symbol"`
	UnRealizedProfit float64      `json:"unRealizedProfit,string"`
	PositionSide     PositionSide `json:"positionSide"`
	UpdateTime       int64        `json:"updateTime"`
}

type PositionInformationV2Params struct {
	RecvWindow int64
}

func (c *Client) PositionInformationV2(symbol string, positionSide PositionSide, opts ...PositionInformationV2Params) ([]*PositionInformationV2, Response, Error) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	params["positionSide"] = positionSide

	if len(opts) > 0 {
		opt := opts[0]
		validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	}

	result, resp, err := doRequest[[]*PositionInformationV2](c, Methods.GET, "/fapi/v1/positionRisk", params, nil, USER_DATA)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

func (c *Client) AllPositionInformationV2(positionSide PositionSide, opts ...PositionInformationV2Params) ([]*PositionInformationV2, Response, Error) {
	params := make(map[string]interface{})
	params["positionSide"] = positionSide

	if len(opts) > 0 {
		opt := opts[0]
		validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	}

	result, resp, err := doRequest[[]*PositionInformationV2](c, Methods.GET, "/fapi/v1/positionRisk", params, nil, USER_DATA)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

////
// Position Information V3
////

type PositionInformationV3 struct {
	// Symbol
	Symbol string `json:"symbol"`
	// Position Side
	PositionSide PositionSide `json:"positionSide"`
	// Position amount, positive for long, negative for short
	PositionAmt float64 `json:"positionAmt,string"`
	// entry price
	EntryPrice float64 `json:"entryPrice,string"`
	// break-even price
	BreakEvenPrice float64 `json:"breakEvenPrice,string"`
	// current mark price
	MarkPrice float64 `json:"markPrice,string"`
	// unrealized profit
	UnRealizedProfit float64 `json:"unRealizedProfit,string"`
	// liquidation price
	LiquidationPrice float64 `json:"liquidationPrice,string"`
	// isolated margin
	IsolatedMargin float64 `json:"isolatedMargin,string"`
	// notional value of position
	Notional float64 `json:"notional,string"`
	// margin asset
	MarginAsset string `json:"marginAsset"`
	// isolated wallet (if isolated position)
	IsolatedWallet float64 `json:"isolatedWallet,string"`
	// initial margin required with current mark price
	InitialMargin float64 `json:"initialMargin,string"`
	// maintenance margin required
	MaintMargin float64 `json:"maintMargin,string"`
	// initial margin required for positions with current mark price
	PositionInitialMargin float64 `json:"positionInitialMargin,string"`
	// initial margin required for open orders with current mark price
	OpenOrderInitialMargin float64 `json:"openOrderInitialMargin,string"`
	// auto-deleverage ranking
	Adl int `json:"adl"`
	// ignore
	BidNotional float64 `json:"bidNotional,string"`
	// ignore
	AskNotional float64 `json:"askNotional,string"`
	// update time
	UpdateTime int64 `json:"updateTime"`
}

type PositionInformationV3Params struct {
	RecvWindow int64
}

func (c *Client) PositionInformationV3(symbol string, positionSide PositionSide, opts ...PositionInformationV3Params) ([]*PositionInformationV3, Response, Error) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	params["positionSide"] = positionSide

	if len(opts) > 0 {
		opt := opts[0]
		validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	}

	result, resp, err := doRequest[[]*PositionInformationV3](c, Methods.GET, "/fapi/v1/positionRisk", params, nil, USER_DATA)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

func (c *Client) AllPositionInformationV3(positionSide PositionSide, opts ...PositionInformationV3Params) ([]*PositionInformationV3, Response, Error) {
	params := make(map[string]interface{})
	params["positionSide"] = positionSide

	if len(opts) > 0 {
		opt := opts[0]
		validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	}

	result, resp, err := doRequest[[]*PositionInformationV3](c, Methods.GET, "/fapi/v1/positionRisk", params, nil, USER_DATA)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

////
// Position ADL Quantile Estimation
////

type AdlQuantileEstimation struct {
	Symbol      string                     `json:"symbol"`
	AdlQuantile AdlQuantileEstimationValue `json:"adlQuantile"`
}

type AdlQuantileEstimationValue struct {
	Long  int `json:"LONG"`
	Short int `json:"SHORT"`
	Both  int `json:"BOTH,omitempty"`
	Hedge int `json:"HEDGE,omitempty"`
}

type AdlQuantileParams struct {
	RecvWindow int64
}

func (c *Client) AllADLQuantileEstimations(opts ...AdlQuantileParams) ([]*AdlQuantileEstimation, Response, Error) {
	params := make(map[string]interface{})
	if len(opts) > 0 {
		opt := opts[0]
		validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	}
	result, resp, err := doRequest[[]*AdlQuantileEstimation](c, Methods.GET, "/fapi/v1/adlQuantile", params, nil, USER_DATA)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

func (c *Client) ADLQuantileEstimations(symbol string, opts ...AdlQuantileParams) ([]*AdlQuantileEstimation, Response, Error) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	if len(opts) > 0 {
		opt := opts[0]
		validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	}
	result, resp, err := doRequest[[]*AdlQuantileEstimation](c, Methods.GET, "/fapi/v1/adlQuantile", params, nil, USER_DATA)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

////
// Position Margin Change History
////

// Get Position Margin Change History (TRADE)
// API Description
// Get Position Margin Change History

// HTTP Request
// GET /fapi/v1/positionMargin/history

// Request Weight
// 1

// Request Parameters
// Name	Type	Mandatory	Description
// symbol	STRING	YES
// type	INT	NO	1: Add position margin，2: Reduce position margin
// startTime	LONG	NO
// endTime	LONG	NO	Default current time if not pass
// limit	INT	NO	Default: 500
// recvWindow	LONG	NO
// timestamp	LONG	YES
// Support querying future histories that are not older than 30 days
// The time between startTime and endTimecan't be more than 30 days
// Response Example
// [
// 	{
// 	  	"symbol": "BTCUSDT",
// 	  	"type": 1,
// 		"deltaType": "USER_ADJUST",
// 		"amount": "23.36332311",
// 	  	"asset": "USDT",
// 	  	"time": 1578047897183,
// 	  	"positionSide": "BOTH"
// 	},
// 	{
// 		"symbol": "BTCUSDT",
// 	  	"type": 1,
// 		"deltaType": "USER_ADJUST",
// 		"amount": "100",
// 	  	"asset": "USDT",
// 	  	"time": 1578047900425,
// 	  	"positionSide": "LONG"
// 	}
// ]

type PositionMarginChangeHistory struct {
	Symbol       string                 `json:"symbol"`
	Type         MarginModificationType `json:"type"`
	DeltaType    string                 `json:"deltaType"`
	Amount       float64                `json:"amount,string"`
	Asset        string                 `json:"asset"`
	Time         int64                  `json:"time"`
	PositionSide PositionSide           `json:"positionSide"`
}

type PositionMarginChangeHistoryParams struct {
	Type       MarginModificationType
	StartTime  int64
	EndTime    int64
	Limit      int
	RecvWindow int64
}

func (c *Client) PositionMarginChangeHistory(symbol string, opt PositionMarginChangeHistoryParams) ([]*PositionMarginChangeHistory, Response, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol
	validation.SetIfNotZero(params, "type", opt.Type)
	validation.SetIfNotZero(params, "startTime", opt.StartTime)
	validation.SetIfNotZero(params, "endTime", opt.EndTime)
	validation.SetIfNotZero(params, "limit", opt.Limit)
	validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)

	result, resp, err := doRequest[[]*PositionMarginChangeHistory](c, Methods.GET, "/fapi/v1/positionMargin/history", params, nil, TRADE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}
