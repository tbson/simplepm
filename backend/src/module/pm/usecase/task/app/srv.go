package app

import (
	"src/common/ctype"
	"src/module/pm"
	"src/module/pm/schema"
	"src/util/dateutil"
	"src/util/dictutil"
	"src/util/numberutil"
	"strings"
)

type Service struct {
	taskRepo            TaskRepo
	taskFieldRepo       TaskFieldRepo
	taskFieldOptionRepo TaskFieldOptionRepo
	taskFieldValueRepo  TaskFieldValueRepo
	taskUserRepo        TaskUserRepo
}

func New(
	taskRepo TaskRepo,
	taskFieldRepo TaskFieldRepo,
	taskFieldOptionRepo TaskFieldOptionRepo,
	taskFieldValueRepo TaskFieldValueRepo,
	taskUserRepo TaskUserRepo,
) Service {
	return Service{
		taskRepo,
		taskFieldRepo,
		taskFieldOptionRepo,
		taskFieldValueRepo,
		taskUserRepo,
	}
}

func (srv Service) Create(structData InputData) (*schema.Task, error) {
	TaskFields := structData.TaskFields
	TaskUsers := structData.TaskUsers
	data := dictutil.StructToDict(structData)
	data["Order"] = srv.getNextTaskOrder(structData.ProjectID)
	delete(data, "TaskFields")
	delete(data, "TaskUsers")
	task, err := srv.taskRepo.Create(data)
	if err != nil {
		return nil, err
	}
	err = srv.syncTaskFieldValues(task.ID, TaskFields)
	if err != nil {
		return nil, err
	}

	err = srv.syncTaskUsers(task.ID, TaskUsers)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (srv Service) Update(
	updateOpts ctype.QueryOpts,
	structData InputData,
	data ctype.Dict,
) (*schema.Task, error) {
	TaskFields := structData.TaskFields
	TaskUsers := structData.TaskUsers
	delete(data, "TaskFields")
	delete(data, "TaskUsers")

	task, err := srv.taskRepo.Update(updateOpts, data)
	if err != nil {
		return nil, err
	}

	err = srv.syncTaskFieldValues(task.ID, TaskFields)
	if err != nil {
		return nil, err
	}

	err = srv.syncTaskUsers(task.ID, TaskUsers)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (srv Service) getNextTaskOrder(ProjectID uint) int {
	opts := ctype.QueryOpts{
		Filters: ctype.Dict{
			"ProjectID": ProjectID,
		},
		Order: "\"order\" DESC",
	}
	tasks, err := srv.taskRepo.List(opts)
	if err != nil {
		return 0
	}
	if len(tasks) == 0 {
		return 1
	}
	return tasks[0].Order + 1
}

func (srv Service) upsertTaskFieldValueText(
	taskID uint,
	taskField schema.TaskField,
	value string,
) error {
	opts := ctype.QueryOpts{
		Filters: ctype.Dict{
			"TaskID":      taskID,
			"TaskFieldID": taskField.ID,
		},
	}
	data := ctype.Dict{
		"TaskID":      taskID,
		"TaskFieldID": taskField.ID,
		"Value":       value,
	}
	_, err := srv.taskFieldValueRepo.UpdateOrCreate(opts, data)
	if err != nil {
		return err
	}
	return nil
}

func (srv Service) upsertTaskFieldValueNumber(
	taskID uint,
	taskField schema.TaskField,
	value string,
) error {
	opts := ctype.QueryOpts{
		Filters: ctype.Dict{
			"TaskID":      taskID,
			"TaskFieldID": taskField.ID,
		},
	}
	numberValue := numberutil.StrToInt(value, 0)
	data := ctype.Dict{
		"TaskID":      taskID,
		"TaskFieldID": taskField.ID,
		"NumberValue": &numberValue,
		"Value":       value,
	}
	_, err := srv.taskFieldValueRepo.UpdateOrCreate(opts, data)
	if err != nil {
		return err
	}
	return nil
}

func (srv Service) upsertTaskFieldValueDate(
	taskID uint,
	taskField schema.TaskField,
	value string,
) error {
	data := ctype.Dict{
		"TaskID":      taskID,
		"TaskFieldID": taskField.ID,
	}
	opts := ctype.QueryOpts{
		Filters: ctype.Dict{
			"TaskID":      taskID,
			"TaskFieldID": taskField.ID,
		},
	}
	dateValue, err := dateutil.StrToDate(value)

	if err == nil {
		data["DateValue"] = &dateValue
		data["Value"] = value
	} else {
		// data["DateValue"] = nil
		data["Value"] = ""
	}
	_, err = srv.taskFieldValueRepo.UpdateOrCreate(opts, data)
	if err != nil {
		return err
	}
	return nil
}

func (srv Service) upsertTaskFieldValueSelect(
	taskID uint,
	taskField schema.TaskField,
	value string,
) error {
	opts := ctype.QueryOpts{
		Filters: ctype.Dict{
			"TaskID":      taskID,
			"TaskFieldID": taskField.ID,
		},
	}
	taskFieldOptionID := numberutil.StrToUint(value, 0)
	data := ctype.Dict{
		"TaskID":      taskID,
		"TaskFieldID": taskField.ID,
		"Value":       value,
	}
	if taskFieldOptionID != 0 {
		data["TaskFieldOptionID"] = &taskFieldOptionID
	}
	_, err := srv.taskFieldValueRepo.UpdateOrCreate(opts, data)
	if err != nil {
		return err
	}
	return nil
}

func (srv Service) createTaskFieldValueSelect(
	taskID uint,
	taskField schema.TaskField,
	value string,
) error {

	taskFieldOptionID := numberutil.StrToUint(value, 0)
	data := ctype.Dict{
		"TaskID":      taskID,
		"TaskFieldID": taskField.ID,
		"Value":       value,
	}
	if taskFieldOptionID != 0 {
		data["TaskFieldOptionID"] = &taskFieldOptionID
	}

	_, err := srv.taskFieldValueRepo.Create(data)
	if err != nil {
		return err
	}
	return nil
}

func (srv Service) upsertTaskFieldValueMultipleSelect(
	taskID uint,
	taskField schema.TaskField,
	value string,
) error {
	values := strings.Split(value, ",")
	for _, value := range values {
		err := srv.createTaskFieldValueSelect(taskID, taskField, value)
		if err != nil {
			return err
		}
	}
	return nil
}

func (srv Service) cleanUpMultipleSelectTaskFieldValues(taskID uint) error {
	opts := ctype.QueryOpts{
		Joins: []string{"TaskField"},
		Filters: ctype.Dict{
			"TaskID":         taskID,
			"TaskField.Type": pm.TASK_FIELD_TYPE_MULTIPLE_SELECT,
		},
	}
	items, err := srv.taskFieldValueRepo.List(opts)
	if err != nil {
		return err
	}
	ids := []uint{}
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	_, err = srv.taskFieldValueRepo.DeleteList(ids)

	return nil
}

func (srv Service) syncTaskFieldValues(
	taskID uint,
	taskFields []TaskFieldData,
) error {
	srv.cleanUpMultipleSelectTaskFieldValues(taskID)
	for _, taskFieldData := range taskFields {
		taskFieldOpts := ctype.QueryOpts{
			Filters:  ctype.Dict{"ID": taskFieldData.TaskFieldID},
			Preloads: []string{"TaskFieldOptions"},
		}
		taskField, err := srv.taskFieldRepo.Retrieve(taskFieldOpts)
		if err != nil {
			return err
		}

		if taskField.Type == pm.TASK_FIELD_TYPE_TEXT {
			err = srv.upsertTaskFieldValueText(taskID, *taskField, taskFieldData.Value)
			if err != nil {
				return err
			}
		} else if taskField.Type == pm.TASK_FIELD_TYPE_NUMBER {
			err = srv.upsertTaskFieldValueNumber(taskID, *taskField, taskFieldData.Value)
			if err != nil {
				return err
			}
		} else if taskField.Type == pm.TASK_FIELD_TYPE_DATE {
			err = srv.upsertTaskFieldValueDate(taskID, *taskField, taskFieldData.Value)
			if err != nil {
				return err
			}
		} else if taskField.Type == pm.TASK_FIELD_TYPE_SELECT {
			err = srv.upsertTaskFieldValueSelect(
				taskID,
				*taskField,
				taskFieldData.Value,
			)
			if err != nil {
				return err
			}
		} else if taskField.Type == pm.TASK_FIELD_TYPE_MULTIPLE_SELECT {
			err = srv.upsertTaskFieldValueMultipleSelect(
				taskID,
				*taskField,
				taskFieldData.Value,
			)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (srv Service) syncTaskUsers(
	taskID uint,
	taskUsers []TaskUserData,
) error {
	opts := ctype.QueryOpts{
		Filters: ctype.Dict{
			"TaskID": taskID,
		},
	}
	_, err := srv.taskUserRepo.DeleteBy(opts)
	if err != nil {
		return err
	}

	for _, taskUserData := range taskUsers {
		data := ctype.Dict{
			"TaskID": taskID,
			"UserID": taskUserData.UserID,
		}
		if taskUserData.GitBranch != nil {
			data["GitBranch"] = taskUserData.GitBranch
		}
		_, err := srv.taskUserRepo.Create(data)
		if err != nil {
			return err
		}
	}
	return nil
}
