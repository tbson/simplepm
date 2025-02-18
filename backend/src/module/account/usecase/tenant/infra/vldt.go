package infra

type InputData struct {
	AuthClientID uint   `json:"auth_client_id" form:"auth_client_id" validate:"required"`
	Uid          string `json:"uid" form:"uid" validate:"required"`
	Title        string `json:"title" form:"title" validate:"required"`
	AvatarStr    string `json:"avatar_str" form:"avatar_str"`
}
