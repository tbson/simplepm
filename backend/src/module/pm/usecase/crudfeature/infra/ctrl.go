package infra

import (
	"net/http"
	"src/common/ctype"
	"src/util/dbutil"
	"src/util/dictutil"
	"src/util/numberutil"
	"src/util/restlistutil"
	"src/util/vldtutil"

	"src/module/abstract/repo/paging"
	"src/module/pm/repo/feature"
	"src/module/pm/repo/task"
	"src/module/pm/schema"

	"github.com/labstack/echo/v4"

	"src/module/pm/usecase/crudfeature/app"
)

type Schema = schema.Feature

var NewRepo = feature.New
var searchableFields = []string{"title", "description"}
var filterableFields = []string{}
var orderableFields = []string{"id", "title", "order"}

func List(c echo.Context) error {
	projectID := numberutil.StrToUint(c.QueryParam("project_id"), 0)
	pager := paging.New[Schema, ListOutput](dbutil.Db(), ListPres)

	options := restlistutil.GetOptions(c, filterableFields, orderableFields)
	options.Order = restlistutil.QueryOrder{Field: "order", Direction: "ASC"}
	options.Filters["project_id"] = projectID
	listResult, err := pager.List(options, searchableFields)
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
	projectID := numberutil.StrToUint(c.QueryParam("project_id"), 0)
	cruder := NewRepo(dbutil.Db())

	structData, err := vldtutil.ValidatePayload(c, InputData{ProjectID: projectID})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	data := dictutil.StructToDict(structData)

	result, err := cruder.Create(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, result)

}

func Update(c echo.Context) error {
	projectID := numberutil.StrToUint(c.QueryParam("project_id"), 0)
	cruder := NewRepo(dbutil.Db())

	structData, fields, err := vldtutil.ValidateUpdatePayload(c, InputData{ProjectID: projectID})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	data := vldtutil.GetDictByFields(structData, fields, []string{})
	id := vldtutil.ValidateId(c.Param("id"))
	updateOptions := ctype.QueryOptions{Filters: ctype.Dict{"ID": id}}
	result, err := cruder.Update(updateOptions, data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, result)
}

func Delete(c echo.Context) error {
	featureRepo := NewRepo(dbutil.Db())
	taskRepo := task.New(dbutil.Db())

	srv := app.New(featureRepo, taskRepo)

	id := vldtutil.ValidateId(c.Param("id"))
	ids, err := srv.Delete(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, ids)
}
