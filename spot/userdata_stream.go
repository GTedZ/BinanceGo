package spot

import (
	"github.com/GTedZ/binancego/internal/json"
	"github.com/GTedZ/binancego/internal/logging"
)

type UserDataWebsocket struct {
	base *WebsocketAPI

	client *Client

	logger logging.Logger

	onMessage func(UserDataEvent)
}

func (c *Client) UserDataStream(cb func(ud UserDataEvent)) (*UserDataWebsocket, Error) {
	return newUserDataWebsocket(c, cb)
}

////
// Private Methods
////

type UserDataSubscribeResponse struct {
	SubscriptionId int `json:"subscriptionId"`
}

func (uws *UserDataWebsocket) subscribe() (*UserDataSubscribeResponse, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	return doWsApiRequest[UserDataSubscribeResponse](uws.base, "userDataStream.subscribe.signature", params, uws.client.apikey, USER_DATA)
}

func (uws *UserDataWebsocket) broadcastOnMessage(ud UserDataEvent) {
	if uws.onMessage != nil {
		uws.onMessage(ud)
	}
}

////
// Public Methods
////

func (uws *UserDataWebsocket) Close() {
	uws.base.Close()
}

////
// Message Handler
////

type UserDataRawEvent struct {
	SubscriptionID int             `json:"subscriptionId"`
	Event          json.RawMessage `json:"event"`
}

type UserDataEvent struct {
	SubscriptionID int               `json:"subscriptionId"`
	EventType      UserDataEventType `json:"-"`
	EventTime      int64             `json:"-"`

	// Exactly one of these will be non-nil
	AccountPosition  *OutboundAccountPositionEvent `json:"-"`
	BalanceUpdate    *BalanceUpdateEvent           `json:"-"`
	ExecutionReport  *ExecutionReportEvent         `json:"-"`
	ListStatus       *ListStatusEvent              `json:"-"`
	StreamTerminated *EventStreamTerminatedEvent   `json:"-"`
	ExternalLock     *ExternalLockUpdateEvent      `json:"-"`
}

// --- Account Update (outboundAccountPosition) ---

type OutboundAccountPositionEvent struct {
	EventType      UserDataEventType `json:"e"`
	EventTime      int64             `json:"E"`
	LastUpdateTime int64             `json:"u"`
	Balances       []UserDataBalance `json:"B"`
}

type UserDataBalance struct {
	Asset  string `json:"a"`
	Free   string `json:"f"`
	Locked string `json:"l"`
}

// --- Balance Update ---

type BalanceUpdateEvent struct {
	EventType UserDataEventType `json:"e"`
	EventTime int64             `json:"E"`
	Asset     string            `json:"a"`
	Delta     string            `json:"d"`
	ClearTime int64             `json:"T"`
}

// --- Execution Report (Order Update) ---

type ExecutionReportEvent struct {
	EventType             UserDataEventType `json:"e"`
	EventTime             int64             `json:"E"`
	Symbol                string            `json:"s"`
	ClientOrderID         string            `json:"c"`
	Side                  string            `json:"S"`
	OrderType             string            `json:"o"`
	TimeInForce           string            `json:"f"`
	Quantity              string            `json:"q"`
	Price                 string            `json:"p"`
	StopPrice             string            `json:"P"`
	IcebergQuantity       string            `json:"F"`
	OrderListID           int64             `json:"g"`
	OriginalClientOrderID string            `json:"C"`
	ExecutionType         string            `json:"x"`
	OrderStatus           string            `json:"X"`
	RejectReason          string            `json:"r"`
	OrderID               int64             `json:"i"`
	LastExecutedQuantity  string            `json:"l"`
	CumulativeFilledQty   string            `json:"z"`
	LastExecutedPrice     string            `json:"L"`
	CommissionAmount      string            `json:"n"`
	CommissionAsset       *string           `json:"N"`
	TransactionTime       int64             `json:"T"`
	TradeID               int64             `json:"t"`
	PreventedMatchID      *int64            `json:"v,omitempty"`
	ExecutionID           int64             `json:"I"`
	IsOnBook              bool              `json:"w"`
	IsMaker               bool              `json:"m"`
	Ignore                bool              `json:"M"`
	OrderCreationTime     int64             `json:"O"`
	CumulativeQuoteQty    string            `json:"Z"`
	LastQuoteQty          string            `json:"Y"`
	QuoteOrderQty         string            `json:"Q"`
	WorkingTime           *int64            `json:"W,omitempty"`
	SelfTradePrevention   string            `json:"V"`

	// --- Conditional fields ---
	TrailingDelta         *int64  `json:"d,omitempty"`
	TrailingTime          *int64  `json:"D,omitempty"`
	StrategyID            *int64  `json:"j,omitempty"`
	StrategyType          *int64  `json:"J,omitempty"`
	PreventedQuantity     *string `json:"A,omitempty"`
	LastPreventedQuantity *string `json:"B,omitempty"`
	TradeGroupID          *int64  `json:"u,omitempty"`
	CounterOrderID        *int64  `json:"U,omitempty"`
	CounterSymbol         *string `json:"Cs,omitempty"`
	PreventedExecQuantity *string `json:"pl,omitempty"`
	PreventedExecPrice    *string `json:"pL,omitempty"`
	PreventedExecQuoteQty *string `json:"pY,omitempty"`
	MatchType             *string `json:"b,omitempty"`
	AllocationID          *int64  `json:"a,omitempty"`
	WorkingFloor          *string `json:"k,omitempty"`
	UsedSor               *bool   `json:"uS,omitempty"`
	PeggedPriceType       *string `json:"gP,omitempty"`
	PeggedOffsetType      *string `json:"gOT,omitempty"`
	PeggedOffsetValue     *int64  `json:"gOV,omitempty"`
	PeggedPrice           *string `json:"gp,omitempty"`
	ExpiryReason          *string `json:"eR,omitempty"`
}

// --- List Status ---

type ListStatusEvent struct {
	EventType         UserDataEventType `json:"e"`
	EventTime         int64             `json:"E"`
	Symbol            string            `json:"s"`
	OrderListID       int64             `json:"g"`
	ContingencyType   string            `json:"c"`
	ListStatusType    string            `json:"l"`
	ListOrderStatus   string            `json:"L"`
	ListRejectReason  string            `json:"r"`
	ListClientOrderID string            `json:"C"`
	TransactionTime   int64             `json:"T"`
	Orders            []ListStatusOrder `json:"O"`
}

type ListStatusOrder struct {
	Symbol        string `json:"s"`
	OrderID       int64  `json:"i"`
	ClientOrderID string `json:"c"`
}

// --- Event Stream Terminated ---

type EventStreamTerminatedEvent struct {
	EventType UserDataEventType `json:"e"`
	EventTime int64             `json:"E"`
}

// --- External Lock Update ---

type ExternalLockUpdateEvent struct {
	EventType       UserDataEventType `json:"e"`
	EventTime       int64             `json:"E"`
	Asset           string            `json:"a"`
	Delta           string            `json:"d"`
	TransactionTime int64             `json:"T"`
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
	socket := &UserDataWebsocket{
		client: client,
		logger: client.logger,

		onMessage: cb,
	}

	base, err := client.WebsocketAPI()
	if err != nil {
		return nil, err
	}

	base.setOnMessage(socket.messageHandler)

	socket.base = base

	if _, _, err := socket.subscribe(); err != nil {
		return nil, err
	}

	return socket, nil
}
