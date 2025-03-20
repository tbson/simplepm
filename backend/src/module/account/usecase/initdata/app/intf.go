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
	Update(opts ctype.QueryOpts, data ctype.Dict) (*schema.User, error)
}

type RoleRepo interface {
	EnsureTenantRoles(ID uint, Uid string) error
	EnsureRolesPems(pemMap ctype.PemMap, opts ctype.QueryOpts) error
	Retrieve(opts ctype.QueryOpts) (*schema.Role, error)
}
