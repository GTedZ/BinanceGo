package Binance

var Errors = errors{
	LibraryCodes: errors_libraryCodes{
		HTTPREQUEST_ERR:       -1,
		SIGNATURE_ERR:         -2,
		RESPONSEBODY_READ_ERR: -3,
		ERROR_PROCESSING_ERR:  -4,
		NOTFOUND_ERR:          -5,
		PARSE_ERR:             -6,
		INVALIDVALUE_ERR:      -7,
	},
}

type errors struct {
	LibraryCodes errors_libraryCodes
}

type errors_libraryCodes struct {
	HTTPREQUEST_ERR       int
	SIGNATURE_ERR         int
	RESPONSEBODY_READ_ERR int
	ERROR_PROCESSING_ERR  int
	NOTFOUND_ERR          int
	PARSE_ERR             int
	INVALIDVALUE_ERR      int
}
