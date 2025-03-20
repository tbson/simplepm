package app

import (
	"src/common/ctype"
	"src/module/account/schema"
)

type TenantRepo interface {
	GetOrCreate(opts ctype.QueryOpts, data ctype.Dict) (*schema.Tenant, error)
}

type UserRepo interface {
	GetOrCreate(opts ctype.QueryOpts, data ctype.Dict) (*schema.User, error)
}
