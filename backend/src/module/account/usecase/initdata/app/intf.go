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
	Update(queryOptions ctype.QueryOptions, data ctype.Dict) (*schema.User, error)
}

type RoleRepo interface {
	EnsureTenantRoles(ID uint, Uid string) error
	EnsureRolesPems(pemMap ctype.PemMap, queryOptions ctype.QueryOptions) error
	Retrieve(queryOptions ctype.QueryOptions) (*schema.Role, error)
}
