package ranges

import "time"

// Period return an array of time.Time comprised between start and end
// every period. (eg. 10h - 11h, 10mn = [ "10h00", "10h10", "10h20", …, "10h50"])
func Period(start, end time.Time, period time.Duration) []time.Time {
	interval := make([]time.Time, 0)
	for curr := start; curr.Before(end); curr = curr.Add(period) {
		interval = append(interval, curr)
	}
	return interval
}
