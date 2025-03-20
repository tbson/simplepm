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

	"src/module/pm/usecase/feature/app"
)

type Schema = schema.Feature

var NewRepo = feature.New
var searchableFields = []string{"title", "description"}
var filterableFields = []string{}
var orderableFields = []string{"id", "title", "order"}

func List(c echo.Context) error {
	projectID := numberutil.StrToUint(c.QueryParam("project_id"), 0)
	pager := paging.New[Schema, ListOutput](dbutil.Db(nil), ListPres)

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
	repo := NewRepo(dbutil.Db(nil))

	id := vldtutil.ValidateId(c.Param("id"))
	opts := ctype.QueryOpts{
		Filters:  ctype.Dict{"id": id},
		Preloads: []string{"Project"},
	}

	result, err := repo.Retrieve(opts)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, DetailPres(*result))
}

func Create(c echo.Context) error {
	projectID := numberutil.StrToUint(c.QueryParam("project_id"), 0)
	repo := NewRepo(dbutil.Db(nil))

	structData, err := vldtutil.ValidatePayload(c, InputData{ProjectID: projectID})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	data := dictutil.StructToDict(structData)

	result, err := repo.Create(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, result)

}

func Update(c echo.Context) error {
	projectID := numberutil.StrToUint(c.QueryParam("project_id"), 0)
	repo := NewRepo(dbutil.Db(nil))

	structData, fields, err := vldtutil.ValidateUpdatePayload(c, InputData{ProjectID: projectID})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	data := vldtutil.GetDictByFields(structData, fields, []string{})
	id := vldtutil.ValidateId(c.Param("id"))
	updateOpts := ctype.QueryOpts{Filters: ctype.Dict{"ID": id}}
	result, err := repo.Update(updateOpts, data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, result)
}

func Delete(c echo.Context) error {
	featureRepo := NewRepo(dbutil.Db(nil))
	taskRepo := task.New(dbutil.Db(nil))

	srv := app.New(featureRepo, taskRepo)

	id := vldtutil.ValidateId(c.Param("id"))
	ids, err := srv.Delete(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, ids)
}
