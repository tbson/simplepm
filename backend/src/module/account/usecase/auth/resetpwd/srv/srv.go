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
	Retrieve(queryOptions ctype.QueryOptions) (*schema.User, error)
	Update(queryOptions ctype.QueryOptions, data ctype.Dict) (*schema.User, error)
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
	getUserOptions := ctype.QueryOptions{
		Filters: ctype.Dict{"Email": email, "TenantID": tenantID},
	}
	user, err := srv.userRepo.Retrieve(getUserOptions)
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
	updateOptions := ctype.QueryOptions{Filters: ctype.Dict{"ID": user.ID}}
	_, err = srv.userRepo.Update(updateOptions, updateData)
	if err != nil {
		return err
	}

	return nil
}
