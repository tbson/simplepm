package app

import (
	"src/common/ctype"
	"src/module/event/schema"
)

type InputData struct {
	Channel   string `json:"channel" form:"channel" validate:"required"`
	TaskID    uint   `json:"task_id" form:"task_id" validate:"required"`
	ProjectID uint   `json:"project_id" form:"project_id" validate:"required"`
	Content   string `json:"content" form:"content" validate:"required"`
}

type SocketAttachment struct {
	FileName string `json:"file_name" form:"file_name"`
	FileType string `json:"file_type" form:"file_type"`
	FileURL  string `json:"file_url" form:"file_url"`
	FileSize int    `json:"file_size" form:"file_size"`
}

type SocketUser struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Color  string `json:"color"`
}

type SocketData struct {
	ID          string             `json:"id" form:"id"`
	Type        string             `json:"type" form:"type"`
	User        SocketUser         `json:"user" form:"user"`
	TaskID      uint               `json:"task_id" form:"task_id"`
	ProjectID   uint               `json:"project_id" form:"project_id"`
	Content     string             `json:"content" form:"content"`
	Attachments []SocketAttachment `json:"attachments" form:"attachments"`
}

type SocketMessage struct {
	Channel string     `json:"channel" form:"channel"`
	Data    SocketData `json:"data" form:"data"`
}

type SocketProvider interface {
	Publish(data interface{}) error
}

type MessageRepo interface {
	List(taskID uint, pageState []byte) ([]schema.Message, []byte, error)
	GetAttachmentMap(messages []schema.Message) (map[string][]schema.Attachment, error)
	Create(message ctype.Dict) (schema.Message, error)
	Update(id string, taskID uint, message schema.Message) (schema.Message, error)
	Delete(id string, taskID uint) error
	CreateAttachment(
		messageID string,
		fileName string,
		fileType string,
		fileURL string,
		fileSize int,
	) (schema.Attachment, error)
}
