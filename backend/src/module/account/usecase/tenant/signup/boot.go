package signup

import (
	"src/module/account/repo/user"
	"src/module/account/usecase/tenant/signup/ctrl"
	"src/module/account/usecase/tenant/signup/srv"
	"src/util/dbutil"
	"src/util/fwutil"

	"src/module/account/repo/role"
	"src/module/account/repo/tenant"
	"src/module/account/srv/auth"
)

var ctrlHandler fwutil.CtrlHandler

func WireCtrl() fwutil.CtrlHandler {
	if ctrlHandler != nil {
		return ctrlHandler
	}

	dbClient := dbutil.Db(nil)

	userRepo := user.New(dbClient)
	tenantRepo := tenant.New(dbClient)
	roleRepo := role.New(dbClient)
	authSrv := auth.New()

	ctrlHandler = ctrl.New(
		srv.New(userRepo, tenantRepo, roleRepo, authSrv),
		dbClient,
	)

	return ctrlHandler
}
