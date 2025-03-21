package infra

import (
	"fmt"
	"src/util/routeutil"

	"src/common/ctype"

	"github.com/labstack/echo/v4"
)

var module = "git"
var featureSet = "gitlab"
var featureSetName = "gitlab"

func RegisterUrls(e *echo.Group, pemMap ctype.PemMap) (*echo.Group, ctype.PemMap) {
	g := e.Group(fmt.Sprintf("/%s/%s", module, featureSet))
	rr := routeutil.RegisterRoute(g, pemMap)
	rr.Public(
		"GET", "/callback", Callback,
	)
	rr.Public(
		"POST", "/webhook", Webhook,
	)
	return e, pemMap
}
