package ctrl

import (
	"src/module/account/usecase/auth/logout/pres"

	"github.com/labstack/echo/v4"
)

type ctrl struct{}

// Logout godoc
// @Summary Logout
// @Description Logout
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} ctype.Dict
// @Failure 400 {object} ctype.Dict
// @Router /account/auth/logout [post]
func (ctrl ctrl) Handler(c echo.Context) error {
	return pres.Logout(c)
}

func New() ctrl {
	return ctrl{}
}
