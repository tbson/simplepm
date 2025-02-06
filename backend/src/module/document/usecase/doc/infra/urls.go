package infra

import (
	"fmt"
	"src/util/routeutil"

	"src/common/ctype"
	"src/common/profiletype"

	"github.com/labstack/echo/v4"
)

var module = "document"
var useCaseGroup = "doc"
var useCaseGroupName = "doc"

func RegisterUrls(e *echo.Group, pemMap ctype.PemMap) (*echo.Group, ctype.PemMap) {
	g := e.Group(fmt.Sprintf("/%s/%s", module, useCaseGroup))
	rr := routeutil.RegisterRoute(g, pemMap)

	rr.Rbac(
		"GET", "/", List,
		[]string{
			profiletype.USER,
			profiletype.ADMIN,
			profiletype.STAFF,
			profiletype.MANAGER,
		},
		fmt.Sprintf("Get %s list", useCaseGroupName),
	)
	rr.Rbac(
		"GET", "/:id", Retrieve,
		[]string{
			profiletype.USER,
			profiletype.ADMIN,
			profiletype.STAFF,
			profiletype.MANAGER,
		},
		fmt.Sprintf("Get %s detail", useCaseGroupName),
	)
	rr.Rbac(
		"POST", "/", Create,
		[]string{
			profiletype.USER,
			profiletype.ADMIN,
			profiletype.STAFF,
			profiletype.MANAGER,
		},
		fmt.Sprintf("Create %s", useCaseGroupName),
	)
	rr.Rbac(
		"PUT", "/:id", Update,
		[]string{
			profiletype.USER,
			profiletype.ADMIN,
			profiletype.STAFF,
			profiletype.MANAGER,
		},
		fmt.Sprintf("Update %s", useCaseGroupName),
	)
	rr.Rbac(
		"DELETE", "/:id", Delete,
		[]string{
			profiletype.USER,
			profiletype.ADMIN,
			profiletype.STAFF,
			profiletype.MANAGER,
		},
		fmt.Sprintf("Delete %s", useCaseGroupName),
	)
	return e, pemMap
}
