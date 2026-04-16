package berror

////
// Interface
////
type Error interface {
	error

	// IsLocal reports whether the error originated locally
	// (network failure, validation, etc.) rather than from the Binance API.
	IsLocal() bool

	// Code returns the error code associated with the error.
	Code() int

	// Message returns a human-readable message describing the error.
	Message() string

	// Unwrap returns the underlying error, if any. If there is no underlying error, Unwrap returns nil.
	// IMPORTANT: `error` is `nil` for any non-local error, since binance simply rejecting the request doesn't have an underlying error in the context of the library.
	Unwrap() error

	// Matches checks if the error code matches the provided code.
	Matches(ErrorCode) bool
}
