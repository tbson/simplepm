package app

import (
	"fmt"
	"src/common/ctype"
	"src/module/account/schema"
	"src/util/numberutil"
)

type Service struct {
	tenantRepo     TenantRepo
	gitAccountRepo GitAccountRepo
	gitRepoRepo    GitRepoRepo
}

func New(
	tenantRepo TenantRepo,
	gitAccountRepo GitAccountRepo,
	gitRepoRepo GitRepoRepo,
) Service {
	return Service{tenantRepo, gitAccountRepo, gitRepoRepo}
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
	repos []GithubRepo,
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

	gitAccountID := gitAccount.ID

	_, err = srv.gitRepoRepo.DeleteBy(ctype.QueryOptions{
		Filters: ctype.Dict{
			"GitAccountID": &gitAccountID,
		},
	})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for _, repo := range repos {
		data := ctype.Dict{
			"GitAccountID": gitAccountID,
			"RepoID":       numberutil.UintToStr(repo.ID),
			"Uid":          repo.FullName,
			"Private":      repo.Private,
		}

		_, err = srv.gitRepoRepo.Create(data)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
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
