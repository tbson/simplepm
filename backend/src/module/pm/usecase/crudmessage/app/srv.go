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
	taskID uint,
) ([]schema.Message, map[string][]schema.Attachment, error) {
	messages, err := srv.messageRepo.List(taskID)
	if err != nil {
		return nil, nil, err
	}
	attachmentMap, err := srv.messageRepo.GetAttachmentMap(messages)
	return messages, attachmentMap, nil
}
