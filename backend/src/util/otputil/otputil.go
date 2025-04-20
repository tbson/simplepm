package otputil

import (
	"src/util/stringutil"
)

func GenerateOtp(otpLength int) string {
	return stringutil.GetRandomString(otpLength)
}
