package infra

import (
	"encoding/json"
	"fmt"
	"net/http"
	"src/common/ctype"

	"github.com/labstack/echo/v4"
)

func Callback(c echo.Context) error {

	var result ctype.Dict
	err := c.Bind(&result)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	jsonString, err := json.MarshalIndent(result, "", "  ")
	fmt.Println("github callback.....................")
	fmt.Println(string(jsonString))
	return c.JSON(http.StatusOK, result)
}

func Webhook(c echo.Context) error {

	var result ctype.Dict
	err := c.Bind(&result)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	jsonString, err := json.MarshalIndent(result, "", "  ")
	fmt.Println("github webhook.....................")
	fmt.Println(string(jsonString))
	return c.JSON(http.StatusOK, result)
}
