package infra

import (
	"src/common/ctype"
	"src/module/pm"
	"src/module/pm/schema"
	"strings"
)

type Feature struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Color       string `json:"color"`
	Order       int    `json:"order"`
}

type TaskField struct {
	TaskFieldID uint   `json:"task_field_id"`
	Type        string `json:"type"`
	Value       string `json:"value"`
}

type Status struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

type ListOutput struct {
	ID          uint        `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Order       int         `json:"order"`
	Status      Status      `json:"status"`
	Feature     Feature     `json:"feature"`
	TaskFields  []TaskField `json:"task_fields"`
}

type TaskFieldOption struct {
	Value       uint                             `json:"value"`
	Label       string                           `json:"label"`
	Description string                           `json:"description"`
	Type        string                           `json:"type"`
	IsStatus    bool                             `json:"is_status"`
	Options     []ctype.SimpleSelectOption[uint] `json:"options"`
}

type DetailOutput = schema.Task

func presItem(item schema.Task) ListOutput {
	result := ListOutput{
		ID:          item.ID,
		Title:       item.Title,
		Description: item.Description,
		Order:       item.Order,
	}

	feature := Feature{
		ID:          item.Feature.ID,
		Title:       item.Feature.Title,
		Description: item.Feature.Description,
		Status:      item.Feature.Status,
		Color:       item.Feature.Color,
		Order:       item.Feature.Order,
	}
	result.Feature = feature

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
