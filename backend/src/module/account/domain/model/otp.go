package model

import (
	"src/common/ctype"
	"src/util/stringutil"
	"time"

	"src/util/i18nmsg"

	"src/util/errutil"
)

type OTP struct {
	code      string
	expiredAt time.Time
}

func NewOTP(length int, lifeMinutes int) OTP {
	code := stringutil.GetRandomString(length)
	now := time.Now()
	expiredAt := now.Add(time.Duration(lifeMinutes) * time.Minute)

	return OTP{
		code:      code,
		expiredAt: expiredAt,
	}
}

func ParseOTP(code string, expiredAt time.Time, length int) (OTP, error) {
	if len(code) != length {
		err := errutil.NewWithArgs(
			i18nmsg.OTPLengthConditionFail,
			ctype.Dict{
				"Value": length,
			},
		)
		return OTP{}, err
	}

	now := time.Now()
	if now.After(expiredAt) {
		err := errutil.New(i18nmsg.OTPExpired)
		return OTP{}, err
	}

	return OTP{
		code:      code,
		expiredAt: expiredAt,
	}, nil
}
