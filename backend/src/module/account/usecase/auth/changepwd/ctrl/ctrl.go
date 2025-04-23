package ctrl

import (
	"net/http"

	"src/common/ctype"
	"src/util/errutil"
	"src/util/vldtutil"

	"github.com/labstack/echo/v4"
)

type SrvProvider interface {
	ChangePwd(userID uint, pwd string) error
}

type ctrl struct {
	Srv SrvProvider
}

type input struct {
	Pwd string `json:"pwd" validate:"required"`
}

func (ctrl ctrl) Handler(c echo.Context) error {
	userID := c.Get("UserID").(uint)
	structData, err := vldtutil.ValidatePayload(c, input{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}

	err = ctrl.Srv.ChangePwd(userID, structData.Pwd)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}

	return c.JSON(http.StatusOK, ctype.Dict{})
}

func New(srv SrvProvider) ctrl {
	return ctrl{Srv: srv}
}
