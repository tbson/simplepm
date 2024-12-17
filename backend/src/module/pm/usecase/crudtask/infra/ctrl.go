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
)

type Schema = schema.Task

var NewRepo = task.New
var folder = "task/avatar"
var searchableFields = []string{"title", "description"}
var filterableFields = []string{"feature_id"}
var orderableFields = []string{"id", "title", "order"}

func Option(c echo.Context) error {
	projectID := numberutil.StrToUint(c.QueryParam("project_id"), 0)
	featureRepo := feature.New(dbutil.Db())
	queryOptions := ctype.QueryOptions{
		Filters: ctype.Dict{"project_id": projectID},
	}
	features, err := featureRepo.List(queryOptions)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	featureOptions := []ctype.SelectOption[uint]{}
	for _, feature := range features {
		featureOptions = append(featureOptions, ctype.SelectOption[uint]{
			Value: feature.ID,
			Label: feature.Title,
		})
	}

	result := ctype.Dict{
		"feature": featureOptions,
	}
	return c.JSON(http.StatusOK, result)
}

func List(c echo.Context) error {
	projectID := numberutil.StrToUint(c.QueryParam("project_id"), 0)
	pager := paging.New[Schema, ListOutput](dbutil.Db(), ListPres)

	options := restlistutil.GetOptions(c, filterableFields, orderableFields)
	options.Filters["project_id"] = projectID
	options.Preloads = []string{"Feature"}
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

	_, data, err := vldtutil.ValidateUpdatePayload(c, InputData{ProjectID: projectID})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	id := vldtutil.ValidateId(c.Param("id"))
	updateOptions := ctype.QueryOptions{Filters: ctype.Dict{"ID": id}}
	result, err := cruder.Update(updateOptions, data)

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
