package app

import (
	"src/common/ctype"
	"src/module/account/schema"
)

type RoleRepo interface {
	Create(data ctype.Dict) (*schema.Role, error)
	Update(id uint, data ctype.Dict) (*schema.Role, error)
}

type CrudRoleRepo interface {
	ListPemByIds(ids []uint) ([]schema.Pem, error)
}
