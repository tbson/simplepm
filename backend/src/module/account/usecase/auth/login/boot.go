package login

import (
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

	appSrv := srv.New(userRepo)
	ctrlHandler = ctrl.New(appSrv)

	return ctrlHandler
}
