package infra

import (
	"net/http"
	"slices"
	"src/common/ctype"
	"src/util/dbutil"
	"src/util/dictutil"
	"src/util/errutil"
	"src/util/restlistutil"
	"src/util/vldtutil"

	"src/module/abstract/repo/paging"
	"src/module/pm"
	"src/module/pm/repo/feature"
	"src/module/pm/repo/project"
	"src/module/pm/repo/taskfield"
	"src/module/pm/repo/taskfieldoption"
	"src/module/pm/schema"

	"src/module/pm/repo/workspace"

	"github.com/labstack/echo/v4"

	"src/module/pm/usecase/crudproject/app"
)

type Schema = schema.Project

var NewRepo = project.New
var folder = "project/avatar"
var searchableFields = []string{"title", "description"}
var filterableFields = []string{"workspace_id", "layout", "status"}
var orderableFields = []string{"id", "title", "order"}

func Option(c echo.Context) error {
	tenantId := c.Get("TenantID").(uint)
	workspaceRepo := workspace.New(dbutil.Db())
	queryOptions := ctype.QueryOptions{
		Filters: ctype.Dict{"tenant_id": tenantId},
	}
	workspaces, err := workspaceRepo.List(queryOptions)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	workspaceOptions := []ctype.SelectOption[uint]{}
	for _, workspace := range workspaces {
		workspaceOptions = append(workspaceOptions, ctype.SelectOption[uint]{
			Value: workspace.ID,
			Label: workspace.Title,
		})
	}

	result := ctype.Dict{
		"workspace": workspaceOptions,
		"layout":    pm.ProjectLayoutOptions,
		"status":    pm.ProjectStatusOptions,
		"task_field": ctype.Dict{
			"type":  pm.TaskFieldTypeOptions,
			"color": pm.TaskFieldColorOptions,
		},
	}
	return c.JSON(http.StatusOK, result)
}

func List(c echo.Context) error {
	tenantId := c.Get("TenantID").(uint)
	pager := paging.New[Schema, ListOutput](dbutil.Db(), ListPres)

	options := restlistutil.GetOptions(c, filterableFields, orderableFields)
	options.Filters["tenant_id"] = tenantId
	options.Preloads = []string{"Workspace"}
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
	tenantId := c.Get("TenantID").(uint)

	db := dbutil.Db()
	tx := db.Begin()
	if tx.Error != nil {
		msg := errutil.New("", []string{tx.Error.Error()})
		return c.JSON(http.StatusBadRequest, msg)
	}

	projectRepo := project.New(tx)
	featureRepo := feature.New(tx)
	taskFieldRepo := taskfield.New(tx)
	taskFieldOptionRepo := taskfieldoption.New(tx)

	srv := app.New(projectRepo, featureRepo, taskFieldRepo, taskFieldOptionRepo)

	structData, err := vldtutil.ValidatePayload(c, InputData{TenantID: tenantId})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	structData.EnsureGitHost()

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

	return c.JSON(http.StatusCreated, MutatePres(*result))

}

func Update(c echo.Context) error {
	tenantId := c.Get("TenantID").(uint)

	db := dbutil.Db()
	tx := db.Begin()
	if tx.Error != nil {
		msg := errutil.New("", []string{tx.Error.Error()})
		return c.JSON(http.StatusBadRequest, msg)
	}

	projectRepo := project.New(tx)
	featureRepo := feature.New(tx)
	taskFieldRepo := taskfield.New(tx)
	taskFieldOptionRepo := taskfieldoption.New(tx)

	srv := app.New(projectRepo, featureRepo, taskFieldRepo, taskFieldOptionRepo)

	structData, fields, err := vldtutil.ValidateUpdatePayload(c, InputData{TenantID: tenantId})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	fieldModifier := []string{}
	if slices.Contains(fields, "git_repo") {
		structData.EnsureGitHost()
		fieldModifier = append(fieldModifier, "git_host")
	}
	data := vldtutil.GetDictByFields(structData, fields, fieldModifier)
	data, err = vldtutil.UploadAndUPdatePayload(c, folder, data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	id := vldtutil.ValidateId(c.Param("id"))
	updateOptions := ctype.QueryOptions{
		Filters: ctype.Dict{"ID": id},
	}
	result, err := srv.Update(updateOptions, data)

	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := tx.Commit().Error; err != nil {
		msg := errutil.New("", []string{err.Error()})
		return c.JSON(http.StatusBadRequest, msg)
	}

	return c.JSON(http.StatusOK, MutatePres(*result))
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
