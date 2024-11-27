package app

import (
	"src/common/ctype"
	"src/module/account/schema"
	"src/util/dateutil"
)

type Service struct {
	crudUserRepo CrudUserRepo
}

func New(crudUserRepo CrudUserRepo) Service {
	return Service{crudUserRepo}
}

func (srv Service) LockUser(id uint, locked bool, lockedReason string) (schema.User, error) {
	data := ctype.Dict{
		"LockedAt":     nil,
		"LockedReason": "",
	}

	if locked {
		data["LockedAt"] = dateutil.Now()
		data["LockedReason"] = lockedReason
	}

	result, err := srv.crudUserRepo.Update(id, data)
	if err != nil {
		return schema.User{}, err
	}
	return *result, nil
}
