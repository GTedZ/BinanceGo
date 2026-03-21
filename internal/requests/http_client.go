package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/GTedZ/binancego/internal/apikey"
	"github.com/GTedZ/binancego/internal/berror"
)

var client = http.Client{}

func do(method Method, baseUrl string, path string, params map[string]interface{}, payload map[string]interface{}, apikeys apikey.KeyPair, securityType RequestSecurityType) (Response, berror.Error) {
	var sortKeys bool

	switch securityType {
	case NONE:
		sortKeys = false

	case API_ONLY, SIGNED:
		sortKeys = true
		currentTimestamp := time.Now().UnixMilli()

		params["timestamp"] = currentTimestamp

	default:
		panic(fmt.Sprintf("Invalid security type '%s'", securityType))
	}

	paramString := createQueryString(params, sortKeys)

	switch securityType {
	case SIGNED:
		if apikeys == nil {
			return nil, berror.NewNotFoundError("An API Key is required for this action.")
		}

		signature, signatureError := apikeys.Sign(paramString)
		if signatureError != nil {
			return nil, berror.NewSignatureError(signatureError)
		}

		paramString += "&" + signature
	}

	fullUrl := baseUrl + path + "?" + paramString

	var payloadReader io.Reader
	if payload != nil {
		buf := new(bytes.Buffer)
		if err := json.NewEncoder(buf).Encode(payload); err != nil {
			return nil, berror.NewNetworkError(err)
		}
		payloadReader = buf
	}

	request, reqError := http.NewRequest(
		string(method),
		fullUrl,
		payloadReader,
	)
	if reqError != nil {
		return nil, berror.NewNetworkError(reqError)
	}

	switch securityType {
	case API_ONLY, SIGNED:
		if apikeys == nil {
			return nil, berror.NewNotFoundError("An API Key is required for this action.")
		}

		request.Header.Add("X-MBX-APIKEY", apikeys.ApiKey())
	}

	startTime := time.Now()
	rawResponse, err := client.Do(request)
	if err != nil {
		return nil, berror.NewNetworkError(err)
	}

	response, read_err := toResponse(rawResponse, startTime)
	if read_err != nil {
		return nil, read_err
	}

	return response, read_err
}
