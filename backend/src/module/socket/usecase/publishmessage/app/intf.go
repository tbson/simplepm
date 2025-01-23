package app

import (
	"src/module/pm/schema"
)

type SocketData struct {
	ID        string `json:"id" form:"id"`
	UserID    uint   `json:"user_id" form:"user_id" validate:"required"`
	TaskID    uint   `json:"task_id" form:"task_id" validate:"required"`
	ProjectID uint   `json:"project_id" form:"project_id" validate:"required"`
	Content   string `json:"content" form:"content" validate:"required"`
}

type SocketMessage struct {
	Channel string     `json:"channel" form:"channel" validate:"required"`
	Data    SocketData `json:"data" form:"data" validate:"required"`
}

type CentrifugoRepo interface {
	Publish(data interface{}) error
}

type MessageRepo interface {
	Create(message schema.Message) (string, error)
}
