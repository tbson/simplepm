package infra

import (
	"net/http"
	"src/module/account/repo/user"
	"src/module/account/usecase/lockuser/app"
	"src/util/dbutil"
	"src/util/vldtutil"

	"github.com/labstack/echo/v4"
)

func LockUser(c echo.Context) error {
	data, err := vldtutil.ValidateUpdatePayload(c, InputData{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	id := vldtutil.ValidateId(c.Param("id"))

	userRepo := user.New(dbutil.Db())
	srv := app.New(userRepo)

	locked := data["Locked"].(bool)
	lockedReason := data["LockedReason"].(string)
	result, err := srv.LockUser(id, locked, lockedReason)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, MutatePres(result))
}
