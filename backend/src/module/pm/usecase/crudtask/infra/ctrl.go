package infra

import (
	"fmt"
	"net/http"
	"src/common/ctype"
	"src/util/dbutil"
	"src/util/errutil"
	"src/util/numberutil"
	"src/util/restlistutil"
	"src/util/vldtutil"

	"src/module/abstract/repo/paging"
	"src/module/pm/repo/feature"
	"src/module/pm/repo/task"
	"src/module/pm/repo/taskfield"
	"src/module/pm/repo/taskfieldoption"
	"src/module/pm/repo/taskfieldvalue"
	"src/module/pm/schema"
	"src/module/pm/usecase/crudtask/app"

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
	taskfieldRepo := taskfield.New(dbutil.Db())
	taskfieldoptionRepo := taskfieldoption.New(dbutil.Db())

	featureQueryOptions := ctype.QueryOptions{
		Filters: ctype.Dict{"ProjectID": projectID},
	}
	features, err := featureRepo.List(featureQueryOptions)
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

	statusQueryOption := ctype.QueryOptions{
		Joins: []string{"TaskField"},
		Filters: ctype.Dict{
			"TaskField.ProjectID": projectID,
			"TaskField.IsStatus":  true,
		},
		Order: fmt.Sprintf("%s.order ASC", schema.TaskFieldOption{}.TableName()),
	}
	status, err := taskfieldoptionRepo.List(statusQueryOption)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	statusOptions := []ctype.SelectOption[uint]{}
	for _, s := range status {
		statusOptions = append(statusOptions, ctype.SelectOption[uint]{
			Value: s.ID,
			Label: s.Title,
		})
	}

	taskFieldQueryOption := ctype.QueryOptions{
		Filters: ctype.Dict{
			"ProjectID": projectID,
		},
		Preloads: []string{"TaskFieldOptions"},
	}
	taskFields, err := taskfieldRepo.List(taskFieldQueryOption)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	taskFieldOptions := []TaskFieldOption{}
	for _, taskField := range taskFields {
		options := []ctype.SimpleSelectOption[uint]{}
		for _, option := range taskField.TaskFieldOptions {
			options = append(options, ctype.SimpleSelectOption[uint]{
				Value: option.ID,
				Label: option.Title,
			})
		}
		taskFieldOptions = append(taskFieldOptions, TaskFieldOption{
			Value:       taskField.ID,
			Label:       taskField.Title,
			Description: taskField.Description,
			Type:        taskField.Type,
			IsStatus:    taskField.IsStatus,
			Options:     options,
		})
	}

	result := ctype.Dict{
		"feature":    featureOptions,
		"status":     statusOptions,
		"task_field": taskFieldOptions,
	}
	return c.JSON(http.StatusOK, result)
}

func List(c echo.Context) error {
	projectID := numberutil.StrToUint(c.QueryParam("project_id"), 0)
	pager := paging.New[Schema, ListOutput](dbutil.Db(), ListPres)

	options := restlistutil.GetOptions(c, filterableFields, orderableFields)
	options.Filters["project_id"] = projectID
	options.Preloads = []string{
		"Feature",
		"TaskFieldValues.TaskField",
		"TaskFieldValues.TaskFieldOption",
	}
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
		Filters: ctype.Dict{"id": id},
		Preloads: []string{
			"Feature",
			"TaskFieldValues.TaskField",
			"TaskFieldValues.TaskFieldOption",
		},
	}

	result, err := cruder.Retrieve(queryOptions)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, DetailPres(*result))
}

func Create(c echo.Context) error {
	projectID := numberutil.StrToUint(c.QueryParam("project_id"), 0)

	db := dbutil.Db()
	tx := db.Begin()
	if tx.Error != nil {
		msg := errutil.New("", []string{tx.Error.Error()})
		return c.JSON(http.StatusBadRequest, msg)
	}

	taskRepo := task.New(tx)
	taskFieldRepo := taskfield.New(tx)
	taskFieldOptionRepo := taskfieldoption.New(tx)
	taskFieldValueRepo := taskfieldvalue.New(tx)

	srv := app.New(taskRepo, taskFieldRepo, taskFieldOptionRepo, taskFieldValueRepo)

	structData, err := vldtutil.ValidatePayload(c, app.InputData{ProjectID: projectID})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	result, err := srv.Create(structData)
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
	projectID := numberutil.StrToUint(c.QueryParam("project_id"), 0)

	db := dbutil.Db()
	tx := db.Begin()
	if tx.Error != nil {
		msg := errutil.New("", []string{tx.Error.Error()})
		return c.JSON(http.StatusBadRequest, msg)
	}
	fmt.Println("case 1")
	taskRepo := task.New(tx)
	taskFieldRepo := taskfield.New(tx)
	taskFieldOptionRepo := taskfieldoption.New(tx)
	taskFieldValueRepo := taskfieldvalue.New(tx)

	fmt.Println("case 2")
	srv := app.New(taskRepo, taskFieldRepo, taskFieldOptionRepo, taskFieldValueRepo)

	fmt.Println("case 3")
	structData, fields, err := vldtutil.ValidateUpdatePayload(
		c, app.InputData{ProjectID: projectID},
	)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	fmt.Println("case 4")
	data := vldtutil.GetDictByFields(structData, fields, []string{})
	id := vldtutil.ValidateId(c.Param("id"))
	updateOptions := ctype.QueryOptions{Filters: ctype.Dict{"ID": id}}
	result, err := srv.Update(updateOptions, structData, data)

	fmt.Println("case 5")
	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusBadRequest, err)
	}

	fmt.Println("case 6")
	if err := tx.Commit().Error; err != nil {
		msg := errutil.New("", []string{err.Error()})
		return c.JSON(http.StatusBadRequest, msg)
	}

	fmt.Println("case 7")
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
