package app

import (
	"src/common/ctype"
	"src/module/account/schema"
)

type UserRepo interface {
	Create(data ctype.Dict) (*schema.User, error)
	Update(updateOptions ctype.QueryOptions, data ctype.Dict) (*schema.User, error)
}

type CrudUserRepo interface {
	ListRoleByIds(ids []uint) ([]schema.Role, error)
}
