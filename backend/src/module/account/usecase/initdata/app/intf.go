package app

import (
	"src/common/ctype"
	"src/module/account/schema"
)

type AuthClientRepo interface {
	GetOrCreate(queryOptions ctype.QueryOptions, data ctype.Dict) (*schema.AuthClient, error)
}

type TenantRepo interface {
	GetOrCreate(queryOptions ctype.QueryOptions, data ctype.Dict) (*schema.Tenant, error)
}

type UserRepo interface {
	GetOrCreate(queryOptions ctype.QueryOptions, data ctype.Dict) (*schema.User, error)
	Update(id uint, data ctype.Dict) (*schema.User, error)
}

type RoleRepo interface {
	EnsureTenantRoles(ID uint, Uid string) error
	Retrieve(queryOptions ctype.QueryOptions) (*schema.Role, error)
}
