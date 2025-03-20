package infra

import (
	"net/http"
	"src/common/ctype"

	"src/common/setting"

	"github.com/labstack/echo/v4"

	"src/util/cookieutil"
	"src/util/tokenutil"
)

func GetAuthJWT(c echo.Context) error {
	clientID := cookieutil.GetValue(c, "session_id")
	clientSecret := setting.CENTRIFUGO_CLIENT_SECRET()
	lifeSpan := setting.CENTRIFUGO_JWT_LIFE_SPAN()
	userID := c.Get("UserID").(uint)
	token, err := tokenutil.GenerateSimpleJWT(clientID, userID, clientSecret, lifeSpan)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, ctype.Dict{"token": token})
}

func GetSubscriptionJWT(c echo.Context) error {
	clientID := cookieutil.GetValue(c, "session_id")
	clientSecret := setting.CENTRIFUGO_CLIENT_SECRET()
	channel := c.QueryParam("channel")
	token, err := tokenutil.GenerateSubscriptionJWT(clientID, clientSecret, channel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, ctype.Dict{"token": token})
}
