package app

import (
	"src/common/ctype"
	"src/module/account/schema"
)

type UserRepo interface {
	Update(updateOpts ctype.QueryOpts, data ctype.Dict) (*schema.User, error)
}
