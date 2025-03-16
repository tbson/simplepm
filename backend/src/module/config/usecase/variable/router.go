package variable

import (
	"fmt"
	"src/util/routeutil"

	"src/common/ctype"
	"src/common/profiletype"

	"github.com/labstack/echo/v4"

	"src/module/config/usecase/variable/create"
	"src/module/config/usecase/variable/list"
	"src/module/config/usecase/variable/option"
	"src/module/config/usecase/variable/remove"
	"src/module/config/usecase/variable/removelist"
	"src/module/config/usecase/variable/retrieve"
	"src/module/config/usecase/variable/update"
)

var module = "config"
var useCaseGroup = "variable"
var useCaseGroupName = "variable"

func RegisterUrls(e *echo.Group, pemMap ctype.PemMap) (*echo.Group, ctype.PemMap) {
	g := e.Group(fmt.Sprintf("/%s/%s", module, useCaseGroup))
	rr := routeutil.RegisterRoute(g, pemMap)

	rr.Private(
		"GET", "/option/", option.WireCtrl().Handler,
	)
	rr.Rbac(
		"GET", "/", list.WireCtrl().Handler,
		[]string{profiletype.ADMIN, profiletype.STAFF},
		fmt.Sprintf("Get %s list", useCaseGroupName),
	)
	rr.Rbac(
		"GET", "/:id", retrieve.WireCtrl().Handler,
		[]string{profiletype.ADMIN, profiletype.STAFF},
		fmt.Sprintf("Get %s detail", useCaseGroupName),
	)
	rr.Rbac(
		"POST", "/", create.WireCtrl().Handler,
		[]string{profiletype.ADMIN},
		fmt.Sprintf("Create %s", useCaseGroupName),
	)
	rr.Rbac(
		"PUT", "/:id", update.WireCtrl().Handler,
		[]string{profiletype.ADMIN},
		fmt.Sprintf("Update %s", useCaseGroupName),
	)
	rr.Rbac(
		"DELETE", "/:id", remove.WireCtrl().Handler,
		[]string{profiletype.ADMIN},
		fmt.Sprintf("Delete %s", useCaseGroupName),
	)
	rr.Rbac(
		"DELETE", "/", removelist.WireCtrl().Handler,
		[]string{profiletype.ADMIN},
		fmt.Sprintf("Delete list %s", useCaseGroupName),
	)
	return e, pemMap
}
