package changepwd

import (
	"src/module/account/repo/user"
	"src/module/account/usecase/auth/changepwd/ctrl"
	"src/module/account/usecase/auth/changepwd/srv"
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

	ctrlHandler = ctrl.New(
		srv.New(userRepo),
	)
	return ctrlHandler
}
