package ctrl

import (
	"fmt"
	"net/http"

	"src/common/ctype"
	"src/util/vldtutil"

	"src/util/errutilnew"

	"github.com/labstack/echo/v4"
)

type SrvProvider interface {
	Login(email string, pwd string, tenantID uint) (ctype.Dict, error)
}

type ctrl struct {
	appSrv SrvProvider
}

type input struct {
	Email string `json:"email" validate:"required"`
	Pwd   string `json:"pwd" validate:"required"`
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
	tenantID := c.Get("TenantID").(uint)
	structData, err := vldtutil.ValidatePayload(c, input{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	cookieData, err := ctrl.appSrv.Login(structData.Email, structData.Pwd, tenantID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutilnew.CustomError).Localize())
	}

	// write cookies here
	fmt.Println(cookieData)

	return c.JSON(http.StatusOK, ctype.Dict{})
}

func New(srv SrvProvider) ctrl {
	return ctrl{appSrv: srv}
}
