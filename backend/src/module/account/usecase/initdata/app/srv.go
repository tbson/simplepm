package app

import (
	"src/common/ctype"
	"src/common/setting"
	"src/module/account/schema"
)

type Service struct {
	tenantRepo TenantRepo
	userRepo   UserRepo
	roleRepo   RoleRepo
}

func New(
	tenantRepo TenantRepo,
	userRepo UserRepo,
	roleRepo RoleRepo,
) Service {
	return Service{tenantRepo, userRepo, roleRepo}
}

func (srv Service) InitData(pemMap ctype.PemMap) error {
	// Init tenant
	opts := ctype.QueryOpts{
		Filters: ctype.Dict{
			"Uid": setting.ADMIN_TEANT_UID(),
		},
	}
	tenantData := ctype.Dict{
		"Uid":   setting.ADMIN_TEANT_UID(),
		"Title": setting.ADMIN_TEANT_TITLE(),
	}
	tenant, err := srv.tenantRepo.GetOrCreate(opts, tenantData)
	if err != nil {
		return err
	}

	// Sync roles for tenant
	err = srv.roleRepo.EnsureTenantRoles(tenant.ID, tenant.Uid)
	if err != nil {
		return err
	}

	// Sync roles and permissions
	opts = ctype.QueryOpts{
		Filters: ctype.Dict{
			"TenantID": tenant.ID,
		},
	}
	err = srv.roleRepo.EnsureRolesPems(pemMap, opts)
	if err != nil {
		return err
	}

	// Assign ADMIN role to user
	opts = ctype.QueryOpts{
		Filters: ctype.Dict{
			"tenant_id": tenant.ID,
			"title":     "ADMIN",
		},
	}
	adminRole, err := srv.roleRepo.Retrieve(opts)
	if err != nil {
		return err
	}

	// Init user
	email := setting.DEFAULT_ADMIN_EMAIL()
	opts = ctype.QueryOpts{
		Filters: ctype.Dict{
			"email": email,
		},
	}
	userData := ctype.Dict{
		"TenantID":   tenant.ID,
		"ExternalID": email,
		"Email":      email,
		"FirstName":  "Admin",
		"LastName":   "Admin",
		"Roles":      []schema.Role{*adminRole},
	}
	user, err := srv.userRepo.GetOrCreate(opts, userData)
	if err != nil {
		return err
	}

	// Update admin attribute
	userData = ctype.Dict{
		"Admin": true,
	}
	updateOpts := ctype.QueryOpts{Filters: ctype.Dict{"ID": user.ID}}
	_, err = srv.userRepo.Update(updateOpts, userData)
	return nil
}
