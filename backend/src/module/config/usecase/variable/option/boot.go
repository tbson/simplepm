package option

import (
	"src/module/config/usecase/variable/option/ctrl"
	"src/util/frameworkutil"
)

var ctrlHandler frameworkutil.CtrlHandler

func WireCtrl() frameworkutil.CtrlHandler {
	if ctrlHandler != nil {
		return ctrlHandler
	}

	ctrlHandler = ctrl.New()
	return ctrlHandler
}
