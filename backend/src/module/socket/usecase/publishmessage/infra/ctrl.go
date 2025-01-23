package infra

import (
	"net/http"
	"src/client/skyllaclient"
	"src/common/ctype"
	"src/util/vldtutil"

	"github.com/labstack/echo/v4"

	"src/module/pm/repo/message"
	"src/module/socket/repo/centrifugo"
	"src/module/socket/usecase/publishmessage/app"
)

func Publish(c echo.Context) error {
	client := skyllaclient.NewClient()
	centrifugoRepo := centrifugo.New()
	messageRepo := message.New(client)

	srv := app.New(centrifugoRepo, messageRepo)

	userID := c.Get("UserID").(uint)
	initData := app.SocketMessage{}
	initData.Data.UserID = userID
	data, err := vldtutil.ValidatePayload(c, initData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = srv.Publish(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, ctype.Dict{})
}
