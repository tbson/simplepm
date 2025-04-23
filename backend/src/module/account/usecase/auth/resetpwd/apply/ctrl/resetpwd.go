package ctrl

import (
	"net/http"

	"src/common/ctype"
	"src/util/errutil"
	"src/util/vldtutil"

	"github.com/labstack/echo/v4"
)

type SrvProvider interface {
	ResetPwd(email string, code string, pwd string, tenantID uint) error
}

type ctrl struct {
	Srv SrvProvider
}

type input struct {
	Email string `json:"email" validate:"required"`
	Code  string `json:"code" validate:"required"`
	Pwd   string `json:"pwd" validate:"required"`
}

func (ctrl ctrl) Handler(c echo.Context) error {
	tenantID := c.Get("TenantID").(uint)
	structData, err := vldtutil.ValidatePayload(c, input{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}

	err = ctrl.Srv.ResetPwd(structData.Email, structData.Code, structData.Pwd, tenantID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}

	return c.JSON(http.StatusOK, ctype.Dict{})
}

func New(srv SrvProvider) ctrl {
	return ctrl{Srv: srv}
}
