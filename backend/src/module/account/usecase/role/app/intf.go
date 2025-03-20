package app

import (
	"src/common/ctype"
	"src/module/account/schema"
)

type RoleRepo interface {
	Create(data ctype.Dict) (*schema.Role, error)
	Update(updateOpts ctype.QueryOpts, data ctype.Dict) (*schema.Role, error)
}

type RoleLocalRepo interface {
	ListPemByIds(ids []uint) ([]schema.Pem, error)
}
