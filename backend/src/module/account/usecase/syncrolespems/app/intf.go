package app

import (
	"src/common/ctype"
	"src/module/account/schema"
)

type SynRolesPemsRepo interface {
	WritePems(pemMap ctype.PemMap) error
}

type RoleRepo interface {
	EnsureTenantRoles(ID uint, Uid string) error
	EnsureRolesPems(pemMap ctype.PemMap, opts ctype.QueryOpts) error
}

type TenantRepo interface {
	List(ctype.QueryOpts) ([]schema.Tenant, error)
}
