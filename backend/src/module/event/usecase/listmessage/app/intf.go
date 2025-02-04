package app

import (
	"src/module/event/schema"
)

type MessageRepo interface {
	List(taskID uint, pageState []byte) ([]schema.Message, []byte, error)
	GetAttachmentMap(messages []schema.Message) (map[string][]schema.Attachment, error)
}
