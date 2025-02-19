package main

import (
	"src/route"
	"src/util/dbutil"
	"src/util/localeutil"

	"src/module/account/repo/role"
	"src/module/account/repo/tenant"
	"src/module/account/usecase/syncrolespems/app"
	"src/module/account/usecase/syncrolespems/infra"

	"github.com/labstack/echo/v4"
)

func main() {
	dbutil.InitDb()
	localeutil.Init("en")

	e := echo.New()
	apiGroup := e.Group("/api/v1")
	_, pemMap := route.CollectRoutes(apiGroup)

	repo := infra.New(dbutil.Db(nil))
	tenantRepo := tenant.New(dbutil.Db(nil))
	roleRepo := role.New(dbutil.Db(nil))

	srv := app.New(repo, roleRepo, tenantRepo)

	srv.SyncRolesPems(pemMap)
}
