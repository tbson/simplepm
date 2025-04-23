package dateutil

import (
	"src/util/errutil"
	"src/util/localeutil"
	"time"
)

func Now() time.Time {
	return time.Now()
}

func Today() time.Time {
	return time.Now().Truncate(24 * time.Hour)
}

func StrToDate(dateStr string) (time.Time, error) {
	emptyTime := time.Time{}
	errObj := errutil.New(localeutil.CanNotParseDateStr)
	if dateStr == "" {
		return emptyTime, errObj
	}
	result, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return emptyTime, errObj
	}
	return result, nil
}

func TimeToStr(t time.Time) string {
	return t.Format("2006-01-02T15:04:05Z07:00")
}
