package srv

import (
	"src/common/ctype"
	"src/module/account/schema"
	"src/util/errutil"
	"src/util/localeutil"
	"src/util/pwdutil"

	"github.com/nicksnyder/go-i18n/v2/i18n"
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
	localizer := localeutil.Get()

	// Check user exists
	userOpts := ctype.QueryOpts{
		Filters: ctype.Dict{"Email": email, "TenantID": tenantID},
	}
	user, err := srv.userRepo.Retrieve(userOpts)
	if err != nil {
		return ctype.Dict{}, err
	}

	// Check user is locked
	if user.LockedAt != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.LockedAccount,
		})
		return ctype.Dict{}, errutil.New("", []string{msg})
	}

	// Check pwd
	ok, err := pwdutil.CheckPwd(pwd, user.Pwd)
	if err != nil {
		return ctype.Dict{}, err
	}

	if !ok {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.InvalidUsernameOrPwd,
		})
		return ctype.Dict{}, errutil.New("", []string{msg})
	}

	return ctype.Dict{}, nil
}
