package authtype

import "time"

type AuthUserInfo struct {
	ID          uint       `json:"id"`
	TenantID    uint       `json:"tenant_id"`
	TenantUid   string     `json:"tenant_uid"`
	Sub         *string    `json:"sub"`
	Admin       bool       `json:"admin"`
	ProfileType string     `json:"profile_type"`
	LockedAt    *time.Time `json:"locked_at"`
	Email       string     `json:"email"`
	FistName    string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	Mobile      *string    `json:"mobile"`
	Avatar      string     `json:"avatar"`
}

type AuthCallbackResult struct {
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
	Realm        string       `json:"realm"`
	UserInfo     AuthUserInfo `json:"user_info"`
}

type SsoUserInfo struct {
	ExternalID string  `json:"external_id"`
	Sub        *string `json:"sub"`
	Email      string  `json:"email"`
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
}

type SsoCallbackResult struct {
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
	Realm        string      `json:"realm"`
	UserInfo     SsoUserInfo `json:"user_info"`
}
