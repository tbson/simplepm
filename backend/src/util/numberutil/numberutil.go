package numberutil

import (
	"strconv"
)

func StrToInt(id string, defaultValue int) int {
	if id == "" {
		return defaultValue
	}
	if id, err := strconv.Atoi(id); err == nil {
		return id
	}
	return defaultValue
}

func StrToUint(id string, defaultValue uint) uint {
	if id == "" {
		return defaultValue
	}
	if id, err := strconv.Atoi(id); err == nil {
		return uint(id)
	}
	return defaultValue
}

func StrToFloat(id string, defaultValue float64) float64 {
	if id == "" {
		return defaultValue
	}
	if id, err := strconv.ParseFloat(id, 64); err == nil {
		return id
	}
	return defaultValue
}

func UintToStr(id uint) string {
	return strconv.Itoa(int(id))
}

func UintToInt(id uint) int {
	return int(id)
}
