package app

import (
	"src/common/ctype"
	"src/module/account/schema"
)

type UserRepo interface {
	Create(data ctype.Dict) (*schema.User, error)
	Update(id uint, data ctype.Dict) (*schema.User, error)
}

type CrudUserRepo interface {
	ListRoleByIds(ids []uint) ([]schema.Role, error)
}
