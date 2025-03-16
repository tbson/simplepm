package refreshtokencheck

import (
	"src/module/account/usecase/auth/refreshtokencheck/ctrl"
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
