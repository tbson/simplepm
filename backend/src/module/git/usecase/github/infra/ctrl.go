package infra

import (
	"encoding/json"
	"fmt"
	"net/http"
	"src/common/ctype"

	"github.com/labstack/echo/v4"
)

func Callback(c echo.Context) error {
	fmt.Println("query params.....................")
	for k, v := range c.QueryParams() {
		fmt.Println(k, v)
	}

	// instalationID := c.QueryParam("installation_id")
	// tenantUID := c.QueryParam("state")

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
	// print all query params
	fmt.Println("query params.....................")
	for k, v := range c.QueryParams() {
		fmt.Println(k, v)
	}

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
