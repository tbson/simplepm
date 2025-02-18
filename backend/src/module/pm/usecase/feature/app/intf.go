package app

import (
	"src/common/ctype"
	"src/module/pm/schema"
)

type FeatureRepo interface {
	Delete(id uint) ([]uint, error)
}

type TaskRepo interface {
	List(queryOptions ctype.QueryOptions) ([]schema.Task, error)
}
