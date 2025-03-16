package ctrl

import (
	"net/http"

	"src/common/ctype"
	"src/util/dictutil"
	"src/util/vldtutil"

	"github.com/labstack/echo/v4"
)

type createInput struct {
	Key         string `json:"key" form:"key" validate:"required"`
	Value       string `json:"value" form:"value"`
	Description string `json:"description" form:"description"`
	DataType    string `json:"data_type" form:"data_type" validate:"required,oneof=STRING INTEGER FLOAT BOOLEAN DATE DATETIME"`
}

type createSrvProvider[S any] interface {
	Create(data ctype.Dict) (*S, error)
}

type createCtrl[S any] struct {
	srv createSrvProvider[S]
}

func (ctrl createCtrl[S]) Handler(c echo.Context) error {
	structData, err := vldtutil.ValidatePayload(c, createInput{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	data := dictutil.StructToDict(structData)

	result, err := ctrl.srv.Create(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, result)
}

func NewCreate[S any](srv createSrvProvider[S]) createCtrl[S] {
	return createCtrl[S]{srv: srv}
}
