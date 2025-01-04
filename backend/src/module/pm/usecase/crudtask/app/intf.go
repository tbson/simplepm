package app

import (
	"src/common/ctype"
	"src/module/pm/schema"
)

type TaskFieldData struct {
	TaskFieldID uint   `json:"task_field_id" form:"task_field_id" validate:"required"`
	Value       string `json:"value" form:"value" validate:"required"`
}

type InputData struct {
	ProjectID   uint            `json:"project_id" form:"project_id" validate:"required"`
	FeatureID   uint            `json:"feature_id" form:"feature_id"`
	Title       string          `json:"title" form:"title" validate:"required"`
	Description string          `json:"description" form:"description"`
	Order       int             `json:"order" form:"order"`
	TaskFields  []TaskFieldData `json:"task_fields" form:"task_fields"`
}

type TaskRepo interface {
	List(queryOptions ctype.QueryOptions) ([]schema.Task, error)
	Create(data ctype.Dict) (*schema.Task, error)
	Update(updateOptions ctype.QueryOptions, data ctype.Dict) (*schema.Task, error)
}

type TaskFieldRepo interface {
	List(queryOptions ctype.QueryOptions) ([]schema.TaskField, error)
	Retrieve(queryOptions ctype.QueryOptions) (*schema.TaskField, error)
}

type TaskFieldOptionRepo interface {
	Retrieve(queryOptions ctype.QueryOptions) (*schema.TaskFieldOption, error)
}

type TaskFieldValueRepo interface {
	Create(data ctype.Dict) (*schema.TaskFieldValue, error)
	Update(updateOptions ctype.QueryOptions, data ctype.Dict) (*schema.TaskFieldValue, error)
	UpdateOrCreate(queryOptions ctype.QueryOptions, data ctype.Dict) (*schema.TaskFieldValue, error)
}
