package frameworkutil

import "github.com/labstack/echo/v4"

type CtrlHandler interface {
	Handler(e echo.Context) error
}
