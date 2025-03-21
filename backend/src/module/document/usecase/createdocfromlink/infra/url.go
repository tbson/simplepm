package infra

import (
	"fmt"
	"src/util/routeutil"

	"src/common/ctype"
	"src/common/profiletype"

	"github.com/labstack/echo/v4"
)

var module = "document"
var featureSet = "doc"
var featureSetName = "create document from link"

func RegisterUrls(e *echo.Group, pemMap ctype.PemMap) (*echo.Group, ctype.PemMap) {
	g := e.Group(fmt.Sprintf("/%s/%s", module, featureSet))
	rr := routeutil.RegisterRoute(g, pemMap)

	rr.Rbac(
		"POST", "/create-doc-from-link/", Create,
		[]string{
			profiletype.USER,
			profiletype.ADMIN,
			profiletype.STAFF,
			profiletype.MANAGER,
		},
		fmt.Sprintf("%s", featureSetName),
	)

	return e, pemMap
}
