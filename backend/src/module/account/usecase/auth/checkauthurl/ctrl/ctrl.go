package ctrl

import (
	"net/http"

	"src/common/ctype"

	"github.com/labstack/echo/v4"
)

type SrvProvider interface {
	GetAuthUrl(tenantUid string, nextParam string) (string, error)
}

type ctrl struct {
	Srv SrvProvider
}

func (ctrl ctrl) Handler(c echo.Context) error {
	tenantUid := c.Param("tenantUid")

	_, err := ctrl.Srv.GetAuthUrl(tenantUid, "")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, ctype.Dict{})
}

func New(srv SrvProvider) ctrl {
	return ctrl{Srv: srv}
}
