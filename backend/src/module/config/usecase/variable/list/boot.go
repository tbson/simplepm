package list

import (
	"src/module/abstract/repo/page"
	"src/module/config/schema"
	"src/module/config/usecase/variable/list/ctrl"
	"src/util/dbutil"
	"src/util/frameworkutil"
)

var ctrlHandler frameworkutil.CtrlHandler

func WireCtrl() frameworkutil.CtrlHandler {
	if ctrlHandler != nil {
		return ctrlHandler
	}

	dbClient := dbutil.Db(nil)

	pageRepo := page.New[schema.Variable](dbClient)

	ctrlHandler = ctrl.New(pageRepo)

	return ctrlHandler
}
