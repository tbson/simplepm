package app

import (
	"src/common/ctype"
	"src/module/pm/schema"
)

type CentrifugoRepo interface {
	Publish(data ctype.SocketMessage) error
}

type MessageRepo interface {
	Create(message schema.Message) (string, error)
}
