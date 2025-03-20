package app

import (
	"src/common/ctype"
	"src/module/pm/schema"
)

type OrderInfoItem struct {
	ID     uint `json:"id"`
	Status uint `json:"status"`
	Order  int  `json:"order"`
}

type InputData struct {
	ProjectID uint            `json:"project_id" validate:"required"`
	Items     []OrderInfoItem `json:"items" validate:"required"`
}

type TaskRepo interface {
	Update(opts ctype.QueryOpts, data ctype.Dict) (*schema.Task, error)
}

type TaskFieldValueRepo interface {
	Retrieve(opts ctype.QueryOpts) (*schema.TaskFieldValue, error)
	Update(opts ctype.QueryOpts, data ctype.Dict) (*schema.TaskFieldValue, error)
}
