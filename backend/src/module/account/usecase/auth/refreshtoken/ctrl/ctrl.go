package ctrl

import (
	"context"
	"net/http"

	"src/common/authtype"
	"src/common/ctype"
	"src/util/vldtutil"

	"github.com/labstack/echo/v4"
)

type SrvProvider interface {
	RefreshToken(
		ctx context.Context,
		realm string,
		refreshToken string,
	) (authtype.SsoCallbackResult, error)
}

type ctrl struct {
	Srv SrvProvider
}

type inputData struct {
	Realm        string `json:"realm" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required"`
}

func (ctrl ctrl) Handler(c echo.Context) error {
	ctx := c.Request().Context()
	structData, err := vldtutil.ValidatePayload(c, inputData{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	_, err = ctrl.Srv.RefreshToken(ctx, structData.Realm, structData.RefreshToken)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, ctype.Dict{})
}

func New(srv SrvProvider) ctrl {
	return ctrl{Srv: srv}
}
