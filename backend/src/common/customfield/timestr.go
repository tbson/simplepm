package customfield

import (
	"encoding/json"
	"fmt"
	"time"
)

// TimeStr is a custom type for time.Time that can be unmarshaled from a string
type TimeStr time.Time

// UnmarshalJSON implements the json.Unmarshaler interface for TimeStr
func (t *TimeStr) UnmarshalJSON(data []byte) error {
	// Remove quotes from string
	s := string(data)
	s = s[1 : len(s)-1] // Remove the surrounding quotes

	// Try parsing with multiple layouts
	formats := []string{
		time.RFC3339,
		"2006-01-02T15:04:05",
		"2006-01-02 15:04:05",
		"2006-01-02",
		"15:04:05",
	}

	var err error
	var parsedTime time.Time

	for _, format := range formats {
		parsedTime, err = time.Parse(format, s)
		if err == nil {
			*t = TimeStr(parsedTime)
			return nil
		}
	}

	// If we reached here, none of the formats worked
	return fmt.Errorf("could not parse time from string: %s", s)
}

// MarshalJSON implements the json.Marshaler interface for TimeStr
func (t *TimeStr) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(*t).Format(time.RFC3339))
}

// String returns the string representation of the time
func (t *TimeStr) String() string {
	if t == nil {
		return ""
	}
	return time.Time(*t).String()
}

// Time returns the time.Time pointer represented by TimeStr
func (t *TimeStr) TimePtr() *time.Time {
	if t == nil {
		return nil
	}
	time := time.Time(*t)
	return &time
}

// Time returns the time.Time represented by TimeStr
func (t *TimeStr) Time() time.Time {
	return time.Time(*t)
}
