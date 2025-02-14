package app

import (
	"src/common/ctype"
	"src/module/account/schema"
)

type Service struct {
	tenantRepo     TenantRepo
	gitAccountRepo GitAccountRepo
}

func New(tenantRepo TenantRepo, gitAccountRepo GitAccountRepo) Service {
	return Service{tenantRepo, gitAccountRepo}
}

func (srv Service) HandleInstallCallback(
	uid string,
	tenantUid string,
) (*schema.GitAccount, error) {
	tenant, err := srv.tenantRepo.Retrieve(ctype.QueryOptions{
		Filters: ctype.Dict{
			"uid": tenantUid,
		},
	})
	if err != nil {
		return nil, err
	}
	tenantID := tenant.ID

	data := ctype.Dict{
		"Uid":      &uid,
		"TenantID": &tenantID,
	}

	result, err := srv.gitAccountRepo.UpdateOrCreate(ctype.QueryOptions{
		Filters: ctype.Dict{
			"Uid": &uid,
		},
	}, data)

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (srv Service) HandleInstallWebhook(
	uid string,
	title string,
	avatar string,
) (*schema.GitAccount, error) {
	data := ctype.Dict{
		"Uid":    &uid,
		"Title":  title,
		"Avatar": avatar,
	}

	gitAccount, err := srv.gitAccountRepo.UpdateOrCreate(ctype.QueryOptions{
		Filters: ctype.Dict{
			"Uid": &uid,
		},
	}, data)
	if err != nil {
		return nil, err
	}

	return gitAccount, nil
}

func (srv Service) HandleUninstallWebhook(
	uid string,
) error {
	_, err := srv.gitAccountRepo.DeleteBy(ctype.QueryOptions{
		Filters: ctype.Dict{
			"Uid": &uid,
		},
	})
	if err != nil {
		return err
	}

	return nil
}
