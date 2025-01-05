package infra

import (
	"net/http"
	"src/common/ctype"
	"src/common/setting"
	"src/util/tokenutil"

	"github.com/labstack/echo/v4"
)

func GetJWT(c echo.Context) error {
	userID := c.Get("UserID").(uint)
	secret := setting.CENTRIFUGO_SECRET
	lifeSpan := setting.CENTRIFUGO_JWT_LIFE_SPAN
	token, err := tokenutil.GenerateSimpleJWT(userID, secret, lifeSpan)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, ctype.Dict{"token": token})
}
