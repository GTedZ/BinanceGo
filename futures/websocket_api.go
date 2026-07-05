package futures

import (
	"time"

	"github.com/GTedZ/binancego/internal/berror"
	"github.com/GTedZ/binancego/internal/json"
	"github.com/GTedZ/binancego/internal/logging"
	"github.com/GTedZ/binancego/internal/requests"
	"github.com/GTedZ/binancego/internal/validation"
	"github.com/GTedZ/binancego/internal/websocket/binance"
)

type WebsocketAPI struct {
	base    baseWebsocket
	baseUrl string

	client *Client

	logger logging.Logger
}

func (c *Client) WebsocketAPI() (*WebsocketAPI, Error) {
	return newWebsocketAPI(c.wssApiBaseUrl, c)
}

////
// Private Methods
////

func (wsapi *WebsocketAPI) setOnMessage(messageHandler func(date []byte)) {
	wsapi.base.SetOnMessage(messageHandler)
}

////
// Public Methods
////

func (wsapi *WebsocketAPI) Close() {
	wsapi.base.Close()
}

////////////////////////////////////////////////////////////////////////////////
// Market data requests
////////////////////////////////////////////////////////////////////////////////

////
// Order Book
////

func (wsapi *WebsocketAPI) OrderBook(symbol string, opts ...OrderBookParams) (*OrderBook, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol

	if len(opts) > 0 {
		opt := opts[0]

		validation.SetIfNotZero(params, "limit", opt.Limit)
	}

	return doWsApiRequest[OrderBook](wsapi, "depth", params, wsapi.client.apikey, NONE)
}

////
// Symbol Price Ticker
////

func (wsapi *WebsocketAPI) PriceTickers() ([]*PriceTicker, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	result, resp, err := doWsApiRequest[[]*PriceTicker](wsapi, "ticker.price", params, wsapi.client.apikey, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

func (wsapi *WebsocketAPI) PriceTicker(symbol string) (*PriceTicker, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol

	return doWsApiRequest[PriceTicker](wsapi, "ticker.price", params, wsapi.client.apikey, NONE)
}

////
// Symbol Order Book Ticker
////

func (wsapi *WebsocketAPI) BookTickers() ([]*OrderBookTicker, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	result, resp, err := doWsApiRequest[[]*OrderBookTicker](wsapi, "ticker.book", params, wsapi.client.apikey, NONE)
	if result == nil {
		return nil, resp, err
	}
	return *result, resp, err
}

func (wsapi *WebsocketAPI) BookTicker(symbol string) (*OrderBookTicker, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol

	return doWsApiRequest[OrderBookTicker](wsapi, "ticker.book", params, wsapi.client.apikey, NONE)
}

////////////////////////////////////////////////////////////////////////////////
// Trading requests
////////////////////////////////////////////////////////////////////////////////

////
// Place Order
////

// PlaceOrder sends a new order over the WebSocket API (method "order.place").
//
// It accepts any orderRequest implementation (e.g. NewLimitBuy, NewMarketSell,
// etc.) — the same request builders used by the REST Client.Order — and submits
// it over the WebSocket connection.
func (wsapi *WebsocketAPI) PlaceOrder(req orderRequest) (*Order, *WsApiResponse, Error) {
	params := req.build()

	return doWsApiRequest[Order](wsapi, "order.place", params, wsapi.client.apikey, TRADE)
}

////
// Modify Order
////

// ModifyOrder amends an existing LIMIT order over the WebSocket API
// (method "order.modify"). It accepts the same modifyOrderRequest used by the
// REST Client.ModifyOrder.
func (wsapi *WebsocketAPI) ModifyOrder(req modifyOrderRequest) (*Order, *WsApiResponse, Error) {
	params := req.build()

	return doWsApiRequest[Order](wsapi, "order.modify", params, wsapi.client.apikey, TRADE)
}

////
// Cancel Order
////

// CancelOrder cancels an active order over the WebSocket API (method "order.cancel").
//
// NOTE: symbol is required by Binance for order cancellation. It is passed
// explicitly here because the REST CancelOrderRequest struct does not carry it.
// TODO: the REST Client.CancelOrder does not send a symbol either; if that is
// intentional (e.g. resolved elsewhere) align this with it, otherwise the REST
// path likely needs the same symbol fix.
func (wsapi *WebsocketAPI) CancelOrder(symbol string, req CancelOrderRequest) (*Order, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol

	validation.SetIfNotZero(params, "orderId", req.OrderId)
	validation.SetIfNotZero(params, "origClientOrderId", req.ClientOrderId)
	validation.SetIfNotZero(params, "recvWindow", req.RecvWindow)

	return doWsApiRequest[Order](wsapi, "order.cancel", params, wsapi.client.apikey, TRADE)
}

////
// Query Order
////

// QueryOrder checks an order's status over the WebSocket API (method "order.status").
func (wsapi *WebsocketAPI) QueryOrder(symbol string, opt OrderParams) (*Order, *WsApiResponse, Error) {
	params := make(map[string]interface{})

	params["symbol"] = symbol

	validation.SetIfNotZero(params, "orderId", opt.OrderId)
	validation.SetIfNotZero(params, "origClientOrderId", opt.OrigClientOrderId)
	validation.SetIfNotZero(params, "recvWindow", opt.RecvWindow)

	return doWsApiRequest[Order](wsapi, "order.status", params, wsapi.client.apikey, USER_DATA)
}

////////////////////////////////////////////////////////////////////////////////
// Account requests
////////////////////////////////////////////////////////////////////////////////

// TODO: The USD-M futures WebSocket API also exposes the account endpoints
// "account.balance" / "v2/account.balance" and "account.status" /
// "v2/account.status". These are intentionally left unimplemented here because
// the corresponding REST response types (futures account balance / account
// information) do not yet exist in this package. Add them alongside the REST
// account endpoints so the WebSocket API can reuse the same response structs,
// mirroring how the market-data and trading methods above reuse existing types.

////
// Response
////

// Status codes in the status field are the same as in HTTP.
// Here are some common status codes that you might encounter:
//
// - 200 indicates a successful response.
//
// - 4XX status codes indicate invalid requests; the issue is on your side.
//
// - 400 – your request failed, see error for the reason.
//
// - 403 – you have been blocked by the Web Application Firewall. This can indicate a rate limit violation or a security block.
//
// - 409 – your request partially failed but also partially succeeded, see error for details.
//
// - 418 – you have been auto-banned for repeated violation of rate limits.
//
// - 429 – you have exceeded API request rate limit, please slow down.
//
// - 5XX status codes indicate internal errors; the issue is on Binance's side.
//
// Important: If a response contains 5xx status code, it does not necessarily mean that your request has failed. Execution status is unknown and the request might have actually succeeded.
// Please use query methods to confirm the status.
// You might also want to establish a new WebSocket connection for that.
type WsApiResponse struct {
	Id         string          `json:"id"`
	Status     int             `json:"status"`
	Error      *WsApiError     `json:"error"`
	Result     json.RawMessage `json:"result"`
	RateLimits []RateLimit     `json:"rateLimits"`
}

type WsApiError struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data json.RawMessage `json:"data"`
}

////
// Request
////

func doWsApiRequest[T any](s *WebsocketAPI, method string, params map[string]interface{}, apikey KeyPair, securityType SecurityType) (*T, *WsApiResponse, Error) {
	payload := make(map[string]interface{})

	var useApikey, isSigned bool

	switch securityType {
	case NONE:
		useApikey = false
		isSigned = false
	case TRADE:
		useApikey = true
		isSigned = true
	case USER_STREAM:
		useApikey = true
		isSigned = false
	case USER_DATA:
		useApikey = true
		isSigned = true
	}

	if useApikey {
		params["apiKey"] = apikey.ApiKey()
	}

	if isSigned {
		params["timestamp"] = time.Now().UnixMilli()
		paramString := requests.CreateQueryString(params, true)
		signature, err := apikey.Sign(paramString)
		if err != nil {
			return nil, nil, berror.NewSignatureError(err)
		}

		params["signature"] = signature
	}

	payload["method"] = method
	payload["params"] = params

	data, err := s.base.SendRequest(payload, websocketRequestTimeout)
	if err != nil {
		return nil, nil, err
	}

	var response *WsApiResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, response, berror.NewParseError(err)
	}

	if response.Status == 200 {
		var result T
		if err := json.Unmarshal(response.Result, &result); err != nil {
			return nil, response, berror.NewParseError(err)
		}

		return &result, response, nil
	}

	return nil, response, berror.NewAPIError(response.Error.Code, response.Error.Msg)
}

////
// Constructor
////

func newWebsocketAPI(baseUrl string, client *Client) (*WebsocketAPI, Error) {
	socket := &WebsocketAPI{
		baseUrl: baseUrl,

		client: client,
		logger: client.logger,
	}

	messageHandler := func(data []byte) {
		client.logger.WARNf("Received message on WebsocketAPI socket when one wasn't expected => %s", string(data))
	}

	base, err := binance.New(baseUrl, messageHandler, nil, nil, client.logger)
	if err != nil {
		return nil, berror.NewNetworkError(err)
	}

	socket.base = base

	return socket, nil
}
