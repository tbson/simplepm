package debugutil

import (
	"fmt"
	"time"
)

func TraceTime(initMessage string) func(string) string {
	start := time.Now()
	fmt.Println(initMessage)
	return func(msg string) string {
		result := msg + ": " + time.Since(start).String()
		fmt.Println(result)
		return result
	}
}
