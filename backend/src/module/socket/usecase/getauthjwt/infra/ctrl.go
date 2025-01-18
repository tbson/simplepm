package infra

import (
	"net/http"
	"src/common/ctype"

	"github.com/labstack/echo/v4"

	"src/module/socket/repo/centrifugo"
)

func GetAuthJWT(c echo.Context) error {
	repo := centrifugo.New()

	userID := c.Get("UserID").(uint)
	token, err := repo.GetAuthJwt(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, ctype.Dict{"token": token})
}
