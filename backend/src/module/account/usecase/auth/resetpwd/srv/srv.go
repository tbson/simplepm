package srv

import (
	"src/common/ctype"
	"src/module/account/schema"
	"src/util/dateutil"
	"src/util/errutil"
	"src/util/localeutil"
	"src/util/pwdutil"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type userProvider interface {
	Retrieve(opts ctype.QueryOpts) (*schema.User, error)
	Update(opts ctype.QueryOpts, data ctype.Dict) (*schema.User, error)
}

type srv struct {
	userRepo userProvider
}

func New(userRepo userProvider) srv {
	return srv{userRepo}
}

func (srv srv) ResetPwd(email string, code string, pwd string, tenantID uint) error {
	localizer := localeutil.Get()

	// Check user exists
	userOpts := ctype.QueryOpts{
		Filters: ctype.Dict{"Email": email, "TenantID": tenantID},
	}
	user, err := srv.userRepo.Retrieve(userOpts)
	if err != nil {
		return err
	}

	// Check reset pwd code
	if user.PwdResetToken != code {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.InvalidResetPwdCode,
		})
		return errutil.New("", []string{msg})
	}

	// Update user pwd
	pwdHash := pwdutil.MakePwd(pwd)
	updateData := ctype.Dict{
		"Pwd":           pwdHash,
		"PwdResetToken": "",
		"PwdResetAt":    dateutil.Now(),
	}
	updateOpts := ctype.QueryOpts{Filters: ctype.Dict{"ID": user.ID}}
	_, err = srv.userRepo.Update(updateOpts, updateData)
	if err != nil {
		return err
	}

	return nil
}
