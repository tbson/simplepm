package refreshtoken

import (
	"src/module/account/usecase/auth/refreshtoken/ctrl"

	"src/module/account/repo/user"
	"src/module/account/usecase/auth/refreshtoken/srv"

	"src/common/setting"
	"src/module/account/domain/srv/authtoken"
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

	srv := srv.New(userRepo, authTokenSrv)
	ctrlHandler = ctrl.New(srv)

	return ctrlHandler
}
