package infra

import (
	"net/http"
	"src/common/ctype"
	"src/module/account/repo/user"
	"src/module/account/usecase/profile/app"
	"src/util/dbutil"
	"src/util/errutil"
	"src/util/vldtutil"

	"github.com/labstack/echo/v4"
)

func GetProfile(c echo.Context) error {
	userID := c.Get("UserID").(uint)
	client := dbutil.Db(nil)
	userRepo := user.New(client)
	user, err := userRepo.Retrieve(ctype.QueryOpts{
		Filters: ctype.Dict{"id": userID},
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}

	return c.JSON(http.StatusOK, user)
}

func UpdateProfile(c echo.Context) error {
	folder := "avatar"
	userID := c.Get("UserID").(uint)

	userRepo := user.New(dbutil.Db(nil))

	srv := app.New(userRepo)

	structData, fields, err := vldtutil.ValidateUpdatePayload(c, InputData{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}

	data := vldtutil.GetDictByFields(structData, fields, []string{})
	data, err = vldtutil.UploadAndUPdatePayload(c, folder, data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}

	result, err := srv.UpdateProfile(userID, data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}

	return c.JSON(http.StatusOK, result)
}

func ChangePassword(c echo.Context) error {
	userID := c.Get("UserID").(uint)

	userRepo := user.New(dbutil.Db(nil))

	srv := app.New(userRepo)

	structData, fields, err := vldtutil.ValidateUpdatePayload(c, InputPassword{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}

	data := vldtutil.GetDictByFields(structData, fields, []string{})
	result, err := srv.ChangePwd(userID, data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}

	return c.JSON(http.StatusOK, result)
}
