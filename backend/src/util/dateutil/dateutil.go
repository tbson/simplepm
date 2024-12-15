package dateutil

import (
	"fmt"
	"time"
)

func Now() time.Time {
	return time.Now()
}

func Today() time.Time {
	return time.Now().Truncate(24 * time.Hour)
}

func TraceTime(initMessage string) func(string) string {
	start := time.Now()
	fmt.Println(initMessage)
	return func(msg string) string {
		result := msg + ": " + time.Since(start).String()
		fmt.Println(result)
		return result
	}
}
