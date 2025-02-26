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
			u.id AS user_id,
			u.avatar AS user_avatar,
			u.color AS user_color,
			TRIM(CONCAT(u.first_name, ' ', u.last_name)) AS user_name,
			tu.task_id AS task_id,
			p.id AS project_id
		FROM
			tasks_users AS tu
		JOIN
			tasks AS t
		ON
			tu.task_id = t.id
		JOIN
			users AS u
		ON
			tu.user_id = u.id
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
