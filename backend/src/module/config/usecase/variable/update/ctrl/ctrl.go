package ctrl

import (
	"src/util/presutil"

	"src/common/ctype"
	"src/util/vldtutil"

	"github.com/labstack/echo/v4"

	"src/module/config/schema"

	"src/module/config/pres"
	"src/module/config/vltd"
)

type srvProvider interface {
	Update(
		opts ctype.QueryOpts,
		structData vltd.UpdateVariableInput,
		fields []string,
	) (*schema.Variable, error)
}

type ctrl struct {
	srv srvProvider
}

// go-openapi doc
// @Summary Update
// @Description Update
// @Tags variable
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Param key body string true "Key"
// @Param value body string false "Value"
// @Param description body string false "Description"
// @Param data_type body string true "Data Type"
// @Success 200 {object} pres.DetailResult
// @Failure 400 {object} ctype.Dict
// @Router /config/variable/{id} [put]
func (ctrl ctrl) Handler(c echo.Context) error {
	resp := presutil.New(c)
	id := c.Param("id")

	structData, fields, err := vldtutil.ValidateUpdatePayload(c, vltd.UpdateVariableInput{})
	if err != nil {
		return resp.Err(err)
	}
	updateOpts := ctype.QueryOpts{Filters: ctype.Dict{"ID": id}}
	result, err := ctrl.srv.Update(updateOpts, structData, fields)

	if err != nil {
		return resp.Err(err)
	}

	return resp.Ok(pres.DetailPres(*result))
}

func New(srv srvProvider) ctrl {
	return ctrl{srv: srv}
}
