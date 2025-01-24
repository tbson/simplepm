package app

import (
	"src/module/pm/schema"
)

type MessageRepo interface {
	List(taskID uint) ([]schema.Message, error)
	GetAttachmentMap(messages []schema.Message) (map[string][]schema.Attachment, error)
}
