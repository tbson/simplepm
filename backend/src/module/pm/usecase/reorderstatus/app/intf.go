package app

import (
	"src/common/ctype"
	"src/module/pm/schema"
)

type InputData struct {
	ProjectID uint   `json:"project_id" validate:"required"`
	IDs       []uint `json:"ids" validate:"required"`
}

type TaskFieldOptionRepo interface {
	GetTableName() string
	List(opts ctype.QueryOpts) ([]schema.TaskFieldOption, error)
	Update(opts ctype.QueryOpts, data ctype.Dict) (*schema.TaskFieldOption, error)
}
