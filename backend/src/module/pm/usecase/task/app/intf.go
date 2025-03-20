package app

import (
	"src/common/ctype"
	"src/module/pm/schema"
)

type TaskFieldData struct {
	TaskFieldID uint   `json:"task_field_id" form:"task_field_id" validate:"required"`
	Value       string `json:"value" form:"value" validate:"required"`
}

type TaskUserData struct {
	UserID    uint    `json:"user_id" form:"user_id" validate:"required"`
	GitBranch *string `json:"git_branch" form:"git_branch"`
}

type InputData struct {
	ProjectID   uint            `json:"project_id" form:"project_id" validate:"required"`
	Title       string          `json:"title" form:"title" validate:"required"`
	Description string          `json:"description" form:"description"`
	TaskFields  []TaskFieldData `json:"task_fields" form:"task_fields"`
	TaskUsers   []TaskUserData  `json:"task_users" form:"task_users"`
}

type TaskRepo interface {
	List(opts ctype.QueryOpts) ([]schema.Task, error)
	Create(data ctype.Dict) (*schema.Task, error)
	Update(updateOpts ctype.QueryOpts, data ctype.Dict) (*schema.Task, error)
}

type TaskFieldRepo interface {
	List(opts ctype.QueryOpts) ([]schema.TaskField, error)
	Retrieve(opts ctype.QueryOpts) (*schema.TaskField, error)
}

type TaskFieldOptionRepo interface {
	Retrieve(opts ctype.QueryOpts) (*schema.TaskFieldOption, error)
}

type TaskFieldValueRepo interface {
	List(opts ctype.QueryOpts) ([]schema.TaskFieldValue, error)
	Create(data ctype.Dict) (*schema.TaskFieldValue, error)
	Update(updateOpts ctype.QueryOpts, data ctype.Dict) (*schema.TaskFieldValue, error)
	UpdateOrCreate(opts ctype.QueryOpts, data ctype.Dict) (*schema.TaskFieldValue, error)
	DeleteList(ids []uint) ([]uint, error)
}

type TaskUserRepo interface {
	List(opts ctype.QueryOpts) ([]schema.TaskUser, error)
	Create(data ctype.Dict) (*schema.TaskUser, error)
	DeleteBy(opts ctype.QueryOpts) ([]uint, error)
}
