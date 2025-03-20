package app

import (
	"src/common/ctype"
)

type Service struct {
	taskFieldRepo TaskFieldRepo
}

func New(taskFieldRepo TaskFieldRepo) Service {
	return Service{taskFieldRepo}
}

func (srv Service) Reorder(data InputData) ([]OrderInfoItem, error) {
	defaultResult := []OrderInfoItem{}
	orderInfo := data.Items
	projectID := data.ProjectID
	for _, info := range orderInfo {
		updateOpts := ctype.QueryOpts{
			Filters: ctype.Dict{
				"ID":        info.ID,
				"ProjectID": projectID,
			},
		}
		data := ctype.Dict{
			"order": info.Order,
		}
		_, err := srv.taskFieldRepo.Update(updateOpts, data)
		if err != nil {
			return defaultResult, err
		}
	}
	return orderInfo, nil
}
