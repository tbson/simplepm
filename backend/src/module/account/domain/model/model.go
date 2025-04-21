package model

type UserInfo struct {
	ID       uint   `json:"id"`
	TenantID uint   `json:"tenant_id"`
	Admin    bool   `json:"admin"`
	FistName string `json:"first_name"`
	LastName string `json:"last_name"`
	Avatar   string `json:"avatar"`
}

func NewUserInfo(id, tenantID uint, admin bool, firstName, lastName, avatar string) UserInfo {
	return UserInfo{
		ID:       id,
		TenantID: tenantID,
		Admin:    admin,
		FistName: firstName,
		LastName: lastName,
		Avatar:   avatar,
	}
}

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

type LoginResult struct {
	TokenPair TokenPair
	UserInfo  UserInfo
}
