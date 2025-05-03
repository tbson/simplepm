package requestresetpwd

import (
	"src/client/emailclient"
	"src/module/account/repo/email"
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

	userRepo := user.New(dbClient)
	emailRepo := email.New(emailClient)

	ctrlHandler = ctrl.New(
		srv.New(userRepo, emailRepo),
	)
	return ctrlHandler
}
