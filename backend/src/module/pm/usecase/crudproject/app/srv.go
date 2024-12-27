package app

import (
	"src/common/ctype"
	"src/module/pm"
	"src/module/pm/schema"
)

type Service struct {
	projectRepo         ProjectRepo
	featureRepo         FeatureRepo
	taskFieldRepo       TaskFieldRepo
	taskFieldOptionRepo TaskFieldOptionRepo
}

func New(
	projectRepo ProjectRepo,
	featureRepo FeatureRepo,
	taskFieldRepo TaskFieldRepo,
	taskFieldOptionRepo TaskFieldOptionRepo,
) Service {
	return Service{
		projectRepo,
		featureRepo,
		taskFieldRepo,
		taskFieldOptionRepo,
	}
}

func (srv Service) Create(data ctype.Dict) (*schema.Project, error) {
	project, err := srv.projectRepo.Create(data)
	if err != nil {
		return nil, err
	}

	featureData := ctype.Dict{
		"ProjectID":   project.ID,
		"Title":       "Default",
		"Description": "",
		"Status":      pm.PROJECT_STATUS_ACTIVE,
		"Default":     true,
		"Order":       0,
	}
	_, err = srv.featureRepo.Create(featureData)
	if err != nil {
		return nil, err
	}

	taskFieldData := ctype.Dict{
		"ProjectID":   project.ID,
		"Title":       "Status",
		"Description": "Task status",
		"Type":        pm.TASK_FIELD_TYPE_SELECT,
		"IsStatus":    true,
		"Order":       1,
	}
	taskField, err := srv.taskFieldRepo.Create(taskFieldData)
	if err != nil {
		return nil, err
	}

	taskFieldOptionDataList := []ctype.Dict{
		{
			"TaskFieldID": taskField.ID,
			"Title":       "To Do",
			"Description": "Task is not started",
			"Color":       pm.TASK_FIELD_OPTION_COLOR_GRAY,
			"Order":       1,
		},
		{
			"TaskFieldID": taskField.ID,
			"Title":       "In Progress",
			"Description": "Task is in progress",
			"Color":       pm.TASK_FIELD_OPTION_COLOR_BLUE,
			"Order":       2,
		},
		{
			"TaskFieldID": taskField.ID,
			"Title":       "Done",
			"Description": "Task is completed",
			"Color":       pm.TASK_FIELD_OPTION_COLOR_GREEN,
			"Order":       3,
		},
	}
	for _, taskFieldOptionData := range taskFieldOptionDataList {
		_, err = srv.taskFieldOptionRepo.Create(taskFieldOptionData)
		if err != nil {
			return nil, err
		}
	}

	return project, nil
}

func (srv Service) Update(
	updateOptions ctype.QueryOptions,
	data ctype.Dict,
) (*schema.Project, error) {
	project, err := srv.projectRepo.Update(updateOptions, data)
	if err != nil {
		return nil, err
	}

	featureOptions := ctype.QueryOptions{
		Filters: ctype.Dict{
			"ProjectID": project.ID,
			"Default":   true,
		},
	}
	featureData := ctype.Dict{
		"Title":       project.Title,
		"Description": project.Description,
	}
	_, err = srv.featureRepo.Update(featureOptions, featureData)
	if err != nil {
		return nil, err
	}
	return project, nil
}
