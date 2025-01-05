package app

import (
	"src/common/ctype"
	"src/util/errutil"
	"src/util/localeutil"

	"github.com/nicksnyder/go-i18n/v2/i18n"
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
	localizer := localeutil.Get()
	taskQueryOptions := ctype.QueryOptions{
		Filters: ctype.Dict{"FeatureID": id},
	}
	tasks, err := srv.taskRepo.List(taskQueryOptions)
	if err != nil {
		return nil, err
	}
	if len(tasks) > 0 {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.FeatureInUse,
		})
		return emptyResult, errutil.New("", []string{msg})
	}
	return srv.featureRepo.Delete(id)
}
