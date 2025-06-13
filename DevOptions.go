package Binance

type DevOpts struct {
	PRINT_HTTP_QUERIES   bool
	PRINT_HTTP_RESPONSES bool

	// Might not log every error returned.
	// It only prints errors specifically handled by the library.
	PRINT_ERRORS bool
	// This is redundant.
	// Any *Error type generated by the library will be logged.
	PRINT_ALL_ERRORS bool

	// Prints all useful activity data and errors.
	// i.e: Forced reconnections, disconnections, etc...
	WS_VERBOSE bool
	// Prints all websocket data.
	// i.e: pings received, pongs sent, etc...
	WS_VERBOSE_FULL bool
	PRINT_WS_ERRORS bool

	// Recommended only for debugging.
	PRINT_WS_MESSAGES bool
}

var DevOptions = DevOpts{
	PRINT_HTTP_QUERIES:   false,
	PRINT_HTTP_RESPONSES: false,

	PRINT_ERRORS:     false,
	PRINT_ALL_ERRORS: false,

	WS_VERBOSE:      false,
	WS_VERBOSE_FULL: false,

	PRINT_WS_ERRORS: true,

	PRINT_WS_MESSAGES: false,
}
