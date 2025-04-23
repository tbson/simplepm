package srv

import (
	"src/common/ctype"
	"src/module/account/schema"
	"src/util/dateutil"
	"src/util/errutil"
	"src/util/localeutil"
	"src/util/pwdutil"
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
		return errutil.New(localeutil.InvalidResetPwdCode)
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
