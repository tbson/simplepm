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
	data, err := vldtutil.ValidatePayload(c, InputData{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	socketMessage := app.SocketMessage{
		Channel: data.Channel,
		Data: app.SocketData{
			ID:        "",
			UserID:    userID,
			TaskID:    data.TaskID,
			ProjectID: data.ProjectID,
			Content:   data.Content,
		},
	}

	resultFiles, err := vldtutil.UploadAndGetMetadata(c, "message")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	messageID, err := srv.Publish(socketMessage)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	for _, file := range resultFiles {
		messageRepo.CreateAttachment(
			messageID,
			file.FileName,
			file.FileType,
			file.FileURL,
		)
	}

	return c.JSON(http.StatusOK, ctype.Dict{})
}
