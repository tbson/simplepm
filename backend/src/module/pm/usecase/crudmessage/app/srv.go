package app

import (
	"src/module/pm/schema"
)

type Service struct {
	messageRepo MessageRepo
}

func New(messageRepo MessageRepo) Service {
	return Service{messageRepo}
}

func (srv Service) List(taskID uint) ([]schema.Message, error) {
	return srv.messageRepo.List(taskID)
}
