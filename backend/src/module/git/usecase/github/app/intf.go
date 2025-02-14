package app

import (
	"src/common/ctype"
	"src/module/account/schema"
)

const GITHUB_CALLBACK_ACTION_INSTALL = "install"
const GITHUB_WEBHOOK_ACTION_CREATED = "created"
const GITHUB_WEBHOOK_ACTION_DELETED = "deleted"

type GithubInstallWebhook struct {
	Action       string `json:"action"`
	Installation struct {
		ID uint `json:"id"`
	}
	Sender struct {
		AvatarURL string `json:"avatar_url"`
		Login     string `json:"login"`
	}
}

type TenantRepo interface {
	Retrieve(queryOptions ctype.QueryOptions) (*schema.Tenant, error)
}

type GitAccountRepo interface {
	UpdateOrCreate(
		queryOptions ctype.QueryOptions,
		data ctype.Dict,
	) (*schema.GitAccount, error)
	DeleteBy(queryOptions ctype.QueryOptions) ([]uint, error)
}
