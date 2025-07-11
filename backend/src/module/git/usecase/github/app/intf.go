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

type TenantRepo interface {
	Retrieve(opts ctype.QueryOpts) (*account.Tenant, error)
}

type GitAccountRepo interface {
	UpdateOrCreate(
		opts ctype.QueryOpts,
		data ctype.Dict,
	) (*account.GitAccount, error)
	DeleteBy(opts ctype.QueryOpts) ([]uint, error)
}

type GitRepoRepo interface {
	Create(data ctype.Dict) (*account.GitRepo, error)
	DeleteBy(opts ctype.QueryOpts) ([]uint, error)
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
	Create(message ctype.Dict) (event.Message, error)
}

type SocketProvider interface {
	Publish(data interface{}) error
}
