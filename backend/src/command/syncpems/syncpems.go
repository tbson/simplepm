package main

import (
	"src/route"
	"src/util/dbutil"

	"src/module/account/repo/role"
	"src/module/account/repo/tenant"
	"src/module/account/usecase/syncrolespems/app"
	"src/module/account/usecase/syncrolespems/infra"

	"github.com/labstack/echo/v4"
)

func main() {
	dbutil.InitDb()

	e := echo.New()
	apiGroup := e.Group("/api/v1")
	_, pemMap := route.CollectRoutes(apiGroup)

	repo := infra.New(dbutil.Db())
	tenantRepo := tenant.New(dbutil.Db())
	roleRepo := role.New(dbutil.Db())

	srv := app.New(repo, roleRepo, tenantRepo)

	srv.SyncRolesPems(pemMap)
}
