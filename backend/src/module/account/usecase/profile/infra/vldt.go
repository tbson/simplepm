package infra

type InputData struct {
	FirstName string `json:"first_name" form:"first_name" validate:"required"`
	LastName  string `json:"last_name" form:"last_name" validate:"required"`
	Mobile    string `json:"mobile" form:"mobile" validate:"required"`
	GithubID  string `json:"github_id" form:"github_id"`
	GitlabID  string `json:"gitlab_id" form:"gitlab_id"`
}

type InputPassword struct {
	Password        string `json:"password" form:"password" validate:"required"`
	PasswordConfirm string `json:"password_confirm" form:"password_confirm" validate:"required"`
}
