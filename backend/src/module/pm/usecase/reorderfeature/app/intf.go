package app

import (
	"src/common/ctype"
	"src/module/pm/schema"
)

type OrderInfoItem struct {
	ID    uint `json:"id"`
	Order int  `json:"order"`
}

type InputData struct {
	ProjectID uint            `json:"project_id" validate:"required"`
	Items     []OrderInfoItem `json:"items" validate:"required"`
}

type FeatureRepo interface {
	Update(queryOptions ctype.QueryOptions, data ctype.Dict) (*schema.Feature, error)
}
