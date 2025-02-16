package infra

type InputData struct {
	FirstName string `json:"first_name" form:"first_name" validate:"required"`
	LastName  string `json:"last_name" form:"last_name" validate:"required"`
	Mobile    string `json:"mobile" form:"mobile" validate:"required"`
}

type InputPassword struct {
	Password        string `json:"password" form:"password" validate:"required"`
	PasswordConfirm string `json:"password_confirm" form:"password_confirm" validate:"required"`
}
