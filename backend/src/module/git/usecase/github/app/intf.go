package app

import (
	"src/common/ctype"
	account "src/module/account/schema"
	event "src/module/event/schema"
	pm "src/module/pm/schema"
)

const GITHUB_CALLBACK_ACTION_INSTALL = "install"
const GITHUB_WEBHOOK_ACTION_CREATED = "created"
const GITHUB_WEBHOOK_ACTION_DELETED = "deleted"
const GITHUB_WEBHOOK_PR_OPENED = "opened"
const GITHUB_WEBHOOK_PR_CLOSED = "closed"

type GithubRepo struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	Private  bool   `json:"private"`
}

type SocketUser struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Color  string `json:"color"`
}

type SocketData struct {
	ID        string                 `json:"id" form:"id"`
	Type      string                 `json:"type" form:"type"`
	User      SocketUser             `json:"user" form:"user"`
	TaskID    uint                   `json:"task_id" form:"task_id"`
	ProjectID uint                   `json:"project_id" form:"project_id"`
	Content   string                 `json:"content" form:"content"`
	GitData   map[string]interface{} `json:"git_data" form:"git_data"`
}

type SocketMessage struct {
	Channel string     `json:"channel" form:"channel"`
	Data    SocketData `json:"data" form:"data"`
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
	Commits     []GithubCommit `json:"commits"`
	PullRequest struct {
		Head struct {
			Ref string `json:"ref"`
		} `json:"head"`
	} `json:"pull_request"`
}

type TaskUser struct {
	TaskID     *uint
	ProjectID  *uint
	UserID     *uint
	UserAvatar *string
	UserName   *string
	UserColor  *string
}

type TenantRepo interface {
	Retrieve(queryOptions ctype.QueryOptions) (*account.Tenant, error)
}

type GitAccountRepo interface {
	UpdateOrCreate(
		queryOptions ctype.QueryOptions,
		data ctype.Dict,
	) (*account.GitAccount, error)
	DeleteBy(queryOptions ctype.QueryOptions) ([]uint, error)
}

type GitRepoRepo interface {
	Create(data ctype.Dict) (*account.GitRepo, error)
	DeleteBy(queryOptions ctype.QueryOptions) ([]uint, error)
}

type GitPushRepo interface {
	Create(data ctype.Dict) (*pm.GitPush, error)
}

type GitCommitRepo interface {
	Create(data ctype.Dict) (*pm.GitCommit, error)
}

type GitRepo interface {
	GetTaskUser(gitRepo string, gitBranch string) (TaskUser, error)
}

type MessageRepo interface {
	Create(message event.Message) (event.Message, error)
}

type CentrifugoRepo interface {
	Publish(data interface{}) error
}
