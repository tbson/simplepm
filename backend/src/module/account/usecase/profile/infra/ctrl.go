package infra

import (
	"net/http"
	"src/common/ctype"
	"src/module/account/repo/iam"
	"src/module/account/repo/user"
	"src/module/account/usecase/profile/app"
	"src/util/dbutil"
	"src/util/ssoutil"
	"src/util/vldtutil"

	"github.com/labstack/echo/v4"
)

func GetProfile(c echo.Context) error {
	userID := c.Get("UserID").(uint)
	client := dbutil.Db()
	userRepo := user.New(client)
	user, err := userRepo.Retrieve(ctype.QueryOptions{
		Filters:  ctype.Dict{"id": userID},
		Preloads: []string{"Tenant.AuthClient"},
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, user)
}

func UpdateProfile(c echo.Context) error {
	folder := "avatar"
	userID := c.Get("UserID").(uint)

	iamRepo := iam.New(ssoutil.Client())
	userRepo := user.New(dbutil.Db())

	srv := app.New(userRepo, iamRepo)

	data, err := vldtutil.ValidateUpdatePayload(c, InputData{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	data, err = vldtutil.UploadAndUPdatePayload(c, folder, data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	result, err := srv.UpdateProfile(userID, data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, result)
}

func ChangePassword(c echo.Context) error {
	userID := c.Get("UserID").(uint)

	iamRepo := iam.New(ssoutil.Client())
	userRepo := user.New(dbutil.Db())

	srv := app.New(userRepo, iamRepo)

	data, err := vldtutil.ValidateUpdatePayload(c, InputPassword{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	result, err := srv.ChangePassword(userID, data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, result)
}
