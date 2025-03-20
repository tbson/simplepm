package ctrl

import (
	"net/http"

	"src/common/ctype"
	"src/util/vldtutil"

	"github.com/labstack/echo/v4"

	"src/module/config/schema"

	"src/module/config/pres"
)

type input struct {
	Key         string `json:"key" form:"key" validate:"required"`
	Value       string `json:"value" form:"value"`
	Description string `json:"description" form:"description"`
	DataType    string `json:"data_type" form:"data_type" validate:"required,oneof=STRING INTEGER FLOAT BOOLEAN DATE DATETIME"`
}

type srvProvider interface {
	Update(opts ctype.QueryOpts, data ctype.Dict) (*schema.Variable, error)
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
	id := vldtutil.ValidateId(c.Param("id"))

	structData, fields, err := vldtutil.ValidateUpdatePayload(c, input{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	data := vldtutil.GetDictByFields(structData, fields, []string{})
	updateOpts := ctype.QueryOpts{Filters: ctype.Dict{"ID": id}}
	result, err := ctrl.srv.Update(updateOpts, data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, pres.DetailPres(*result))
}

func New(srv srvProvider) ctrl {
	return ctrl{srv: srv}
}
