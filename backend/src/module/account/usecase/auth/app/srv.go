package app

import (
	"context"
	"src/common/ctype"
	"src/util/dictutil"
	"src/util/errutil"
	"src/util/localeutil"
	"src/util/ssoutil"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Service struct {
	userRepo UserRepo
	iamRepo  IamRepo
	authRepo AuthRepo
}

func New(userRepo UserRepo, iamRepo IamRepo, authRepo AuthRepo) Service {
	return Service{userRepo, iamRepo, authRepo}
}

func (srv Service) parseTenantUidFromState(state string) (string, error) {
	localizer := localeutil.Get()
	stateData, err := ssoutil.DecodeState(state)
	if err != nil {
		return "", err
	}

	tenantUid, ok := stateData["tenantUid"].(string)
	if !ok {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.InvalidState,
		})
		return "", errutil.New("", []string{msg})
	}
	return tenantUid, nil
}

func (srv Service) GetAuthUrl(tenantUid string) (string, error) {
	state := ctype.Dict{
		"tenantUid": tenantUid,
	}
	authClientInfo, err := srv.authRepo.GetAuthClientFromTenantUid(tenantUid)
	if err != nil {
		return "", err
	}

	realm := authClientInfo.Realm
	clientId := authClientInfo.ClientID

	url := srv.iamRepo.GetAuthUrl(realm, clientId, state)

	return url, nil
}

func (srv Service) GetLogoutUrl(tenantUid string) (string, error) {
	authClientInfo, err := srv.authRepo.GetAuthClientFromTenantUid(tenantUid)
	if err != nil {
		return "", err
	}

	realm := authClientInfo.Realm
	clientId := authClientInfo.ClientID

	url := srv.iamRepo.GetLogoutUrl(realm, clientId)

	return url, nil
}

func (srv Service) HandleCallback(
	ctx context.Context,
	state string,
	code string,
) (ssoutil.TokensAndClaims, error) {
	tenantUid, err := srv.parseTenantUidFromState(state)
	if err != nil {
		return ssoutil.TokensAndClaims{}, err
	}

	authClientInfo, err := srv.authRepo.GetAuthClientFromTenantUid(tenantUid)
	if err != nil {
		return ssoutil.TokensAndClaims{}, err
	}
	tenantID := authClientInfo.TenantID
	realm := authClientInfo.Realm
	clientId := authClientInfo.ClientID
	clientSecret := authClientInfo.ClientSecret

	tokensAndClaims, err := srv.iamRepo.ValidateCallback(
		ctx, realm, clientId, clientSecret, code,
	)

	if err != nil {
		return ssoutil.TokensAndClaims{}, err
	}

	userInfo := tokensAndClaims.UserInfo
	user, err := srv.authRepo.GetTenantUser(tenantID, userInfo.Email)
	if err != nil {
		userData := dictutil.StructToDict(userInfo)
		delete(userData, "ID")
		delete(userData, "ProfileType")
		delete(userData, "TenantUid")
		userData["TenantID"] = tenantID
		newUser, err := srv.userRepo.Create(userData)
		user = AuthUserResult{
			ID:    newUser.ID,
			Admin: newUser.Admin,
		}
		if err != nil {
			return tokensAndClaims, err
		}
	} else {
		localizer := localeutil.Get()
		if user.LockedAt != nil {
			srv.iamRepo.Logout(clientId, clientSecret, realm, tokensAndClaims.RefreshToken)
			msg := localizer.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: localeutil.LockedAccount,
			})
			return tokensAndClaims, errutil.New("", []string{msg})
		}

		userData := dictutil.StructToDict(userInfo)
		delete(userData, "ID")
		delete(userData, "ProfileType")
		delete(userData, "TenantUid")
		userData["TenantID"] = tenantID
		_, err = srv.userRepo.Update(user.ID, userData)
		if err != nil {
			return tokensAndClaims, err
		}
	}

	if user.Admin {
		tokensAndClaims.UserInfo.ProfileType = "admin"
	} else {
		tokensAndClaims.UserInfo.ProfileType = "user"
	}
	tokensAndClaims.UserInfo.TenantUid = tenantUid
	tokensAndClaims.UserInfo.ID = user.ID
	return tokensAndClaims, nil
}

func (srv Service) RefreshToken(
	ctx context.Context,
	realm string,
	refreshToken string,
) (ssoutil.TokensAndClaims, error) {
	sub, err := srv.iamRepo.GetSub(refreshToken, realm)
	if err != nil {
		return ssoutil.TokensAndClaims{}, err
	}

	authClientInfo, err := srv.authRepo.GetAuthClientFromSub(sub)
	if err != nil {
		return ssoutil.TokensAndClaims{}, err
	}

	clientId := authClientInfo.ClientID
	clientSecret := authClientInfo.ClientSecret

	tokensAndClaims, err := srv.iamRepo.RefreshToken(
		ctx, realm, refreshToken, clientId, clientSecret,
	)
	if err != nil {
		return ssoutil.TokensAndClaims{}, err
	}

	return tokensAndClaims, nil
}
