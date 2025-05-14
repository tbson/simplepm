package ctrl

import (
	"src/util/presutil"

	"src/common/ctype"
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
	resp := presutil.New(c)
	userID := c.Get("UserID").(uint)
	structData, err := vldtutil.ValidatePayload(c, input{})
	if err != nil {
		return resp.Err(err)
	}

	err = ctrl.Srv.ChangePwd(userID, structData.Pwd)
	if err != nil {
		return resp.Err(err)
	}

	return resp.Ok(ctype.Dict{})
}

func New(srv SrvProvider) ctrl {
	return ctrl{Srv: srv}
}
