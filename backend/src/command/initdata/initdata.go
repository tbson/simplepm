package main

import (
	"src/common/ctype"
	"src/common/setting"
	"src/module/account/repo/authclient"
	"src/module/account/repo/tenant"
	"src/module/account/repo/user"
	"src/util/dbutil"
)

func main() {
	dbutil.InitDb()
	db := dbutil.Db()
	authClientRepo := authclient.New(db)
	tenantRepo := tenant.New(db)
	userRepo := user.New(db)

	queryOptions := ctype.QueryOptions{
		Filters: ctype.Dict{
			"Uid": setting.KEYCLOAK_DEFAULT_CLIENT_ID,
		},
	}
	authClientData := ctype.Dict{
		"Uid":         setting.KEYCLOAK_DEFAULT_CLIENT_ID,
		"Description": "Default client",
		"Secret":      setting.KEYCLOAK_DEFAULT_CLIENT_SECRET,
		"Partition":   setting.KEYCLOAK_DEFAULT_REALM,
		"Default":     true,
	}

	authClient, err := authClientRepo.GetOrCreate(queryOptions, authClientData)
	if err != nil {
		panic(err)
	}

	queryOptions = ctype.QueryOptions{
		Filters: ctype.Dict{
			"Uid": "default",
		},
	}
	tenantData := ctype.Dict{
		"AuthClientID": authClient.ID,
		"Uid":          setting.ADMIN_TEANT_UID,
		"Title":        "Admin",
	}
	tenant, err := tenantRepo.GetOrCreate(queryOptions, tenantData)
	if err != nil {
		panic(err)
	}

	email := "admin@local.dev"
	queryOptions = ctype.QueryOptions{
		Filters: ctype.Dict{
			"email": email,
		},
	}
	userData := ctype.Dict{
		"TenantID":   tenant.ID,
		"ExternalID": email,
		"Email":      email,
		"FirstName":  "Admin",
		"LastName":   "Admin",
		"Admin":      true,
	}
	_, err = userRepo.GetOrCreate(queryOptions, userData)
	if err != nil {
		panic(err)
	}

}
