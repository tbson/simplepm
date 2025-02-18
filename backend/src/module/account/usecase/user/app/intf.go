package app

import (
	"src/common/ctype"
	"src/module/account/schema"
)

type UserRepo interface {
	Create(data ctype.Dict) (*schema.User, error)
	Update(updateOptions ctype.QueryOptions, data ctype.Dict) (*schema.User, error)
}

type UserLocalRepo interface {
	ListRoleByIds(ids []uint) ([]schema.Role, error)
}
