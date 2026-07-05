package futures

import (
	"strings"
	"sync"
	"time"

	"github.com/GTedZ/binancego/internal/berror"
	"github.com/GTedZ/binancego/internal/json"
	"github.com/GTedZ/binancego/internal/logging"
	"github.com/GTedZ/binancego/internal/websocket/binance"
)

const (
	websocketRequestTimeout = 4 * time.Second
)

type websocket[T any] struct {
	base      baseWebsocket
	baseUrl   string
	baseRoute WebsocketRoute

	logger logging.Logger

	streams struct {
		mu   sync.Mutex
		data map[string]struct{}
	}

	onMessage func(data T)
}

func (s *websocket[T]) buildUrl() string {
	streams := s.readSubscriptions()

	// build the full URL with the base URL, route, and streams
	fullUrl := s.baseUrl + string(s.baseRoute) + "/stream?streams=" + strings.Join(streams, "/")

	return fullUrl
}

////
// Broadcasts
////

func (s *websocket[T]) broadcastOnMessage(data T) {
	if s.onMessage != nil {
		s.onMessage(data)
	}
}

////
// Requests
////

func (s *websocket[T]) subscribe(streams []string) Error {
	payload := make(map[string]interface{})

	payload["method"] = "SUBSCRIBE"
	payload["params"] = streams

	_, err := s.base.SendRequest(payload, websocketRequestTimeout)
	if err != nil {
		return err
	}

	s.streams.mu.Lock()
	defer s.streams.mu.Unlock()

	for _, stream := range streams {
		s.streams.data[stream] = struct{}{}
	}

	return nil
}

func (s *websocket[T]) unsubscribe(streams []string) Error {
	payload := make(map[string]interface{})

	payload["method"] = "UNSUBSCRIBE"
	payload["params"] = streams

	_, err := s.base.SendRequest(payload, websocketRequestTimeout)
	if err != nil {
		return err
	}
	s.streams.mu.Lock()
	defer s.streams.mu.Unlock()

	for _, stream := range streams {
		delete(s.streams.data, stream)
	}

	return nil
}

func (s *websocket[T]) listSubscriptions() ([]string, Error) {
	payload := make(map[string]interface{})

	payload["method"] = "LIST_SUBSCRIPTIONS"

	streams, err := doWsRequest[[]string](s, payload)
	if err != nil {
		return nil, err
	}

	s.streams.mu.Lock()
	defer s.streams.mu.Unlock()

	// Reset the internal streams set so that it's synced with binance
	clear(s.streams.data)

	for _, stream := range *streams {
		s.streams.data[stream] = struct{}{}
	}

	return *streams, nil
}

func (s *websocket[T]) readSubscriptions() []string {
	s.streams.mu.Lock()
	defer s.streams.mu.Unlock()

	keys := make([]string, 0, len(s.streams.data))
	for k := range s.streams.data {
		keys = append(keys, k)
	}

	return keys
}

////
// Embedded public methods
////

func (s *websocket[T]) ListSubscriptions() ([]string, Error) {
	return s.listSubscriptions()
}

func (s *websocket[T]) ReadSubscriptions() []string {
	return s.readSubscriptions()
}

func (s *websocket[T]) Close() {
	s.base.Close()
}

////
// Message Handler
////

type wsMessageEnvelope struct {
	Code   *int            `json:"code"`
	Msg    *string         `json:"msg"`
	Stream *string         `json:"stream"`
	Data   json.RawMessage `json:"data"`
}

func (s *websocket[T]) messageHandler(data []byte) {
	var me wsMessageEnvelope
	if err := json.Unmarshal(data, &me); err != nil {
		s.logger.ERRORf("Failed to parse WS message envelope: %s", err.Error())
		return
	}

	if me.Code != nil {
		s.logger.WARNf("Error message received in socket => code: %d, msg: %s", *me.Code, *me.Msg)
		return
	}

	var result T
	var rawData json.RawMessage

	if me.Stream != nil {
		rawData = me.Data
	} else {
		s.logger.WARNf("Received non-combined stream message")
		rawData = data
	}

	if err := json.Unmarshal(rawData, &result); err != nil {
		s.logger.ERRORf("Failed to parse WS payload into type T => %s", err.Error())
		return
	}

	s.broadcastOnMessage(result)
}

////
// Request Handler
////

type response struct {
	Id     string          `json:"id"`
	Result json.RawMessage `json:"result"`
}

func doWsRequest[R any, T any](s *websocket[T], payload map[string]interface{}) (*R, Error) {
	data, err := s.base.SendRequest(payload, websocketRequestTimeout)
	if err != nil {
		return nil, err
	}

	var response response
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, berror.NewParseError(err)
	}

	var result R
	if err := json.Unmarshal(response.Result, &result); err != nil {
		return nil, berror.NewParseError(err)
	}

	return &result, nil
}

////
// Constructor
////

func newWebsocket[T any](baseUrl string, websocketRoute WebsocketRoute, streams []string, onMessage func(data T), logger logging.Logger) (*websocket[T], Error) {
	socket := &websocket[T]{
		baseUrl: baseUrl,

		baseRoute: websocketRoute,

		logger: logger,

		onMessage: onMessage,
	}
	socket.streams.data = make(map[string]struct{})

	for _, stream := range streams {
		socket.streams.data[stream] = struct{}{}
	}

	base, err := binance.New(socket.buildUrl(), socket.messageHandler, nil, nil, logger)
	if err != nil {
		return nil, berror.NewNetworkError(err)
	}

	socket.base = base

	return socket, nil
}
