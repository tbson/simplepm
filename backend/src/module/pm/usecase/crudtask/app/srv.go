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
}

func New(
	taskRepo TaskRepo,
	taskFieldRepo TaskFieldRepo,
	taskFieldOptionRepo TaskFieldOptionRepo,
	taskFieldValueRepo TaskFieldValueRepo,
) Service {
	return Service{
		taskRepo,
		taskFieldRepo,
		taskFieldOptionRepo,
		taskFieldValueRepo,
	}
}

func (srv Service) upsertTaskFieldValueText(
	taskID uint,
	taskField schema.TaskField,
	value string,
) error {
	queryOptions := ctype.QueryOptions{
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
	_, err := srv.taskFieldValueRepo.UpdateOrCreate(queryOptions, data)
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
	queryOptions := ctype.QueryOptions{
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
	_, err := srv.taskFieldValueRepo.UpdateOrCreate(queryOptions, data)
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
	queryOptions := ctype.QueryOptions{
		Filters: ctype.Dict{
			"TaskID":      taskID,
			"TaskFieldID": taskField.ID,
		},
	}
	dateValue, err := dateutil.StrToDate(value)
	data := ctype.Dict{
		"TaskID":      taskID,
		"TaskFieldID": taskField.ID,
	}
	if err == nil {
		data["DateValue"] = &dateValue
		data["Value"] = value
	}
	_, err = srv.taskFieldValueRepo.UpdateOrCreate(queryOptions, data)
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
	queryOptions := ctype.QueryOptions{
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
	_, err := srv.taskFieldValueRepo.UpdateOrCreate(queryOptions, data)
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
		err := srv.upsertTaskFieldValueSelect(taskID, taskField, value)
		if err != nil {
			return err
		}
	}
	return nil
}

func (srv Service) upsertTaskFieldValues(
	taskID uint,
	taskFields []TaskFieldData,
) error {
	for _, taskFieldData := range taskFields {
		taskFieldQueryOptions := ctype.QueryOptions{
			Filters:  ctype.Dict{"ID": taskFieldData.TaskFieldID},
			Preloads: []string{"TaskFieldOptions"},
		}
		taskField, err := srv.taskFieldRepo.Retrieve(taskFieldQueryOptions)
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
			err = srv.upsertTaskFieldValueSelect(taskID, *taskField, taskFieldData.Value)
			if err != nil {
				return err
			}
		} else if taskField.Type == pm.TASK_FIELD_TYPE_MULTIPLE_SELECT {
			err = srv.upsertTaskFieldValueMultipleSelect(taskID, *taskField, taskFieldData.Value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (srv Service) Create(structData InputData) (*schema.Task, error) {
	TaskFields := structData.TaskFields
	data := dictutil.StructToDict(structData)
	delete(data, "TaskFields")
	task, err := srv.taskRepo.Create(data)
	if err != nil {
		return nil, err
	}
	err = srv.upsertTaskFieldValues(task.ID, TaskFields)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (srv Service) Update(
	updateOptions ctype.QueryOptions,
	structData InputData,
	data ctype.Dict,
) (*schema.Task, error) {
	TaskFields := structData.TaskFields
	delete(data, "TaskFields")
	task, err := srv.taskRepo.Update(updateOptions, data)
	if err != nil {
		return nil, err
	}
	err = srv.upsertTaskFieldValues(task.ID, TaskFields)
	if err != nil {
		return nil, err
	}

	return task, nil
}
