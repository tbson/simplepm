package app

import (
	"fmt"
	"src/module/aws/repo/s3"
	"src/module/event/schema"
)

type Service struct {
	centrifugoRepo CentrifugoRepo
	messageRepo    MessageRepo
}

func New(centrifugoRepo CentrifugoRepo, messageRepo MessageRepo) Service {
	return Service{centrifugoRepo, messageRepo}
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

func (srv Service) Create(
	socketMessage SocketMessage,
	files []s3.FileInfo,
) (string, error) {
	socketAttachments := []SocketAttachment{}
	messageData := schema.Message{
		UserID:     socketMessage.Data.User.ID,
		TaskID:     socketMessage.Data.TaskID,
		ProjectID:  socketMessage.Data.ProjectID,
		Content:    socketMessage.Data.Content,
		UserName:   socketMessage.Data.User.Name,
		UserAvatar: socketMessage.Data.User.Avatar,
		UserColor:  socketMessage.Data.User.Color,
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
			file.FileSize,
		)
		if err != nil {
			return "", err
		}
		socketAttachment := SocketAttachment{
			FileName: attachment.FileName,
			FileType: attachment.FileType,
			FileURL:  attachment.FileURL,
			FileSize: attachment.FileSize,
		}
		socketAttachments = append(socketAttachments, socketAttachment)
	}
	socketMessage.Data.ID = message.ID
	socketMessage.Data.Attachments = socketAttachments
	err = srv.centrifugoRepo.Publish(socketMessage)
	if err != nil {
		return "", err
	}

	return message.ID, nil
}

func (srv Service) Update(
	id string,
	taskID uint,
	socketMessage SocketMessage,
) (string, error) {
	messageData := schema.Message{
		Content: socketMessage.Data.Content,
	}
	message, err := srv.messageRepo.Update(id, taskID, messageData)
	if err != nil {
		fmt.Println("case 1")
		return "", err
	}
	err = srv.centrifugoRepo.Publish(socketMessage)
	if err != nil {
		fmt.Println("case 2")
		return "", err
	}

	return message.ID, nil
}

func (srv Service) Delete(id string, taskID uint, socketMessage SocketMessage) error {
	err := srv.messageRepo.Delete(id, taskID)
	if err != nil {
		fmt.Println("case 1")
		return err
	}
	err = srv.centrifugoRepo.Publish(socketMessage)
	if err != nil {
		fmt.Println("case 2")
		return err
	}

	return nil
}
