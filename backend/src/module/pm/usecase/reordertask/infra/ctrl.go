package infra

import (
	"net/http"
	"src/util/dbutil"
	"src/util/errutil"
	"src/util/vldtutil"

	"src/module/pm/repo/task"
	"src/module/pm/repo/taskfieldvalue"

	"src/module/pm/usecase/reordertask/app"

	"github.com/labstack/echo/v4"
)

func Reorder(c echo.Context) error {
	db := dbutil.Db(nil)
	tx := db.Begin()
	if tx.Error != nil {
		msg := errutil.NewRaw(tx.Error.Error())
		return c.JSON(http.StatusBadRequest, msg)
	}

	data, err := vldtutil.ValidatePayload(c, app.InputData{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}

	taskRepo := task.New(tx)
	taskFieldValueRepo := taskfieldvalue.New(tx)
	srv := app.New(taskRepo, taskFieldValueRepo)

	result, err := srv.Reorder(data)
	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}
	if err := tx.Commit().Error; err != nil {
		msg := errutil.NewRaw(err.Error())
		return c.JSON(http.StatusBadRequest, msg)
	}

	return c.JSON(http.StatusOK, result)
}
