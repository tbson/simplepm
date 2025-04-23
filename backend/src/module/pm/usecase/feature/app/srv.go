package app

import (
	"src/common/ctype"
	"src/util/errutil"
	"src/util/i18nmsg"
)

type Service struct {
	featureRepo FeatureRepo
	taskRepo    TaskRepo
}

func New(featureRepo FeatureRepo, taskRepo TaskRepo) Service {
	return Service{featureRepo, taskRepo}
}

func (srv Service) Delete(id uint) ([]uint, error) {
	emptyResult := []uint{}
	taskQueryOpts := ctype.QueryOpts{
		Filters: ctype.Dict{"FeatureID": id},
	}
	tasks, err := srv.taskRepo.List(taskQueryOpts)
	if err != nil {
		return nil, err
	}
	if len(tasks) > 0 {
		return emptyResult, errutil.New(i18nmsg.FeatureInUse)
	}
	return srv.featureRepo.Delete(id)
}
