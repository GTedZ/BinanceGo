package Binance

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/GTedZ/binancego/lib"
)

type RequestClient struct {
	binance *Binance

	client *http.Client
}

func (requestClient *RequestClient) init(binance *Binance) {
	requestClient.binance = binance
	requestClient.client = &http.Client{}
}

type Response struct {
	Status     string // e.g. "200 OK"
	StatusCode int    // e.g. 200
	Proto      string // e.g. "HTTP/1.0"
	ProtoMajor int    // e.g. 1
	ProtoMinor int    // e.g. 0

	// This is added by the 'Binance-Go' library
	// It is simply the elapsed time between sending the request and receiving the response
	Latency int64

	// Header maps header keys to values. If the response had multiple
	// headers with the same key, they may be concatenated, with comma
	// delimiters.  (RFC 7230, section 3.2.2 requires that multiple headers
	// be semantically equivalent to a comma-delimited sequence.) When
	// Header values are duplicated by other fields in this struct (e.g.,
	// ContentLength, TransferEncoding, Trailer), the field values are
	// authoritative.
	//
	// Keys in the map are canonicalized (see CanonicalHeaderKey).
	Header http.Header

	Body []byte

	// ContentLength records the length of the associated content. The
	// value -1 indicates that the length is unknown. Unless Request.Method
	// is "HEAD", values >= 0 indicate that the given number of bytes may
	// be read from Body.
	ContentLength int64

	// Contains transfer encodings from outer-most to inner-most. Value is
	// nil, means that "identity" encoding is used.
	TransferEncoding []string

	// Close records whether the header directed that the connection be
	// closed after reading Body. The value is advice for clients: neither
	// ReadResponse nor Response.Write ever closes a connection.
	Close bool

	// Uncompressed reports whether the response was sent compressed but
	// was decompressed by the http package. When true, reading from
	// Body yields the uncompressed content instead of the compressed
	// content actually set from the server, ContentLength is set to -1,
	// and the "Content-Length" and "Content-Encoding" fields are deleted
	// from the responseHeader. To get the original response from
	// the server, set Transport.DisableCompression to true.
	Uncompressed bool

	// Trailer maps trailer keys to values in the same
	// format as Header.
	//
	// The Trailer initially contains only nil values, one for
	// each key specified in the server's "Trailer" header
	// value. Those values are not added to Header.
	//
	// Trailer must not be accessed concurrently with Read calls
	// on the Body.
	//
	// After Body.Read has returned io.EOF, Trailer will contain
	// any trailer values sent by the server.
	Trailer http.Header

	// Request is the request that was sent to obtain this Response.
	// Request's Body is nil (having already been consumed).
	// This is only populated for Client requests.
	Request *http.Request

	// TLS contains information about the TLS connection on which the
	// response was received. It is nil for unencrypted responses.
	// The pointer is shared between responses and should not be
	// modified.
	TLS *tls.ConnectionState
}

// # Fetches the current used weight returned the request.
//
// interval: "1m", "3m", "1d", "1W", "1M", or simply ""
//
// But most common and only one used as of writing this is "1m"
//
// Returns an error if the header is not found
func (resp *Response) GetUsedWeight(interval string) (int64, error) {
	key := "X-Mbx-Used-Weight"
	if interval != "" {
		key += "-" + interval
	}

	strValue := resp.Header.Get(key)

	if strValue == "" {
		return 0, fmt.Errorf("no Used Weight was found for this interval")
	}

	// Parses the value to int64
	value, err := strconv.ParseInt(strValue, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("Error parsing header value: %s", err.Error())
	}

	return value, nil
}

// # Extracts the used weight and the request time
//
// If the used weight EXCEEDS "weightLimit"
//
// # It will wait until the next reset time before returning from the function call
//
// In short, after each request, call this function, if the returned error is nil, you're free to continue with your next request
//
// Usual interval used by binance is "1m"
func (resp *Response) WaitUsedWeight(interval string, weightLimit int64, nextRequestWeight ...int64) error {
	kline_interval, _, err := Utils.GetIntervalFromString(interval)
	if err != nil {
		return err
	}

	var usedWeight int64

	usedWeight, err = resp.GetUsedWeight(interval)
	if err != nil {
		return err
	}

	if len(nextRequestWeight) != 0 {
		usedWeight += nextRequestWeight[0]
	}

	if usedWeight > weightLimit {
		total_intervalValue := kline_interval.Value

		requestTime, err := resp.GetRequestTime()
		if err != nil {
			return err
		}

		unixMilli := requestTime.UnixMilli()
		millis_toWait := unixMilli - (unixMilli % total_intervalValue)

		time.Sleep(time.Duration(millis_toWait) * time.Millisecond)
	}

	return nil
}

func (resp *Response) GetRequestTime() (time.Time, error) {
	key := "Date"

	strValue := resp.Header.Get(key)

	if strValue == "" {
		return time.Now(), fmt.Errorf("no Date header was found for this request")
	}

	parsedTime, err := time.Parse(time.RFC1123, strValue)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return time.Now(), lib.LocalError(Errors.LibraryCodes.PARSE_ERR, "There was an error parsing the date from request headers")
	}

	return parsedTime, nil
}

func (resp *Response) GetLatency() (latency int64, err error) {
	if resp == nil {
		return 0, fmt.Errorf("cannot read latency from nil response")
	}
	return resp.Latency, nil
}

func (requestClient *RequestClient) readResponseBody(rawResponse *http.Response) (*Response, error) {
	var resp Response

	data, err := io.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	resp.Body = data
	resp.Close = rawResponse.Close
	resp.ContentLength = rawResponse.ContentLength
	resp.Header = rawResponse.Header
	resp.Proto = rawResponse.Proto
	resp.ProtoMajor = rawResponse.ProtoMajor
	resp.ProtoMinor = rawResponse.ProtoMinor
	resp.Request = rawResponse.Request
	resp.Status = rawResponse.Status
	resp.StatusCode = rawResponse.StatusCode
	resp.TLS = rawResponse.TLS
	resp.Trailer = rawResponse.Trailer
	resp.TransferEncoding = rawResponse.TransferEncoding
	resp.Uncompressed = rawResponse.Uncompressed

	return &resp, nil
}

func (requestClient *RequestClient) processResponse(rawResponse *http.Response, latency int64) (*Response, *Error) {
	resp, err := requestClient.readResponseBody(rawResponse)
	if err != nil {
		return nil, lib.LocalError(Errors.LibraryCodes.RESPONSEBODY_READ_ERR, err.Error())
	}
	resp.Latency = latency

	if resp.StatusCode >= 400 {
		err, unmarshallErr := lib.BinanceError(resp.StatusCode, resp.Body)
		if unmarshallErr != nil {
			return resp, lib.LocalError(Errors.LibraryCodes.ERROR_PROCESSING_ERR, unmarshallErr.Error())
		}

		return resp, err
	}

	return resp, nil
}

//

func (requestClient *RequestClient) Unsigned(method string, baseURL string, URL string, params map[string]interface{}) (*Response, *Error) {
	if params == nil {
		panic("'params' map must be initialized")
	}

	var err error
	var rawResponse *http.Response

	paramString := Utils.CreateQueryString(params, false)

	fullQuery := baseURL + URL + "?" + paramString

	startTime := time.Now().UnixMilli()
	switch method {
	case Constants.Methods.GET:
		rawResponse, err = http.Get(fullQuery)

	default:
		panic(fmt.Sprintf("Method passed to Unsigned Request function is invalid, received: '%s'\nSupported methods are ('%s', '%s', '%s', '%s', '%s')", method, Constants.Methods.GET, Constants.Methods.POST, Constants.Methods.PUT, Constants.Methods.PATCH, Constants.Methods.DELETE))
	}
	if err != nil {
		return nil, lib.LocalError(Errors.LibraryCodes.HTTPREQUEST_ERR, err.Error())
	}
	defer rawResponse.Body.Close()
	latency := time.Now().UnixMilli() - startTime

	resp, respErr := requestClient.processResponse(rawResponse, latency)

	Logger.DEBUG(fmt.Sprintf("%s %s: %s\n", resp.Request.Method, resp.Status, fullQuery))
	Logger.DEBUG(fmt.Sprintf("%s %s: %s =>\nResponse: %s\n", resp.Request.Method, resp.Status, fullQuery, string(resp.Body)))

	return resp, respErr
}

func (requestClient *RequestClient) APIKEY_only(method string, baseURL string, URL string, params map[string]interface{}) (*Response, *Error) {
	if params == nil {
		panic("'params' map must be initialized")
	}

	paramString := Utils.CreateQueryString(params, false)

	startTime := time.Now().UnixMilli()
	fullQuery := baseURL + URL + "?" + paramString

	req, err := http.NewRequest(method, fullQuery, nil)
	if err != nil {
		return nil, lib.LocalError(Errors.LibraryCodes.HTTPREQUEST_ERR, err.Error())
	}

	APIKEY := requestClient.binance.API.GetAPIKEY()
	req.Header.Set("X-MBX-APIKEY", APIKEY)

	rawResponse, err := requestClient.client.Do(req)
	if err != nil {
		Logger.DEBUG("[VERBOSE] Request error", err)
		localErr := lib.LocalError(Errors.LibraryCodes.HTTPREQUEST_ERR, err.Error())
		return nil, localErr
	}
	defer rawResponse.Body.Close()
	latency := time.Now().UnixMilli() - startTime

	resp, respErr := requestClient.processResponse(rawResponse, latency)

	Logger.DEBUG(fmt.Sprintf("%s %s: %s\n", resp.Request.Method, resp.Status, fullQuery))
	Logger.DEBUG(fmt.Sprintf("%s %s: %s =>\nResponse: %s\n", resp.Request.Method, resp.Status, fullQuery, string(resp.Body)))

	return resp, respErr
}

func (requestClient *RequestClient) Signed(method string, baseURL string, URL string, params map[string]interface{}) (*Response, *Error) {
	if params == nil {
		panic("'params' map must be initialized")
	}

	params["timestamp"] = time.Now().UnixMilli() + requestClient.binance.Opts.timestamp_offset

	if params["recvWindow"] == nil {

		if requestClient.binance.Opts.recvWindow != 0 {
			params["recvWindow"] = requestClient.binance.Opts.recvWindow
		}

	}

	paramString := Utils.CreateQueryString(params, true)
	fmt.Println(paramString)

	APIKEY := requestClient.binance.API.GetAPIKEY()

	signature, err := requestClient.binance.API.Sign(paramString)
	if err != nil {
		return nil, lib.LocalError(Errors.LibraryCodes.SIGNATURE_ERR, err.Error())
	}

	startTime := time.Now().UnixMilli()
	fullQuery := baseURL + URL + "?" + paramString + "&signature=" + signature

	req, err := http.NewRequest(method, fullQuery, nil)
	if err != nil {
		return nil, lib.LocalError(Errors.LibraryCodes.HTTPREQUEST_ERR, err.Error())
	}

	req.Header.Set("X-MBX-APIKEY", APIKEY)

	rawResponse, err := requestClient.client.Do(req)
	if err != nil {
		localErr := lib.LocalError(Errors.LibraryCodes.HTTPREQUEST_ERR, err.Error())
		return nil, localErr
	}
	defer rawResponse.Body.Close()
	latency := time.Now().UnixMilli() - startTime

	resp, respErr := requestClient.processResponse(rawResponse, latency)

	Logger.DEBUG(fmt.Sprintf("%s %s: %s\n", resp.Request.Method, resp.Status, fullQuery))
	Logger.DEBUG(fmt.Sprintf("%s %s: %s =>\nResponse: %s\n", resp.Request.Method, resp.Status, fullQuery, string(resp.Body)))

	return resp, respErr
}
