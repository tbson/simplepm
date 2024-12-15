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
