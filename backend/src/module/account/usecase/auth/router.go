package auth

import (
	"fmt"
	"src/util/routeutil"

	"src/common/ctype"

	"github.com/labstack/echo/v4"

	"src/module/account/usecase/auth/changepwd"
	"src/module/account/usecase/auth/checkrefreshtoken"
	"src/module/account/usecase/auth/login"
	"src/module/account/usecase/auth/logout"
	"src/module/account/usecase/auth/refreshtoken"
	"src/module/account/usecase/auth/requestresetpwd"
	"src/module/account/usecase/auth/resetpwd"
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
		"POST", "/logout/", logout.WireCtrl().Handler,
	)
	rr.Public(
		"PUT", "/change-pwd", changepwd.WireCtrl().Handler,
	)
	rr.Private(
		"POST", "/reset-pwd/request", requestresetpwd.WireCtrl().Handler,
	)
	rr.Public(
		"POST", "/reset-pwd/apply", resetpwd.WireCtrl().Handler,
	)
	rr.Private(
		"GET", "/refresh-token/check", checkrefreshtoken.WireCtrl().Handler,
	)
	rr.Public(
		"POST", "/refresh-token", refreshtoken.WireCtrl().Handler,
	)
	return e, pemMap
}
