package ctrl

import (
	"net/http"

	"src/common/ctype"
	"src/module/account"
	"src/util/vldtutil"

	"github.com/labstack/echo/v4"
)

type SrvProvider interface {
	RefreshToken(refreshToken string) (account.TokenPair, error)
}

type ctrl struct {
	Srv SrvProvider
}

type input struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

func (ctrl ctrl) Handler(c echo.Context) error {
	structData, err := vldtutil.ValidatePayload(c, input{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	_, err = ctrl.Srv.RefreshToken(structData.RefreshToken)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, ctype.Dict{})
}

func New(srv SrvProvider) ctrl {
	return ctrl{Srv: srv}
}
