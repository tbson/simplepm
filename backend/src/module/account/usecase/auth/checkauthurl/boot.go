package checkauthurl

import (
	"src/module/account/repo/iam"
	"src/module/account/usecase/auth/checkauthurl/ctrl"
	"src/module/account/usecase/auth/checkauthurl/repo"
	"src/module/account/usecase/auth/checkauthurl/srv"
	"src/util/dbutil"
	"src/util/frameworkutil"
	"src/util/ssoutil"
)

var ctrlHandler frameworkutil.CtrlHandler

func WireCtrl() frameworkutil.CtrlHandler {
	if ctrlHandler != nil {
		return ctrlHandler
	}

	dbClient := dbutil.Db(nil)
	ssoClient := ssoutil.Client()

	localDataRepo := repo.New(dbClient)
	iamRepo := iam.New(ssoClient)

	ctrlHandler = ctrl.New(
		srv.New(localDataRepo, iamRepo),
	)
	return ctrlHandler
}
