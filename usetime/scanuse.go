package usetime

import(
	"time"
	"fmt"
)

// Just add "defer usetime.TotalTime(youstr, time.Now())" 
// Before your script runs
func TotalTime(youstr string, start time.Time){
	fmt.Println("%s: %.3fs", youstr, time.Since(start).Seconds())
}
