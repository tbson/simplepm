package repo

import (
	"src/common/ctype"
	"src/module/account/repo/tenant"

	"src/module/account/usecase/auth/checkauthurl/model"

	"gorm.io/gorm"
)

type repo struct {
	client *gorm.DB
}

func New(client *gorm.DB) repo {
	return repo{
		client: client,
	}
}

func (r repo) GetAuthClientFromTenantUid(
	tenantUid string,
) (model.AuthClientInfo, error) {
	repo := tenant.New(r.client)
	queryOptions := ctype.QueryOptions{
		Filters:  ctype.Dict{"uid": tenantUid},
		Preloads: []string{"AuthClient"},
	}
	tenant, err := repo.Retrieve(queryOptions)
	if err != nil {
		return model.AuthClientInfo{}, err
	}

	return model.AuthClientInfo{
		TenantID:     tenant.ID,
		Realm:        tenant.AuthClient.Partition,
		ClientID:     tenant.AuthClient.Uid,
		ClientSecret: tenant.AuthClient.Secret,
	}, nil
}
