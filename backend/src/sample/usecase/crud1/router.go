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
	rr.RbacNew(
		"GET", "/", list.WireCtrl,
		[]string{profiletype.ADMIN, profiletype.STAFF},
		fmt.Sprintf("Get %s list", featureSetName),
	)
	rr.RbacNew(
		"GET", "/:id", retrieve.WireCtrl,
		[]string{profiletype.ADMIN, profiletype.STAFF},
		fmt.Sprintf("Get %s detail", featureSetName),
	)
	rr.RbacNew(
		"POST", "/", create.WireCtrl,
		[]string{profiletype.ADMIN},
		fmt.Sprintf("Create %s", featureSetName),
	)
	rr.RbacNew(
		"PUT", "/:id", update.WireCtrl,
		[]string{profiletype.ADMIN},
		fmt.Sprintf("Update %s", featureSetName),
	)
	rr.RbacNew(
		"DELETE", "/:id", remove.WireCtrl,
		[]string{profiletype.ADMIN},
		fmt.Sprintf("Delete %s", featureSetName),
	)
	rr.RbacNew(
		"DELETE", "/", removelist.WireCtrl,
		[]string{profiletype.ADMIN},
		fmt.Sprintf("Delete list %s", featureSetName),
	)
	return e, pemMap
}
