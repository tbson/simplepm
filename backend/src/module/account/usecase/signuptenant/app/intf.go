package app

import (
	"src/common/ctype"
	"src/module/account/schema"
)

type AuthClientRepo interface {
	Retrieve(queryOptions ctype.QueryOptions) (*schema.AuthClient, error)
}

type TenantRepo interface {
	Retrieve(queryOptions ctype.QueryOptions) (*schema.Tenant, error)
	Create(data ctype.Dict) (*schema.Tenant, error)
}

type RoleRepo interface {
	Retrieve(queryOptions ctype.QueryOptions) (*schema.Role, error)
	EnsureTenantRoles(ID uint, Uid string) error
	EnsureRolesPems(pemMap ctype.PemMap, queryOptions ctype.QueryOptions) error
}

type UserRepo interface {
	Create(data ctype.Dict) (*schema.User, error)
}

type IamRepo interface {
	GetAdminAccessToken() (string, error)
	CreateUser(
		accessToken string,
		realm string,
		email string,
		firstName string,
		lastName string,
		mobile *string,
	) (string, error)
	SetPassword(
		accessToken string,
		sub string,
		realm string,
		password string,
	) error
	SendVerifyEmail(
		accessToken string,
		clientID string,
		sub string,
		realm string,
	) error
}
