package tenant

import (
	"fmt"
	"src/util/routeutil"

	"src/common/ctype"

	"github.com/labstack/echo/v4"

	"src/module/account/usecase/tenant/signup"
)

var module = "account"
var featureSet = "tenant"
var featureSetName = "tenant"

func RegisterUrls(e *echo.Group, pemMap ctype.PemMap) (*echo.Group, ctype.PemMap) {
	g := e.Group(fmt.Sprintf("/%s/%s", module, featureSet))
	rr := routeutil.RegisterRoute(g, pemMap)

	rr.Public(
		"POST", "/signup", signup.WireCtrl().Handler,
	)

	return e, pemMap
}
