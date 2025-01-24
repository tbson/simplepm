package app

import (
	"src/module/pm/schema"
)

type MessageRepo interface {
	List(taskID uint) ([]schema.Message, error)
}
