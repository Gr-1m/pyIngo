package usetime

import (
	"fmt"
	"time"
)

// Just add "defer TotalTime(youstr, time.Now())"
// Before your script Runs
func TotalTime(youstr string, start time.Time) {
	fmt.Printf("%s: %.3fs\n", youstr, time.Since(start).Seconds())
}
