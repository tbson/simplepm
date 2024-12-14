package main

import (
	"src/module/account/repo/authclient"
	"src/module/account/repo/role"
	"src/module/account/repo/tenant"
	"src/module/account/repo/user"
	"src/module/account/usecase/initdata/app"
	"src/route"
	"src/util/dbutil"
	"src/util/localeutil"

	"github.com/labstack/echo/v4"
)

func main() {
	localeutil.Init("en")
	dbutil.InitDb()
	db := dbutil.Db()
	tx := db.Begin()
	if tx.Error != nil {
		panic(tx.Error.Error())
	}

	e := echo.New()
	apiGroup := e.Group("/api/v1")
	_, pemMap := route.CollectRoutes(apiGroup)

	authClientRepo := authclient.New(tx)
	tenantRepo := tenant.New(tx)
	userRepo := user.New(tx)
	roleRepo := role.New(tx)

	srv := app.New(authClientRepo, tenantRepo, userRepo, roleRepo)
	err := srv.InitData(pemMap)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	if err := tx.Commit().Error; err != nil {
		panic(err.Error())
	}
}
