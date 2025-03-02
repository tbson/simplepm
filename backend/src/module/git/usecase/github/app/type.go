package app

import (
	"src/common/customfield"
)

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

type RepoInput struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	Private  bool   `json:"private"`
}

type CommitInput struct {
	ID        string `json:"id"`
	URL       string `json:"url"`
	Message   string `json:"message"`
	Committer struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Username string `json:"username"`
	} `json:"committer"`
}

type PullRequestInput struct {
	ID         customfield.StringDigit `json:"id"`
	Title      string                  `json:"title"`
	FromBranch string                  `json:"from_branch"`
	ToBranch   string                  `json:"to_branch"`
	URL        string                  `json:"url"`
	MergedAt   *customfield.TimeStr    `json:"merged_at"`
	State      string                  `json:"state"`
	Head       struct {
		Ref string `json:"ref"`
	} `json:"head"`
	Base struct {
		Ref string `json:"ref"`
	} `json:"base"`
}

type GithubWebhook struct {
	Ref          string `json:"ref"`
	Action       string `json:"action"`
	Installation struct {
		ID uint `json:"id"`
	}
	Repositories []RepoInput   `json:"repositories"`
	Repository   RepoInput     `json:"repository"`
	Commits      []CommitInput `json:"commits"`
	Sender       struct {
		AvatarURL string `json:"avatar_url"`
		Login     string `json:"login"`
	}
	PullRequest PullRequestInput `json:"pull_request"`
}

type TaskUser struct {
	TaskID     *uint
	ProjectID  *uint
	UserID     *uint
	UserAvatar *string
	UserName   *string
	UserColor  *string
}
