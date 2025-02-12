package app

import (
	"src/common/ctype"
	"src/module/document/repo/net"
	"src/module/document/schema"
)

type InputData struct {
	UserID      uint   `json:"user_id" validate:"required"`
	TaskID      uint   `json:"task_id" validate:"required"`
	Link        string `json:"link" validate:"required"`
	Type        string `json:"type"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type DocRepo interface {
	Create(data ctype.Dict) (*schema.Doc, error)
}

type NetRepo interface {
	GetHTMLMeta(link string) (net.HTMLMeta, error)
}
