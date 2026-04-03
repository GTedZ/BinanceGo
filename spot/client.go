package spot

import (
	"github.com/GTedZ/binancego/internal/logging"
)

type Client struct {
	apikey KeyPair

	baseUrl           string
	baseUrlMarketData string

	logger logging.Logger
}

//

func new() *Client {
	return &Client{
		apikey:            nil,
		baseUrl:           Constants.BaseUrls.Api,
		baseUrlMarketData: Constants.BaseUrlsMarketData.DataApiVision,

		logger: logging.NewNilLogger(),
	}
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
