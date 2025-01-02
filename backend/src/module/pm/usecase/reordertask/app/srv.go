package app

import (
	"src/common/ctype"
	"src/util/numberutil"
)

type Service struct {
	taskRepo           TaskRepo
	taskFieldValueRepo TaskFieldValueRepo
}

func New(taskRepo TaskRepo, taskFieldValueRepo TaskFieldValueRepo) Service {
	return Service{taskRepo, taskFieldValueRepo}
}

func (srv Service) updateStatus(taskId uint, status uint) error {
	queryOptions := ctype.QueryOptions{
		Joins: []string{"TaskField"},
		Filters: ctype.Dict{
			"TaskID":             taskId,
			"TaskField.IsStatus": true,
		},
	}
	taskFieldValue, err := srv.taskFieldValueRepo.Retrieve(queryOptions)
	if err != nil {
		return err
	}
	taskFieldValueID := taskFieldValue.ID

	updateOptions := ctype.QueryOptions{
		Filters: ctype.Dict{
			"ID": taskFieldValueID,
		},
	}

	data := ctype.Dict{
		"TaskFieldOptionID": status,
		"Value":             numberutil.UintToStr(status),
	}
	_, err = srv.taskFieldValueRepo.Update(updateOptions, data)
	if err != nil {
		return err
	}
	return nil
}

func (srv Service) Reorder(data InputData) ([]OrderInfoItem, error) {
	defaultResult := []OrderInfoItem{}
	orderInfo := data.Items
	projectID := data.ProjectID
	for _, info := range orderInfo {
		status := info.Status
		updateOptions := ctype.QueryOptions{
			Filters: ctype.Dict{
				"ID":        info.ID,
				"ProjectID": projectID,
			},
		}
		data := ctype.Dict{
			"order": info.Order,
		}
		_, err := srv.taskRepo.Update(updateOptions, data)
		if err != nil {
			return defaultResult, err
		}

		err = srv.updateStatus(info.ID, status)
		if err != nil {
			return defaultResult, err
		}
	}
	return orderInfo, nil
}
