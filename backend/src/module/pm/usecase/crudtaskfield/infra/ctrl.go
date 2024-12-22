package infra

import (
	"net/http"
	"src/common/ctype"
	"src/util/dbutil"
	"src/util/restlistutil"
	"src/util/vldtutil"

	"src/module/abstract/repo/paging"
	"src/module/pm/repo/taskfield"
	"src/module/pm/repo/taskfieldoption"
	"src/module/pm/schema"

	"github.com/labstack/echo/v4"

	"src/module/pm/usecase/crudtaskfield/app"
)

type Schema = schema.TaskField

var NewRepo = taskfield.New
var folder = "taskField/avatar"
var searchableFields = []string{"title", "description"}
var filterableFields = []string{"project_id"}
var orderableFields = []string{"id", "title", "order"}

func List(c echo.Context) error {
	if err := vldtutil.CheckRequiredFilter(c, "project_id"); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	pager := paging.New[Schema, ListOutput](dbutil.Db(), ListPres)

	options := restlistutil.GetOptions(c, filterableFields, orderableFields)
	options.Order = restlistutil.QueryOrder{Field: "order", Direction: "ASC"}

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
		Filters:  ctype.Dict{"id": id},
		Preloads: []string{"TaskFieldOptions"},
	}

	result, err := cruder.Retrieve(queryOptions)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, DetailPres(*result))
}

func Create(c echo.Context) error {
	taskFieldRepo := taskfield.New(dbutil.Db())
	taskFieldOptionRepo := taskfieldoption.New(dbutil.Db())
	srv := app.New(taskFieldRepo, taskFieldOptionRepo)

	structData, err := vldtutil.ValidatePayload(c, app.InputData{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	result, err := srv.Create(structData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, result)

}

func Update(c echo.Context) error {
	taskFieldRepo := taskfield.New(dbutil.Db())
	taskFieldOptionRepo := taskfieldoption.New(dbutil.Db())
	srv := app.New(taskFieldRepo, taskFieldOptionRepo)

	structData, fields, err := vldtutil.ValidateUpdatePayload(c, app.InputData{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	data := vldtutil.GetDictByFields(structData, fields, []string{})
	options := structData.TaskFieldOptions

	id := vldtutil.ValidateId(c.Param("id"))
	updateOptions := ctype.QueryOptions{
		Filters: ctype.Dict{"ID": id},
	}
	result, err := srv.Update(updateOptions, data, options)

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
