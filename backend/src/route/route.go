package route

import (
	"src/common/ctype"
	auth "src/module/account/usecase/auth"
	authclient "src/module/account/usecase/authclient/infra"
	lockuser "src/module/account/usecase/lockuser/infra"
	profile "src/module/account/usecase/profile/infra"
	role "src/module/account/usecase/role/infra"
	signuptenant "src/module/account/usecase/signuptenant/infra"
	tenant "src/module/account/usecase/tenant/infra"
	user "src/module/account/usecase/user/infra"
	variable "src/module/config/usecase/variable/infra"
	createdocfromlink "src/module/document/usecase/createdocfromlink/infra"
	doc "src/module/document/usecase/doc/infra"
	docattachment "src/module/document/usecase/docattachment/infra"
	message "src/module/event/usecase/message/infra"
	github "src/module/git/usecase/github/infra"
	gitlab "src/module/git/usecase/gitlab/infra"
	feature "src/module/pm/usecase/feature/infra"
	project "src/module/pm/usecase/project/infra"
	reorderfeature "src/module/pm/usecase/reorderfeature/infra"
	reorderstatus "src/module/pm/usecase/reorderstatus/infra"
	reordertask "src/module/pm/usecase/reordertask/infra"
	reordertaskfield "src/module/pm/usecase/reordertaskfield/infra"
	task "src/module/pm/usecase/task/infra"
	taskfield "src/module/pm/usecase/taskfield/infra"
	workspace "src/module/pm/usecase/workspace/infra"
	jwt "src/module/socket/usecase/jwt/infra"

	"github.com/labstack/echo/v4"
)

func CollectRoutes(e *echo.Group) (*echo.Group, ctype.PemMap) {
	pemMap := ctype.PemMap{}
	e, pemMap = variable.RegisterUrls(e, pemMap)
	e, pemMap = authclient.RegisterUrls(e, pemMap)
	e, pemMap = tenant.RegisterUrls(e, pemMap)
	e, pemMap = role.RegisterUrls(e, pemMap)
	e, pemMap = user.RegisterUrls(e, pemMap)
	e, pemMap = workspace.RegisterUrls(e, pemMap)
	e, pemMap = project.RegisterUrls(e, pemMap)
	e, pemMap = taskfield.RegisterUrls(e, pemMap)
	e, pemMap = auth.RegisterUrls(e, pemMap)
	e, pemMap = profile.RegisterUrls(e, pemMap)
	e, pemMap = lockuser.RegisterUrls(e, pemMap)
	e, pemMap = signuptenant.RegisterUrls(e, pemMap)
	e, pemMap = reordertaskfield.RegisterUrls(e, pemMap)
	e, pemMap = task.RegisterUrls(e, pemMap)
	e, pemMap = feature.RegisterUrls(e, pemMap)
	e, pemMap = reorderfeature.RegisterUrls(e, pemMap)
	e, pemMap = reordertask.RegisterUrls(e, pemMap)
	e, pemMap = reorderstatus.RegisterUrls(e, pemMap)
	e, pemMap = message.RegisterUrls(e, pemMap)
	e, pemMap = doc.RegisterUrls(e, pemMap)
	e, pemMap = docattachment.RegisterUrls(e, pemMap)
	e, pemMap = createdocfromlink.RegisterUrls(e, pemMap)
	e, pemMap = jwt.RegisterUrls(e, pemMap)
	e, pemMap = github.RegisterUrls(e, pemMap)
	e, pemMap = gitlab.RegisterUrls(e, pemMap)
	return e, pemMap
}
