package ctrl

import (
	"src/util/presutil"

	"src/util/vldtutil"

	"src/module/account/domain/model"

	"src/module/account/usecase/auth/login/pres"

	"github.com/labstack/echo/v4"
)

type SrvProvider interface {
	Login(email string, pwd string, tenantID uint) (model.AuthResult, error)
}

type ctrl struct {
	appSrv SrvProvider
}

type input struct {
	Email      string `json:"email" validate:"required"`
	Pwd        string `json:"pwd" validate:"required"`
	ClientType string `json:"client_type" validate:"required,oneof=web app"`
}

// Login godoc
// @Summary Login
// @Description Login
// @Tags auth
// @Accept json
// @Produce json
// @Param email body string true "Email"
// @Param pwd body string true "Password"
// @Success 200 {object} ctype.Dict
// @Failure 400 {object} ctype.Dict
// @Router /account/auth/login [post]
func (ctrl ctrl) Handler(c echo.Context) error {
	resp := presutil.New(c)
	next := c.QueryParam("next")
	tenantID := c.Get("TenantID").(uint)
	structData, err := vldtutil.ValidatePayload(c, input{})
	if err != nil {
		return resp.Err(err)
	}

	loginResult, err := ctrl.appSrv.Login(structData.Email, structData.Pwd, tenantID)
	if err != nil {
		return resp.Err(err)
	}
	return pres.Login(c, loginResult, next, structData.ClientType)
}

func New(srv SrvProvider) ctrl {
	return ctrl{appSrv: srv}
}
