package srv

import (
	"src/common/ctype"
	"src/module/account/schema"
	"src/util/pwdutil"
)

type userProvider interface {
	Update(opts ctype.QueryOpts, data ctype.Dict) (*schema.User, error)
}

type srv struct {
	userRepo userProvider
}

func New(userRepo userProvider) srv {
	return srv{userRepo}
}

func (srv srv) ChangePwd(userID uint, pwd string) error {
	// Update user pwd
	pwdHash := pwdutil.MakePwd(pwd)
	updateData := ctype.Dict{
		"Pwd": pwdHash,
	}
	updateOpts := ctype.QueryOpts{Filters: ctype.Dict{"ID": userID}}
	_, err := srv.userRepo.Update(updateOpts, updateData)
	if err != nil {
		return err
	}

	return nil
}
