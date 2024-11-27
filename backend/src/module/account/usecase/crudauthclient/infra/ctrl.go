package infra

import (
	"net/http"
	"src/common/ctype"
	"src/util/dbutil"
	"src/util/restlistutil"
	"src/util/vldtutil"

	"src/module/abstract/repo/paging"
	"src/module/account/repo/authclient"
	"src/module/account/schema"

	"github.com/labstack/echo/v4"
)

type Schema = schema.AuthClient

var NewRepo = authclient.New

var searchableFields = []string{"uid", "description", "partition"}
var filterableFields = []string{}
var orderableFields = []string{"id", "uid"}

func List(c echo.Context) error {
	pager := paging.New[Schema, ListOutput](dbutil.Db(), ListPres)

	options := restlistutil.GetOptions(c, filterableFields, orderableFields)
	listResult, err := pager.Paging(options, searchableFields)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, listResult)
}

func Retrieve(c echo.Context) error {
	cruder := NewRepo(dbutil.Db())

	id := vldtutil.ValidateId(c.Param("id"))
	queryOptions := ctype.QueryOptions{
		Filters: ctype.Dict{"id": id},
	}

	result, err := cruder.Retrieve(queryOptions)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, DetailPres(*result))
}

func Create(c echo.Context) error {
	cruder := NewRepo(dbutil.Db())
	data, err := vldtutil.ValidatePayload(c, InputData{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	result, err := cruder.Create(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, result)

}

func Update(c echo.Context) error {
	cruder := NewRepo(dbutil.Db())

	data, err := vldtutil.ValidateUpdatePayload(c, InputData{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	id := vldtutil.ValidateId(c.Param("id"))
	result, err := cruder.Update(id, data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, result)
}

func Delete(c echo.Context) error {
	cruder := NewRepo(dbutil.Db())

	id := vldtutil.ValidateId(c.Param("id"))
	ids, err := cruder.Delete(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, ids)
}

func DeleteList(c echo.Context) error {
	cruder := NewRepo(dbutil.Db())

	ids := vldtutil.ValidateIds(c.QueryParam("ids"))
	ids, err := cruder.DeleteList(ids)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, ids)
}
