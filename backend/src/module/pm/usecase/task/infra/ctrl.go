package infra

import (
	"fmt"
	"net/http"
	"src/common/ctype"
	"src/util/dbutil"
	"src/util/dictutil"
	"src/util/errutil"
	"src/util/numberutil"
	"src/util/restlistutil"
	"src/util/vldtutil"

	"src/module/abstract/repo/paging"
	"src/module/account/repo/user"
	account "src/module/account/schema"
	"src/module/pm/repo/feature"
	"src/module/pm/repo/project"
	"src/module/pm/repo/task"
	"src/module/pm/repo/taskfield"
	"src/module/pm/repo/taskfieldoption"
	"src/module/pm/repo/taskfieldvalue"
	"src/module/pm/repo/taskuser"
	"src/module/pm/schema"
	"src/module/pm/usecase/task/app"

	"src/client/queueclient"
	"src/queue"

	"github.com/labstack/echo/v4"
)

type Schema = schema.Task

var NewRepo = task.New
var folder = "task/avatar"
var searchableFields = []string{"title", "description"}
var filterableFields = []string{"feature_id"}
var orderableFields = []string{"id", "title", "order"}

func Option(c echo.Context) error {
	tenantID := c.Get("TenantID").(uint)
	projectID := numberutil.StrToUint(c.QueryParam("project_id"), 0)
	projectRepo := project.New(dbutil.Db(nil))
	featureRepo := feature.New(dbutil.Db(nil))
	taskfieldRepo := taskfield.New(dbutil.Db(nil))
	taskfieldoptionRepo := taskfieldoption.New(dbutil.Db(nil))
	userRepo := user.New(dbutil.Db(nil))

	projectQueryOpts := ctype.QueryOpts{
		Filters: ctype.Dict{"ID": projectID},
	}
	project, err := projectRepo.Retrieve(projectQueryOpts)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	featureQueryOpts := ctype.QueryOpts{
		Filters: ctype.Dict{"ProjectID": projectID},
		Order:   "\"order\" ASC",
	}
	features, err := featureRepo.List(featureQueryOpts)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	featureOpts := []ctype.SelectOption[uint]{}
	for _, feature := range features {
		featureOpts = append(featureOpts, ctype.SelectOption[uint]{
			Value: feature.ID,
			Label: feature.Title,
		})
	}

	statusOpts := ctype.QueryOpts{
		Joins: []string{"TaskField"},
		Filters: ctype.Dict{
			"TaskField.ProjectID": projectID,
			"TaskField.IsStatus":  true,
		},
		Order: fmt.Sprintf("%s.order ASC", schema.TaskFieldOption{}.TableName()),
	}
	status, err := taskfieldoptionRepo.List(statusOpts)
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

	taskFieldQueryOption := ctype.QueryOpts{
		Filters: ctype.Dict{
			"ProjectID": projectID,
		},
		Preloads: []string{"TaskFieldOptions"},
	}
	taskFields, err := taskfieldRepo.List(taskFieldQueryOption)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	taskFieldOpts := []TaskFieldOption{}
	for _, taskField := range taskFields {
		options := []ctype.SimpleSelectOption[uint]{}
		for _, option := range taskField.TaskFieldOptions {
			options = append(options, ctype.SimpleSelectOption[uint]{
				Value: option.ID,
				Label: option.Title,
			})
		}
		taskFieldOpts = append(taskFieldOpts, TaskFieldOption{
			Value:       taskField.ID,
			Label:       taskField.Title,
			Description: taskField.Description,
			Type:        taskField.Type,
			IsStatus:    taskField.IsStatus,
			Options:     options,
		})
	}

	userQueryOption := ctype.QueryOpts{
		Filters: ctype.Dict{
			"TenantID": tenantID,
		},
	}
	users, err := userRepo.List(userQueryOption)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	userOpts := []ctype.SelectOption[uint]{}
	for _, u := range users {
		userOpts = append(userOpts, ctype.SelectOption[uint]{
			Value: u.ID,
			Label: u.Email,
			// Label: fmt.Sprintf("%s %s", u.FirstName, u.LastName),
		})
	}

	result := ctype.Dict{
		"project_info": ctype.Dict{
			"id":    project.ID,
			"title": project.Title,
		},
		"feature":    featureOpts,
		"status":     statusOpts,
		"task_field": taskFieldOpts,
		"user":       userOpts,
	}
	return c.JSON(http.StatusOK, result)
}

func List(c echo.Context) error {
	projectID := numberutil.StrToUint(c.QueryParam("project_id"), 0)
	pager := paging.New[Schema, ListOutput](dbutil.Db(nil), ListPres)

	options := restlistutil.GetOptions(c, filterableFields, orderableFields)
	options.Filters["project_id"] = projectID
	options.Preloads = []string{
		"TaskFieldValues.TaskField",
		"TaskFieldValues.TaskFieldOption",
		"TaskUsers.User",
	}
	options.Order = restlistutil.QueryOrder{Field: "order", Direction: "ASC"}
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
		Filters: ctype.Dict{"id": id},
		Preloads: []string{
			"Project",
			"TaskFieldValues.TaskField",
			"TaskFieldValues.TaskFieldOption",
			"TaskUsers",
		},
	}

	result, err := repo.Retrieve(opts)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, DetailPres(*result))
}

func Create(c echo.Context) error {
	user := c.Get("User").(*account.User)
	tenantID := c.Get("TenantID").(uint)

	db := dbutil.Db(nil)
	tx := db.Begin()
	if tx.Error != nil {
		msg := errutil.New("", []string{tx.Error.Error()})
		return c.JSON(http.StatusBadRequest, msg)
	}

	taskRepo := task.New(tx)
	taskFieldRepo := taskfield.New(tx)
	taskFieldOptionRepo := taskfieldoption.New(tx)
	taskFieldValueRepo := taskfieldvalue.New(tx)
	taskUserRepo := taskuser.New(tx)

	srv := app.New(
		taskRepo,
		taskFieldRepo,
		taskFieldOptionRepo,
		taskFieldValueRepo,
		taskUserRepo,
	)

	structData, err := vldtutil.ValidatePayload(c, app.InputData{})
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

	// Publish to queue
	client := queueclient.NewClient()
	client.Publish(queue.LOG_CREATE_TASK, ctype.Dict{
		"tenant_id":      tenantID,
		"project_id":     result.ProjectID,
		"task_id":        result.ID,
		"user_id":        user.ID,
		"user_full_name": user.FullName(),
		"source_id":      result.ID,
		"source_title":   result.Title,
		"value":          dictutil.StructToDict(structData),
	})

	return c.JSON(http.StatusCreated, MutatePres(*result))

}

func Update(c echo.Context) error {
	user := c.Get("User").(*account.User)
	tenantID := c.Get("TenantID").(uint)

	projectID := numberutil.StrToUint(c.QueryParam("project_id"), 0)

	db := dbutil.Db(nil)
	tx := db.Begin()
	if tx.Error != nil {
		msg := errutil.New("", []string{tx.Error.Error()})
		return c.JSON(http.StatusBadRequest, msg)
	}
	taskRepo := task.New(tx)
	taskFieldRepo := taskfield.New(tx)
	taskFieldOptionRepo := taskfieldoption.New(tx)
	taskFieldValueRepo := taskfieldvalue.New(tx)
	taskUserRepo := taskuser.New(tx)

	srv := app.New(
		taskRepo,
		taskFieldRepo,
		taskFieldOptionRepo,
		taskFieldValueRepo,
		taskUserRepo,
	)

	structData, fields, err := vldtutil.ValidateUpdatePayload(
		c, app.InputData{ProjectID: projectID},
	)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	data := vldtutil.GetDictByFields(structData, fields, []string{})
	id := vldtutil.ValidateId(c.Param("id"))
	updateOpts := ctype.QueryOpts{Filters: ctype.Dict{"ID": id}}
	result, err := srv.Update(updateOpts, structData, data)

	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := tx.Commit().Error; err != nil {
		msg := errutil.New("", []string{err.Error()})
		return c.JSON(http.StatusBadRequest, msg)
	}

	// Publish to queue
	client := queueclient.NewClient()
	client.Publish(queue.LOG_EDIT_TASK, ctype.Dict{
		"tenant_id":      tenantID,
		"project_id":     result.ProjectID,
		"task_id":        result.ID,
		"user_id":        user.ID,
		"user_full_name": user.FullName(),
		"source_id":      result.ID,
		"source_title":   result.Title,
		"value":          dictutil.StructToDict(structData),
	})

	return c.JSON(http.StatusOK, MutatePres(*result))
}

func Delete(c echo.Context) error {
	user := c.Get("User").(*account.User)
	tenantID := c.Get("TenantID").(uint)

	repo := NewRepo(dbutil.Db(nil))

	id := vldtutil.ValidateId(c.Param("id"))

	result, err := repo.Retrieve(ctype.QueryOpts{Filters: ctype.Dict{"id": id}})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	ids, err := repo.Delete(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Publish to queue
	client := queueclient.NewClient()
	client.Publish(queue.LOG_DELETE_TASK, ctype.Dict{
		"tenant_id":      tenantID,
		"project_id":     result.ProjectID,
		"task_id":        result.ID,
		"user_id":        user.ID,
		"user_full_name": user.FullName(),
		"source_id":      result.ID,
		"source_title":   result.Title,
		"value":          ctype.Dict{},
	})

	return c.JSON(http.StatusOK, ids)
}
