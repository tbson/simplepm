package resputil

import (
	"net/http"

	"src/util/errutil"

	"github.com/labstack/echo/v4"
)

type resp struct {
	c echo.Context
}

func errResp(c echo.Context, err error, errCode int) error {
	return c.JSON(errCode, err.(*errutil.CustomError).Localize())
}

func New(c echo.Context) *resp {
	return &resp{c: c}
}

func (r resp) Ok(result interface{}) error {
	return r.c.JSON(http.StatusOK, result)
}

func (r resp) Err(err error) error {
	return errResp(r.c, err, http.StatusBadRequest)
}

func (r resp) Err401(err error) error {
	return errResp(r.c, err, http.StatusUnauthorized)
}

func (r resp) Err403(err error) error {
	return errResp(r.c, err, http.StatusForbidden)
}

func (r resp) Err404(err error) error {
	return errResp(r.c, err, http.StatusNotFound)
}
