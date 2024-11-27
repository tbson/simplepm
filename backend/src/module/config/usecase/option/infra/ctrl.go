package infra

import (
	"net/http"

	"src/common/ctype"
	"src/module/config"

	"github.com/labstack/echo/v4"
)

func Option(c echo.Context) error {
	result := ctype.Dict{
		"data_type": config.VariableDataTypeOptions,
	}
	return c.JSON(http.StatusOK, result)
}
