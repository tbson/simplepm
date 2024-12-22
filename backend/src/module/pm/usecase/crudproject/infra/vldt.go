package infra

import (
	"src/module/pm"
	"strings"
)

type InputData struct {
	TenantID    uint   `json:"tenant_id" form:"tenant_id" validate:"required"`
	Title       string `json:"title" form:"title" validate:"required"`
	Description string `json:"description" form:"description"`
	Layout      string `json:"layout" form:"layout" validate:"required"`
	Status      string `json:"status" form:"status" validate:"required"`
	WorkspaceID *uint  `json:"workspace_id" form:"workspace_id"`
	GitRepo     string `json:"git_repo" form:"git_repo"`
	GitHost     string `json:"git_host" form:"git_host"`
	Order       int    `json:"order" form:"order"`
}

func (i *InputData) EnsureGitHost() {
	if strings.Contains(i.GitRepo, "github.com") {
		i.GitHost = pm.PROJECT_REPO_TYPE_GITHUB
	} else if strings.Contains(i.GitRepo, "gitlab.com") {
		i.GitHost = pm.PROJECT_REPO_TYPE_GITLAB
	} else {
		i.GitHost = ""
	}
}
