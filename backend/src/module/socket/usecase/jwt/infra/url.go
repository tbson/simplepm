package infra

import (
	"fmt"
	"src/util/routeutil"

	"src/common/ctype"

	"github.com/labstack/echo/v4"
)

var module = "socket"
var featureSet = "jwt"
var featureSetName = "JWT"

func RegisterUrls(e *echo.Group, pemMap ctype.PemMap) (*echo.Group, ctype.PemMap) {
	g := e.Group(fmt.Sprintf("/%s/%s", module, featureSet))
	rr := routeutil.RegisterRoute(g, pemMap)

	rr.Private(
		"GET", "/auth/", GetAuthJWT,
	)
	rr.Private(
		"GET", "/subscription/", GetSubscriptionJWT,
	)
	return e, pemMap
}
