package app

import (
	"src/common/ctype"
	"src/module/pm"
	"src/module/pm/schema"
)

type Service struct {
	projectRepo ProjectRepo
	featureRepo FeatureRepo
}

func New(projectRepo ProjectRepo, featureRepo FeatureRepo) Service {
	return Service{projectRepo, featureRepo}
}

func (srv Service) Create(data ctype.Dict) (*schema.Project, error) {
	project, err := srv.projectRepo.Create(data)
	if err != nil {
		return nil, err
	}

	featureData := ctype.Dict{
		"ProjectID":   project.ID,
		"Title":       project.Title,
		"Description": project.Description,
		"Status":      pm.PROJECT_STATUS_ACTIVE,
		"Default":     true,
		"Order":       0,
	}
	_, err = srv.featureRepo.Create(featureData)
	if err != nil {
		return nil, err
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
