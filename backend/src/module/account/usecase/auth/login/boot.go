package login

import (
	"src/module/account/repo/user"
	"src/module/account/usecase/auth/login/ctrl"
	"src/module/account/usecase/auth/login/srv"
	"src/util/dbutil"
	"src/util/frameworkutil"
)

var ctrlHandler frameworkutil.CtrlHandler

func WireCtrl() frameworkutil.CtrlHandler {
	if ctrlHandler != nil {
		return ctrlHandler
	}

	dbClient := dbutil.Db(nil)

	userRepo := user.New(dbClient)

	ctrlHandler = ctrl.New(
		srv.New(userRepo),
	)
	return ctrlHandler
}
