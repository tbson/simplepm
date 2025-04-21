package login

import (
	"src/common/setting"
	"src/module/account/domain/srv/authtoken"
	"src/module/account/domain/srv/pwdpolicy"
	"src/module/account/repo/user"
	"src/module/account/usecase/auth/login/ctrl"
	"src/module/account/usecase/auth/login/srv"
	"src/util/dbutil"
	"src/util/fwutil"
)

var ctrlHandler fwutil.CtrlHandler

func WireCtrl() fwutil.CtrlHandler {
	if ctrlHandler != nil {
		return ctrlHandler
	}

	dbClient := dbutil.Db(nil)

	userRepo := user.New(dbClient)

	tokenSettings := setting.AUTH_TOKEN_SETTINGS()
	authTokenSrv := authtoken.New(
		tokenSettings.AccessTokenSecret,
		tokenSettings.RefreshTokenSecret,
		tokenSettings.AccessTokenLifetime,
		tokenSettings.RefreshTokenLifetime,
	)

	pwdPolicy := pwdpolicy.New()

	appSrv := srv.New(userRepo, authTokenSrv, pwdPolicy)
	ctrlHandler = ctrl.New(appSrv)

	return ctrlHandler
}
