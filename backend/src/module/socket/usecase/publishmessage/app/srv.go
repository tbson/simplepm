package app

import (
	"src/module/aws/repo/s3"
	"src/module/pm/schema"
)

type Service struct {
	centrifugoRepo CentrifugoRepo
	messageRepo    MessageRepo
}

func New(centrifugoRepo CentrifugoRepo, messageRepo MessageRepo) Service {
	return Service{centrifugoRepo, messageRepo}
}

func (srv Service) Publish(
	socketMessage SocketMessage,
	files []s3.FileInfo,
) (string, error) {
	socketAttachments := []SocketAttachment{}
	messageData := schema.Message{
		UserID:    socketMessage.Data.User.ID,
		TaskID:    socketMessage.Data.TaskID,
		ProjectID: socketMessage.Data.ProjectID,
		Content:   socketMessage.Data.Content,
	}
	message, err := srv.messageRepo.Create(messageData)
	if err != nil {
		return "", err
	}

	for _, file := range files {
		attachment, err := srv.messageRepo.CreateAttachment(
			message.ID,
			file.FileName,
			file.FileType,
			file.FileURL,
		)
		if err != nil {
			return "", err
		}
		socketAttachment := SocketAttachment{
			FileName: attachment.FileName,
			FileType: attachment.FileType,
			FileURL:  attachment.FileURL,
		}
		socketAttachments = append(socketAttachments, socketAttachment)
	}
	socketMessage.Data.Attachments = socketAttachments

	socketMessage.Data.ID = message.ID
	err = srv.centrifugoRepo.Publish(socketMessage)
	if err != nil {
		return "", err
	}

	return message.ID, nil
}
