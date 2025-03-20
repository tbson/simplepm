package app

import (
	"src/common/ctype"
	"src/module/account/schema"
)

type TenantRepo interface {
	Create(data ctype.Dict) (*schema.Tenant, error)
	Update(updateOpts ctype.QueryOpts, data ctype.Dict) (*schema.Tenant, error)
}
