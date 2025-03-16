package ctrl

import (
	"net/http"

	"src/common/ctype"
	"src/util/vldtutil"

	"github.com/labstack/echo/v4"
)

type retrieveSrvProvider[S any] interface {
	Retrieve(queryOptions ctype.QueryOptions) (*S, error)
}

type retrieveCtrl[S any] struct {
	srv retrieveSrvProvider[S]
}

func (ctrl retrieveCtrl[S]) Handler(c echo.Context) error {

	id := vldtutil.ValidateId(c.Param("id"))
	queryOptions := ctype.QueryOptions{
		Filters: ctype.Dict{"id": id},
	}
	result, err := ctrl.srv.Retrieve(queryOptions)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, result)
}

func NewRetrieve[S any](srv retrieveSrvProvider[S]) retrieveCtrl[S] {
	return retrieveCtrl[S]{srv: srv}
}
