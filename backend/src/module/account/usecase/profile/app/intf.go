package app

import (
	"src/common/ctype"
	"src/module/account/schema"
)

type UserRepo interface {
	Retrieve(opts ctype.QueryOptions) (*schema.User, error)
	Update(updateOptions ctype.QueryOptions, data ctype.Dict) (*schema.User, error)
}

type AuthSrv interface {
	SetPwd(userID uint, pwd string) error
}
