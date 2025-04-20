package srv

import (
	"fmt"
	"src/common/ctype"
	"src/module/account/schema"
	"src/util/errutilnew"
	"src/util/localeutil"
	"src/util/pwdutil"
	"time"

	"src/module/account/domain/srv/pwdpolicy"
)

type userProvider interface {
	Retrieve(opts ctype.QueryOpts) (*schema.User, error)
}

type srv struct {
	userRepo userProvider
}

func New(userRepo userProvider) srv {
	return srv{userRepo}
}

func (srv srv) Login(email string, pwd string, tenantID uint) (ctype.Dict, error) {
	pwdPolicy := pwdpolicy.New()

	// Check user exists
	userOpts := ctype.QueryOpts{
		Filters: ctype.Dict{"Email": email, "TenantID": tenantID},
	}
	user, err := srv.userRepo.Retrieve(userOpts)
	if err != nil {
		return ctype.Dict{}, err
	}

	// Check pwd policy
	var lastResetPwd *time.Time
	if user.PwdResetAt != nil {
		lastResetPwd = user.PwdResetAt
	} else {
		lastResetPwd = &user.CreatedAt
	}
	fmt.Println("lastResetPwd", lastResetPwd)
	err = pwdPolicy.CheckOnValidation(pwd, *lastResetPwd, 0)
	if err != nil {
		return ctype.Dict{}, err
	}

	// Check user is locked
	if user.LockedAt != nil {
		return ctype.Dict{}, errutilnew.NewSimple(localeutil.LockedAccount)
	}

	// Check pwd
	err = pwdutil.CheckPwd(pwd, user.Pwd)
	if err != nil {
		return ctype.Dict{}, err
	}

	return ctype.Dict{}, nil
}
