package spot

import (
	"fmt"

	"github.com/GTedZ/binancego/internal/validation"
)

// Account
type Account struct {
	MakerCommission  int64 `json:"makerCommission"`
	TakerCommission  int64 `json:"takerCommission"`
	BuyerCommission  int64 `json:"buyerCommission"`
	SellerCommission int64 `json:"sellerCommission"`

	CommissionRates CommissionRates `json:"commissionRates"`

	CanTrade    bool `json:"canTrade"`
	CanWithdraw bool `json:"canWithdraw"`
	CanDeposit  bool `json:"canDeposit"`

	Brokered                   bool `json:"brokered"`
	RequireSelfTradePrevention bool `json:"requireSelfTradePrevention"`
	PreventSor                 bool `json:"preventSor"`

	UpdateTime  int64  `json:"updateTime"`
	AccountType string `json:"accountType"`

	Balances []Balance `json:"balances"`

	Permissions []Permission `json:"permissions"`

	UID int64 `json:"uid"`
}

type CommissionRates struct {
	Maker  float64 `json:"maker,string"`
	Taker  float64 `json:"taker,string"`
	Buyer  float64 `json:"buyer,string"`
	Seller float64 `json:"seller,string"`
}

type Balance struct {
	Asset  string  `json:"asset"`
	Free   float64 `json:"free,string"`
	Locked float64 `json:"locked,string"`
}

type AccountParams struct {
	OmitZeroBalances bool
	RecvWindow       int
}

func (c *Client) Account(opts ...AccountParams) (*Account, Response, Error) {
	params := make(map[string]interface{})

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "omitZeroBalances", opt.OmitZeroBalances)
		validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	}

	return doRequest[Account](c, Methods.GET, "/api/v3/account", params, nil, USER_DATA)
}

// Order
type QueryOrderParams struct {
	OrderID           int64
	OrigClientOrderID string
	RecvWindow        int
}

func (c *Client) QueryOrder(symbol string, opts ...QueryOrderParams) (*Order, Response, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "orderId", opt.OrderID)
		validation.SetIfNotZero(params, "origClientOrderId", opt.OrigClientOrderID)
		validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	}

	return doRequest[Order](c, Methods.GET, "/api/v3/order", params, nil, USER_DATA)
}

// Open Orders

type OpenOrdersParams struct {
	Symbol     string
	RecvWindow int
}

func (c *Client) OpenOrders(opts ...OpenOrdersParams) ([]*Order, Response, Error) {
	params := make(map[string]interface{})

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "symbol", opt.Symbol)
		validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	}

	result, resp, err := doRequest[[]*Order](c, Methods.GET, "/api/v3/openOrders", params, nil, USER_DATA)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

type OpenOrdersBySymbolParams struct {
	RecvWindow int
}

func (c *Client) OpenOrdersBySymbol(symbol string, opts ...OpenOrdersBySymbolParams) ([]*Order, Response, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	}

	result, resp, err := doRequest[[]*Order](c, Methods.GET, "/api/v3/openOrders", params, nil, USER_DATA)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

// All Orders

type AllOrdersParams struct {
	OrderID    int64
	StartTime  int64
	EndTime    int64
	Limit      int
	RecvWindow int
}

func (c *Client) AllOrders(symbol string, opts ...AllOrdersParams) ([]*Order, Response, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "orderId", opt.OrderID)
		validation.SetIfNotZero(params, "startTime", opt.StartTime)
		validation.SetIfNotZero(params, "endTime", opt.EndTime)
		validation.SetIfNotZero(params, "limit", opt.Limit)
		validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	}

	result, resp, err := doRequest[[]*Order](c, Methods.GET, "/api/v3/allOrders", params, nil, USER_DATA)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

// Order List

type OrderList struct {
	OrderListID int64 `json:"orderListId"`

	ContingencyType ContingencyType `json:"contingencyType"`
	ListStatusType  ListStatusType  `json:"listStatusType"`
	ListOrderStatus ListOrderStatus `json:"listOrderStatus"`

	ListClientOrderID string `json:"listClientOrderId"`

	TransactionTime int64 `json:"transactionTime"`

	Symbol string `json:"symbol"`

	Orders []OrderListItem `json:"orders"`
}

type OrderListItem struct {
	Symbol string `json:"symbol"`

	OrderId       int64  `json:"orderId"`
	ClientOrderId string `json:"clientOrderId"`
}

type OrderListByOrderIdParams struct {
	RecvWindow int
}

func (c *Client) OrderListByOrderListId(orderListID int64, opts ...OrderListByOrderIdParams) (*OrderList, Response, Error) {
	params := map[string]interface{}{
		"orderListId": orderListID,
	}

	if len(opts) > 0 {
		opt := opts[0]
		validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	}

	return doRequest[OrderList](c, Methods.GET, "/api/v3/orderList", params, nil, USER_DATA)
}

type OrderListByClientOrderIDParams struct {
	RecvWindow int
}

func (c *Client) OrderListByClientOrderId(clientOrderId string, opts ...OrderListByClientOrderIDParams) (*OrderList, Response, Error) {
	params := map[string]interface{}{
		"origClientOrderId": clientOrderId,
	}

	if len(opts) > 0 {
		opt := opts[0]
		validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	}

	return doRequest[OrderList](c, Methods.GET, "/api/v3/orderList", params, nil, USER_DATA)
}

// Query All Order Lists

type AllOrderListsParams struct {
	FromID     int64
	StartTime  int64
	EndTime    int64
	Limit      int
	RecvWindow int
}

func (p AllOrderListsParams) Validate() error {
	if p.FromID != 0 && (p.StartTime != 0 || p.EndTime != 0) {
		return fmt.Errorf("fromId cannot be used with startTime or endTime")
	}

	if p.StartTime != 0 && p.EndTime != 0 {
		if p.EndTime-p.StartTime > 24*60*60*1000 {
			return fmt.Errorf("time range cannot exceed 24 hours")
		}
	}

	return nil
}

func (c *Client) AllOrderLists(opts ...AllOrderListsParams) ([]*OrderList, Response, Error) {
	params := make(map[string]interface{})

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "fromId", opt.FromID)
		validation.SetIfNotZero(params, "startTime", opt.StartTime)
		validation.SetIfNotZero(params, "endTime", opt.EndTime)
		validation.SetIfNotZero(params, "limit", opt.Limit)
		validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	}

	result, resp, err := doRequest[[]*OrderList](c, Methods.GET, "/api/v3/allOrderList", params, nil, USER_DATA)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

// Open Order Lists

type OpenOrderListsParams struct {
	RecvWindow int
}

func (c *Client) OpenOrderLists(opts ...OpenOrderListsParams) ([]*OrderList, Response, Error) {
	params := make(map[string]interface{})

	if len(opts) > 0 {
		opt := opts[0]
		validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	}

	result, resp, err := doRequest[[]*OrderList](c, Methods.GET, "/api/v3/openOrderList", params, nil, USER_DATA)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

// Account Trades

type AccountTradesParam struct {
	OrderID    int64
	StartTime  int64
	EndTime    int64
	FromID     int64
	Limit      int
	RecvWindow int
}

func (c *Client) Trades(symbol string, opts ...AccountTradesParam) ([]*Trade, Response, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "orderId", opt.OrderID)
		validation.SetIfNotZero(params, "startTime", opt.StartTime)
		validation.SetIfNotZero(params, "endTime", opt.EndTime)
		validation.SetIfNotZero(params, "fromId", opt.FromID)
		validation.SetIfNotZero(params, "limit", opt.Limit)
		validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	}

	result, resp, err := doRequest[[]*Trade](c, Methods.GET, "/api/v3/myTrades", params, nil, USER_DATA)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

// Unfilled Order Count

type UnfilledOrderCount struct {
	RateLimitType RateLimitType `json:"rateLimitType"`

	Interval    RateLimitInterval `json:"interval"`
	IntervalNum int64             `json:"intervalNum"`

	Limit int `json:"limit"`
	Count int `json:"count"`
}

type UnfilledOrderCountParams struct {
	RecvWindow int
}

func (c *Client) UnfilledOrderCount(opts ...UnfilledOrderCountParams) ([]*UnfilledOrderCount, Response, Error) {
	params := make(map[string]interface{})

	if len(opts) > 0 {
		opt := opts[0]
		validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	}

	result, resp, err := doRequest[[]*UnfilledOrderCount](c, Methods.GET, "/api/v3/rateLimit/order", params, nil, USER_DATA)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

// Prevented Matches

type PreventedMatch struct {
	Symbol string `json:"symbol"`

	PreventedMatchID int64 `json:"preventedMatchId"`

	TakerOrderID int64 `json:"takerOrderId"`

	MakerSymbol  string `json:"makerSymbol"`
	MakerOrderID int64  `json:"makerOrderId"`

	TradeGroupID int64 `json:"tradeGroupId"`

	SelfTradePreventionMode STPMode `json:"selfTradePreventionMode"`

	Price                  float64 `json:"price,string"`
	MakerPreventedQuantity float64 `json:"makerPreventedQuantity,string"`

	TransactTime int64 `json:"transactTime"`
}

type PreventedMatchesParams struct {
	PreventedMatchID     int64
	OrderID              int64
	FromPreventedMatchID int64
	Limit                int
	RecvWindow           int
}

func (c *Client) PreventedMatches(symbol string, opts ...PreventedMatchesParams) ([]*PreventedMatch, Response, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "preventedMatchId", opt.PreventedMatchID)
		validation.SetIfNotZero(params, "orderId", opt.OrderID)
		validation.SetIfNotZero(params, "fromPreventedMatchId", opt.FromPreventedMatchID)
		validation.SetIfNotZero(params, "limit", opt.Limit)
		validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	}

	result, resp, err := doRequest[[]*PreventedMatch](c, Methods.GET, "/api/v3/myPreventedMatches", params, nil, USER_DATA)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

// Allocations

type Allocation struct {
	Symbol string `json:"symbol"`

	AllocationID   int64          `json:"allocationId"`
	AllocationType AllocationType `json:"allocationType"`

	OrderID     int64 `json:"orderId"`
	OrderListID int64 `json:"orderListId"`

	Price    float64 `json:"price,string"`
	Qty      float64 `json:"qty,string"`
	QuoteQty float64 `json:"quoteQty,string"`

	Commission      float64 `json:"commission,string"`
	CommissionAsset string  `json:"commissionAsset"`

	Time int64 `json:"time"`

	IsBuyer     bool `json:"isBuyer"`
	IsMaker     bool `json:"isMaker"`
	IsAllocator bool `json:"isAllocator"`
}

type AllocationsParams struct {
	StartTime        int64
	EndTime          int64
	FromAllocationID int64
	OrderID          int64
	Limit            int
	RecvWindow       int
}

func (c *Client) Allocations(symbol string, opts ...AllocationsParams) ([]*Allocation, Response, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "startTime", opt.StartTime)
		validation.SetIfNotZero(params, "endTime", opt.EndTime)
		validation.SetIfNotZero(params, "fromAllocationId", opt.FromAllocationID)
		validation.SetIfNotZero(params, "orderId", opt.OrderID)
		validation.SetIfNotZero(params, "limit", opt.Limit)
		validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	}

	result, resp, err := doRequest[[]*Allocation](c, Methods.GET, "/api/v3/myAllocations", params, nil, USER_DATA)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

// Commissions

type AccountCommission struct {
	Symbol string `json:"symbol"`

	StandardCommission CommissionRates `json:"standardCommission"`
	SpecialCommission  CommissionRates `json:"specialCommission"`
	TaxCommission      CommissionRates `json:"taxCommission"`

	Discount CommissionDiscount `json:"discount"`
}

// EffectiveMakerRate calculates the effective maker rate by applying the discounts if applied
func (a AccountCommission) EffectiveMakerRate() float64 {
	if a.Discount.EnabledForAccount && a.Discount.EnabledForSymbol {
		return a.StandardCommission.Maker * (1 - a.Discount.Discount)
	}
	return a.StandardCommission.Maker
}

// EffectiveTakerRate calculates the effective maker rate by applying the discounts if applied
func (a AccountCommission) EffectiveTakerRate() float64 {
	if a.Discount.EnabledForAccount && a.Discount.EnabledForSymbol {
		return a.StandardCommission.Taker * (1 - a.Discount.Discount)
	}
	return a.StandardCommission.Taker
}

type CommissionDiscount struct {
	EnabledForAccount bool    `json:"enabledForAccount"`
	EnabledForSymbol  bool    `json:"enabledForSymbol"`
	DiscountAsset     string  `json:"discountAsset"`
	Discount          float64 `json:"discount,string"`
}

type AccountCommissionParams struct {
	RecvWindow int
}

func (c *Client) AccountCommission(symbol string, opts ...AccountCommissionParams) (*AccountCommission, Response, Error) {
	params := map[string]interface{}{
		"symbol": symbol,
	}

	if len(opts) > 0 {
		opt := opts[0]
		validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	}

	return doRequest[AccountCommission](c, Methods.GET, "/api/v3/account/commission", params, nil, USER_DATA)
}

// Order Amendments

type OrderAmendment struct {
	Symbol string `json:"symbol"`

	OrderID     int64 `json:"orderId"`
	ExecutionID int64 `json:"executionId"`

	OrigClientOrderID string `json:"origClientOrderId"`
	NewClientOrderID  string `json:"newClientOrderId"`

	OrigQty float64 `json:"origQty,string"`
	NewQty  float64 `json:"newQty,string"`

	Time int64 `json:"time"`
}

type OrderAmendmentsParams struct {
	FromExecutionID int64
	Limit           int
	RecvWindow      int
}

func (c *Client) OrderAmendments(symbol string, orderID int64, opts ...OrderAmendmentsParams) ([]*OrderAmendment, Response, Error) {
	params := map[string]interface{}{
		"symbol":  symbol,
		"orderId": orderID,
	}

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "fromExecutionId", opt.FromExecutionID)
		validation.SetIfNotZero(params, "limit", opt.Limit)
		validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	}

	result, resp, err := doRequest[[]*OrderAmendment](c, Methods.GET, "/api/v3/order/amendments", params, nil, USER_DATA)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

// Relevant Filters

type AccountFilters struct {
	ExchangeFilters ExchangeFilters `json:"exchangeFilters"`
	SymbolFilters   SymbolFilters   `json:"symbolFilters"`
	AssetFilters    AssetFilters    `json:"assetFilters"`
}

type AssetFilters struct {
	MAX_ASSET *AssetFilterMaxAsset
}

type AssetFilterMaxAsset struct {
	FilterType string  `json:"filterType"`
	Asset      string  `json:"asset"`
	Limit      float64 `json:"limit,string"`
}

type AccountFiltersParams struct {
	RecvWindow int
}

// AccountFilters retrieves the list of filters relevant to an account on a given symbol.
//
// This is the only endpoint that shows if an account has MAX_ASSET filters applied to it.
func (c *Client) AccountFilters(symbol string, opts ...AccountFiltersParams) (*AccountFilters, Response, Error) {
	params := map[string]interface{}{
		"symbol": symbol,
	}

	if len(opts) > 0 {
		opt := opts[0]
		validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)
	}

	return doRequest[AccountFilters](c, Methods.GET, "/api/v3/myFilters", params, nil, USER_DATA)
}
