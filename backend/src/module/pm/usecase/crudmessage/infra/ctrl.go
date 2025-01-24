package infra

import (
	"fmt"
	"net/http"
	"src/client/skyllaclient"
	"src/module/account/schema"
	"src/module/pm/repo/message"
	"src/util/numberutil"

	"src/module/pm/usecase/crudmessage/app"

	"github.com/labstack/echo/v4"
)

func List(c echo.Context) error {
	user := c.Get("User").(*schema.User)
	taskID := numberutil.StrToUint(c.QueryParam("task_id"), 0)

	client := skyllaclient.NewClient()
	messageRepo := message.New(client)
	srv := app.New(messageRepo)

	messages, attachmentMap, err := srv.List(taskID)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, ListPres(messages, attachmentMap, *user))
}
