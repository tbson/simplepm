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

func (srv Service) List(
	taskID uint, pageState []byte,
) ([]schema.Message, []byte, map[string][]schema.Attachment, error) {
	messages, nextPageState, err := srv.messageRepo.List(taskID, pageState)
	if err != nil {
		return nil, nil, nil, err
	}
	attachmentMap, err := srv.messageRepo.GetAttachmentMap(messages)
	return messages, nextPageState, attachmentMap, nil
}
