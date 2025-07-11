package infra

import (
	"net/http"
	"src/module/account/repo/user"
	"src/module/account/usecase/lockuser/app"
	"src/util/dbutil"
	"src/util/errutil"
	"src/util/vldtutil"

	"github.com/labstack/echo/v4"
)

func LockUser(c echo.Context) error {
	structData, fields, err := vldtutil.ValidateUpdatePayload(c, InputData{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}

	data := vldtutil.GetDictByFields(structData, fields, []string{})
	id := vldtutil.ValidateId(c.Param("id"))

	userRepo := user.New(dbutil.Db(nil))
	srv := app.New(userRepo)

	locked := data["Locked"].(bool)
	lockedReason := data["LockedReason"].(string)
	result, err := srv.LockUser(id, locked, lockedReason)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}

	return c.JSON(http.StatusOK, MutatePres(result))
}
