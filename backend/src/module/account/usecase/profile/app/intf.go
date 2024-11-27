package app

import (
	"src/common/ctype"
	"src/module/account/schema"
)

type UserRepo interface {
	Retrieve(opts ctype.QueryOptions) (*schema.User, error)
	Update(id uint, data ctype.Dict) (*schema.User, error)
}

type IamRepo interface {
	GetAdminAccessToken() (string, error)
	UpdateUser(accessToken string, realm string, sub string, data ctype.Dict) error
	SetPassword(accessToken string, sub string, realm string, password string) error
}
