package infra

import (
	"net/http"
	"src/module/document/schema"
	"src/util/dbutil"
	"src/util/vldtutil"

	"src/module/document/repo/doc"
	"src/module/document/repo/net"

	"src/module/document/usecase/createdocfromlink/app"

	"github.com/labstack/echo/v4"
)

type Schema = schema.Doc

func Create(c echo.Context) error {
	userID := c.Get("UserID").(uint)

	docRepo := doc.New(dbutil.Db(nil))
	netRepo := net.New(http.DefaultClient)
	srv := app.New(docRepo, netRepo)

	structData, err := vldtutil.ValidatePayload(c, app.InputData{UserID: userID})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	result, err := srv.Create(structData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, result)
}
