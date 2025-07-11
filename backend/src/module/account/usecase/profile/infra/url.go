package infra

import (
	"fmt"
	"src/util/routeutil"

	"src/common/ctype"

	"github.com/labstack/echo/v4"
)

var module = "account"
var featureSet = "profile"
var featureSetName = "profile"

func RegisterUrls(e *echo.Group, pemMap ctype.PemMap) (*echo.Group, ctype.PemMap) {
	g := e.Group(fmt.Sprintf("/%s/%s", module, featureSet))
	rr := routeutil.RegisterRoute(g, pemMap)

	rr.Private(
		"GET", "/", GetProfile,
	)
	rr.Private(
		"PUT", "/", UpdateProfile,
	)
	rr.Private(
		"PUT", "/password/", ChangePassword,
	)
	return e, pemMap
}
