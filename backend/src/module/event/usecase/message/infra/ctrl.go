package infra

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"src/client/scyllaclient"
	"src/common/ctype"
	"src/util/errutil"
	"src/util/numberutil"
	"src/util/vldtutil"

	"github.com/labstack/echo/v4"

	"src/module/account/schema"
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
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}
	return c.JSON(http.StatusOK, ListPres(messages, pageState, attachmentMap, *user))
}

func Create(c echo.Context) error {
	client := scyllaclient.NewClient()
	centrifugoRepo := centrifugo.New()
	messageRepo := message.New(client)

	srv := app.New(centrifugoRepo, messageRepo)

	user := c.Get("User").(*schema.User)

	data, err := vldtutil.ValidatePayload(c, app.InputData{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}
	channel := data.Channel
	socketUser := app.SocketUser{
		ID:     user.ID,
		Name:   user.FullName(),
		Avatar: user.Avatar,
		Color:  user.Color,
	}

	files, err := vldtutil.UploadAndGetMetadata(c, folder)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}

	_, err = srv.Create(data, files, socketUser, channel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
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

	data, _, err := vldtutil.ValidateUpdatePayload(c, app.InputData{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}

	_, err = srv.Update(id, taskID, data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
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

	data, _, err := vldtutil.ValidateUpdatePayload(c, app.InputData{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}

	err = srv.Delete(id, taskID, data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}

	return c.JSON(http.StatusOK, ctype.Dict{})
}
