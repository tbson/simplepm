package infra

import (
	"go/doc"
	"net/http"
	"src/common/ctype"
	"src/module/document/schema"

	"github.com/labstack/echo/v4"
)

type Schema = schema.Doc

var NewRepo = doc.New

func Create(c echo.Context) error {
	result := ctype.Dict{}
	return c.JSON(http.StatusCreated, result)
}
