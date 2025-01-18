package route

import (
	"src/common/ctype"
	auth "src/module/account/usecase/auth/infra"
	crudauthclient "src/module/account/usecase/crudauthclient/infra"
	crudrole "src/module/account/usecase/crudrole/infra"
	crudtenant "src/module/account/usecase/crudtenant/infra"
	cruduser "src/module/account/usecase/cruduser/infra"
	lockuser "src/module/account/usecase/lockuser/infra"
	profile "src/module/account/usecase/profile/infra"
	signuptenant "src/module/account/usecase/signuptenant/infra"
	crudvariable "src/module/config/usecase/crudvariable/infra"
	configoption "src/module/config/usecase/option/infra"
	github "src/module/git/usecase/github/infra"
	gitlab "src/module/git/usecase/gitlab/infra"
	crudfeature "src/module/pm/usecase/crudfeature/infra"
	crudproject "src/module/pm/usecase/crudproject/infra"
	crudtask "src/module/pm/usecase/crudtask/infra"
	crudtaskfield "src/module/pm/usecase/crudtaskfield/infra"
	crudworkspace "src/module/pm/usecase/crudworkspace/infra"
	reorderfeature "src/module/pm/usecase/reorderfeature/infra"
	reorderstatus "src/module/pm/usecase/reorderstatus/infra"
	reordertask "src/module/pm/usecase/reordertask/infra"
	reordertaskfield "src/module/pm/usecase/reordertaskfield/infra"
	getauthjwt "src/module/socket/usecase/getauthjwt/infra"
	publishmessage "src/module/socket/usecase/publishmessage/infra"

	"github.com/labstack/echo/v4"
)

func CollectRoutes(e *echo.Group) (*echo.Group, ctype.PemMap) {
	pemMap := ctype.PemMap{}
	e, pemMap = configoption.RegisterUrls(e, pemMap)
	e, pemMap = crudvariable.RegisterUrls(e, pemMap)
	e, pemMap = crudauthclient.RegisterUrls(e, pemMap)
	e, pemMap = crudtenant.RegisterUrls(e, pemMap)
	e, pemMap = crudrole.RegisterUrls(e, pemMap)
	e, pemMap = cruduser.RegisterUrls(e, pemMap)
	e, pemMap = crudworkspace.RegisterUrls(e, pemMap)
	e, pemMap = crudproject.RegisterUrls(e, pemMap)
	e, pemMap = crudtaskfield.RegisterUrls(e, pemMap)
	e, pemMap = auth.RegisterUrls(e, pemMap)
	e, pemMap = profile.RegisterUrls(e, pemMap)
	e, pemMap = lockuser.RegisterUrls(e, pemMap)
	e, pemMap = signuptenant.RegisterUrls(e, pemMap)
	e, pemMap = reordertaskfield.RegisterUrls(e, pemMap)
	e, pemMap = crudtask.RegisterUrls(e, pemMap)
	e, pemMap = crudfeature.RegisterUrls(e, pemMap)
	e, pemMap = reorderfeature.RegisterUrls(e, pemMap)
	e, pemMap = reordertask.RegisterUrls(e, pemMap)
	e, pemMap = reorderstatus.RegisterUrls(e, pemMap)
	e, pemMap = getauthjwt.RegisterUrls(e, pemMap)
	e, pemMap = publishmessage.RegisterUrls(e, pemMap)
	e, pemMap = github.RegisterUrls(e, pemMap)
	e, pemMap = gitlab.RegisterUrls(e, pemMap)
	return e, pemMap
}
