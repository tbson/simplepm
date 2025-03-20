package refreshtokencheck

import (
	"src/module/account/usecase/auth/refreshtokencheck/ctrl"
	"src/util/fwutil"
)

var ctrlHandler fwutil.CtrlHandler

func WireCtrl() fwutil.CtrlHandler {
	if ctrlHandler != nil {
		return ctrlHandler
	}

	ctrlHandler := ctrl.New()
	return ctrlHandler
}
