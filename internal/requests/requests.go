package requests

import (
	"github.com/GTedZ/binancego/internal/apikey"
	"github.com/GTedZ/binancego/internal/berror"
)

func Unsigned(method Method, baseUrl string, path string, params map[string]interface{}, payload map[string]interface{}) (Response, berror.Error) {
	return do(method, baseUrl, path, params, payload, nil, NONE)
}

func ApiOnly(method Method, baseUrl string, path string, params map[string]interface{}, payload map[string]interface{}, apikey apikey.KeyPair) (Response, berror.Error) {
	return do(method, baseUrl, path, params, payload, apikey, API_ONLY)
}

func Signed(method Method, baseUrl string, path string, params map[string]interface{}, payload map[string]interface{}, apikey apikey.KeyPair) (Response, berror.Error) {
	return do(method, baseUrl, path, params, payload, apikey, SIGNED)
}
