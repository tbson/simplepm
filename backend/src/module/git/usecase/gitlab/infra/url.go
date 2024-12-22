package infra

import (
	"fmt"
	"src/util/routeutil"

	"src/common/ctype"

	"github.com/labstack/echo/v4"
)

var module = "git"
var useCaseGroup = "gitlab"
var useCaseGroupName = "gitlab"

func RegisterUrls(e *echo.Group, pemMap ctype.PemMap) (*echo.Group, ctype.PemMap) {
	g := e.Group(fmt.Sprintf("/%s/%s", module, useCaseGroup))
	rr := routeutil.RegisterRoute(g, pemMap)
	rr.Public(
		"GET", "/callback", Callback,
	)
	rr.Public(
		"POST", "/webhook", Webhook,
	)
	return e, pemMap
}
