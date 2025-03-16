package ctrl

import (
	"net/http"

	"src/util/restlistutil"
	"src/util/vldtutil"

	"github.com/labstack/echo/v4"
)

type listSrvProvider[S any] interface {
	Paging(
		options restlistutil.ListOptions,
		searchFields []string,
	) (restlistutil.ListRestfulResult[S], error)
}

type listCtrl[S any] struct {
	srv listSrvProvider[S]
}

var searchableFields = []string{"key", "value", "description"}
var filterableFields = []string{"data_type"}
var orderableFields = []string{"id", "key"}

func (ctrl listCtrl[S]) Handler(c echo.Context) error {

	if err := vldtutil.CheckRequiredFilter(c, "tenant_id"); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	options := restlistutil.GetOptions(c, filterableFields, orderableFields)

	result, err := ctrl.srv.Paging(options, searchableFields)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, result)
}

func NewList[S any](srv listSrvProvider[S]) listCtrl[S] {
	return listCtrl[S]{srv: srv}
}
