package app

import (
	"src/module/pm/schema"
)

type socketData struct {
	ID        string `json:"id"`
	UserID    uint   `json:"user_id"`
	TaskID    uint   `json:"task_id"`
	ProjectID uint   `json:"project_id"`
	Content   string `json:"content"`
}

type SocketMessage struct {
	Channel string     `json:"channel" form:"channel" validate:"required"`
	Data    socketData `json:"data" form:"data" validate:"required"`
}

type CentrifugoRepo interface {
	Publish(data interface{}) error
}

type MessageRepo interface {
	Create(message schema.Message) (string, error)
}
