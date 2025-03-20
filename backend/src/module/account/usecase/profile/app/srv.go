package app

import (
	"src/common/ctype"
	"src/module/account/schema"
	"src/util/errutil"
	"src/util/localeutil"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Service struct {
	userRepo UserRepo
	authSrv  AuthSrv
}

func New(userRepo UserRepo, authSrv AuthSrv) Service {
	return Service{userRepo, authSrv}
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
	result := ctype.Dict{}
	localizer := localeutil.Get()
	if data["Pwd"].(string) != data["PwdConfirm"].(string) {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.PasswordsNotMatch,
		})
		return result, errutil.New("", []string{msg})
	}
	pwd := data["Pwd"].(string)

	user, err := srv.userRepo.Retrieve(ctype.QueryOpts{
		Filters: ctype.Dict{"id": userID},
	})
	if err != nil {
		return result, err
	}

	err = srv.authSrv.SetPwd(user.ID, pwd)
	if err != nil {
		return result, err
	}
	return result, nil
}
