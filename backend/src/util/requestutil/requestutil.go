package requestutil

import (
	"mime/multipart"
	"strings"

	"src/util/stringutil"

	"github.com/labstack/echo/v4"
)

func GetFileHeaderMap(c echo.Context) map[string][]*multipart.FileHeader {
	result := map[string][]*multipart.FileHeader{}
	if !strings.Contains(c.Request().Header.Get("Content-Type"), "multipart/form-data") {
		return result
	}

	form, err := c.MultipartForm()
	if err != nil {
		return result
	}

	// Add form data to keys map
	for k, v := range form.File {
		result[stringutil.ToCamelCase(k)] = v
	}

	return result
}
