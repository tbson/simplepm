package ctrl

import (
	"net/http"

	"src/common/ctype"

	"github.com/labstack/echo/v4"
)

type ctrl struct{}

func (ctrl ctrl) Handler(c echo.Context) error {
	return c.JSON(http.StatusOK, ctype.Dict{})
}

func New() ctrl {
	return ctrl{}
}
