package list

import (
	"src/module/abstract/repo/page"
	"src/module/config/repo/variable"
	"src/module/config/schema"
	"src/module/config/usecase/variable/crud/ctrl"
	"src/util/dbutil"
	"src/util/frameworkutil"
)

type ctrlMap struct {
	List       frameworkutil.CtrlHandler
	Retrieve   frameworkutil.CtrlHandler
	Create     frameworkutil.CtrlHandler
	Update     frameworkutil.CtrlHandler
	Delete     frameworkutil.CtrlHandler
	DeleteList frameworkutil.CtrlHandler
}

var ctrls *ctrlMap

func WireCtrl() ctrlMap {
	if ctrls != nil {
		return *ctrls
	}

	dbClient := dbutil.Db(nil)

	pageRepo := page.New[schema.Variable](dbClient)
	variableRepo := variable.New(dbClient)

	ctrls = &ctrlMap{
		List:       ctrl.NewList[schema.Variable](pageRepo),
		Retrieve:   ctrl.NewRetrieve[schema.Variable](variableRepo),
		Create:     ctrl.NewCreate[schema.Variable](variableRepo),
		Update:     ctrl.NewUpdate[schema.Variable](variableRepo),
		Delete:     ctrl.NewDelete[schema.Variable](variableRepo),
		DeleteList: ctrl.NewDeleteList[schema.Variable](variableRepo),
	}

	return *ctrls
}
