package ctrl

import (
	"net/http"

	"src/util/errutil"
	"src/util/restlistutil"

	"src/module/config/schema"

	"src/module/config/pres"

	"github.com/labstack/echo/v4"
)

type srvProvider interface {
	Paging(
		options restlistutil.ListOptions,
		searchFields []string,
	) (restlistutil.ListRestfulResult[schema.Variable], error)
}

type ctrl struct {
	srv srvProvider
}

var searchableFields = []string{"key", "value", "description"}
var filterableFields = []string{"data_type"}
var orderableFields = []string{"id", "key"}

// go-openapi doc
// @Summary List
// @Description List
// @Tags variable
// @Accept json
// @Produce json
// @Param tenant_id query string true "Tenant ID"
// @Param key query string false "Key"
// @Param value query string false "Value"
// @Param description query string false "Description"
// @Param data_type query string false "Data Type"
// @Param order query string false "Order"
// @Param page query string false "Page"
// @Param page_size query string false "Page Size"
// @Success 200 {object} pres.PageResult
// @Failure 400 {object} ctype.Dict
// @Router /config/variable [get]
func (ctrl ctrl) Handler(c echo.Context) error {
	pagingOptions := restlistutil.GetOptions(c, filterableFields, orderableFields)

	result, err := ctrl.srv.Paging(pagingOptions, searchableFields)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}

	return c.JSON(http.StatusOK, pres.PagePres(result))
}

func New(srv srvProvider) ctrl {
	return ctrl{srv: srv}
}
