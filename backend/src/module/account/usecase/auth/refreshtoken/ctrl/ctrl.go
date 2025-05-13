package ctrl

import (
	"src/util/presutil"

	"src/module/account/domain/model"
	"src/module/account/usecase/auth/refreshtoken/pres"
	"src/module/account/usecase/auth/refreshtoken/pres/cookie"
	"src/util/cookieutil"
	"src/util/vldtutil"

	"github.com/labstack/echo/v4"
)

type SrvProvider interface {
	RefreshToken(refreshToken string) (model.TokenPair, error)
}

type ctrl struct {
	Srv SrvProvider
}

type input struct {
	// RefreshToken string `json:"refresh_token" validate:"required"`
	ClientType string `json:"client_type" validate:"required,oneof=web app"`
}

func (ctrl ctrl) Handler(c echo.Context) error {
	resp := presutil.New(c)
	structData, err := vldtutil.ValidatePayload(c, input{})
	if err != nil {
		return resp.Err(err)
	}
	refreshToken := cookieutil.GetValue(c, "refresh_token")
	tokenPair, err := ctrl.Srv.RefreshToken(refreshToken)
	if err != nil {
		return resp.Err(err)
	}

	if structData.ClientType == model.CLIENT_TYPE_WEB {
		return cookie.RefreshTokenPres(c, tokenPair)
	}

	return pres.RefreshToken(c, tokenPair, structData.ClientType)
}

func New(srv SrvProvider) ctrl {
	return ctrl{Srv: srv}
}
