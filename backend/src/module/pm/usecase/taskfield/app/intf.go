package app

import (
	"src/common/ctype"
	"src/module/pm/schema"
)

const FE_OPTION_EXISTING = "EXISTING"
const FE_OPTION_CREATED = "CREATED"
const FE_OPTION_UPDATED = "UPDATED"
const FE_OPTION_DELETED = "DELETED"

type FeTaskFieldOption struct {
	ID          uint   `json:"id"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Color       string `json:"color" validate:"required"`
	Order       int    `json:"order" validate:"required"`
	FeStatus    string `json:"fe_status" validate:"required"`
}

type InputData struct {
	ProjectID        uint                `json:"project_id" form:"project_id" validate:"required"`
	Title            string              `json:"title" form:"title" validate:"required"`
	Description      string              `json:"description" form:"description"`
	Type             string              `json:"type" form:"type" validate:"required"`
	Order            int                 `json:"order" form:"order"`
	TaskFieldOptions []FeTaskFieldOption `json:"task_field_options"`
}

type TaskFieldRepo interface {
	Create(data ctype.Dict) (*schema.TaskField, error)
	Update(queryOptions ctype.QueryOptions, data ctype.Dict) (*schema.TaskField, error)
}

type TaskFieldOptionRepo interface {
	Create(data ctype.Dict) (*schema.TaskFieldOption, error)
	Update(queryOptions ctype.QueryOptions, data ctype.Dict) (*schema.TaskFieldOption, error)
	Delete(id uint) ([]uint, error)
}
