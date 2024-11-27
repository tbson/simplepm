package infra

type InputData struct {
	Uid       string  `json:"uid" form:"uid" validate:"required"`
	Title     string  `json:"title" form:"title" validate:"required"`
	Email     string  `json:"email" form:"email" validate:"required"`
	Mobile    *string `json:"mobile" form:"mobile"`
	FirstName string  `json:"first_name" form:"first_name" validate:"required"`
	LastName  string  `json:"last_name" form:"last_name" validate:"required"`
	Password  string  `json:"password" form:"password" validate:"required"`
}
