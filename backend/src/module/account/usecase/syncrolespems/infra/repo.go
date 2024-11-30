package infra

import (
	"slices"
	"src/common/ctype"
	"src/common/profiletype"
	"src/module/account/repo/pem"

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

func getAdmin(profileTypes []string) bool {
	// if profileTypes ONLY contains ADMIN, return true
	if len(profileTypes) == 1 && profileTypes[0] == profiletype.ADMIN {
		return true
	}

	// if profileTypes ONLY contains STAFF, return true
	if len(profileTypes) == 1 && profileTypes[0] == profiletype.STAFF {
		return true
	}

	// if profileTypes contains ADMIN and STAFF, return true
	if slices.Contains(
		profileTypes,
		profiletype.ADMIN,
	) && slices.Contains(
		profileTypes,
		profiletype.STAFF,
	) {
		return true
	}

	return false
}

func (r Repo) WritePems(pemMap ctype.PemMap) error {
	pemRepo := pem.New(r.client)
	for _, pemData := range pemMap {
		filterOptions := ctype.QueryOptions{
			Filters: ctype.Dict{
				"module": pemData.Module,
				"action": pemData.Action,
			},
		}
		data := ctype.Dict{
			"Title":  pemData.Title,
			"Module": pemData.Module,
			"Action": pemData.Action,
			"Admin":  getAdmin(pemData.ProfileTypes),
		}

		_, err := pemRepo.GetOrCreate(filterOptions, data)

		if err != nil {
			panic(err)
		}
	}
	return nil
}
