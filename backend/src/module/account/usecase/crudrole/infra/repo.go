package infra

import (
	"src/module/account/schema"
	"src/util/errutil"

	"gorm.io/gorm"
)

type Repo struct {
	client *gorm.DB
}

func New(client *gorm.DB) Repo {
	return Repo{client: client}
}

func (r Repo) ListPemByIds(ids []uint) ([]schema.Pem, error) {
	db := r.client
	var items []schema.Pem
	result := db.Where("id IN ?", ids).Find(&items)
	err := result.Error
	if err != nil {
		return items, errutil.NewGormError(err)
	}
	return items, err
}
