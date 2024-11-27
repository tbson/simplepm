package infra

import (
	"fmt"
	"src/util/routeutil"

	"src/common/ctype"

	"github.com/labstack/echo/v4"
)

var module = "account"
var useCaseGroup = "signup-tenant"
var useCaseGroupName = "signup tenant"

func RegisterUrls(e *echo.Group, pemMap ctype.PemMap) (*echo.Group, ctype.PemMap) {
	g := e.Group(fmt.Sprintf("/%s/%s", module, useCaseGroup))
	rr := routeutil.RegisterRoute(g, pemMap)

	rr.Public(
		"POST", "/", SignupTenant,
	)
	return e, pemMap
}
