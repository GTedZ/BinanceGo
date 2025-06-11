package Binance

import "sync"

const MILLISECOND = 1
const SECOND = 1000 * MILLISECOND
const MINUTE = 60 * SECOND
const HOUR = 60 * MINUTE
const DAY = 24 * HOUR
const WEEK = 7 * DAY

var INTERVALS_mu sync.Mutex
var STATIC_INTERVAL_CHARS = map[rune]int64{
	'x': MILLISECOND,
	's': SECOND,
	'm': MINUTE,
	'h': HOUR,
	'd': DAY,
}
var COMPLEX_INTERVALS = struct {
	WEEK  rune
	MONTH rune
	YEAR  rune
}{
	WEEK:  'w',
	MONTH: 'M',
	YEAR:  'Y',
}

var Constants = struct {
	Methods Methods
}{
	Methods: Methods{
		GET:    "GET",
		POST:   "POST",
		PUT:    "PUT",
		PATCH:  "PATCH",
		DELETE: "DELETE",
	},
}

type Methods struct {
	GET    string
	POST   string
	PUT    string
	PATCH  string
	DELETE string
}
