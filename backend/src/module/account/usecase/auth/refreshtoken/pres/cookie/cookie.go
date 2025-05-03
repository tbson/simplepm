package cookie

import (
	"net/http"
	"src/common/ctype"
	"src/util/cookieutil"

	"src/module/account/domain/model"

	"github.com/labstack/echo/v4"
)

func RefreshTokenPres(c echo.Context, tokenPair model.TokenPair) error {
	accessToken := tokenPair.AccessToken
	refreshToken := tokenPair.RefreshToken
	accessTokenCookie := cookieutil.NewAccessTokenCookie(accessToken)
	refreshTokenCookie := cookieutil.NewRefreshTokenCookie(refreshToken)
	c.SetCookie(accessTokenCookie)
	c.SetCookie(refreshTokenCookie)

	return c.JSON(http.StatusOK, ctype.Dict{})
}
