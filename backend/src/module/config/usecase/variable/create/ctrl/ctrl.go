package ctrl

import (
	"net/http"

	"src/util/errutil"
	"src/util/presutil"
	"src/util/vldtutil"

	"github.com/labstack/echo/v4"

	"src/module/config/schema"

	"src/module/config/pres"
	"src/module/config/vltd"
)

type srvProvider interface {
	Create(structData vltd.CreateVariableInput) (*schema.Variable, error)
}

type ctrl struct {
	srv srvProvider
}

// go-openapi doc
// @Summary Create
// @Description Create
// @Tags variable
// @Accept json
// @Produce json
// @Param key body string true "Key"
// @Param value body string false "Value"
// @Param description body string false "Description"
// @Param data_type body string true "Data Type"
// @Success 200 {object} pres.DetailResult
// @Failure 400 {object} ctype.Dict
// @Router /config/variable [post]
func (ctrl ctrl) Handler(c echo.Context) error {
	resp := presutil.New(c)
	structData, err := vldtutil.ValidatePayload(c, vltd.CreateVariableInput{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}

	result, err := ctrl.srv.Create(structData)
	if err != nil {
		return resp.Err(err)
	}
	return resp.Ok(pres.DetailPres(*result))
}

func New(srv srvProvider) ctrl {
	return ctrl{srv: srv}
}
