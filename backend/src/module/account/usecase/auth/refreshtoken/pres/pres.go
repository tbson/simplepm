package pres

import (
	"net/http"
	"src/common/ctype"
	"src/util/cookieutil"

	"src/module/account/domain/model"

	"github.com/labstack/echo/v4"
)

func RefreshToken(c echo.Context, tokenPair model.TokenPair, clientType string) error {
	if clientType == model.CLIENT_TYPE_WEB {
		return cookiePres(c, tokenPair)
	}
	return jsonPres(c, tokenPair)
}

func cookiePres(c echo.Context, tokenPair model.TokenPair) error {
	accessToken := tokenPair.AccessToken
	refreshToken := tokenPair.RefreshToken
	accessTokenCookie := cookieutil.SetAccessTokenCookie(accessToken)
	refreshTokenCookie := cookieutil.SetRefreshTokenCookie(refreshToken)
	c.SetCookie(accessTokenCookie)
	c.SetCookie(refreshTokenCookie)

	return c.JSON(http.StatusOK, ctype.Dict{})
}

func jsonPres(c echo.Context, tokenPair model.TokenPair) error {
	data := map[string]interface{}{
		"access_token":  tokenPair.AccessToken,
		"refresh_token": tokenPair.RefreshToken,
	}

	return c.JSON(http.StatusOK, data)
}
