package infra

import (
	"net/http"
	"src/common/ctype"
	"src/util/dbutil"
	"src/util/errutil"
	"src/util/restlistutil"
	"src/util/vldtutil"

	"src/module/abstract/repo/paging"
	"src/module/account/repo/authclient"
	"src/module/account/repo/role"
	"src/module/account/repo/tenant"
	"src/module/account/schema"
	"src/module/account/usecase/crudtenant/app"

	"github.com/labstack/echo/v4"
)

type Schema = schema.Tenant

var NewRepo = tenant.New
var folder = "tenant/avatar"
var searchableFields = []string{"uid", "title"}
var filterableFields = []string{}
var orderableFields = []string{"id", "uid"}

func Option(c echo.Context) error {
	authClientRepo := authclient.New(dbutil.Db())
	authClients, err := authClientRepo.List(ctype.QueryOptions{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	var authClientOptions []ctype.SelectOption[uint] = make(
		[]ctype.SelectOption[uint],
		len(authClients),
	)
	for i, authClient := range authClients {
		authClientOptions[i] = ctype.SelectOption[uint]{
			Value:       authClient.ID,
			Label:       authClient.Uid,
			Description: authClient.Partition,
		}
	}

	result := ctype.Dict{
		"auth_client": authClientOptions,
	}

	return c.JSON(http.StatusOK, result)
}

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
	db := dbutil.Db()
	tx := db.Begin()
	if tx.Error != nil {
		msg := errutil.New("", []string{tx.Error.Error()})
		return c.JSON(http.StatusBadRequest, msg)
	}

	cruder := NewRepo(tx)
	roleRepo := role.New(tx)

	srv := app.New(cruder, roleRepo)

	data, err := vldtutil.ValidatePayload(c, InputData{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	data, err = vldtutil.UploadAndUPdatePayload(c, folder, data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	result, err := srv.Create(data)

	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := tx.Commit().Error; err != nil {
		msg := errutil.New("", []string{err.Error()})
		return c.JSON(http.StatusBadRequest, msg)
	}

	return c.JSON(http.StatusCreated, result)

}

func Update(c echo.Context) error {
	cruder := NewRepo(dbutil.Db())

	data, err := vldtutil.ValidateUpdatePayload(c, InputData{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	data, err = vldtutil.UploadAndUPdatePayload(c, folder, data)
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
