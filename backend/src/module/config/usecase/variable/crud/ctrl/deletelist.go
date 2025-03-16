package ctrl

import (
	"net/http"

	"src/util/vldtutil"

	"github.com/labstack/echo/v4"
)

type deleteListSrvProvider[S any] interface {
	DeleteList(ids []uint) ([]uint, error)
}

type deleteListCtrl[S any] struct {
	srv deleteListSrvProvider[S]
}

func (ctrl deleteListCtrl[S]) Handler(c echo.Context) error {
	ids := vldtutil.ValidateIds(c.QueryParam("ids"))

	ids, err := ctrl.srv.DeleteList(ids)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, ids)
}

func NewDeleteList[S any](srv deleteListSrvProvider[S]) deleteListCtrl[S] {
	return deleteListCtrl[S]{srv: srv}
}
