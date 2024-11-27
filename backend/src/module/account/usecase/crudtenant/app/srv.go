package app

import (
	"src/common/ctype"
	"src/module/account/schema"
)

type Service struct {
	tenantRepo TenantRepo
	roleRepo   RoleRepo
}

func New(tenantRepo TenantRepo, roleRepo RoleRepo) Service {
	return Service{tenantRepo, roleRepo}
}

func (srv Service) Create(data ctype.Dict) (*schema.Tenant, error) {
	tenant, err := srv.tenantRepo.Create(data)
	if err != nil {
		return nil, err
	}

	err = srv.roleRepo.EnsureTenantRoles(tenant.ID, tenant.Uid)
	if err != nil {
		return nil, err
	}

	return tenant, nil
}
