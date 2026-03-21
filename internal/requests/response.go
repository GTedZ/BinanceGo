package requests

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/GTedZ/binancego/internal/berror"
)

type Response interface {
	Status() string
	StatusCode() int
	Proto() string
	ProtoMajor() int
	ProtoMinor() int

	Headers() map[string][]string
	Body() []byte
	ContentLength() int64
	TransferEncoding() []string
	Close() bool
	Uncompressed() bool
	Trailer() http.Header
	Request() *http.Request
	TLS() *tls.ConnectionState

	// Custom

	StartTime() time.Time
	EndTime() time.Time

	// Latency is used to get the elapsed time between sending the request and receiving the response headers.
	//
	// Not to be confused with Duration() which measures the elapsed time between sending the request and reading the full response.
	Latency() time.Duration

	// Duration is used to get the elapsed time between sending the request and receiving the full response.
	Duration() time.Duration

	GetRequestTime() time.Time
}

type HttpResponse struct {
	// The local start time before sending the request
	requestStartTime time.Time
	// The local time that we received the headers
	responseHeadersTime time.Time
	// The local end time after executing the request and reading the body
	requestEndTime time.Time

	status     string // e.g. "200 OK"
	statusCode int    // e.g. 200
	proto      string // e.g. "HTTP/1.0"
	protoMajor int    // e.g. 1
	protoMinor int    // e.g. 0

	// header maps header keys to values. If the response had multiple
	// headers with the same key, they may be concatenated, with comma
	// delimiters.  (RFC 7230, section 3.2.2 requires that multiple headers
	// be semantically equivalent to a comma-delimited sequence.) When
	// header values are duplicated by other fields in this struct (e.g.,
	// ContentLength, TransferEncoding, Trailer), the field values are
	// authoritative.
	//
	// Keys in the map are canonicalized (see CanonicalHeaderKey).
	header http.Header

	body []byte

	// contentLength records the length of the associated content. The
	// value -1 indicates that the length is unknown. Unless Request.Method
	// is "HEAD", values >= 0 indicate that the given number of bytes may
	// be read from Body.
	contentLength int64

	// Contains transfer encodings from outer-most to inner-most. Value is
	// nil, means that "identity" encoding is used.
	transferEncoding []string

	// close records whether the header directed that the connection be
	// closed after reading Body. The value is advice for clients: neither
	// ReadResponse nor Response.Write ever closes a connection.
	close bool

	// uncompressed reports whether the response was sent compressed but
	// was decompressed by the http package. When true, reading from
	// Body yields the uncompressed content instead of the compressed
	// content actually set from the server, ContentLength is set to -1,
	// and the "Content-Length" and "Content-Encoding" fields are deleted
	// from the responseHeader. To get the original response from
	// the server, set Transport.DisableCompression to true.
	uncompressed bool

	// trailer maps trailer keys to values in the same
	// format as Header.
	//
	// The trailer initially contains only nil values, one for
	// each key specified in the server's "trailer" header
	// value. Those values are not added to Header.
	//
	// trailer must not be accessed concurrently with Read calls
	// on the Body.
	//
	// After Body.Read has returned io.EOF, trailer will contain
	// any trailer values sent by the server.
	trailer http.Header

	// request is the request that was sent to obtain this Response.
	// request's Body is nil (having already been consumed).
	// This is only populated for Client requests.
	request *http.Request

	// tls contains information about the tls connection on which the
	// response was received. It is nil for unencrypted responses.
	// The pointer is shared between responses and should not be
	// modified.
	tls *tls.ConnectionState
}

func (r *HttpResponse) Status() string {
	return r.status
}

func (r *HttpResponse) StatusCode() int {
	return r.statusCode
}

func (r *HttpResponse) Proto() string {
	return r.proto
}

func (r *HttpResponse) ProtoMajor() int {
	return r.protoMajor
}

func (r *HttpResponse) ProtoMinor() int {
	return r.protoMinor
}

func (r *HttpResponse) ContentLength() int64 {
	return r.contentLength
}

func (r *HttpResponse) TransferEncoding() []string {
	return r.transferEncoding
}

func (r *HttpResponse) Close() bool {
	return r.close
}

func (r *HttpResponse) Uncompressed() bool {
	return r.uncompressed
}

func (r *HttpResponse) Trailer() http.Header {
	return r.trailer
}

func (r *HttpResponse) Request() *http.Request {
	return r.request
}

func (r *HttpResponse) TLS() *tls.ConnectionState {
	return r.tls
}

func (r *HttpResponse) Headers() map[string][]string {
	return r.header
}

func (r *HttpResponse) Body() []byte {
	return r.body
}

func (r *HttpResponse) StartTime() time.Time {
	return r.requestStartTime
}

func (r *HttpResponse) EndTime() time.Time {
	return r.requestEndTime
}

func (r *HttpResponse) Latency() time.Duration {
	return r.responseHeadersTime.Sub(r.requestStartTime)
}

func (r *HttpResponse) Duration() time.Duration {
	return r.requestEndTime.Sub(r.requestStartTime)
}

func (r *HttpResponse) GetRequestTime() time.Time {
	strValue := r.header.Get("Date")

	if strValue == "" {
		return time.Now()
	}

	parsedTime, err := time.Parse(time.RFC1123, strValue)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return time.Now()
	}

	return parsedTime
}

// //
// "public" handler
// //
func toResponse(rawResponse *http.Response, startTime time.Time) (*HttpResponse, berror.Error) {
	var resp HttpResponse

	resp.requestStartTime = startTime
	resp.responseHeadersTime = time.Now()

	data, err := io.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, berror.NewResponseReadError(err)
	}

	resp.requestEndTime = time.Now()

	resp.body = data
	resp.close = rawResponse.Close
	resp.contentLength = rawResponse.ContentLength
	resp.header = rawResponse.Header
	resp.proto = rawResponse.Proto
	resp.protoMajor = rawResponse.ProtoMajor
	resp.protoMinor = rawResponse.ProtoMinor

	resp.request = rawResponse.Request
	resp.status = rawResponse.Status
	resp.statusCode = rawResponse.StatusCode
	resp.tls = rawResponse.TLS
	resp.trailer = rawResponse.Trailer
	resp.transferEncoding = rawResponse.TransferEncoding
	resp.uncompressed = rawResponse.Uncompressed

	return &resp, nil
}
