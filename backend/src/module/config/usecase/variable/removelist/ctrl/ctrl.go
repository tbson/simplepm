package ctrl

import (
	"net/http"

	"src/util/vldtutil"

	"src/module/config/pres"

	"github.com/labstack/echo/v4"
)

type srvProvider interface {
	DeleteList(ids []uint) ([]uint, error)
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
	ids := vldtutil.ValidateIds(c.QueryParam("ids"))

	result, err := ctrl.srv.DeleteList(ids)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, pres.DeletePres(result))
}

func New(srv srvProvider) ctrl {
	return ctrl{srv: srv}
}
