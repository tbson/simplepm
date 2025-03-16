package auth

import (
	"fmt"
	"src/util/routeutil"

	"src/common/ctype"

	"github.com/labstack/echo/v4"

	"src/module/account/usecase/auth/changepwd"
	"src/module/account/usecase/auth/login"
	"src/module/account/usecase/auth/refreshtoken"
	"src/module/account/usecase/auth/refreshtokencheck"
	"src/module/account/usecase/auth/requestresetpwd"
	"src/module/account/usecase/auth/resetpwd"
)

var module = "account"
var useCaseGroup = "auth"
var useCaseGroupName = "auth"

func RegisterUrls(e *echo.Group, pemMap ctype.PemMap) (*echo.Group, ctype.PemMap) {
	g := e.Group(fmt.Sprintf("/%s/%s", module, useCaseGroup))
	rr := routeutil.RegisterRoute(g, pemMap)

	rr.Public(
		"POST", "/login", login.WireCtrl().Handler,
	)
	rr.Public(
		"PUT", "/change-pwd", changepwd.WireCtrl().Handler,
	)
	rr.Public(
		"POST", "/reset-pwd", resetpwd.WireCtrl().Handler,
	)
	rr.Public(
		"POST", "/request-reset-pwd", requestresetpwd.WireCtrl().Handler,
	)
	rr.Public(
		"POST", "/refresh-token", refreshtoken.WireCtrl().Handler,
	)
	rr.Public(
		"GET", "/refresh-token-check", refreshtokencheck.WireCtrl().Handler,
	)
	return e, pemMap
}
