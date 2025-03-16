package repo

import (
	"src/common/ctype"
	"src/module/account/repo/user"
	"src/util/errutil"
	"src/util/localeutil"

	"src/module/account/usecase/auth/checkauthurl/model"

	"github.com/nicksnyder/go-i18n/v2/i18n"
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

func (r repo) GetAuthClientFromSub(sub string) (model.AuthClientInfo, error) {
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
		return model.AuthClientInfo{}, err
	}

	if user.LockedAt != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.LockedAccount,
		})
		return model.AuthClientInfo{}, errutil.New("", []string{msg})
	}

	return model.AuthClientInfo{
		TenantID:     user.TenantID,
		Realm:        user.Tenant.AuthClient.Partition,
		ClientID:     user.Tenant.AuthClient.Uid,
		ClientSecret: user.Tenant.AuthClient.Secret,
	}, nil
}
