package infra

import (
	"net/http"
	"src/common/ctype"
	"src/util/dbutil"
	"src/util/restlistutil"
	"src/util/vldtutil"

	"src/module/abstract/repo/paging"
	"src/module/account/repo/role"
	"src/module/account/repo/user"
	"src/module/account/schema"
	"src/module/account/usecase/cruduser/app"

	"github.com/labstack/echo/v4"
)

type Schema = schema.User

var NewRepo = user.New

var searchableFields = []string{"uid", "description", "partition"}
var filterableFields = []string{}
var orderableFields = []string{"id", "uid"}

func Option(c echo.Context) error {
	tenantId := c.Get("TenantID").(uint)
	roleRepo := role.New(dbutil.Db())
	queryOptions := ctype.QueryOptions{
		Filters: ctype.Dict{"tenant_id": tenantId},
		Order:   "title ASC",
	}
	items, err := roleRepo.List(queryOptions)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	roleOptions := []ctype.SelectOption[uint]{}
	for _, item := range items {
		roleOptions = append(roleOptions, ctype.SelectOption[uint]{
			Value: item.ID,
			Label: item.Title,
		})
	}
	result := ctype.Dict{
		"role": roleOptions,
	}
	return c.JSON(http.StatusOK, result)
}

func List(c echo.Context) error {
	tenantId := c.Get("TenantID").(uint)
	pager := paging.New[Schema, ListOutput](dbutil.Db(), ListPres)

	options := restlistutil.GetOptions(c, filterableFields, orderableFields)
	options.Filters["tenant_id"] = tenantId
	options.Preloads = []string{"Roles"}
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
		Preloads: []string{"Roles"},
	}

	result, err := cruder.Retrieve(queryOptions)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, DetailPres(*result))
}

func Create(c echo.Context) error {
	tenantId := c.Get("TenantID").(uint)
	cruder := NewRepo(dbutil.Db())
	data, err := vldtutil.ValidatePayload(c, InputData{TenantID: tenantId})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	result, err := cruder.Create(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, MutatePres(*result))

}

func Update(c echo.Context) error {
	tenantId := c.Get("TenantID").(uint)
	userRepo := NewRepo(dbutil.Db())
	crudUserRepo := New(dbutil.Db())

	srv := app.New(userRepo, crudUserRepo)

	data, err := vldtutil.ValidateUpdatePayload(c, InputData{TenantID: tenantId})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	id := vldtutil.ValidateId(c.Param("id"))
	result, err := srv.Update(id, data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, MutatePres(result))
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
