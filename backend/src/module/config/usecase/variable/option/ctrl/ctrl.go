package ctrl

import (
	"src/util/resputil"

	"src/module/config"

	"src/module/config/pres"

	"github.com/labstack/echo/v4"
)

type ctrl struct{}

// go-openapi doc
// @Summary Option
// @Description Option
// @Tags variable
// @Accept json
// @Produce json
// @Success 200 {object} pres.OptionResult
// @Router /config/variable/option/ [get]
func (ctrl ctrl) Handler(c echo.Context) error {
	resp := resputil.New(c)

	result := pres.OptionPres(config.VariableDataTypeOptions)
	return resp.Ok(result)
}

func New() ctrl {
	return ctrl{}
}
