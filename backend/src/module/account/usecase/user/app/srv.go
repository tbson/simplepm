package app

import (
	"src/common/ctype"
	"src/module/account/schema"
)

type Service struct {
	UserRepo      UserRepo
	UserLocalRepo UserLocalRepo
}

func New(UserRepo UserRepo, UserLocalRepo UserLocalRepo) Service {
	return Service{UserRepo, UserLocalRepo}
}

func (srv Service) Create(data ctype.Dict) (schema.User, error) {
	emptyResult := schema.User{}
	roleIds := data["RoleIDs"].([]uint)
	delete(data, "RoleIDs")
	roles, err := srv.UserLocalRepo.ListRoleByIds(roleIds)
	if err != nil {
		return emptyResult, err
	}
	data["Roles"] = roles

	result, err := srv.UserRepo.Create(data)
	if err != nil {
		return emptyResult, err
	}
	return *result, nil
}

func (srv Service) Update(updateOpts ctype.QueryOpts, data ctype.Dict) (schema.User, error) {
	emptyResult := schema.User{}
	roleIds := data["RoleIDs"].([]uint)
	delete(data, "RoleIDs")
	roles, err := srv.UserLocalRepo.ListRoleByIds(roleIds)
	if err != nil {
		return emptyResult, err
	}
	data["Roles"] = roles

	result, err := srv.UserRepo.Update(updateOpts, data)
	if err != nil {
		return emptyResult, err
	}
	return *result, nil
}
