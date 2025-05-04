package ctrl

import (
	"net/http"

	"src/util/errutil"
	"src/util/vldtutil"

	"src/module/config/pres"

	"github.com/labstack/echo/v4"
)

type srvProvider interface {
	Delete(id uint) ([]uint, error)
}

type ctrl struct {
	srv srvProvider
}

// go-openapi doc
// @Summary Delete
// @Description Delete
// @Tags variable
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} pres.DeleteResult
// @Failure 400 {object} ctype.Dict
// @Router /config/variable/{id} [delete]
func (ctrl ctrl) Handler(c echo.Context) error {
	id := vldtutil.ValidateId(c.Param("id"))

	result, err := ctrl.srv.Delete(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}

	return c.JSON(http.StatusOK, pres.DeletePres(result))
}

func New(srv srvProvider) ctrl {
	return ctrl{srv: srv}
}
