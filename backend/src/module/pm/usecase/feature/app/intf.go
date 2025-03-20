package app

import (
	"src/common/ctype"
	"src/module/pm/schema"
)

type FeatureRepo interface {
	Delete(id uint) ([]uint, error)
}

type TaskRepo interface {
	List(opts ctype.QueryOpts) ([]schema.Task, error)
}
