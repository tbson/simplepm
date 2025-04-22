package routeutil

import (
	"fmt"
	"reflect"
	"runtime"
	"src/common/ctype"
	"src/middleware"
	"src/util/fwutil"
	"strings"

	"github.com/labstack/echo/v4"
)

type RuteDefaultHandlerFunc func(
	string,
	string,
	echo.HandlerFunc,
) ctype.PemMap

type RuteRbacHandlerFunc func(
	string,
	string,
	echo.HandlerFunc,
	[]string,
	string,
) ctype.PemMap

type RuteRbacHandlerFuncNew func(
	string,
	string,
	func() fwutil.CtrlHandler,
	[]string,
	string,
) ctype.PemMap

type HandleMap struct {
	Public  RuteDefaultHandlerFunc
	Private RuteDefaultHandlerFunc
	Rbac    RuteRbacHandlerFunc
	RbacNew RuteRbacHandlerFuncNew
}

var pemMap *ctype.PemMap

func getFnPath(fn interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
}

func SetPemMap(pemMapParam *ctype.PemMap) {
	pemMap = pemMapParam
}

func GetPemMap() ctype.PemMap {
	return *pemMap
}

func GetHandlerInfo(ctrl interface{}) (string, string) {
	fnPath := getFnPath(ctrl)
	fmt.Printf("%+v\n", fnPath)
	arrResult := strings.Split(fnPath, ".")
	module := arrResult[0]
	action := arrResult[1]

	arrModule := strings.Split(module, "/")
	module = arrModule[len(arrModule)-2]
	return module, action
}

func GetHandlerInfoNew(ctrl interface{}) (string, string) {
	fnPath := getFnPath(ctrl)
	fmt.Printf("%+v\n", fnPath)
	arrResult := strings.Split(fnPath, ".")
	modulePath := arrResult[0]

	arrModulePath := strings.Split(modulePath, "/")
	action := arrModulePath[len(arrModulePath)-1]
	module := arrModulePath[len(arrModulePath)-2]
	return module, action
}

func RegisterRoute(group *echo.Group, pemMap ctype.PemMap) HandleMap {
	return HandleMap{
		Public: func(
			verb string,
			path string,
			ctrl echo.HandlerFunc,
		) ctype.PemMap {
			verbs := []string{verb}
			group.Match(verbs, path, ctrl)
			return pemMap
		},
		Private: func(
			verb string,
			path string,
			ctrl echo.HandlerFunc,
		) ctype.PemMap {
			verbs := []string{verb}
			module, action := GetHandlerInfo(ctrl)
			group.Match(verbs, path, ctrl, middleware.AuthMiddleware(module, action, false))
			return pemMap
		},
		Rbac: func(
			verb string,
			path string,
			ctrl echo.HandlerFunc,
			profileTypes []string,
			title string,
		) ctype.PemMap {
			verbs := []string{verb}
			module, action := GetHandlerInfo(ctrl)
			key := module + "." + action
			role := ctype.Pem{
				ProfileTypes: profileTypes,
				Title:        title,
				Module:       module,
				Action:       action,
			}
			pemMap[key] = role
			group.Match(verbs, path, ctrl, middleware.AuthMiddleware(module, action, true))
			return pemMap
		},
		RbacNew: func(
			verb string,
			path string,
			ctrl func() fwutil.CtrlHandler,
			profileTypes []string,
			title string,
		) ctype.PemMap {
			verbs := []string{verb}
			module, action := GetHandlerInfoNew(ctrl)
			key := module + "." + action
			role := ctype.Pem{
				ProfileTypes: profileTypes,
				Title:        title,
				Module:       module,
				Action:       action,
			}
			pemMap[key] = role
			group.Match(verbs, path, ctrl().Handler, middleware.AuthMiddleware(module, action, true))
			return pemMap
		},
	}
}
