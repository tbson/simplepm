package ctrl

import (
	"fmt"
	"net/http"

	"src/module/account/domain/model"
	"src/module/account/usecase/auth/refreshtoken/pres/cookie"
	"src/module/account/usecase/auth/refreshtoken/pres/json"
	"src/util/cookieutil"
	"src/util/errutil"
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
	structData, err := vldtutil.ValidatePayload(c, input{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}
	refreshToken := cookieutil.GetValue(c, "refresh_token")
	fmt.Println("refreshToken", refreshToken)
	tokenPair, err := ctrl.Srv.RefreshToken(refreshToken)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}

	if structData.ClientType == "web" {
		return cookie.RefreshTokenPres(c, tokenPair)
	}

	return json.RefreshTokenPres(c, tokenPair)
}

func New(srv SrvProvider) ctrl {
	return ctrl{Srv: srv}
}
