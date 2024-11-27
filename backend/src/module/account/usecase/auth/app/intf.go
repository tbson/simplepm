package app

import (
	"context"
	"src/common/ctype"
	"src/module/account/schema"
	"src/util/ssoutil"
	"time"
)

type PemModulesActionsMap map[string][]string

type AuthUserResult struct {
	ID       uint
	Admin    bool
	LockedAt *time.Time
	Sub      *string
}

type AuthClientInfo struct {
	TenantID     uint
	Realm        string
	ClientID     string
	ClientSecret string
}

type UserRepo interface {
	Create(data ctype.Dict) (*schema.User, error)
	Update(id uint, data ctype.Dict) (*schema.User, error)
}

type IamRepo interface {
	GetAuthUrl(realm string, clientId string, state ctype.Dict) string
	GetLogoutUrl(realm string, clientId string) string
	ValidateCallback(
		ctx context.Context,
		realm string,
		clientId string,
		clientSecret string,
		code string,
	) (ssoutil.TokensAndClaims, error)
	GetSub(tokenStr string, realm string) (string, error)
	RefreshToken(
		ctx context.Context,
		realm string,
		refreshToken string,
		clientId string,
		clientSecret string,
	) (ssoutil.TokensAndClaims, error)
	Logout(clientId string, clientSecret string, realm string, refreshToken string) error
}

type AuthRepo interface {
	GetTenantUser(tenantID uint, email string) (AuthUserResult, error)
	GetAuthClientFromTenantUid(tenantUid string) (AuthClientInfo, error)
	GetAuthClientFromSub(tenantUid string) (AuthClientInfo, error)
	GetPemModulesActionsMap(userId uint) (PemModulesActionsMap, error)
}
