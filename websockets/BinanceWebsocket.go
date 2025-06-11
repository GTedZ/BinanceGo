package websockets

import (
	"strings"

	jsoniter "github.com/json-iterator/go"
)

type BinanceWebsocket struct {
	base *reconnectingPrivateMessageWebsocket

	baseURL    string
	streams    []string
	isCombined bool

	OnMessage   func(messageType int, msg []byte)
	OnReconnect func()
}

type CombinedStream_MSG struct {
	Stream string              `json:"stream"`
	Data   jsoniter.RawMessage `json:"data"`
}

////

func (socket *BinanceWebsocket) onMessage(messageType int, msg []byte) {
	if socket.isCombined {
		var tempData CombinedStream_MSG
		err := jsoniter.Unmarshal(msg, &tempData)
		if err != nil {
			Logger.ERROR("ERR PARSING COMBINED MESSAGE!!!", err)
			return
		}
		msg = tempData.Data
	}

	if socket.OnMessage != nil {
		socket.OnMessage(messageType, msg)
	}
}

func (socket *BinanceWebsocket) onReconnect() {
	if socket.OnReconnect != nil {
		socket.OnReconnect()
	}
}

func (socket *BinanceWebsocket) buildURL(streams []string) string {
	queryURL := ""

	socket.isCombined = len(streams) > 1

	if len(streams) != 0 {
		if len(streams) == 1 {
			queryURL += "/ws/" + streams[0]
		} else {
			queryURL += "/stream?streams=" + strings.Join(streams, "/")
		}
	}

	fullStreamURL := socket.baseURL + queryURL

	return fullStreamURL
}

//// Public Methods

func (socket *BinanceWebsocket) GetStreams() []string {
	return socket.streams
}

func (socket *BinanceWebsocket) SetStreams(streams []string) {
	socket.streams = streams
}

func (socket *BinanceWebsocket) Close() {
	socket.base.Close()
}

func (socket *BinanceWebsocket) SendPrivateMessage(message map[string]interface{}, timeout_sec ...int) (response []byte, hasTimedOut bool, err error) {
	return socket.base.SendPrivateMessage(message, timeout_sec...)
}

func (socket *BinanceWebsocket) ListSubscriptions(timeout_sec ...int) (subscriptions []string, err error) {
	request := make(map[string]interface{})
	request["method"] = "LIST_SUBSCRIPTIONS"
	resp, _, err := socket.SendPrivateMessage(request, timeout_sec...)
	if err != nil {
		return nil, err
	}

	var response struct {
		Id     string   `json:"id"`
		Result []string `json:"result"`
	}
	err = jsoniter.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}

	socket.streams = response.Result
	socket.UpdateStreams()

	return response.Result, err
}

func (socket *BinanceWebsocket) UpdateStreams() {
	socket.streams = utils.RemoveDuplicates(socket.streams)

	fullStreamURL := socket.buildURL(socket.streams)
	socket.base.SetURL(fullStreamURL)
}

////

// This function WILL block until a successful connection can be established
func CreateBinanceWebsocket(baseURL string, streams []string) *BinanceWebsocket {
	var socket = &BinanceWebsocket{
		baseURL: baseURL,
		streams: streams,
	}

	fullStreamURL := socket.buildURL(streams)

	socket.base = createReconnectingPrivateMessageWebsocket(fullStreamURL, "id")

	socket.base.OnMessage = socket.onMessage
	socket.base.OnReconnect = socket.onReconnect

	return socket
}
