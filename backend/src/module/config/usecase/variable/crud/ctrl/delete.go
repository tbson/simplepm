package ctrl

import (
	"net/http"

	"src/util/vldtutil"

	"github.com/labstack/echo/v4"
)

type deleteSrvProvider[S any] interface {
	Delete(id uint) ([]uint, error)
}

type deleteCtrl[S any] struct {
	srv deleteSrvProvider[S]
}

func (ctrl deleteCtrl[S]) Handler(c echo.Context) error {
	id := vldtutil.ValidateId(c.Param("id"))

	ids, err := ctrl.srv.Delete(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, ids)
}

func NewDelete[S any](srv deleteSrvProvider[S]) deleteCtrl[S] {
	return deleteCtrl[S]{srv: srv}
}
