package app

import (
	"src/common/ctype"
	"src/module/account/schema"
)

type Service struct {
	Repo         UserRepo
	crudUserRepo CrudUserRepo
}

func New(Repo UserRepo, crudUserRepo CrudUserRepo) Service {
	return Service{Repo, crudUserRepo}
}

func (srv Service) Create(data ctype.Dict) (schema.User, error) {
	emptyResult := schema.User{}
	roleIds := data["RoleIDs"].([]uint)
	delete(data, "RoleIDs")
	roles, err := srv.crudUserRepo.ListRoleByIds(roleIds)
	if err != nil {
		return emptyResult, err
	}
	data["Roles"] = roles

	result, err := srv.Repo.Create(data)
	if err != nil {
		return emptyResult, err
	}
	return *result, nil
}

func (srv Service) Update(id uint, data ctype.Dict) (schema.User, error) {
	emptyResult := schema.User{}
	roleIds := data["RoleIDs"].([]uint)
	delete(data, "RoleIDs")
	roles, err := srv.crudUserRepo.ListRoleByIds(roleIds)
	if err != nil {
		return emptyResult, err
	}
	data["Roles"] = roles

	result, err := srv.Repo.Update(id, data)
	if err != nil {
		return emptyResult, err
	}
	return *result, nil
}
