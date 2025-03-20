package app

import (
	"src/common/ctype"
	"src/common/setting"
	"src/module/account/schema"
)

type Service struct {
	tenantRepo TenantRepo
	userRepo   UserRepo
}

type InitDataResult struct {
	TenantID schema.Tenant
	UserID   schema.User
}

func New(
	tenantRepo TenantRepo,
	userRepo UserRepo,
) Service {
	return Service{tenantRepo, userRepo}
}

func (srv Service) InitData() (InitDataResult, error) {
	// Init tenant
	queryOptions := ctype.QueryOptions{
		Filters: ctype.Dict{
			"Uid": setting.ADMIN_TEANT_UID(),
		},
	}
	tenantData := ctype.Dict{
		"Uid":   setting.TEST_TEANT_UID(),
		"Title": setting.TETS_TEANT_TITLE(),
	}
	tenant, err := srv.tenantRepo.GetOrCreate(queryOptions, tenantData)
	if err != nil {
		return InitDataResult{}, err
	}

	// Init user
	email := setting.TEST_USER_EMAIL_OWNER()
	queryOptions = ctype.QueryOptions{
		Filters: ctype.Dict{
			"email": email,
		},
	}

	userData := ctype.Dict{
		"TenantID":  tenant.ID,
		"Email":     email,
		"FirstName": "Owner",
		"LastName":  "Owner",
		"Pwd":       setting.TEST_USER_PASSWORD(),
	}

	user, err := srv.userRepo.GetOrCreate(queryOptions, userData)
	if err != nil {
		return InitDataResult{}, err
	}

	return InitDataResult{
		TenantID: *tenant,
		UserID:   *user,
	}, nil
}
