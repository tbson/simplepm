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
	iamRepo  IamRepo
}

func New(userRepo UserRepo, iamRepo IamRepo) Service {
	return Service{userRepo, iamRepo}
}

func (srv Service) UpdateProfile(userID uint, data ctype.Dict) (*schema.User, error) {
	user, err := srv.userRepo.Retrieve(ctype.QueryOptions{
		Filters:  ctype.Dict{"id": userID},
		Preloads: []string{"Tenant.AuthClient"},
	})
	if err != nil {
		return nil, err
	}
	sub := user.Sub
	realm := user.Tenant.AuthClient.Partition

	accessToken, err := srv.iamRepo.GetAdminAccessToken()
	if err != nil {
		return nil, err
	}

	err = srv.iamRepo.UpdateUser(accessToken, realm, *sub, data)
	if err != nil {
		return nil, err
	}

	userResult, err := srv.userRepo.Update(userID, data)
	if err != nil {
		return nil, err
	}
	return userResult, nil
}

func (srv Service) ChangePassword(userID uint, data ctype.Dict) (ctype.Dict, error) {
	result := ctype.Dict{}
	localizer := localeutil.Get()
	if data["Password"].(string) != data["PasswordConfirm"].(string) {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.PasswordsNotMatch,
		})
		return result, errutil.New("", []string{msg})
	}
	password := data["Password"].(string)

	user, err := srv.userRepo.Retrieve(ctype.QueryOptions{
		Filters:  ctype.Dict{"id": userID},
		Preloads: []string{"Tenant.AuthClient"},
	})
	if err != nil {
		return result, err
	}
	sub := user.Sub
	realm := user.Tenant.AuthClient.Partition

	accessToken, err := srv.iamRepo.GetAdminAccessToken()
	if err != nil {
		return result, err
	}

	err = srv.iamRepo.SetPassword(accessToken, *sub, realm, password)
	if err != nil {
		return result, err
	}
	return result, nil
}
