package infra

import (
	"net/http"
	"src/util/dbutil"
	"src/util/errutil"
	"src/util/vldtutil"

	"src/module/pm/repo/feature"

	"src/module/pm/usecase/reorderfeature/app"

	"github.com/labstack/echo/v4"
)

func Reorder(c echo.Context) error {
	db := dbutil.Db()
	tx := db.Begin()
	if tx.Error != nil {
		msg := errutil.New("", []string{tx.Error.Error()})
		return c.JSON(http.StatusBadRequest, msg)
	}

	data, err := vldtutil.ValidatePayload(c, app.InputData{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	repo := feature.New(tx)
	srv := app.New(repo)

	result, err := srv.Reorder(data)
	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := tx.Commit().Error; err != nil {
		msg := errutil.New("", []string{err.Error()})
		return c.JSON(http.StatusBadRequest, msg)
	}

	return c.JSON(http.StatusOK, result)
}
