package dbutil

import (
	"context"
	"fmt"
	"src/common/setting"
	"src/util/errutil"
	"src/util/testutil"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	account "src/module/account/schema"
	config "src/module/config/schema"
	document "src/module/document/schema"
	event "src/module/event/schema"
	pm "src/module/pm/schema"
)

const DEFAULT_PAGE_SIZE = 10

var db *gorm.DB
var e error

func RegisterModels() []interface{} {
	return []interface{}{
		&config.Variable{},
		&account.Tenant{},
		&account.User{},
		&account.Role{},
		&account.Pem{},
		&account.GitAccount{},
		&account.GitRepo{},
		&pm.Workspace{},
		&pm.WorkspaceUser{},
		&pm.Project{},
		&pm.ProjectUser{},
		&pm.TaskUser{},
		&pm.TaskField{},
		&pm.TaskFieldOption{},
		&pm.Feature{},
		&pm.Task{},
		&pm.TaskFieldValue{},
		&pm.GitPush{},
		&pm.GitCommit{},
		&document.Doc{},
		&document.DocAttachment{},
		&event.Change{},
	}
}

func InitDb() {
	host := setting.DB_HOST()
	user := setting.DB_USER()
	password := setting.DB_PASSWORD()
	dbName := setting.DB_NAME()
	port := setting.DB_PORT()
	timeZone := setting.TIME_ZONE()
	if testutil.IsTest() {
		dbName = dbName + "_test"
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		host, user, password, dbName, port, timeZone)
	db, e = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if e != nil {
		panic(e)
	}
}

func Db(ctx *context.Context) *gorm.DB {
	if ctx == nil {
		return db
	}
	return db.WithContext(*ctx)
}

func WithTx(db *gorm.DB, fn func(*gorm.DB) error) error {
	tx := db.Begin()
	if tx.Error != nil {
		return errutil.NewRaw(tx.Error.Error())
	}

	// Ensure rollback on panic
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Ensure rollback on error
	if err := fn(tx); err != nil {
		tx.Rollback()
		return errutil.NewRaw(err.Error())
	}

	// Return the error from the commit operation
	if commitErr := tx.Commit().Error; commitErr != nil {
		return errutil.NewRaw(commitErr.Error())
	}
	return nil
}

func WithTxWithValue[T any](db *gorm.DB, fn func(*gorm.DB) (T, error)) (T, error) {
	var result T
	tx := db.Begin()
	if tx.Error != nil {
		return result, errutil.NewRaw(tx.Error.Error())
	}

	// Ensure rollback on panic
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Ensure rollback on error
	var err error
	result, err = fn(tx)
	if err != nil {
		tx.Rollback()
		return result, errutil.NewRaw(err.Error())
	}

	// Return the error from the commit operation
	if commitErr := tx.Commit().Error; commitErr != nil {
		return result, errutil.NewRaw(commitErr.Error())
	}

	return result, nil
}
