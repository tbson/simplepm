package app

import (
	"src/common/ctype"
)

type Service struct {
	featureRepo FeatureRepo
}

func New(featureRepo FeatureRepo) Service {
	return Service{featureRepo}
}

func (srv Service) Reorder(data InputData) ([]OrderInfoItem, error) {
	defaultResult := []OrderInfoItem{}
	orderInfo := data.Items
	projectID := data.ProjectID
	for _, info := range orderInfo {
		updateOptions := ctype.QueryOptions{
			Filters: ctype.Dict{
				"ID":        info.ID,
				"ProjectID": projectID,
			},
		}
		data := ctype.Dict{
			"order": info.Order,
		}
		_, err := srv.featureRepo.Update(updateOptions, data)
		if err != nil {
			return defaultResult, err
		}
	}
	return orderInfo, nil
}
