package infra

import (
	"net/http"
	"src/client/skyllaclient"
	"src/module/pm/repo/message"
	"src/util/numberutil"

	"github.com/labstack/echo/v4"
)

func List(c echo.Context) error {
	client := skyllaclient.NewClient()
	repo := message.New(client)
	// get taskID from query parameter
	taskIDStr := c.QueryParam("task_id")
	// convert it to uint
	taskID := numberutil.StrToUint(taskIDStr, 0)
	result, err := repo.List(taskID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, result)
}
