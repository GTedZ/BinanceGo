package futures

import (
	"fmt"

	"github.com/GTedZ/binancego/internal/berror"
	"github.com/GTedZ/binancego/internal/json"
	"github.com/GTedZ/binancego/internal/requests"
)

func doRequest[T any](c *Client, method RequestMethod, path string, params map[string]interface{}, payload map[string]interface{}, securityType SecurityType) (*T, Response, Error) {
	resp, err := c.MakeRequest(method, path, params, payload, securityType)
	if err != nil {
		return nil, resp, err
	}

	var result T
	parseErr := json.Unmarshal(resp.Body(), &result)
	if parseErr != nil {
		return nil, nil, berror.NewParseError(parseErr)
	}

	return &result, resp, nil
}

// Used to execute client requests.
// The Base URL used is the Spot Client's URL.
// If the securityType is `NONE`, it will default to the special URL provided by binance (currently `https://data-api.binance.vision`) to speed up query speed.
// Can be overridden via the `.SetBaseUrlMarketData()`
func (c *Client) MakeRequest(method RequestMethod, path string, params map[string]interface{}, payload map[string]interface{}, securityType SecurityType) (Response, Error) {

	switch securityType {
	case NONE:
		return requests.Unsigned(method, c.baseUrl, path, params, nil)

	case USER_STREAM:
		return requests.ApiOnly(method, c.baseUrl, path, params, payload, c.apikey)

	case TRADE, USER_DATA:
		return requests.Signed(method, c.baseUrl, path, params, payload, c.apikey)

	default:
		return nil, berror.NewInvalidValueError(fmt.Sprintf("Security Type passed to Request function is invalid, received: '%s'\nSupported methods are ('%s', '%s', '%s', '%s')", securityType, NONE, USER_STREAM, TRADE, USER_DATA))
	}
}
