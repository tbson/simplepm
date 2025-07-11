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
	"src/module/pm"
	"src/module/pm/repo/feature"
	"src/module/pm/repo/project"
	"src/module/pm/repo/taskfield"
	"src/module/pm/repo/taskfieldoption"
	"src/module/pm/schema"

	"src/module/account/repo/gitrepo"
	"src/module/pm/repo/workspace"

	"github.com/labstack/echo/v4"

	"src/module/pm/usecase/project/app"
)

type Schema = schema.Project

var NewRepo = project.New
var folder = "project/avatar"
var searchableFields = []string{"title", "description"}
var filterableFields = []string{"workspace_id", "layout", "status"}
var orderableFields = []string{"id", "title", "order"}

func Option(c echo.Context) error {
	tenantId := c.Get("TenantID").(uint)
	workspaceRepo := workspace.New(dbutil.Db(nil))
	gitRepoRepo := gitrepo.New(dbutil.Db(nil))
	opts := ctype.QueryOpts{
		Filters: ctype.Dict{"tenant_id": tenantId},
	}
	workspaces, err := workspaceRepo.List(opts)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}
	workspaceOptions := []ctype.SelectOption[uint]{}
	for _, workspace := range workspaces {
		workspaceOptions = append(workspaceOptions, ctype.SelectOption[uint]{
			Value: workspace.ID,
			Label: workspace.Title,
		})
	}

	gitRepoQueryOpts := ctype.QueryOpts{
		Joins:   []string{"GitAccount"},
		Filters: ctype.Dict{"GitAccount.TenantID": tenantId},
	}

	gitRepos, err := gitRepoRepo.List(gitRepoQueryOpts)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}
	gitRepoData := []ctype.SimpleSelectOption[string]{}

	for _, gitRepo := range gitRepos {
		gitRepoData = append(gitRepoData, ctype.SimpleSelectOption[string]{
			Value: gitRepo.Uid,
			Label: gitRepo.Uid,
		})
	}

	result := ctype.Dict{
		"workspace": workspaceOptions,
		"layout":    pm.ProjectLayoutOptions,
		"status":    pm.ProjectStatusOptions,
		"task_field": ctype.Dict{
			"type": pm.TaskFieldTypeOptions,
		},
		"git_repo": gitRepoData,
	}
	return c.JSON(http.StatusOK, result)
}

func Bookmark(c echo.Context) error {
	tenantId := c.Get("TenantID").(uint)
	repo := NewRepo(dbutil.Db(nil))

	opts := ctype.QueryOpts{
		Filters: ctype.Dict{"TenantID": tenantId},
	}
	result, err := repo.List(opts)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}
	return c.JSON(http.StatusOK, ListBookmarkPres(result))
}

func List(c echo.Context) error {
	tenantId := c.Get("TenantID").(uint)
	pager := paging.New[Schema, ListOutput](dbutil.Db(nil), ListPres)

	options := restlistutil.GetOptions(c, filterableFields, orderableFields)
	options.Filters["tenant_id"] = tenantId
	options.Preloads = []string{"Workspace"}
	listResult, err := pager.Paging(options, searchableFields)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}

	return c.JSON(http.StatusOK, listResult)
}

func Retrieve(c echo.Context) error {
	repo := NewRepo(dbutil.Db(nil))

	id := vldtutil.ValidateId(c.Param("id"))
	opts := ctype.QueryOpts{
		Filters: ctype.Dict{"id": id},
	}

	result, err := repo.Retrieve(opts)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, DetailPres(*result))
}

func Create(c echo.Context) error {
	ctx := c.Request().Context()
	tenantId := c.Get("TenantID").(uint)

	db := dbutil.Db(&ctx)
	tx := db.Begin()
	if tx.Error != nil {
		msg := errutil.NewRaw(tx.Error.Error())
		return c.JSON(http.StatusBadRequest, msg)
	}

	projectRepo := project.New(tx)
	featureRepo := feature.New(tx)
	taskFieldRepo := taskfield.New(tx)
	taskFieldOptionRepo := taskfieldoption.New(tx)

	srv := app.New(projectRepo, featureRepo, taskFieldRepo, taskFieldOptionRepo)

	structData, err := vldtutil.ValidatePayload(c, InputData{TenantID: tenantId})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}

	data := dictutil.StructToDict(structData)
	data, err = vldtutil.UploadAndUPdatePayload(c, folder, data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}

	result, err := srv.Create(data)
	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}

	if err := tx.Commit().Error; err != nil {
		msg := errutil.NewRaw(err.Error())
		return c.JSON(http.StatusBadRequest, msg)
	}

	return c.JSON(http.StatusCreated, MutatePres(*result))

}

func Update(c echo.Context) error {
	ctx := c.Request().Context()
	tenantId := c.Get("TenantID").(uint)

	db := dbutil.Db(&ctx)
	tx := db.Begin()
	if tx.Error != nil {
		msg := errutil.NewRaw(tx.Error.Error())
		return c.JSON(http.StatusBadRequest, msg)
	}

	projectRepo := project.New(tx)
	featureRepo := feature.New(tx)
	taskFieldRepo := taskfield.New(tx)
	taskFieldOptionRepo := taskfieldoption.New(tx)

	srv := app.New(projectRepo, featureRepo, taskFieldRepo, taskFieldOptionRepo)

	structData, fields, err := vldtutil.ValidateUpdatePayload(c, InputData{TenantID: tenantId})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}
	fieldModifier := []string{}

	data := vldtutil.GetDictByFields(structData, fields, fieldModifier)
	data, err = vldtutil.UploadAndUPdatePayload(c, folder, data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}

	id := vldtutil.ValidateId(c.Param("id"))
	updateOpts := ctype.QueryOpts{
		Filters: ctype.Dict{"ID": id},
	}
	result, err := srv.Update(updateOpts, data)

	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}

	if err := tx.Commit().Error; err != nil {
		msg := errutil.NewRaw(err.Error())
		return c.JSON(http.StatusBadRequest, msg)
	}

	return c.JSON(http.StatusOK, MutatePres(*result))
}

func Delete(c echo.Context) error {
	ctx := c.Request().Context()
	repo := NewRepo(dbutil.Db(&ctx))

	id := vldtutil.ValidateId(c.Param("id"))
	ids, err := repo.Delete(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}

	return c.JSON(http.StatusOK, ids)
}

func DeleteList(c echo.Context) error {
	ctx := c.Request().Context()
	repo := NewRepo(dbutil.Db(&ctx))

	ids := vldtutil.ValidateIds(c.QueryParam("ids"))
	ids, err := repo.DeleteList(ids)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
	}

	return c.JSON(http.StatusOK, ids)
}
