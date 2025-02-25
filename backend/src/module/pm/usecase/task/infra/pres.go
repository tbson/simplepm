package infra

import (
	"src/common/ctype"
	"src/module/pm"
	"src/module/pm/repo/task"
	"src/module/pm/schema"
	"src/util/dbutil"
	"strings"
)

type Project struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Order       int    `json:"order"`
}

type TaskField struct {
	TaskFieldID uint   `json:"task_field_id"`
	Type        string `json:"type"`
	Value       string `json:"value"`
}

type TaskUser struct {
	ID        uint    `json:"id"`
	UserID    uint    `json:"user_id"`
	Avatar    string  `json:"avatar"`
	GitBranch *string `json:"git_branch"`
}

type Status struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

type TaskFieldOption struct {
	Value       uint                             `json:"value"`
	Label       string                           `json:"label"`
	Description string                           `json:"description"`
	Type        string                           `json:"type"`
	IsStatus    bool                             `json:"is_status"`
	Options     []ctype.SimpleSelectOption[uint] `json:"options"`
}

type ListOutput struct {
	ID          uint        `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Order       int         `json:"order"`
	Status      Status      `json:"status"`
	Project     Project     `json:"project"`
	TaskFields  []TaskField `json:"task_fields"`
	TaskUsers   []TaskUser  `json:"task_users"`
}

type DetailOutput = schema.Task

func presItem(item schema.Task) ListOutput {
	result := ListOutput{
		ID:          item.ID,
		Title:       item.Title,
		Description: item.Description,
		Order:       item.Order,
	}

	project := Project{
		ID:          item.Project.ID,
		Title:       item.Project.Title,
		Description: item.Project.Description,
		Status:      item.Project.Status,
		Order:       item.Project.Order,
	}
	result.Project = project

	taskFields := []TaskField{}
	multipleSelectFieldMap := map[uint][]string{}
	for _, taskFieldValue := range item.TaskFieldValues {
		taskField := taskFieldValue.TaskField
		if taskField.Type == pm.TASK_FIELD_TYPE_MULTIPLE_SELECT {
			if _, ok := multipleSelectFieldMap[taskField.ID]; !ok {
				multipleSelectFieldMap[taskField.ID] = []string{}
			}
			multipleSelectFieldMap[taskField.ID] = append(
				multipleSelectFieldMap[taskField.ID], taskFieldValue.Value,
			)
		} else {
			taskFields = append(taskFields, TaskField{
				TaskFieldID: taskField.ID,
				Type:        taskField.Type,
				Value:       taskFieldValue.Value,
			})
		}
		if taskField.IsStatus {
			result.Status = Status{
				ID:    taskFieldValue.TaskFieldOption.ID,
				Title: taskFieldValue.TaskFieldOption.Title,
			}
		}
	}
	for taskFieldID, values := range multipleSelectFieldMap {
		taskFields = append(taskFields, TaskField{
			TaskFieldID: taskFieldID,
			Type:        pm.TASK_FIELD_TYPE_MULTIPLE_SELECT,
			Value:       strings.Join(values, ","),
		})
	}
	result.TaskFields = taskFields

	taskUsers := []TaskUser{}
	for _, taskUser := range item.TaskUsers {
		taskUsers = append(taskUsers, TaskUser{
			ID:        taskUser.ID,
			UserID:    taskUser.UserID,
			GitBranch: taskUser.GitBranch,
			Avatar:    taskUser.User.Avatar,
		})
	}
	result.TaskUsers = taskUsers

	return result
}

func ListPres(items []schema.Task) []ListOutput {
	result := make([]ListOutput, 0)
	for _, item := range items {
		result = append(result, presItem(item))
	}
	return result
}

func DetailPres(item schema.Task) ListOutput {
	return presItem(item)
}

func MutatePres(item schema.Task) ListOutput {
	taskRepo := task.New(dbutil.Db(nil))
	queryOptions := ctype.QueryOptions{
		Filters: ctype.Dict{
			"id": item.ID,
		},
		Preloads: []string{
			"Project",
			"TaskFieldValues.TaskField",
			"TaskFieldValues.TaskFieldOption",
		},
	}
	task, _ := taskRepo.Retrieve(queryOptions)
	return presItem(*task)
}
