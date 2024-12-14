package app

import (
	"src/common/ctype"
	"src/module/pm/schema"
	"src/util/dictutil"
)

type Service struct {
	taskFieldRepo       TaskFieldRepo
	taskFieldOptionRepo TaskFieldOptionRepo
}

func New(taskFieldRepo TaskFieldRepo, taskFieldOptionRepo TaskFieldOptionRepo) Service {
	return Service{taskFieldRepo, taskFieldOptionRepo}
}

func (srv Service) syncOptions(taskFieldID uint, options []FeTaskFieldOption) error {
	for _, optionStruct := range options {
		status := optionStruct.FeStatus
		id := optionStruct.ID

		data := dictutil.StructToDict(optionStruct)
		data["TaskFieldID"] = taskFieldID
		delete(data, "ID")
		delete(data, "FeStatus")

		if status == FE_OPTION_CREATED {
			_, err := srv.taskFieldOptionRepo.Create(data)
			if err != nil {
				return err
			}
		}
		if status == FE_OPTION_UPDATED {
			queryOptions := ctype.QueryOptions{
				Filters: ctype.Dict{"ID": id},
			}
			_, err := srv.taskFieldOptionRepo.Update(queryOptions, data)
			if err != nil {
				return err
			}
		}
		if status == FE_OPTION_DELETED {
			_, err := srv.taskFieldOptionRepo.Delete(id)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (srv Service) Create(structData InputData) (*schema.TaskField, error) {
	defaultResult := schema.TaskField{}
	options := structData.TaskFieldOptions

	data := dictutil.StructToDict(structData)
	delete(data, "TaskFieldOptions")

	result, err := srv.taskFieldRepo.Create(data)
	if err != nil {
		return &defaultResult, err
	}
	srv.syncOptions(result.ID, options)
	return result, nil
}

func (srv Service) Update(
	queryOptions ctype.QueryOptions,
	data ctype.Dict,
	options []FeTaskFieldOption,
) (*schema.TaskField, error) {
	defaultResult := schema.TaskField{}
	delete(data, "TaskFieldOptions")
	result, err := srv.taskFieldRepo.Update(queryOptions, data)

	if err != nil {
		return &defaultResult, err
	}

	srv.syncOptions(result.ID, options)
	return result, nil
}
