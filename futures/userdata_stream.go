package futures

import (
	"time"

	"github.com/GTedZ/binancego/internal/berror"
	"github.com/GTedZ/binancego/internal/json"
	"github.com/GTedZ/binancego/internal/logging"
	"github.com/GTedZ/binancego/internal/websocket/binance"
)

type UserDataWebsocket struct {
	base baseWebsocket

	client *Client

	logger logging.Logger

	listenKey string

	// closed on Close() to stop the keepalive goroutine.
	stopKeepAlive chan struct{}

	onMessage func(UserDataEvent)
}

func (c *Client) UserDataStream(cb func(ud UserDataEvent)) (*UserDataWebsocket, Error) {
	return newUserDataWebsocket(c, cb)
}

////
// ListenKey management
////

type listenKeyResponse struct {
	ListenKey string `json:"listenKey"`
}

// StartUserDataStream creates a new listenKey (POST /fapi/v1/listenKey).
//
// The stream is valid for 60 minutes; use KeepAliveUserDataStream to extend it.
func (c *Client) StartUserDataStream() (string, Response, Error) {
	result, resp, err := doRequest[listenKeyResponse](c, Methods.POST, "/fapi/v1/listenKey", nil, nil, USER_STREAM)
	if err != nil {
		return "", resp, err
	}
	return result.ListenKey, resp, nil
}

// KeepAliveUserDataStream extends the validity of the current listenKey by
// another 60 minutes (PUT /fapi/v1/listenKey).
func (c *Client) KeepAliveUserDataStream() (Response, Error) {
	return c.MakeRequest(Methods.PUT, "/fapi/v1/listenKey", nil, nil, USER_STREAM)
}

// CloseUserDataStream closes the current listenKey (DELETE /fapi/v1/listenKey).
func (c *Client) CloseUserDataStream() (Response, Error) {
	return c.MakeRequest(Methods.DELETE, "/fapi/v1/listenKey", nil, nil, USER_STREAM)
}

////
// Private Methods
////

func (uws *UserDataWebsocket) broadcastOnMessage(ud UserDataEvent) {
	if uws.onMessage != nil {
		uws.onMessage(ud)
	}
}

func (uws *UserDataWebsocket) keepAliveLoop() {
	ticker := time.NewTicker(30 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-uws.stopKeepAlive:
			return
		case <-ticker.C:
			if _, err := uws.client.KeepAliveUserDataStream(); err != nil {
				uws.logger.ERRORf("Failed to keepalive futures user data stream: %s", err.Error())
			}
		}
	}
}

////
// Public Methods
////

func (uws *UserDataWebsocket) Close() {
	close(uws.stopKeepAlive)
	uws.base.Close()

	if _, err := uws.client.CloseUserDataStream(); err != nil {
		uws.logger.ERRORf("Failed to close futures user data stream: %s", err.Error())
	}
}

////
// Events
////

type UserDataEvent struct {
	EventType UserDataEventType `json:"-"`
	EventTime int64             `json:"-"`

	// Exactly one of these will be non-nil
	ListenKeyExpired              *ListenKeyExpiredEvent              `json:"-"`
	MarginCall                    *MarginCallEvent                    `json:"-"`
	AccountUpdate                 *AccountUpdateEvent                 `json:"-"`
	OrderTradeUpdate              *OrderTradeUpdateEvent              `json:"-"`
	TradeLite                     *TradeLiteEvent                     `json:"-"`
	AccountConfigUpdate           *AccountConfigUpdateEvent           `json:"-"`
	StrategyUpdate                *StrategyUpdateEvent                `json:"-"`
	GridUpdate                    *GridUpdateEvent                    `json:"-"`
	ConditionalOrderTriggerReject *ConditionalOrderTriggerRejectEvent `json:"-"`
}

// --- listenKeyExpired ---

type ListenKeyExpiredEvent struct {
	EventType UserDataEventType `json:"e"`
	EventTime int64             `json:"E"`
	// TODO: some payload versions also include the expired "listenKey" value.
	ListenKey string `json:"listenKey"`
}

// --- MARGIN_CALL ---

type MarginCallEvent struct {
	EventType UserDataEventType `json:"e"`
	EventTime int64             `json:"E"`
	// Cross wallet balance. Only pushed with crossed position margin call.
	CrossWalletBalance float64              `json:"cw,string"`
	Positions          []MarginCallPosition `json:"p"`
}

type MarginCallPosition struct {
	Symbol            string       `json:"s"`
	PositionSide      PositionSide `json:"ps"`
	PositionAmount    float64      `json:"pa,string"`
	MarginType        MarginType   `json:"mt"`
	IsolatedWallet    float64      `json:"iw,string"`
	MarkPrice         float64      `json:"mp,string"`
	UnrealizedPnl     float64      `json:"up,string"`
	MaintenanceMargin float64      `json:"mm,string"`
}

// --- ACCOUNT_UPDATE ---

type AccountUpdateEvent struct {
	EventType       UserDataEventType `json:"e"`
	EventTime       int64             `json:"E"`
	TransactionTime int64             `json:"T"`
	UpdateData      AccountUpdateData `json:"a"`
}

type AccountUpdateData struct {
	// Event reason type (e.g. ORDER, FUNDING_FEE, DEPOSIT, WITHDRAW, ...)
	Reason    string                  `json:"m"`
	Balances  []AccountUpdateBalance  `json:"B"`
	Positions []AccountUpdatePosition `json:"P"`
}

type AccountUpdateBalance struct {
	Asset              string  `json:"a"`
	WalletBalance      float64 `json:"wb,string"`
	CrossWalletBalance float64 `json:"cw,string"`
	// Balance change except PnL and commission
	BalanceChange float64 `json:"bc,string"`
}

type AccountUpdatePosition struct {
	Symbol         string  `json:"s"`
	PositionAmount float64 `json:"pa,string"`
	EntryPrice     float64 `json:"ep,string"`
	BreakevenPrice float64 `json:"bep,string"`
	// (Pre-fee) accumulated realized
	AccumulatedRealized float64      `json:"cr,string"`
	UnrealizedPnl       float64      `json:"up,string"`
	MarginType          string       `json:"mt"`
	IsolatedWallet      float64      `json:"iw,string"`
	PositionSide        PositionSide `json:"ps"`
}

// --- ORDER_TRADE_UPDATE ---

type OrderTradeUpdateEvent struct {
	EventType       UserDataEventType     `json:"e"`
	EventTime       int64                 `json:"E"`
	TransactionTime int64                 `json:"T"`
	Order           OrderTradeUpdateOrder `json:"o"`
}

type OrderTradeUpdateOrder struct {
	Symbol        string      `json:"s"`
	ClientOrderID string      `json:"c"`
	Side          OrderSide   `json:"S"`
	OrderType     OrderType   `json:"o"`
	TimeInForce   TimeInForce `json:"f"`
	OrigQuantity  float64     `json:"q,string"`
	OrigPrice     float64     `json:"p,string"`
	AvgPrice      float64     `json:"ap,string"`
	StopPrice     float64     `json:"sp,string"`
	// Execution type (NEW, CANCELED, CALCULATED, EXPIRED, TRADE, AMENDMENT)
	ExecutionType string      `json:"x"`
	OrderStatus   OrderStatus `json:"X"`
	OrderID       int64       `json:"i"`
	// Order last filled quantity
	LastFilledQty float64 `json:"l,string"`
	// Order filled accumulated quantity
	FilledAccumulatedQty float64 `json:"z,string"`
	// Last filled price
	LastFilledPrice float64 `json:"L,string"`
	CommissionAsset string  `json:"N"`
	Commission      float64 `json:"n,string"`
	// Order trade time
	TradeTime int64 `json:"T"`
	TradeID   int64 `json:"t"`
	// Bids notional
	BidsNotional float64 `json:"b,string"`
	// Asks notional
	AsksNotional float64 `json:"a,string"`
	IsMaker      bool    `json:"m"`
	IsReduceOnly bool    `json:"R"`
	// Stop price working type
	WorkingType   WorkingType  `json:"wt"`
	OrigOrderType OrderType    `json:"ot"`
	PositionSide  PositionSide `json:"ps"`
	// If Close-All (only pushed with conditional order)
	IsCloseAll bool `json:"cp"`
	// Activation price (only pushed with TRAILING_STOP_MARKET order)
	ActivationPrice float64 `json:"AP,string"`
	// Callback rate (only pushed with TRAILING_STOP_MARKET order)
	CallbackRate float64 `json:"cr,string"`
	PriceProtect bool    `json:"pP"`
	// Realized profit of the trade
	RealizedProfit          float64    `json:"rp,string"`
	SelfTradePreventionMode STPMode    `json:"V"`
	PriceMatch              PriceMatch `json:"pm"`
	// GTD order auto cancel time
	GoodTillDate int64 `json:"gtd"`
}

// --- TRADE_LITE ---

type TradeLiteEvent struct {
	EventType       UserDataEventType `json:"e"`
	EventTime       int64             `json:"E"`
	TransactionTime int64             `json:"T"`
	Symbol          string            `json:"s"`
	OrigQuantity    float64           `json:"q,string"`
	OrigPrice       float64           `json:"p,string"`
	IsMaker         bool              `json:"m"`
	ClientOrderID   string            `json:"c"`
	Side            OrderSide         `json:"S"`
	LastFilledPrice float64           `json:"L,string"`
	LastFilledQty   float64           `json:"l,string"`
	TradeID         int64             `json:"t"`
	OrderID         int64             `json:"i"`
}

// --- ACCOUNT_CONFIG_UPDATE ---

type AccountConfigUpdateEvent struct {
	EventType       UserDataEventType `json:"e"`
	EventTime       int64             `json:"E"`
	TransactionTime int64             `json:"T"`
	// Present when a symbol's leverage changed.
	Leverage *AccountConfigLeverage `json:"ac"`
	// Present when the account's Multi-Assets mode changed.
	MultiAssets *AccountConfigMultiAssets `json:"ai"`
}

type AccountConfigLeverage struct {
	Symbol   string `json:"s"`
	Leverage int    `json:"l"`
}

type AccountConfigMultiAssets struct {
	MultiAssetsMode bool `json:"j"`
}

// --- STRATEGY_UPDATE ---

type StrategyUpdateEvent struct {
	EventType       UserDataEventType `json:"e"`
	EventTime       int64             `json:"E"`
	TransactionTime int64             `json:"T"`
	Strategy        StrategyUpdate    `json:"su"`
}

type StrategyUpdate struct {
	StrategyID     int64  `json:"si"`
	StrategyType   string `json:"st"`
	StrategyStatus string `json:"ss"`
	Symbol         string `json:"s"`
	UpdateTime     int64  `json:"ut"`
	OpCode         int    `json:"c"`
}

// --- GRID_UPDATE ---

type GridUpdateEvent struct {
	EventType       UserDataEventType `json:"e"`
	EventTime       int64             `json:"E"`
	TransactionTime int64             `json:"T"`
	Grid            GridUpdate        `json:"gu"`
}

type GridUpdate struct {
	StrategyID     int64   `json:"si"`
	StrategyType   string  `json:"st"`
	StrategyStatus string  `json:"ss"`
	Symbol         string  `json:"s"`
	RealizedPnl    float64 `json:"r,string"`
	UnmatchedAvg   float64 `json:"up,string"`
	UpdateTime     int64   `json:"ut"`
}

// --- CONDITIONAL_ORDER_TRIGGER_REJECT ---

type ConditionalOrderTriggerRejectEvent struct {
	EventType       UserDataEventType                  `json:"e"`
	EventTime       int64                              `json:"E"`
	TransactionTime int64                              `json:"T"`
	Order           ConditionalOrderTriggerRejectOrder `json:"or"`
}

type ConditionalOrderTriggerRejectOrder struct {
	Symbol       string `json:"s"`
	OrderID      int64  `json:"i"`
	RejectReason string `json:"r"`
}

////
// Message Handler
////

func (uws *UserDataWebsocket) messageHandler(data []byte) {
	var result UserDataEvent

	if err := json.Unmarshal(data, &result); err != nil {
		uws.logger.ERRORf("failed to parse User Data Event Stream Message: %v | payload=%s", err, string(data))
		return
	}

	uws.broadcastOnMessage(result)
}

////
// Constructor
////

func newUserDataWebsocket(client *Client, cb func(UserDataEvent)) (*UserDataWebsocket, Error) {
	listenKey, _, err := client.StartUserDataStream()
	if err != nil {
		return nil, err
	}

	socket := &UserDataWebsocket{
		client: client,
		logger: client.logger,

		listenKey:     listenKey,
		stopKeepAlive: make(chan struct{}),

		onMessage: cb,
	}

	url := client.wssBaseUrl + "/ws/" + listenKey

	base, wsErr := binance.New(url, socket.messageHandler, nil, nil, client.logger)
	if wsErr != nil {
		return nil, berror.NewNetworkError(wsErr)
	}

	socket.base = base

	go socket.keepAliveLoop()

	return socket, nil
}
