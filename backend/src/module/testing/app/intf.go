package app

import (
	"src/common/ctype"
	"src/module/account/schema"
)

type TenantRepo interface {
	GetOrCreate(queryOptions ctype.QueryOptions, data ctype.Dict) (*schema.Tenant, error)
}

type UserRepo interface {
	GetOrCreate(queryOptions ctype.QueryOptions, data ctype.Dict) (*schema.User, error)
}
