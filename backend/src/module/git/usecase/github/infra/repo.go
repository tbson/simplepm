package infra

import (
	"src/module/pm"

	"src/module/git/usecase/github/app"

	"gorm.io/gorm"
)

type Repo struct {
	client *gorm.DB
}

func New(client *gorm.DB) Repo {
	return Repo{
		client: client,
	}
}

func (r Repo) GetTaskUser(gitRepo string, gitBranch string) (app.TaskUser, error) {
	// localizer := localeutil.Get()
	result := app.TaskUser{}
	gitHost := pm.PROJECT_REPO_TYPE_GITHUB
	sql := `
		SELECT
			tu.user_id,
			tu.task_id
		FROM
			tasks_users AS tu
		JOIN
			tasks AS t
		ON
			tu.task_id = t.id
		JOIN
			projects AS p
		ON
			t.project_id = p.id
		WHERE
			p.git_host = @gitHost
			AND p.git_repo = @gitRepo
			AND tu.git_branch = @gitBranch;
	`
	params := map[string]interface{}{
		"gitHost":   gitHost,
		"gitRepo":   gitRepo,
		"gitBranch": gitBranch,
	}

	err := r.client.Raw(sql, params).Find(&result).Error
	if err != nil {
		return result, err
	}
	/*
		if result.TaskID == nil || result.UserID == nil {
			msg := localizer.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: localeutil.TaskUserNotFound,
			})
			return result, errutil.New("", []string{msg})
		}
	*/
	return result, nil
}
