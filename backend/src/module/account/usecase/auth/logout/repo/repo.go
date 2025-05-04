package repo

import (
	"slices"

	"src/module/account/repo/user"

	"src/common/ctype"

	"src/util/stringutil"

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

func (r Repo) GetPemModulesActionsMap(userId uint) (map[string][]string, error) {
	repo := user.New(r.client)

	queryOptions := ctype.QueryOpts{
		Filters: ctype.Dict{"id": userId},
		Preloads: []string{
			"Roles.Pems",
		},
	}
	user, err := repo.Retrieve(queryOptions)
	if err != nil {
		return nil, err
	}

	result := make(map[string][]string)
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
