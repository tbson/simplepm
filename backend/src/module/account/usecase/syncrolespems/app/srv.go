package app

import (
	"src/common/ctype"
)

type Service struct {
	repo       SynRolesPemsRepo
	roleRepo   RoleRepo
	tenantRepo TenantRepo
}

func New(syncRolesPemsRepo SynRolesPemsRepo, roleRepo RoleRepo, tenantRepo TenantRepo) Service {
	return Service{syncRolesPemsRepo, roleRepo, tenantRepo}
}

func (srv Service) ensureAllTenantRoles() error {

	tenants, err := srv.tenantRepo.List(ctype.QueryOptions{})
	if err != nil {
		return err
	}

	for _, tenant := range tenants {
		srv.roleRepo.EnsureTenantRoles(tenant.ID, tenant.Uid)
	}

	return nil
}

func (srv Service) SyncRolesPems(pemMap ctype.PemMap) error {
	srv.repo.WritePems(pemMap)
	srv.ensureAllTenantRoles()
	srv.repo.EnsureRolesPems(pemMap)
	return nil
}
