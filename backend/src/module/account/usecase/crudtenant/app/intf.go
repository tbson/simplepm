package app

import (
	"src/common/ctype"
	"src/module/account/schema"
)

type TenantRepo interface {
	Create(data ctype.Dict) (*schema.Tenant, error)
}

type RoleRepo interface {
	EnsureTenantRoles(tenantID uint, tenantUid string) error
}
