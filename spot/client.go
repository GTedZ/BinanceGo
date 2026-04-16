package spot

import (
	"github.com/GTedZ/binancego/internal/apikey"
	"github.com/GTedZ/binancego/internal/logging"
)

type Client struct {
	apikey KeyPair

	baseUrl           string
	baseUrlMarketData string

	wssBaseUrlMarketData string
	wssBaseUrl           string

	wssApiBaseUrl string

	logger logging.Logger

	Websocket *wsEndpoints
}

//

func new() *Client {
	c := &Client{
		apikey: &apikey.NilKeyPair{},

		baseUrl:           Constants.BaseUrls.Api,
		baseUrlMarketData: Constants.BaseUrlsMarketData.DataApiVision,

		wssBaseUrlMarketData: Constants.WssBaseUrlsMarketData.DataStreamBinanceVision,
		wssBaseUrl:           Constants.WssBaseUrls.Stream9443,

		wssApiBaseUrl: Constants.WssApiBaseUrls.WsApi,

		logger: logging.NewNilLogger(),
	}

	c.Websocket = newWsEndpoints(c)

	return c
}

func New(opts ...ClientOption) (*Client, Error) {
	c := new()

	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func NewReadClient() *Client {
	return new()
}
