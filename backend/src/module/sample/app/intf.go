package app

import (
	"src/common/ctype"
	"src/module/account/schema"
)

type TenantRepo interface {
	Create(data ctype.Dict) (*schema.Tenant, error)
	Update(updateOptions ctype.QueryOptions, data ctype.Dict) (*schema.Tenant, error)
}
