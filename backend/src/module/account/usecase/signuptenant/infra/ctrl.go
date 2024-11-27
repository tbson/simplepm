package infra

import (
	"net/http"
	"src/common/ctype"
	"src/util/dbutil"
	"src/util/errutil"
	"src/util/ssoutil"
	"src/util/vldtutil"

	"src/module/account/repo/authclient"
	"src/module/account/repo/iam"
	"src/module/account/repo/role"
	"src/module/account/repo/tenant"
	"src/module/account/repo/user"
	"src/module/account/usecase/signuptenant/app"

	"github.com/labstack/echo/v4"
)

func SignupTenant(c echo.Context) error {
	data, err := vldtutil.ValidatePayload(c, InputData{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// for tenant
	uid := data["Uid"].(string)
	title := data["Title"].(string)

	// for user
	email := data["Email"].(string)
	mobile := data["Mobile"].(*string)
	firstName := data["FirstName"].(string)
	lastName := data["LastName"].(string)
	password := data["Password"].(string)
	admin := true

	db := dbutil.Db()
	tx := db.Begin()
	if tx.Error != nil {
		msg := errutil.New("", []string{tx.Error.Error()})
		return c.JSON(http.StatusBadRequest, msg)
	}

	authClientRepo := authclient.New(tx)
	tenantRepo := tenant.New(tx)
	userRepo := user.New(tx)
	roleRepo := role.New(tx)
	iamRepo := iam.New(ssoutil.Client())

	srv := app.New(authClientRepo, tenantRepo, userRepo, roleRepo, iamRepo)

	err = srv.SignupTenant(
		uid,
		title,
		email,
		mobile,
		firstName,
		lastName,
		password,
		admin,
	)

	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := tx.Commit().Error; err != nil {
		msg := errutil.New("", []string{err.Error()})
		return c.JSON(http.StatusBadRequest, msg)
	}

	return c.JSON(http.StatusOK, ctype.Dict{})
}
