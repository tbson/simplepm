package app

import (
	"src/common/ctype"
	"src/module/account/schema"
)

type CrudUserRepo interface {
	Update(id uint, data ctype.Dict) (*schema.User, error)
}
