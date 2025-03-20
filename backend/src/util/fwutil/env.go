package fwutil

import (
	"os"
	"src/util/numberutil"
)

func Env(key string, defaultVal string) string {
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

func BoolEnv(key string, defaultVal bool) bool {
	val, ok := os.LookupEnv(key)
	if ok {
		return val == "true"
	}
	return defaultVal
}
