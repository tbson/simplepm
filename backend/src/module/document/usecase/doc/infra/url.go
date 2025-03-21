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
var featureSetName = "doc"

func RegisterUrls(e *echo.Group, pemMap ctype.PemMap) (*echo.Group, ctype.PemMap) {
	g := e.Group(fmt.Sprintf("/%s/%s", module, featureSet))
	rr := routeutil.RegisterRoute(g, pemMap)

	rr.Rbac(
		"GET", "/", List,
		[]string{
			profiletype.USER,
			profiletype.ADMIN,
			profiletype.STAFF,
			profiletype.MANAGER,
		},
		fmt.Sprintf("Get %s list", featureSetName),
	)
	rr.Rbac(
		"GET", "/:id", Retrieve,
		[]string{
			profiletype.USER,
			profiletype.ADMIN,
			profiletype.STAFF,
			profiletype.MANAGER,
		},
		fmt.Sprintf("Get %s detail", featureSetName),
	)
	rr.Rbac(
		"POST", "/", Create,
		[]string{
			profiletype.USER,
			profiletype.ADMIN,
			profiletype.STAFF,
			profiletype.MANAGER,
		},
		fmt.Sprintf("Create %s", featureSetName),
	)
	rr.Rbac(
		"PUT", "/:id", Update,
		[]string{
			profiletype.USER,
			profiletype.ADMIN,
			profiletype.STAFF,
			profiletype.MANAGER,
		},
		fmt.Sprintf("Update %s", featureSetName),
	)
	rr.Rbac(
		"DELETE", "/:id", Delete,
		[]string{
			profiletype.USER,
			profiletype.ADMIN,
			profiletype.STAFF,
			profiletype.MANAGER,
		},
		fmt.Sprintf("Delete %s", featureSetName),
	)
	return e, pemMap
}
