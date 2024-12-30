package dateutil

import (
	"time"
)

func Now() time.Time {
	return time.Now()
}

func Today() time.Time {
	return time.Now().Truncate(24 * time.Hour)
}

func StrToDate(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02", dateStr)
}
