package srv

import (
	"src/common/ctype"
	"src/module/account/schema"

	"gorm.io/gorm"
)

type userProvider interface {
	Create(data ctype.Dict) (*schema.User, error)
	WithTx(*gorm.DB)
}

type tenantProvider interface {
	Create(data ctype.Dict) (*schema.Tenant, error)
	WithTx(*gorm.DB)
}

type roleProvider interface {
	Retrieve(opts ctype.QueryOpts) (*schema.Role, error)
	EnsureTenantRoles(ID uint, Uid string) error
	EnsureRolesPems(pemMap ctype.PemMap, opts ctype.QueryOpts) error
	WithTx(*gorm.DB)
}

type authProvider interface {
	SetPwd(userID uint, pwd string) error
	SendVerifyEmail(userID uint) error
	// WithTx(*gorm.DB)
}

type srv struct {
	userRepo   userProvider
	tenantRepo tenantProvider
	roleRepo   roleProvider
	authSrv    authProvider
}

func New(
	userRepo userProvider,
	tenantRepo tenantProvider,
	roleRepo roleProvider,
	authSrv authProvider,
) srv {
	return srv{userRepo, tenantRepo, roleRepo, authSrv}
}

func (srv srv) WithTx(tx *gorm.DB) {
	srv.userRepo.WithTx(tx)
	srv.tenantRepo.WithTx(tx)
	srv.roleRepo.WithTx(tx)
	// srv.authSrv.WithTx(tx)
}

func (srv srv) Signup(
	pemMap ctype.PemMap,
	uid string,
	title string,
	email string,
	mobile *string,
	firstName string,
	lastName string,
	pwd string,
) error {
	// create tenant
	tenantData := ctype.Dict{
		"Uid":   uid,
		"Title": title,
	}

	tenant, err := srv.tenantRepo.Create(tenantData)
	if err != nil {
		return err
	}

	// ensure tenant roles
	err = srv.roleRepo.EnsureTenantRoles(tenant.ID, tenant.Uid)
	if err != nil {
		return err
	}

	// Sync roles and permissions
	opts := ctype.QueryOpts{
		Filters: ctype.Dict{
			"TenantID": tenant.ID,
		},
	}
	err = srv.roleRepo.EnsureRolesPems(pemMap, opts)
	if err != nil {
		return err
	}

	// get MANAGER role
	roleOpts := ctype.QueryOpts{
		Filters: ctype.Dict{
			"TenantID": tenant.ID,
			"Title":    "MANAGER",
		},
	}
	role, err := srv.roleRepo.Retrieve(roleOpts)
	if err != nil {
		return err
	}

	// create user
	userData := ctype.Dict{
		"TenantID":  tenant.ID,
		"Email":     email,
		"Mobile":    mobile,
		"FirstName": firstName,
		"LastName":  lastName,
		"Roles":     []schema.Role{*role},
	}

	user, err := srv.userRepo.Create(userData)
	if err != nil {
		return err
	}

	// set pwd
	err = srv.authSrv.SetPwd(user.ID, pwd)
	if err != nil {
		return err
	}

	// send verify email
	err = srv.authSrv.SendVerifyEmail(user.ID)
	if err != nil {
		return err
	}

	return nil
}
