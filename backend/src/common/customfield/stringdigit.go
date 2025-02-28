package customfield

import "encoding/json"

type StringDigit string

func (s *StringDigit) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as a number first
	var num json.Number
	if err := json.Unmarshal(data, &num); err == nil {
		*s = StringDigit(num.String())
		return nil
	}

	// If not a number, try as a string
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	*s = StringDigit(str)
	return nil
}

func (s StringDigit) String() string {
	return string(s)
}
