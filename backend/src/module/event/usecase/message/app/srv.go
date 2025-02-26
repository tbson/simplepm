package app

import (
	"src/module/aws/repo/s3"
	"src/module/event"
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
	data InputData,
	files []s3.FileInfo,
	socketUser SocketUser,
	channel string,
) (string, error) {
	messageType := event.MESSAGE_CREATED
	socketAttachments := []SocketAttachment{}
	messageData := schema.Message{
		TaskID:     data.TaskID,
		ProjectID:  data.ProjectID,
		Content:    data.Content,
		Type:       messageType,
		UserID:     socketUser.ID,
		UserName:   socketUser.Name,
		UserAvatar: socketUser.Avatar,
		UserColor:  socketUser.Color,
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

	socketMessage := SocketMessage{
		Channel: data.Channel,
		Data: SocketData{
			ID:          message.ID,
			Type:        messageType,
			User:        socketUser,
			TaskID:      data.TaskID,
			ProjectID:   data.ProjectID,
			Content:     data.Content,
			Attachments: socketAttachments,
		},
	}
	err = srv.centrifugoRepo.Publish(socketMessage)
	if err != nil {
		return "", err
	}

	return message.ID, nil
}

func (srv Service) Update(
	id string,
	taskID uint,
	data InputData,
) (string, error) {
	messageType := event.MESSAGE_UPDATED
	messageData := schema.Message{
		Content: data.Content,
	}
	message, err := srv.messageRepo.Update(id, taskID, messageData)
	if err != nil {
		return "", err
	}

	socketMessage := SocketMessage{
		Channel: data.Channel,
		Data: SocketData{
			ID:      id,
			Type:    messageType,
			Content: data.Content,
		},
	}
	err = srv.centrifugoRepo.Publish(socketMessage)
	if err != nil {
		return "", err
	}

	return message.ID, nil
}

func (srv Service) Delete(id string, taskID uint, data InputData) error {
	messageType := event.MESSAGE_DELETED
	err := srv.messageRepo.Delete(id, taskID)
	if err != nil {
		return err
	}

	socketMessage := SocketMessage{
		Channel: data.Channel,
		Data: SocketData{
			ID:   id,
			Type: messageType,
		},
	}
	err = srv.centrifugoRepo.Publish(socketMessage)
	if err != nil {
		return err
	}

	return nil
}
