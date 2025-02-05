package infra

import (
	"fmt"
	"src/util/routeutil"

	"src/common/ctype"

	"github.com/labstack/echo/v4"
)

var module = "event"
var useCaseGroup = "message"
var useCaseGroupName = "message"

func RegisterUrls(e *echo.Group, pemMap ctype.PemMap) (*echo.Group, ctype.PemMap) {
	g := e.Group(fmt.Sprintf("/%s/%s", module, useCaseGroup))
	rr := routeutil.RegisterRoute(g, pemMap)

	rr.Private(
		"GET", "/", List,
	)
	rr.Private(
		"POST", "/", Create,
	)
	return e, pemMap
}
