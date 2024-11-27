package app

import (
	"src/common/ctype"
	"src/module/account/schema"
)

type Service struct {
	roleRepo     RoleRepo
	crudRoleRepo CrudRoleRepo
}

func New(roleRepo RoleRepo, crudRoleRepo CrudRoleRepo) Service {
	return Service{roleRepo, crudRoleRepo}
}

func (srv Service) Create(data ctype.Dict) (schema.Role, error) {
	emptyResult := schema.Role{}
	pemIds := data["PemIDs"].([]uint)
	delete(data, "PemIDs")
	pems, err := srv.crudRoleRepo.ListPemByIds(pemIds)
	if err != nil {
		return emptyResult, err
	}
	data["Pems"] = pems

	result, err := srv.roleRepo.Create(data)
	if err != nil {
		return emptyResult, err
	}
	return *result, nil
}

func (srv Service) Update(id uint, data ctype.Dict) (schema.Role, error) {
	emptyResult := schema.Role{}
	pemIds := data["PemIDs"].([]uint)
	delete(data, "PemIDs")
	pems, err := srv.crudRoleRepo.ListPemByIds(pemIds)
	if err != nil {
		return emptyResult, err
	}
	data["Pems"] = pems

	result, err := srv.roleRepo.Update(id, data)
	if err != nil {
		return emptyResult, err
	}
	return *result, nil
}
