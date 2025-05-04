package requestresetpwd

import (
	"src/adapter/email"
	"src/module/account/repo/user"
	"src/module/account/usecase/auth/requestresetpwd/ctrl"
	"src/module/account/usecase/auth/requestresetpwd/srv"
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
	emailAdapter := email.New()

	ctrlSrv := srv.New(
		userRepo,
		emailAdapter,
	)

	ctrlHandler = ctrl.New(ctrlSrv)
	return ctrlHandler
}
