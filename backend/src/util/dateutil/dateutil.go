package dateutil

import (
	"src/util/errutil"
	"src/util/localeutil"
	"time"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func Now() time.Time {
	return time.Now()
}

func Today() time.Time {
	return time.Now().Truncate(24 * time.Hour)
}

func StrToDate(dateStr string) (time.Time, error) {
	emptyTime := time.Time{}
	localizer := localeutil.Get()
	msg := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: localeutil.CanNotParseDateStr,
	})
	if dateStr == "" {
		return emptyTime, errutil.New("", []string{msg})
	}
	result, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return emptyTime, errutil.New("", []string{msg})
	}
	return result, nil
}

func TimeToStr(t time.Time) string {
	return t.Format("2006-01-02T15:04:05Z07:00")
}
