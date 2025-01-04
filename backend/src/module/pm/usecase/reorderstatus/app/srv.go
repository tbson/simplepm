package app

import (
	"fmt"
	"src/common/ctype"
)

type Service struct {
	taskFieldOptionRepo TaskFieldOptionRepo
}

func New(taskFieldOptionRepo TaskFieldOptionRepo) Service {
	return Service{taskFieldOptionRepo}
}

func (srv Service) Reorder(data InputData) ([]uint, error) {
	defaultResult := []uint{}
	ids := data.IDs
	projectID := data.ProjectID
	for index, id := range ids {
		order := index + 1
		updateOptions := ctype.QueryOptions{
			Joins: []string{"TaskField"},
			Filters: ctype.Dict{
				fmt.Sprintf("%s.ID", srv.taskFieldOptionRepo.GetTableName()): id,
				"TaskField.ProjectID": projectID,
			},
		}
		data := ctype.Dict{
			"order": order,
		}
		_, err := srv.taskFieldOptionRepo.Update(updateOptions, data)
		if err != nil {
			return defaultResult, err
		}
	}
	return defaultResult, nil
}
