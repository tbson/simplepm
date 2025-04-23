package infra

import (
	"net/http"
	"src/common/ctype"
	"src/util/dbutil"
	"src/util/dictutil"
	"src/util/errutil"
	"src/util/restlistutil"
	"src/util/vldtutil"

	"src/module/abstract/repo/paging"
	"src/module/document/repo/doc"
	"src/module/document/schema"

	"github.com/labstack/echo/v4"
)

type Schema = schema.Doc

var NewRepo = doc.New
var folder = "doc/avatar"
var searchableFields = []string{"title", "description"}
var filterableFields = []string{}
var orderableFields = []string{"id", "title", "order"}

func List(c echo.Context) error {
	taskID := vldtutil.ValidateId(c.QueryParam("task_id"))
	pager := paging.New[Schema, ListOutput](dbutil.Db(nil), ListPres)

	options := restlistutil.GetOptions(c, filterableFields, orderableFields)
	options.Filters["task_id"] = taskID
	listResult, err := pager.Paging(options, searchableFields)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}

	return c.JSON(http.StatusOK, listResult)
}

func Retrieve(c echo.Context) error {
	repo := NewRepo(dbutil.Db(nil))

	id := vldtutil.ValidateId(c.Param("id"))
	opts := ctype.QueryOpts{
		Filters: ctype.Dict{"id": id},
	}

	result, err := repo.Retrieve(opts)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, DetailPres(*result))
}

func Create(c echo.Context) error {
	userID := c.Get("UserID").(uint)
	repo := NewRepo(dbutil.Db(nil))
	structData, err := vldtutil.ValidatePayload(c, InputData{UserID: userID})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}

	data := dictutil.StructToDict(structData)
	data, err = vldtutil.UploadAndUPdatePayload(c, folder, data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}

	result, err := repo.Create(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}

	return c.JSON(http.StatusCreated, result)

}

func Update(c echo.Context) error {
	repo := NewRepo(dbutil.Db(nil))

	structData, fields, err := vldtutil.ValidateUpdatePayload(c, InputData{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}

	data := vldtutil.GetDictByFields(structData, fields, []string{})
	data, err = vldtutil.UploadAndUPdatePayload(c, folder, data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}

	id := vldtutil.ValidateId(c.Param("id"))
	updateOpts := ctype.QueryOpts{Filters: ctype.Dict{"ID": id}}
	result, err := repo.Update(updateOpts, data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}

	return c.JSON(http.StatusOK, result)
}

func Delete(c echo.Context) error {
	repo := NewRepo(dbutil.Db(nil))

	id := vldtutil.ValidateId(c.Param("id"))
	ids, err := repo.Delete(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}

	return c.JSON(http.StatusOK, ids)
}
