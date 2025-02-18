package infra

import (
	"src/common/ctype"

	"gorm.io/gorm"
)

type Repo struct {
	client *gorm.DB
}

func New(client *gorm.DB) Repo {
	return Repo{
		client: client,
	}
}

func (r Repo) SampleMethod(tenantID uint, email string) (ctype.Dict, error) {
	return ctype.Dict{}, nil
}
