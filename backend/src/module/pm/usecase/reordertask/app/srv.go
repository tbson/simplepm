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
	opts := ctype.QueryOpts{
		Joins: []string{"TaskField"},
		Filters: ctype.Dict{
			"TaskID":             taskId,
			"TaskField.IsStatus": true,
		},
	}
	taskFieldValue, err := srv.taskFieldValueRepo.Retrieve(opts)
	if err != nil {
		return err
	}
	taskFieldValueID := taskFieldValue.ID

	updateOpts := ctype.QueryOpts{
		Filters: ctype.Dict{
			"ID": taskFieldValueID,
		},
	}

	data := ctype.Dict{
		"TaskFieldOptionID": status,
		"Value":             numberutil.UintToStr(status),
	}
	_, err = srv.taskFieldValueRepo.Update(updateOpts, data)
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
		updateOpts := ctype.QueryOpts{
			Filters: ctype.Dict{
				"ID":        info.ID,
				"ProjectID": projectID,
			},
		}
		data := ctype.Dict{
			"order": info.Order,
		}
		_, err := srv.taskRepo.Update(updateOpts, data)
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
