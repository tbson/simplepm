package app

import (
	"src/common/ctype"
	"src/module/pm/schema"
)

type ProjectRepo interface {
	Create(data ctype.Dict) (*schema.Project, error)
	Update(updateOpts ctype.QueryOpts, data ctype.Dict) (*schema.Project, error)
}

type FeatureRepo interface {
	Create(data ctype.Dict) (*schema.Feature, error)
	Update(updateOpts ctype.QueryOpts, data ctype.Dict) (*schema.Feature, error)
}

type TaskFieldRepo interface {
	Create(data ctype.Dict) (*schema.TaskField, error)
}

type TaskFieldOptionRepo interface {
	Create(data ctype.Dict) (*schema.TaskFieldOption, error)
}
