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
var featureSet = "variable"
var featureSetName = "variable"

func RegisterUrls(e *echo.Group, pemMap ctype.PemMap) (*echo.Group, ctype.PemMap) {
	g := e.Group(fmt.Sprintf("/%s/%s", module, featureSet))
	rr := routeutil.RegisterRoute(g, pemMap)

	rr.Private(
		"GET", "/option/", option.WireCtrl().Handler,
	)
	rr.Rbac(
		"GET", "/", list.WireCtrl().Handler,
		[]string{profiletype.ADMIN, profiletype.STAFF},
		fmt.Sprintf("Get %s list", featureSetName),
	)
	rr.Rbac(
		"GET", "/:id", retrieve.WireCtrl().Handler,
		[]string{profiletype.ADMIN, profiletype.STAFF},
		fmt.Sprintf("Get %s detail", featureSetName),
	)
	rr.Rbac(
		"POST", "/", create.WireCtrl().Handler,
		[]string{profiletype.ADMIN},
		fmt.Sprintf("Create %s", featureSetName),
	)
	rr.Rbac(
		"PUT", "/:id", update.WireCtrl().Handler,
		[]string{profiletype.ADMIN},
		fmt.Sprintf("Update %s", featureSetName),
	)
	rr.Rbac(
		"DELETE", "/:id", remove.WireCtrl().Handler,
		[]string{profiletype.ADMIN},
		fmt.Sprintf("Delete %s", featureSetName),
	)
	rr.Rbac(
		"DELETE", "/", removelist.WireCtrl().Handler,
		[]string{profiletype.ADMIN},
		fmt.Sprintf("Delete list %s", featureSetName),
	)
	return e, pemMap
}
