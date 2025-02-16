package app

import (
	"src/common/ctype"
	"src/module/account/schema"
)

const GITHUB_CALLBACK_ACTION_INSTALL = "install"
const GITHUB_WEBHOOK_ACTION_CREATED = "created"
const GITHUB_WEBHOOK_ACTION_DELETED = "deleted"

type GithubRepo struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	Private  bool   `json:"private"`
}

type GithubCommit struct {
	ID        string `json:"id"`
	URL       string `json:"url"`
	Message   string `json:"message"`
	Committer struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Username string `json:"username"`
	} `json:"committer"`
}

type GithubWebhook struct {
	Ref          string `json:"ref"`
	Action       string `json:"action"`
	Installation struct {
		ID uint `json:"id"`
	}
	Repositories []GithubRepo `json:"repositories"`
	Repository   GithubRepo   `json:"repository"`
	Sender       struct {
		AvatarURL string `json:"avatar_url"`
		Login     string `json:"login"`
	}
	Pusher struct {
		Name string `json:"name"`
	}
	Commits []GithubCommit `json:"commits"`
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

type GitRepoRepo interface {
	Create(data ctype.Dict) (*schema.GitRepo, error)
	DeleteBy(queryOptions ctype.QueryOptions) ([]uint, error)
}
