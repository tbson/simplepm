package auth

import (
	"fmt"
	"src/util/routeutil"

	"src/common/ctype"

	"github.com/labstack/echo/v4"

	"src/module/account/usecase/auth/changepwd"
	"src/module/account/usecase/auth/login"
	refreshtokenapply "src/module/account/usecase/auth/refreshtoken/apply"
	refreshtokencheck "src/module/account/usecase/auth/refreshtoken/check"
	resetpwdapply "src/module/account/usecase/auth/resetpwd/apply"
	resetpwdrequest "src/module/account/usecase/auth/resetpwd/request"
)

var module = "account"
var featureSet = "auth"
var featureSetName = "auth"

func RegisterUrls(e *echo.Group, pemMap ctype.PemMap) (*echo.Group, ctype.PemMap) {
	g := e.Group(fmt.Sprintf("/%s/%s", module, featureSet))
	rr := routeutil.RegisterRoute(g, pemMap)

	rr.Public(
		"POST", "/login/", login.WireCtrl().Handler,
	)
	rr.Public(
		"PUT", "/change-pwd", changepwd.WireCtrl().Handler,
	)
	rr.Public(
		"POST", "/reset-pwd/request", resetpwdrequest.WireCtrl().Handler,
	)
	rr.Public(
		"POST", "/reset-pwd/apply", resetpwdapply.WireCtrl().Handler,
	)
	rr.Public(
		"GET", "/refresh-token/check", refreshtokencheck.WireCtrl().Handler,
	)
	rr.Public(
		"POST", "/refresh-token/apply", refreshtokenapply.WireCtrl().Handler,
	)
	return e, pemMap
}
