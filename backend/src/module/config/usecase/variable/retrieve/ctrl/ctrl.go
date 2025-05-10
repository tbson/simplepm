package ctrl

import (
	"src/util/resputil"

	"src/common/ctype"
	"src/util/vldtutil"

	"src/module/config/schema"

	"src/module/config/pres"

	"github.com/labstack/echo/v4"
)

type srvProvider interface {
	Retrieve(opts ctype.QueryOpts) (*schema.Variable, error)
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
	resp := resputil.New(c)

	id := vldtutil.ValidateId(c.Param("id"))
	opts := ctype.QueryOpts{
		Filters: ctype.Dict{"id": id},
	}
	result, err := ctrl.srv.Retrieve(opts)
	if err != nil {
		return resp.Err404(err)
	}

	return resp.Ok(pres.DetailPres(*result))
}

func New(srv srvProvider) ctrl {
	return ctrl{srv: srv}
}
