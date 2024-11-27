package app

import (
	"src/common/ctype"
	"src/module/account/schema"
)

type Service struct {
	authClientRepo AuthClientRepo
	tenantRepo     TenantRepo
	userRepo       UserRepo
	roleRepo       RoleRepo
	iamRepo        IamRepo
}

func New(
	authClient AuthClientRepo,
	tenantRepo TenantRepo,
	userRepo UserRepo,
	roleRepo RoleRepo,
	iamRepo IamRepo,
) Service {
	return Service{
		authClient,
		tenantRepo,
		userRepo,
		roleRepo,
		iamRepo,
	}
}

func (srv Service) SignupTenant(
	uid string,
	title string,
	email string,
	mobile *string,
	firstName string,
	lastName string,
	password string,
	admin bool,
) error {
	// get default auth client
	authClientOptions := ctype.QueryOptions{
		Filters: ctype.Dict{
			"Default": true,
		},
	}
	authClient, err := srv.authClientRepo.Retrieve(authClientOptions)
	if err != nil {
		return err
	}

	clientID := authClient.Uid
	realm := authClient.Partition

	// create tenant
	tenantData := ctype.Dict{
		"AuthClientID": authClient.ID,
		"Uid":          uid,
		"Title":        title,
	}

	tentant, err := srv.tenantRepo.Create(tenantData)
	if err != nil {
		return err
	}

	// ensure tenant roles
	err = srv.roleRepo.EnsureTenantRoles(tentant.ID, tentant.Uid)
	if err != nil {
		return err
	}

	// get MANAGER role
	roleOptions := ctype.QueryOptions{
		Filters: ctype.Dict{
			"TenantID": tentant.ID,
			"Title":    "MANAGER",
		},
	}
	role, err := srv.roleRepo.Retrieve(roleOptions)
	if err != nil {
		return err
	}

	// get admin access token

	accessToken, err := srv.iamRepo.GetAdminAccessToken()
	if err != nil {
		return err
	}

	// create KeyCloak user
	sub, err := srv.iamRepo.CreateUser(
		accessToken,
		realm,
		email,
		firstName,
		lastName,
		mobile,
	)

	// set password
	err = srv.iamRepo.SetPassword(accessToken, sub, realm, password)
	if err != nil {
		return err
	}

	// create user
	userData := ctype.Dict{
		"TenantID":  tentant.ID,
		"Email":     email,
		"Mobile":    mobile,
		"FirstName": firstName,
		"LastName":  lastName,
		"Roles":     []schema.Role{*role},
	}

	_, err = srv.userRepo.Create(userData)
	if err != nil {
		return err
	}

	// send verify email
	err = srv.iamRepo.SendVerifyEmail(accessToken, clientID, sub, realm)
	if err != nil {
		return err
	}

	return nil
}
