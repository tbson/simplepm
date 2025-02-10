package infra

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"src/client/scyllaclient"
	"src/common/ctype"
	"src/util/numberutil"
	"src/util/vldtutil"
	"strings"

	"github.com/labstack/echo/v4"

	"src/module/account/schema"
	"src/module/event"
	"src/module/event/repo/centrifugo"
	"src/module/event/repo/message"
	"src/module/event/usecase/message/app"
)

var folder = "message"

func List(c echo.Context) error {
	user := c.Get("User").(*schema.User)
	taskID := numberutil.StrToUint(c.QueryParam("task_id"), 0)

	client := scyllaclient.NewClient()
	centrifugoRepo := centrifugo.New()
	messageRepo := message.New(client)
	srv := app.New(centrifugoRepo, messageRepo)

	pageStateParam := c.QueryParam("page_state")
	var pageState []byte
	if pageStateParam != "" {
		decoded, err := base64.StdEncoding.DecodeString(pageStateParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid page_state",
			})
		}
		pageState = decoded
	}

	messages, pageState, attachmentMap, err := srv.List(taskID, pageState)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, ListPres(messages, pageState, attachmentMap, *user))
}

func Create(c echo.Context) error {
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
			ID:   "",
			Type: event.CREATE_MESSAGE,
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

	files, err := vldtutil.UploadAndGetMetadata(c, folder)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	_, err = srv.Create(socketMessage, files)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, ctype.Dict{})
}

func Update(c echo.Context) error {
	client := scyllaclient.NewClient()
	centrifugoRepo := centrifugo.New()
	messageRepo := message.New(client)

	srv := app.New(centrifugoRepo, messageRepo)

	id := c.Param("id")
	taskID := vldtutil.ValidateId(c.Param("task_id"))

	structData, _, err := vldtutil.ValidateUpdatePayload(c, InputData{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	socketMessage := app.SocketMessage{
		Channel: structData.Channel,
		Data: app.SocketData{
			ID:      id,
			Type:    event.UPDATE_MESSAGE,
			Content: structData.Content,
		},
	}

	_, err = srv.Update(id, taskID, socketMessage)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, ctype.Dict{})
}

func Delete(c echo.Context) error {
	client := scyllaclient.NewClient()
	centrifugoRepo := centrifugo.New()
	messageRepo := message.New(client)

	srv := app.New(centrifugoRepo, messageRepo)

	id := c.Param("id")
	taskID := vldtutil.ValidateId(c.Param("task_id"))

	fmt.Println("id", id)
	fmt.Println("task_id", taskID)

	structData, _, err := vldtutil.ValidateUpdatePayload(c, InputData{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	socketMessage := app.SocketMessage{
		Channel: structData.Channel,
		Data: app.SocketData{
			ID:   id,
			Type: event.DELETE_MESSAGE,
		},
	}

	err = srv.Delete(id, taskID, socketMessage)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, ctype.Dict{})
}
