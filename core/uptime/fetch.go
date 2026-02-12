package uptime

import "time"

var startTime = time.Now()

func FetchUptime() float64 {
	return time.Since(startTime).Seconds()
}
