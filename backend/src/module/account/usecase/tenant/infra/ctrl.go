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
	"src/module/account/repo/role"
	"src/module/account/repo/tenant"
	"src/module/account/schema"
	"src/module/account/usecase/tenant/app"

	"github.com/labstack/echo/v4"
)

type Schema = schema.Tenant

var NewRepo = tenant.New
var folder = "tenant/avatar"
var searchableFields = []string{"uid", "title"}
var filterableFields = []string{}
var orderableFields = []string{"id", "uid"}

func Option(c echo.Context) error {
	result := ctype.Dict{}

	return c.JSON(http.StatusOK, result)
}

func List(c echo.Context) error {
	pager := paging.New[Schema, ListOutput](dbutil.Db(nil), ListPres)

	options := restlistutil.GetOptions(c, filterableFields, orderableFields)
	listResult, err := pager.Paging(options, searchableFields)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, listResult)
}

func Retrieve(c echo.Context) error {
	user := c.Get("User").(*schema.User)
	repo := NewRepo(dbutil.Db(nil))

	id := vldtutil.ValidateId(c.Param("id"))
	if id == 0 {
		id = user.TenantID
	}
	queryOptions := ctype.QueryOptions{
		Filters:  ctype.Dict{"id": id},
		Preloads: []string{"GitAccounts"},
	}

	result, err := repo.Retrieve(queryOptions)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, DetailPres(*result))
}

func Create(c echo.Context) error {
	db := dbutil.Db(nil)
	tx := db.Begin()
	if tx.Error != nil {
		msg := errutil.New("", []string{tx.Error.Error()})
		return c.JSON(http.StatusBadRequest, msg)
	}

	repo := NewRepo(tx)
	roleRepo := role.New(tx)

	srv := app.New(repo, roleRepo)

	structData, err := vldtutil.ValidatePayload(c, InputData{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	data := dictutil.StructToDict(structData)
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
	repo := NewRepo(dbutil.Db(nil))

	structData, fields, err := vldtutil.ValidateUpdatePayload(c, InputData{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	data := vldtutil.GetDictByFields(structData, fields, []string{})
	data, err = vldtutil.UploadAndUPdatePayload(c, folder, data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

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
