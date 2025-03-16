package ctrl

import (
	"net/http"

	"src/common/ctype"
	"src/util/vldtutil"

	"github.com/labstack/echo/v4"
)

type updateInput struct {
	Key         string `json:"key" form:"key" validate:"required"`
	Value       string `json:"value" form:"value"`
	Description string `json:"description" form:"description"`
	DataType    string `json:"data_type" form:"data_type" validate:"required,oneof=STRING INTEGER FLOAT BOOLEAN DATE DATETIME"`
}

type updateSrvProvider[S any] interface {
	Update(queryOptions ctype.QueryOptions, data ctype.Dict) (*S, error)
}

type updateCtrl[S any] struct {
	srv updateSrvProvider[S]
}

func (ctrl updateCtrl[S]) Handler(c echo.Context) error {
	id := vldtutil.ValidateId(c.Param("id"))

	structData, fields, err := vldtutil.ValidateUpdatePayload(c, updateInput{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	data := vldtutil.GetDictByFields(structData, fields, []string{})
	updateOptions := ctype.QueryOptions{Filters: ctype.Dict{"ID": id}}
	result, err := ctrl.srv.Update(updateOptions, data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, result)
}

func NewUpdate[S any](srv updateSrvProvider[S]) updateCtrl[S] {
	return updateCtrl[S]{srv: srv}
}
