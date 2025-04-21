package fwutil

import (
	"os"
	"src/util/numberutil"
)

func StrEnv(key string, defaultVal string) string {
	val, ok := os.LookupEnv(key)
	if ok {
		return val
	}
	return defaultVal
}

func IntEnv(key string, defaultVal int) int {
	val, ok := os.LookupEnv(key)
	if ok {
		return numberutil.StrToInt(val, defaultVal)
	}
	return defaultVal
}

func UintEnv(key string, defaultVal uint) uint {
	val, ok := os.LookupEnv(key)
	if ok {
		return numberutil.StrToUint(val, defaultVal)
	}
	return defaultVal
}

func FloatEnv(key string, defaultVal float64) float64 {
	val, ok := os.LookupEnv(key)
	if ok {
		return numberutil.StrToFloat(val, defaultVal)
	}
	return defaultVal
}

func BoolEnv(key string, defaultVal bool) bool {
	val, ok := os.LookupEnv(key)
	if ok {
		return val == "true"
	}
	return defaultVal
}
