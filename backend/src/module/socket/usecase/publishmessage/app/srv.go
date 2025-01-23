package app

import (
	"src/module/pm/schema"
)

type Service struct {
	centrifugoRepo CentrifugoRepo
	messageRepo    MessageRepo
}

func New(centrifugoRepo CentrifugoRepo, messageRepo MessageRepo) Service {
	return Service{centrifugoRepo, messageRepo}
}

func (srv Service) Publish(data SocketMessage) (string, error) {
	message := schema.Message{
		UserID:    data.Data.UserID,
		TaskID:    data.Data.TaskID,
		ProjectID: data.Data.ProjectID,
		Content:   data.Data.Content,
	}
	id, err := srv.messageRepo.Create(message)
	if err != nil {
		return "", err
	}
	data.Data.ID = id
	err = srv.centrifugoRepo.Publish(data)
	if err != nil {
		return "", err
	}
	return id, nil
}
