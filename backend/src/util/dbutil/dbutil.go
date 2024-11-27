package dbutil

import (
	"fmt"
	"src/common/setting"
	"src/util/testutil"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	account "src/module/account/schema"
	config "src/module/config/schema"
)

const DEFAULT_PAGE_SIZE = 10

var db *gorm.DB
var e error

func RegisterModels() []interface{} {
	return []interface{}{
		&config.Variable{},
		&account.Tenant{},
		&account.AuthClient{},
		&account.User{},
		&account.Role{},
		&account.Pem{},
	}
}

func InitDb() {
	host := setting.DB_HOST
	user := setting.DB_USER
	password := setting.DB_PASSWORD
	dbName := setting.DB_NAME
	port := setting.DB_PORT
	timeZone := setting.TIME_ZONE
	if testutil.IsTest() {
		dbName = dbName + "_test"
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", host, user, password, dbName, port, timeZone)
	db, e = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if e != nil {
		panic(e)
	}
}

func Db() *gorm.DB {
	return db
}
