package refreshtoken

import (
	"src/module/account/usecase/auth/refreshtoken/ctrl"

	"src/module/account/repo/user"
	"src/module/account/usecase/auth/refreshtoken/srv"

	"src/module/account/srv/auth"
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
	authSrv := auth.New()

	ctrlHandler = ctrl.New(
		srv.New(authSrv, userRepo),
	)

	return ctrlHandler
}
