package app

import (
	"context"
	"src/common/authtype"
	"src/common/ctype"
	"src/module/account/schema"
	"src/util/dateutil"
	"src/util/dictutil"
	"src/util/errutil"
	"src/util/localeutil"
	"src/util/pwdutil"
	"src/util/ssoutil"
	"src/util/stringutil"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Service struct {
	userRepo  UserRepo
	roleRepo  RoleRepo
	iamRepo   IamRepo
	authRepo  AuthRepo
	emailRepo EmailRepo
}

func New(
	userRepo UserRepo,
	roleRepo RoleRepo,
	iamRepo IamRepo,
	authRepo AuthRepo,
	emailRepo EmailRepo,
) Service {
	return Service{userRepo, roleRepo, iamRepo, authRepo, emailRepo}
}

func (srv Service) parseState(state string) (StateData, error) {
	result := StateData{}
	localizer := localeutil.Get()
	stateData, err := ssoutil.DecodeState(state)
	if err != nil {
		return result, err
	}

	tenantUid, uidErr := stateData["tenantUid"].(string)
	next, _ := stateData["next"].(string)
	if !uidErr {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.InvalidState,
		})
		return result, errutil.New("", []string{msg})
	}
	result.TenantUid = tenantUid
	result.Next = next
	return result, nil
}

func (srv Service) GetAuthUrl(tenantUid string, nextParam string) (string, error) {
	state := ctype.Dict{
		"tenantUid": tenantUid,
	}
	if nextParam != "" {
		state["next"] = nextParam
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

func (srv Service) GetLogoutUrl(tenantUid string, idToken string) (string, error) {
	authClientInfo, err := srv.authRepo.GetAuthClientFromTenantUid(tenantUid)
	if err != nil {
		return "", err
	}

	realm := authClientInfo.Realm
	clientId := authClientInfo.ClientID

	url := srv.iamRepo.GetLogoutUrl(realm, clientId, idToken)

	return url, nil
}

func (srv Service) HandleCallback(
	ctx context.Context,
	state string,
	code string,
) (authtype.AuthCallbackResult, error) {
	blankResult := authtype.AuthCallbackResult{}
	stateData, err := srv.parseState(state)
	tenantUid := stateData.TenantUid
	next := stateData.Next
	if err != nil {
		return blankResult, err
	}

	authClientInfo, err := srv.authRepo.GetAuthClientFromTenantUid(tenantUid)
	if err != nil {
		return blankResult, err
	}
	tenantID := authClientInfo.TenantID
	realm := authClientInfo.Realm
	clientId := authClientInfo.ClientID
	clientSecret := authClientInfo.ClientSecret

	tokensAndClaims, err := srv.iamRepo.ValidateCallback(
		ctx, realm, clientId, clientSecret, code,
	)

	if err != nil {
		return blankResult, err
	}

	userInfo := tokensAndClaims.UserInfo
	user, err := srv.authRepo.GetTenantUser(tenantID, userInfo.Email)
	if err != nil {
		roleOptions := ctype.QueryOptions{
			Filters: ctype.Dict{
				"TenantID": tenantID,
				"Title":    "USER",
			},
		}
		role, err := srv.roleRepo.Retrieve(roleOptions)
		if err != nil {
			return blankResult, err
		}

		userData := dictutil.StructToDict(userInfo)
		userData["TenantID"] = tenantID
		userData["Roles"] = []schema.Role{*role}

		_, err = srv.userRepo.Create(userData)
		if err != nil {
			return blankResult, err
		}
		user, _ = srv.authRepo.GetTenantUser(tenantID, userInfo.Email)
	} else {
		localizer := localeutil.Get()
		if user.LockedAt != nil {
			srv.iamRepo.Logout(
				clientId,
				clientSecret,
				realm,
				tokensAndClaims.RefreshToken,
			)
			msg := localizer.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: localeutil.LockedAccount,
			})
			return blankResult, errutil.New("", []string{msg})
		}

		userData := dictutil.StructToDict(userInfo)
		updateOptions := ctype.QueryOptions{Filters: ctype.Dict{"ID": user.ID}}
		_, err = srv.userRepo.Update(updateOptions, userData)
		if err != nil {
			return blankResult, err
		}
	}
	result := authtype.AuthCallbackResult{
		AccessToken:  tokensAndClaims.AccessToken,
		RefreshToken: tokensAndClaims.RefreshToken,
		IDToken:      tokensAndClaims.IDToken,
		Realm:        tokensAndClaims.Realm,
		UserInfo:     user,
		Next:         next,
	}
	return result, nil
}

func (srv Service) RefreshToken(
	ctx context.Context,
	realm string,
	refreshToken string,
) (authtype.SsoCallbackResult, error) {
	sub, err := srv.iamRepo.GetSub(refreshToken, realm)
	if err != nil {
		return authtype.SsoCallbackResult{}, err
	}

	authClientInfo, err := srv.authRepo.GetAuthClientFromSub(sub)
	if err != nil {
		return authtype.SsoCallbackResult{}, err
	}

	clientId := authClientInfo.ClientID
	clientSecret := authClientInfo.ClientSecret

	tokensAndClaims, err := srv.iamRepo.RefreshToken(
		ctx, realm, refreshToken, clientId, clientSecret,
	)
	if err != nil {
		return authtype.SsoCallbackResult{}, err
	}

	return tokensAndClaims, nil
}

func (srv Service) RequestResetPassword(email string, tenantUid string) error {
	// Check user exists
	getUserOptions := ctype.QueryOptions{
		Filters: ctype.Dict{"Email": email, "Tenant.UID": tenantUid},
		Joins:   []string{"Tenant"},
	}
	user, err := srv.userRepo.Retrieve(getUserOptions)
	if err != nil {
		return err
	}

	// Generate reset password token
	code := stringutil.GetRandomString(6)

	// update user reset password token
	updateOptions := ctype.QueryOptions{Filters: ctype.Dict{"ID": user.ID}}
	updateData := ctype.Dict{
		"PwdResetToken": code,
	}
	_, err = srv.userRepo.Update(updateOptions, updateData)
	if err != nil {
		return err
	}

	// Send email containing reset password token
	to := user.Email
	subject := "Reset Password"
	body := ctype.EmailBody{
		HmtlPath: "emails/reset-pwd.html",
		Data: ctype.Dict{
			"Name": user.FullName(),
			"Code": code,
		},
	}
	srv.emailRepo.SendEmailAsync(to, subject, body)
	return nil
}

func (srv Service) ResetPassword(
	email string,
	code string,
	password string,
	tenantUid string,
) error {
	localizer := localeutil.Get()

	// Check user exists
	getUserOptions := ctype.QueryOptions{
		Filters: ctype.Dict{"Email": email, "Tenant.UID": tenantUid},
		Joins:   []string{"Tenant"},
	}
	user, err := srv.userRepo.Retrieve(getUserOptions)
	if err != nil {
		return err
	}

	// Check reset password code
	if user.PwdResetToken != code {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.InvalidResetPwdCode,
		})
		return errutil.New("", []string{msg})
	}

	// Update user password
	pwdHash, err := pwdutil.MakePwd(password)
	if err != nil {
		return err
	}
	updateData := ctype.Dict{
		"Password":      pwdHash,
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
