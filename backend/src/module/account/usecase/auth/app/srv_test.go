package app

import (
	"src/util/dbutil"
	"testing"

	"src/module/account/repo/tenant"
	"src/module/account/repo/user"
	testingApp "src/module/testing/app"

	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	dbutil.InitDb()
	dbClient := dbutil.Db(nil)

	tenantRepo := tenant.New(dbClient)
	userRepo := user.New(dbClient)

	testtingSrv := testingApp.New(tenantRepo, userRepo)
	testtingSrv.InitData()

	m.Run()

	cleanup(dbClient)
}

func cleanup(dbClient *gorm.DB) {
	dbClient.Exec("TRUNCATE TABLE tenants")
	dbClient.Exec("TRUNCATE TABLE users")
}
