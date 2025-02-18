package app

import (
	"src/common/ctype"
	"src/module/account/schema"
)

type UserRepo interface {
	Update(updateOptions ctype.QueryOptions, data ctype.Dict) (*schema.User, error)
}
