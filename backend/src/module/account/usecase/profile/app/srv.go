package app

import (
	"src/common/ctype"
	"src/module/account/schema"
	"src/util/errutil"
	"src/util/localeutil"

	"src/module/account/domain/srv/pwdpolicy"

	"src/util/pwdutil"

	"github.com/nicksnyder/go-i18n/v2/i18n"
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
	localizer := localeutil.Get()
	if data["Pwd"].(string) != data["PwdConfirm"].(string) {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.PasswordsNotMatch,
		})
		return result, errutil.New("", []string{msg})
	}
	pwd := data["Pwd"].(string)

	err := pwdPolicy.CheckOnCreation(pwd, []string{})
	if err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.PasswordPolicyError,
		})
		return result, errutil.New("", []string{msg})
	}

	user, err := srv.userRepo.Retrieve(ctype.QueryOpts{
		Filters: ctype.Dict{"id": userID},
	})
	if err != nil {
		return result, err
	}

	pwdHash := pwdutil.MakePwd(pwd)
	if pwdHash == "" {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.PasswordHashError,
		})
		return result, errutil.New("", []string{msg})
	}

	queryOpts := ctype.QueryOpts{Filters: ctype.Dict{"ID": user.ID}}
	updateData := ctype.Dict{"Pwd": pwdHash}

	_, err = srv.userRepo.Update(queryOpts, updateData)
	if err != nil {
		return result, err
	}
	return result, nil
}
