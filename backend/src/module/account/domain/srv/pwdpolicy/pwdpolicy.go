package pwdpolicy

import (
	"src/common/ctype"
	"src/util/errutilnew"
	"src/util/localeutil"
	"time"

	"src/common/setting"
)

var maxResetPwdPeriodDays = setting.MAX_RESET_PWD_PERIOD_DAYS()
var maxFailedAttempts = setting.MAX_PWD_FAILED_ATTEMPTS()
var lastPwdsCheck = setting.LAST_PWDS_CHECK()

type pwdHistory []string

type srv struct{}

func New() *srv {
	return &srv{}
}

func (s srv) CheckOnCreation(pwd string, pwdHistory pwdHistory) error {
	complexityErr := checkPwdComplexityPolicy(pwd)
	uniquenessErr := checkPwdUniquenessPolicy(pwd, pwdHistory)

	err := errutilnew.NewEmpty()
	if complexityErr != nil {
		err.Merge(complexityErr.(*errutilnew.CustomError))
	}
	if uniquenessErr != nil {
		err.Merge(uniquenessErr.(*errutilnew.CustomError))
	}

	if len(err.Errors) > 0 {
		return err
	}
	return nil
}

func (s srv) CheckOnValidation(pwd string, lastResetPwd time.Time, failedAttempts int) error {
	rotationErr := checkPwdRotationPolicy(lastResetPwd)
	failedAttempsErr := checkFailedAttemptsPolicy(pwd, failedAttempts)

	err := errutilnew.NewEmpty()
	if rotationErr != nil {
		err.Merge(rotationErr.(*errutilnew.CustomError))
	}
	if failedAttempsErr != nil {
		err.Merge(failedAttempsErr.(*errutilnew.CustomError))
	}
	if len(err.Errors) > 0 {
		return err
	}
	return nil
}

func checkPwdComplexityPolicy(pwd string) error {
	// to be implemented
	return nil
}

func checkPwdUniquenessPolicy(pwd string, pwdHistory pwdHistory) error {
	// to be implemented
	return nil
}

func checkPwdRotationPolicy(lastResetPwd time.Time) error {
	if int(time.Now().Sub(lastResetPwd).Hours()/24) <= maxResetPwdPeriodDays {
		return nil
	}

	return errutilnew.NewSimpleWithArgs(
		localeutil.RotatePwdPolicyFail,
		ctype.Dict{
			"Value": maxResetPwdPeriodDays,
		},
	)
}

func checkFailedAttemptsPolicy(pwd string, failedAttempts int) error {
	// to be implemented
	return nil
}
