package infra

import (
	"slices"
	"src/common/authtype"
	"src/common/ctype"
	"src/module/account/repo/tenant"
	"src/module/account/repo/user"
	"src/util/errutil"
	"src/util/localeutil"
	"src/util/stringutil"

	"src/module/account/usecase/auth/app"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"gorm.io/gorm"
)

type Repo struct {
	client *gorm.DB
}

func New(client *gorm.DB) Repo {
	return Repo{
		client: client,
	}
}

func (r Repo) GetTenantUser(
	tenantID uint,
	email string,
) (authtype.AuthUserInfo, error) {
	repo := user.New(r.client)
	queryOptions := ctype.QueryOptions{
		Filters: ctype.Dict{
			"tenant_id": tenantID,
			"email":     email,
		},
		Preloads: []string{"Tenant"},
	}
	user, err := repo.Retrieve(queryOptions)
	if err != nil {
		return authtype.AuthUserInfo{}, err
	}
	profileType := "user"
	if user.Admin {
		profileType = "admin"
	}
	result := authtype.AuthUserInfo{
		ID:          user.ID,
		TenantID:    user.TenantID,
		TenantUid:   user.Tenant.Uid,
		Sub:         user.Sub,
		Admin:       user.Admin,
		ProfileType: profileType,
		LockedAt:    user.LockedAt,
		Email:       user.Email,
		FistName:    user.FirstName,
		LastName:    user.LastName,
		Mobile:      user.Mobile,
		Avatar:      user.Avatar,
	}
	return result, err
}

func (r Repo) GetAuthClientFromTenantUid(tenantUid string) (app.AuthClientInfo, error) {
	repo := tenant.New(r.client)
	queryOptions := ctype.QueryOptions{
		Filters:  ctype.Dict{"uid": tenantUid},
		Preloads: []string{"AuthClient"},
	}
	tenant, err := repo.Retrieve(queryOptions)
	if err != nil {
		return app.AuthClientInfo{}, err
	}

	return app.AuthClientInfo{
		TenantID:     tenant.ID,
		Realm:        tenant.AuthClient.Partition,
		ClientID:     tenant.AuthClient.Uid,
		ClientSecret: tenant.AuthClient.Secret,
	}, nil
}

func (r Repo) GetPemModulesActionsMap(userId uint) (app.PemModulesActionsMap, error) {
	repo := user.New(r.client)

	queryOptions := ctype.QueryOptions{
		Filters: ctype.Dict{"id": userId},
		Preloads: []string{
			"Roles.Pems",
		},
	}
	user, err := repo.Retrieve(queryOptions)
	if err != nil {
		return nil, err
	}

	result := make(app.PemModulesActionsMap)
	for _, role := range user.Roles {
		for _, pem := range role.Pems {
			module := stringutil.ToSnakeCase(pem.Module)
			action := stringutil.ToSnakeCase(pem.Action)
			if _, ok := result[module]; !ok {
				result[module] = make([]string, 0)
			}
			if !slices.Contains(result[module], action) {
				result[module] = append(result[module], action)
			}
		}
	}

	return result, nil
}

func (r Repo) GetAuthClientFromSub(sub string) (app.AuthClientInfo, error) {
	localizer := localeutil.Get()
	repo := user.New(r.client)
	queryOptions := ctype.QueryOptions{
		Filters: ctype.Dict{"sub": sub},
		Preloads: []string{
			"Tenant.AuthClient",
		},
	}
	user, err := repo.Retrieve(queryOptions)
	if err != nil {
		return app.AuthClientInfo{}, err
	}

	if user.LockedAt != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.LockedAccount,
		})
		return app.AuthClientInfo{}, errutil.New("", []string{msg})
	}

	return app.AuthClientInfo{
		TenantID:     user.TenantID,
		Realm:        user.Tenant.AuthClient.Partition,
		ClientID:     user.Tenant.AuthClient.Uid,
		ClientSecret: user.Tenant.AuthClient.Secret,
	}, nil
}
