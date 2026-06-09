package futures

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

func WithWebsocketBaseUrl(url string) ClientOption {
	return func(c *Client) Error {
		url, err := validateAndNormalizeBaseUrl(url)
		if err != nil {
			return err
		}

		c.wssBaseUrl = url
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
