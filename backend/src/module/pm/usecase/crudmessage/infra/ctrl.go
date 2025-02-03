package infra

import (
	"fmt"
	"net/http"
	"src/client/scyllaclient"
	"src/module/account/schema"
	"src/module/pm/repo/message"
	"src/util/numberutil"

	"src/module/pm/usecase/crudmessage/app"

	"github.com/labstack/echo/v4"

	"encoding/base64"
)

func List(c echo.Context) error {
	user := c.Get("User").(*schema.User)
	taskID := numberutil.StrToUint(c.QueryParam("task_id"), 0)

	client := scyllaclient.NewClient()
	messageRepo := message.New(client)
	srv := app.New(messageRepo)

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
