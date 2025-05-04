package srv

import (
	"src/common/ctype"
	"src/module/account/schema"
	"src/util/errutil"
	"src/util/i18nmsg"
	"src/util/pwdutil"
	"time"

	"src/module/account/domain/model"
)

type authTokenProvider interface {
	GenerateTokenPair(userID uint) (model.TokenPair, error)
}

type pwdPolicyProvider interface {
	CheckOnValidation(pwd string, lastResetPwd time.Time, failedAttempts int) error
}

type userProvider interface {
	Retrieve(opts ctype.QueryOpts) (*schema.User, error)
}

type srv struct {
	userRepo     userProvider
	authTokenSrv authTokenProvider
	pwdPolicySrv pwdPolicyProvider
}

func New(
	userRepo userProvider,
	authTokenSrv authTokenProvider,
	pwdPolicySrv pwdPolicyProvider,
) srv {
	return srv{userRepo, authTokenSrv, pwdPolicySrv}
}

func (srv srv) Login(email string, pwd string, tenantID uint) (model.AuthResult, error) {
	result := model.AuthResult{}
	// Check user exists
	userOpts := ctype.QueryOpts{
		Filters: ctype.Dict{"Email": email, "TenantID": tenantID},
	}
	user, err := srv.userRepo.Retrieve(userOpts)
	if err != nil {
		return result, err
	}

	// Check pwd policy
	var lastResetPwd *time.Time
	if user.PwdResetAt != nil {
		lastResetPwd = user.PwdResetAt
	} else {
		lastResetPwd = &user.CreatedAt
	}
	err = srv.pwdPolicySrv.CheckOnValidation(pwd, *lastResetPwd, 0)
	if err != nil {
		return result, err
	}

	// Check user is locked
	if user.LockedAt != nil {
		return result, errutil.New(i18nmsg.LockedAccount)
	}

	// Check pwd
	err = pwdutil.CheckPwd(pwd, user.Pwd)
	if err != nil {
		return result, err
	}

	// Generate token pair
	tokenPair, err := srv.authTokenSrv.GenerateTokenPair(user.ID)
	userInfo := model.NewUserInfo(
		user.ID,
		user.TenantID,
		user.Admin,
		user.FirstName,
		user.LastName,
		user.Avatar,
	)
	result.TokenPair = tokenPair
	result.UserInfo = userInfo

	return result, nil
}
