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
	List(queryOptions ctype.QueryOptions) ([]schema.TaskFieldOption, error)
	Update(queryOptions ctype.QueryOptions, data ctype.Dict) (*schema.TaskFieldOption, error)
}
