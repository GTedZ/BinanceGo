package reconnecting

import (
	"math"
	"time"
)

const (
	InitialBackoff = 100 * time.Millisecond
	MaxBackoff     = 5000 * time.Millisecond
	BackoffFactor  = 2.0
)

// computeBackoffDuration returns the duration to wait before the given retry attempt.
// tries is the number of failed attempts so far.
func computeBackoffDuration(tries int) time.Duration {
	if tries < 1 {
		return InitialBackoff
	}

	tries = min(tries, 50)

	d := float64(InitialBackoff) * math.Pow(BackoffFactor, float64(tries-1))
	if d >= float64(MaxBackoff) {
		return MaxBackoff
	}

	return time.Duration(d)
}
