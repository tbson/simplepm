package ctrl

import (
	"net/http"

	"src/common/ctype"
	"src/util/vldtutil"

	"src/module/config/schema"

	"src/module/config/pres"

	"github.com/labstack/echo/v4"
)

type srvProvider interface {
	Retrieve(queryOptions ctype.QueryOptions) (*schema.Variable, error)
}

type ctrl struct {
	srv srvProvider
}

// go-openapi doc
// @Summary Retrieve
// @Description Retrieve
// @Tags variable
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} pres.DetailResult
// @Failure 400 {object} ctype.Dict
// @Router /config/variable/{id} [get]
func (ctrl ctrl) Handler(c echo.Context) error {

	id := vldtutil.ValidateId(c.Param("id"))
	queryOptions := ctype.QueryOptions{
		Filters: ctype.Dict{"id": id},
	}
	result, err := ctrl.srv.Retrieve(queryOptions)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, pres.DetailPres(*result))
}

func New(srv srvProvider) ctrl {
	return ctrl{srv: srv}
}
