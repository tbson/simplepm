package infra

import (
	"net/http"
	"src/common/ctype"
	"src/util/dbutil"
	"src/util/dictutil"
	"src/util/restlistutil"
	"src/util/vldtutil"

	"src/module/abstract/repo/paging"
	"src/module/config"
	"src/module/config/repo/variable"
	"src/module/config/schema"

	"github.com/labstack/echo/v4"
)

type Schema = schema.Variable

var NewRepo = variable.New

var searchableFields = []string{"key", "value", "description"}
var filterableFields = []string{"data_type"}
var orderableFields = []string{"id", "key"}

func Option(c echo.Context) error {
	result := ctype.Dict{
		"data_type": config.VariableDataTypeOptions,
	}
	return c.JSON(http.StatusOK, result)
}

// List godoc
//
//	@Summary		Get list of variables
//	@Description	Get list of variables with filtering, sorting and paging
//	@Tags			config
//	@Accept			json
//	@Produce		json
//	@Param			q	query	string		false	"Search string"
//	@Param			page	query	int		false	"Page number"
//	@Param			order	query	int		false	"Order by id, key"
//	@Param			data_type	query	string	false	"Filter by data type"
//	@Success		200	{object}	restlistutil.ListRestfulResult[schema.Variable]
//	@Failure		404	{object}	map[string]interface{}
//	@Router			/api/v1/config/variable/ [get]
func List(c echo.Context) error {
	pager := paging.New[Schema, ListOutput](dbutil.Db(nil), ListPres)

	if err := vldtutil.CheckRequiredFilter(c, "tenant_id"); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	options := restlistutil.GetOptions(c, filterableFields, orderableFields)
	listResult, err := pager.Paging(options, searchableFields)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, listResult)
}

func Retrieve(c echo.Context) error {
	repo := NewRepo(dbutil.Db(nil))

	id := vldtutil.ValidateId(c.Param("id"))
	queryOptions := ctype.QueryOptions{
		Filters: ctype.Dict{"id": id},
	}

	result, err := repo.Retrieve(queryOptions)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, DetailPres(*result))
}

func Create(c echo.Context) error {
	repo := NewRepo(dbutil.Db(nil))
	structData, err := vldtutil.ValidatePayload(c, InputData{})
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
	repo := NewRepo(dbutil.Db(nil))

	structData, fields, err := vldtutil.ValidateUpdatePayload(c, InputData{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	data := vldtutil.GetDictByFields(structData, fields, []string{})
	id := vldtutil.ValidateId(c.Param("id"))
	updateOptions := ctype.QueryOptions{Filters: ctype.Dict{"ID": id}}
	result, err := repo.Update(updateOptions, data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, result)
}

func Delete(c echo.Context) error {
	repo := NewRepo(dbutil.Db(nil))

	id := vldtutil.ValidateId(c.Param("id"))
	ids, err := repo.Delete(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, ids)
}

func DeleteList(c echo.Context) error {
	repo := NewRepo(dbutil.Db(nil))

	ids := vldtutil.ValidateIds(c.QueryParam("ids"))
	ids, err := repo.DeleteList(ids)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, ids)
}
