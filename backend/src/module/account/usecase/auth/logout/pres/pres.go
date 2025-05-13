package pres

import (
	"src/common/ctype"
	"src/util/cookieutil"
	"src/util/presutil"

	"github.com/labstack/echo/v4"
)

func Logout(c echo.Context) error {
	resp := presutil.New(c)
	accessTokenCookie := cookieutil.SetAccessTokenCookie("")
	refreshTokenCookie := cookieutil.SetRefreshTokenCookie("")
	sessionIDCookie := cookieutil.SetSessionIDCookie("")
	c.SetCookie(accessTokenCookie)
	c.SetCookie(refreshTokenCookie)
	c.SetCookie(sessionIDCookie)
	return resp.Ok(ctype.Dict{})
}
