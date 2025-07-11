package ctrl

import (
	"src/util/presutil"

	"src/util/vldtutil"

	"src/module/config/pres"

	"github.com/labstack/echo/v4"
)

type srvProvider interface {
	DeleteList(ids []string) ([]string, error)
}

type ctrl struct {
	srv srvProvider
}

// go-openapi doc
// @Summary Delete List
// @Description Delete List
// @Tags variable
// @Accept json
// @Produce json
// @Param ids query string true "IDs"
// @Success 200 {object} pres.DeleteResult
// @Failure 400 {object} ctype.Dict
// @Router /config/variable [delete]
func (ctrl ctrl) Handler(c echo.Context) error {
	resp := presutil.New(c)

	ids := vldtutil.ValidateStrIds(c.QueryParam("ids"))

	result, err := ctrl.srv.DeleteList(ids)

	if err != nil {
		return resp.Err(err)
	}

	return resp.Ok(pres.DeletePres(result))
}

func New(srv srvProvider) ctrl {
	return ctrl{srv: srv}
}
