package infra

type InputData struct {
	Uid       string `json:"uid" form:"uid" validate:"required"`
	Title     string `json:"title" form:"title" validate:"required"`
	AvatarStr string `json:"avatar_str" form:"avatar_str"`
}
