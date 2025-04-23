package app

import (
	"src/common/ctype"
	"src/module/account/schema"
	"src/util/errutil"
	"src/util/i18nmsg"

	"src/module/account/domain/srv/pwdpolicy"

	"src/util/pwdutil"
)

type Service struct {
	userRepo UserRepo
}

func New(userRepo UserRepo) Service {
	return Service{userRepo}
}

func (srv Service) UpdateProfile(userID uint, data ctype.Dict) (*schema.User, error) {
	user, err := srv.userRepo.Retrieve(ctype.QueryOpts{
		Filters: ctype.Dict{"id": userID},
	})
	if err != nil {
		return nil, err
	}

	updateOpts := ctype.QueryOpts{Filters: ctype.Dict{"ID": user.ID}}
	userResult, err := srv.userRepo.Update(updateOpts, data)
	if err != nil {
		return nil, err
	}
	return userResult, nil
}

func (srv Service) ChangePwd(userID uint, data ctype.Dict) (ctype.Dict, error) {
	pwdPolicy := pwdpolicy.New()
	result := ctype.Dict{}
	if data["Pwd"].(string) != data["PwdConfirm"].(string) {
		return result, errutil.New(i18nmsg.PasswordsNotMatch)
	}
	pwd := data["Pwd"].(string)

	err := pwdPolicy.CheckOnCreation(pwd, []string{})
	if err != nil {
		return result, errutil.New(i18nmsg.PasswordPolicyError)
	}

	user, err := srv.userRepo.Retrieve(ctype.QueryOpts{
		Filters: ctype.Dict{"id": userID},
	})
	if err != nil {
		return result, err
	}

	pwdHash := pwdutil.MakePwd(pwd)
	if pwdHash == "" {
		return result, errutil.New(i18nmsg.PasswordHashError)
	}

	queryOpts := ctype.QueryOpts{Filters: ctype.Dict{"ID": user.ID}}
	updateData := ctype.Dict{"Pwd": pwdHash}

	_, err = srv.userRepo.Update(queryOpts, updateData)
	if err != nil {
		return result, err
	}
	return result, nil
}
