package app

import (
	"src/module/event/schema"
)

type SocketAttachment struct {
	FileName string `json:"file_name" form:"file_name" validate:"required"`
	FileType string `json:"file_type" form:"file_type" validate:"required"`
	FileURL  string `json:"file_url" form:"file_url" validate:"required"`
	FileSize int    `json:"file_size" form:"file_size" validate:"required"`
}

type SocketUser struct {
	ID     uint   `json:"id" validate:"required"`
	Name   string `json:"name" validate:"required"`
	Avatar string `json:"avatar"`
	Color  string `json:"color"`
}

type SocketData struct {
	ID          string             `json:"id" form:"id"`
	User        SocketUser         `json:"user" form:"user" validate:"required"`
	TaskID      uint               `json:"task_id" form:"task_id" validate:"required"`
	ProjectID   uint               `json:"project_id" form:"project_id" validate:"required"`
	Content     string             `json:"content" form:"content" validate:"required"`
	Attachments []SocketAttachment `json:"attachments" form:"attachments"`
}

type SocketMessage struct {
	Channel string     `json:"channel" form:"channel" validate:"required"`
	Data    SocketData `json:"data" form:"data" validate:"required"`
}

type CentrifugoRepo interface {
	Publish(data interface{}) error
}

type MessageRepo interface {
	Create(message schema.Message) (schema.Message, error)
	CreateAttachment(
		messageID string,
		fileName string,
		fileType string,
		fileURL string,
		fileSize int,
	) (schema.Attachment, error)
}
