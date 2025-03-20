package option

import (
	"src/module/config/usecase/variable/option/ctrl"
	"src/util/fwutil"
)

var ctrlHandler fwutil.CtrlHandler

func WireCtrl() fwutil.CtrlHandler {
	if ctrlHandler != nil {
		return ctrlHandler
	}

	ctrlHandler = ctrl.New()
	return ctrlHandler
}
