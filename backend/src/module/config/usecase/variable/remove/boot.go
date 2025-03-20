package remove

import (
	"src/module/config/repo/variable"
	"src/module/config/usecase/variable/remove/ctrl"
	"src/util/dbutil"
	"src/util/fwutil"
)

var ctrlHandler fwutil.CtrlHandler

func WireCtrl() fwutil.CtrlHandler {
	if ctrlHandler != nil {
		return ctrlHandler
	}

	dbClient := dbutil.Db(nil)

	variableRepo := variable.New(dbClient)

	ctrlHandler = ctrl.New(variableRepo)

	return ctrlHandler
}
