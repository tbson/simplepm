package app

import (
	"src/common/ctype"
	"src/module/account/schema"
)

type UserRepo interface {
	Retrieve(opts ctype.QueryOpts) (*schema.User, error)
	Update(updateOpts ctype.QueryOpts, data ctype.Dict) (*schema.User, error)
}

type AuthSrv interface {
	SetPwd(userID uint, pwd string) error
}
