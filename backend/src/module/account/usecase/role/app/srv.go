package app

import (
	"src/common/ctype"
	"src/module/account/schema"
)

type Service struct {
	roleRepo      RoleRepo
	roleLocalRepo RoleLocalRepo
}

func New(roleRepo RoleRepo, roleLocalRepo RoleLocalRepo) Service {
	return Service{roleRepo, roleLocalRepo}
}

func (srv Service) Create(data ctype.Dict) (schema.Role, error) {
	emptyResult := schema.Role{}
	pemIds := data["PemIDs"].([]uint)
	delete(data, "PemIDs")
	pems, err := srv.roleLocalRepo.ListPemByIds(pemIds)
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

func (srv Service) Update(updateOpts ctype.QueryOpts, data ctype.Dict) (schema.Role, error) {
	emptyResult := schema.Role{}
	pemIds := data["PemIDs"].([]uint)
	delete(data, "PemIDs")
	pems, err := srv.roleLocalRepo.ListPemByIds(pemIds)
	if err != nil {
		return emptyResult, err
	}
	data["Pems"] = pems

	result, err := srv.roleRepo.Update(updateOpts, data)
	if err != nil {
		return emptyResult, err
	}
	return *result, nil
}
