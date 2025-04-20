package apply

import (
	"src/module/account/usecase/auth/refreshtoken/apply/ctrl"

	"src/module/account/repo/user"
	"src/module/account/usecase/auth/refreshtoken/apply/srv"

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

	srv := srv.New(userRepo)
	ctrlHandler = ctrl.New(srv)

	return ctrlHandler
}
