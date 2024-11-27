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

func (r Repo) ListRoleByIds(ids []uint) ([]schema.Role, error) {
	db := r.client
	var items []schema.Role
	result := db.Where("id IN ?", ids).Find(&items)
	err := result.Error
	if err != nil {
		return items, errutil.NewGormError(err)
	}
	return items, err
}
