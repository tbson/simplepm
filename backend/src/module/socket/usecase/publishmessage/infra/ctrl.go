package infra

import (
	"net/http"
	"src/common/ctype"
	"src/util/vldtutil"

	"github.com/labstack/echo/v4"

	"src/module/socket/repo/centrifugo"
	"src/util/tokenutil"
)

func Publish(c echo.Context) error {
	repo := centrifugo.New()
	userID := c.Get("UserID").(uint)
	uuid := tokenutil.GenerateUUID()
	initData := ctype.SocketMessage{}
	initData.Data.UserID = userID
	initData.Data.ID = uuid
	data, err := vldtutil.ValidatePayload(c, initData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = repo.Publish(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, ctype.Dict{})
}
