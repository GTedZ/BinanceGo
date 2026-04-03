package spot

import (
	"net/url"
	"strings"

	"github.com/GTedZ/binancego/internal/berror"
	"github.com/GTedZ/binancego/internal/logging"
)

// //
// Client builder
// //
type ClientOption func(*Client) Error

func WithAPIKey(keypair KeyPair) ClientOption {
	return func(c *Client) Error {
		c.apikey = keypair
		return nil
	}
}

func WithBaseURL(url string) ClientOption {
	return func(c *Client) Error {
		url, err := validateAndNormalizeBaseUrl(url)
		if err != nil {
			return err
		}

		c.baseUrl = url
		return nil
	}
}

func WithMarketDataURL(url string) ClientOption {
	return func(c *Client) Error {
		url, err := validateAndNormalizeBaseUrl(url)
		if err != nil {
			return err
		}

		c.baseUrlMarketData = url
		return nil
	}
}

func validateAndNormalizeBaseUrl(raw string) (string, Error) {
	u, err := url.Parse(raw)
	if err != nil {
		return raw, berror.NewInvalidValueError("Invalid base URL")
	}

	if u.Scheme == "" || u.Host == "" {
		return raw, berror.NewInvalidValueError("Invalid base URL")
	}

	// remove trailing slash from path
	u.Path = strings.TrimSuffix(u.Path, "/")

	// remove query/fragment if present
	u.RawQuery = ""
	u.Fragment = ""

	return u.String(), nil
}

func WithLogger(logger logging.Logger) ClientOption {
	return func(c *Client) Error {
		c.logger = logger
		return nil
	}
}

////
// ExchangeInfo builder
////

type exchangeInfoParams map[string]interface{}

type ExchangeInfoOption func(exchangeInfoParams) Error

func WithSymbols(s []string) ExchangeInfoOption {
	return func(eip exchangeInfoParams) Error {
		eip["symbols"] = s
		return nil
	}
}

func WithPermissions(perms []Permission) ExchangeInfoOption {
	return func(eip exchangeInfoParams) Error {
		eip["permissions"] = perms
		return nil
	}
}

func WithSymbolStatus(s SymbolStatus) ExchangeInfoOption {
	return func(eip exchangeInfoParams) Error {
		eip["symbolStatus"] = s
		return nil
	}
}

func WithShowPermissionSets(value bool) ExchangeInfoOption {
	return func(eip exchangeInfoParams) Error {
		eip["showPermissionSets"] = value
		return nil
	}
}

func buildExchangeInfoParams(opts ...ExchangeInfoOption) (exchangeInfoParams, Error) {
	params := make(exchangeInfoParams)
	for _, builderFunc := range opts {
		err := builderFunc(params)
		if err != nil {
			return nil, err
		}
	}

	return params, nil
}
