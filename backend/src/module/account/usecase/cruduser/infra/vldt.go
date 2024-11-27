package infra

import "src/common/ctype"

type InputData struct {
	TenantID    uint       `json:"tenant_id" form:"tenant_id" validate:"required"`
	TenantTmpID *uint      `json:"tenant_tmp_id" form:"tenant_tmp_id"`
	Uid         string     `json:"uid" form:"uid" validate:"required"`
	Email       string     `json:"email" form:"email" validate:"required"`
	Mobile      *string    `json:"mobile" form:"mobile"`
	FirstName   string     `json:"first_name" form:"first_name"`
	LastName    string     `json:"last_name" form:"last_name"`
	AvatarStr   string     `json:"avatar_str" form:"avatar_str"`
	ExtraInfo   ctype.Dict `json:"extra_info" form:"extra_info"`
	RoleIDs     []uint     `json:"role_ids" form:"role_ids" validate:"required"`
	// Admin       bool       `json:"admin" form:"admin"`
}
