package infra

import (
	"net/http"
	"src/client/scyllaclient"
	"src/common/ctype"
	"src/util/vldtutil"
	"strings"

	"github.com/labstack/echo/v4"

	"src/module/account/schema"
	"src/module/pm/repo/message"
	"src/module/socket/repo/centrifugo"
	"src/module/socket/usecase/publishmessage/app"
)

func Publish(c echo.Context) error {
	client := scyllaclient.NewClient()
	centrifugoRepo := centrifugo.New()
	messageRepo := message.New(client)

	srv := app.New(centrifugoRepo, messageRepo)

	user := c.Get("User").(*schema.User)
	userFullName := user.FirstName + " " + user.LastName
	userFullName = strings.TrimSpace(userFullName)
	data, err := vldtutil.ValidatePayload(c, InputData{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	socketMessage := app.SocketMessage{
		Channel: data.Channel,
		Data: app.SocketData{
			ID: "",
			User: app.SocketUser{
				ID:     user.ID,
				Name:   userFullName,
				Avatar: user.Avatar,
				Color:  user.Color,
			},
			TaskID:      data.TaskID,
			ProjectID:   data.ProjectID,
			Content:     data.Content,
			Attachments: []app.SocketAttachment{},
		},
	}

	files, err := vldtutil.UploadAndGetMetadata(c, "message")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	_, err = srv.Publish(socketMessage, files)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, ctype.Dict{})
}
