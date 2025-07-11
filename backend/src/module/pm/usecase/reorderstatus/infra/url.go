package infra

import (
	"fmt"
	"src/util/routeutil"

	"src/common/ctype"
	"src/common/profiletype"

	"github.com/labstack/echo/v4"
)

var module = "pm"
var featureSet = "task"
var featureSetName = "status"

func RegisterUrls(e *echo.Group, pemMap ctype.PemMap) (*echo.Group, ctype.PemMap) {
	g := e.Group(fmt.Sprintf("/%s/%s", module, featureSet))
	rr := routeutil.RegisterRoute(g, pemMap)

	rr.Rbac(
		"PUT", "/reorder-status/", Reorder,
		[]string{profiletype.ADMIN, profiletype.STAFF, profiletype.MANAGER},
		fmt.Sprintf("Reorder %s item", featureSetName),
	)
	return e, pemMap
}
