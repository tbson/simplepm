package ctrl

import (
	"net/http"

	"src/util/vldtutil"

	"src/util/errutilnew"

	"src/module/account/domain/model"

	"src/module/account/usecase/auth/login/pres/cookie"
	"src/module/account/usecase/auth/login/pres/json"

	"github.com/labstack/echo/v4"
)

type SrvProvider interface {
	Login(email string, pwd string, tenantID uint) (model.LoginResult, error)
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
	next := c.QueryParam("next")
	tenantID := c.Get("TenantID").(uint)
	structData, err := vldtutil.ValidatePayload(c, input{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	loginResult, err := ctrl.appSrv.Login(structData.Email, structData.Pwd, tenantID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutilnew.CustomError).Localize())
	}

	if structData.ClientType == "web" {
		return cookie.LoginPres(c, loginResult, next)
	}

	return json.LoginPres(c, loginResult, next)
}

func New(srv SrvProvider) ctrl {
	return ctrl{appSrv: srv}
}
