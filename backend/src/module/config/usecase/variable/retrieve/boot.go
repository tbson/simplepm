package retrieve

import (
	"src/module/config/repo/variable"
	"src/module/config/usecase/variable/retrieve/ctrl"
	"src/util/dbutil"
	"src/util/frameworkutil"
)

var ctrlHandler frameworkutil.CtrlHandler

func WireCtrl() frameworkutil.CtrlHandler {
	if ctrlHandler != nil {
		return ctrlHandler
	}

	dbClient := dbutil.Db(nil)

	variableRepo := variable.New(dbClient)

	ctrlHandler = ctrl.New(variableRepo)

	return ctrlHandler
}
