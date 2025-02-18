package app

import (
	"src/common/ctype"
	"src/module/account/schema"
	"src/util/dateutil"
)

type Service struct {
	userRepo UserRepo
}

func New(userRepo UserRepo) Service {
	return Service{userRepo}
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
	updateOptions := ctype.QueryOptions{Filters: ctype.Dict{"ID": id}}

	result, err := srv.userRepo.Update(updateOptions, data)
	if err != nil {
		return schema.User{}, err
	}
	return *result, nil
}
