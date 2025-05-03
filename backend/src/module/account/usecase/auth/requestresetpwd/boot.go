package requestresetpwd

import (
	"src/client/emailclient"
	"src/module/account/extsrv/email"
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
	emailClient, err := emailclient.NewClient()
	if err != nil {
		panic(err)
	}

	ctrlSrv := srv.New(
		user.New(dbClient),
		email.New(emailClient),
	)

	ctrlHandler = ctrl.New(ctrlSrv)
	return ctrlHandler
}
