package infra

import (
	"net/http"
	"src/common/ctype"
	"src/util/dbutil"
	"src/util/restlistutil"
	"src/util/vldtutil"

	"src/module/abstract/repo/paging"
	"src/module/account/repo/pem"
	"src/module/account/repo/role"
	"src/module/account/schema"
	"src/module/account/usecase/crudrole/app"

	"github.com/labstack/echo/v4"
)

type Schema = schema.Role

var NewRepo = role.New
var folder = "role/avatar"
var searchableFields = []string{"uid", "title"}
var filterableFields = []string{"tenant_id"}
var orderableFields = []string{"id", "uid"}

func Option(c echo.Context) error {
	admin := c.Get("Admin").(bool)
	pemRepo := pem.New(dbutil.Db())
	queryOptions := ctype.QueryOptions{
		Filters: ctype.Dict{},
		Order:   "module ASC",
	}
	if !admin {
		queryOptions.Filters["admin"] = false
	}
	items, err := pemRepo.List(queryOptions)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	pemOptions := []ctype.SelectOption[uint]{}
	for _, item := range items {
		pemOptions = append(pemOptions, ctype.SelectOption[uint]{
			Value:       item.ID,
			Label:       item.Module,
			Description: item.Title,
		})
	}
	result := ctype.Dict{
		"pem": pemOptions,
	}
	return c.JSON(http.StatusOK, result)
}

func List(c echo.Context) error {
	tenantId := c.Get("TenantID").(uint)
	pager := paging.New[Schema, ListOutput](dbutil.Db(), ListPres)

	options := restlistutil.GetOptions(c, filterableFields, orderableFields)
	options.Filters["tenant_id"] = tenantId
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
		Filters:  ctype.Dict{"id": id},
		Preloads: []string{"Pems"},
	}

	result, err := cruder.Retrieve(queryOptions)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, DetailPres(*result))
}

func Create(c echo.Context) error {
	tenantId := c.Get("TenantID").(uint)
	roleRepo := NewRepo(dbutil.Db())
	crudRoleRepo := New(dbutil.Db())

	srv := app.New(roleRepo, crudRoleRepo)

	data, err := vldtutil.ValidatePayload(c, InputData{TenantID: tenantId})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	result, err := srv.Create(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, result)

}

func Update(c echo.Context) error {
	tenantId := c.Get("TenantID").(uint)
	roleRepo := NewRepo(dbutil.Db())
	crudRoleRepo := New(dbutil.Db())

	srv := app.New(roleRepo, crudRoleRepo)

	id := vldtutil.ValidateId(c.Param("id"))
	data, err := vldtutil.ValidateUpdatePayload(c, InputData{TenantID: tenantId})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	result, err := srv.Update(id, data)

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
