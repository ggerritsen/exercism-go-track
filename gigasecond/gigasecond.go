// gigasecond provides functionality around time.Time modifications related to giga seconds
package gigasecond

import "time"

// AddGigasecond returns the time 1 gigasecond after `t`.
// A gigasecond is 10^9 (1,000,000,000) seconds.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1000000000 * time.Second)
}
