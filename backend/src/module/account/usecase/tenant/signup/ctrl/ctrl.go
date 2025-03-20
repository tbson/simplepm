package ctrl

import (
	"net/http"

	"src/common/ctype"
	"src/util/dbutil"
	"src/util/routeutil"
	"src/util/vldtutil"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type srvProvider interface {
	Signup(
		pemMap ctype.PemMap,
		uid string,
		title string,
		email string,
		mobile *string,
		firstName string,
		lastName string,
		pwd string,

	) error
	WithTx(tx *gorm.DB)
}

type ctrl struct {
	srv      srvProvider
	dbClient *gorm.DB
}

type input struct {
	UID       string  `json:"uid" validate:"required"`
	Title     string  `json:"title" validate:"required"`
	Email     string  `json:"email" validate:"required"`
	Mobile    *string `json:"mobile"`
	FirstName string  `json:"first_name" validate:"required"`
	LastName  string  `json:"last_name" validate:"required"`
	Pwd       string  `json:"pwd" validate:"required"`
}

// Handler godoc
// @Summary Signup
// @Description Signup
// @Tags Tenant
// @Accept json
// @Produce json
// @Param uid body string true "UID"
// @Param title body string true "Title"
// @Param email body string true "Email"
// @Param mobile body string false "Mobile"
// @Param first_name body string true "First Name"
// @Param last_name body string true "Last Name"
// @Param pwd body string true "Password"
// @Success 200 {object} ctype.Dict
// @Router /tenant/signup [post]
func (ctrl ctrl) Handler(c echo.Context) error {
	pemMap := routeutil.GetPemMap()
	data, err := vldtutil.ValidatePayload(c, input{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// for tenant
	uid := data.UID
	title := data.Title

	// for user
	email := data.Email
	mobile := data.Mobile
	firstName := data.FirstName
	lastName := data.LastName
	pwd := data.Pwd

	err = dbutil.WithTx(ctrl.dbClient, func(tx *gorm.DB) error {
		ctrl.srv.WithTx(tx)
		return ctrl.srv.Signup(
			pemMap,
			uid,
			title,
			email,
			mobile,
			firstName,
			lastName,
			pwd,
		)
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, ctype.Dict{})
}

func New(srv srvProvider, dbClient *gorm.DB) ctrl {
	return ctrl{srv: srv, dbClient: dbClient}
}
