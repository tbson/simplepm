package app

import (
	"src/common/ctype"
	"src/module/account/schema"
)

type Service struct {
	tenantRepo TenantRepo
}

func New(tenantRepo TenantRepo) Service {
	return Service{tenantRepo}
}

func (srv Service) Create(data ctype.Dict) (*schema.Tenant, error) {
	tenant, err := srv.tenantRepo.Create(data)
	if err != nil {
		return nil, err
	}
	return tenant, nil
}

func (srv Service) Update(updateOpts ctype.QueryOpts, data ctype.Dict) (*schema.Tenant, error) {
	tenant, err := srv.tenantRepo.Update(updateOpts, data)
	if err != nil {
		return nil, err
	}
	return tenant, nil
}
