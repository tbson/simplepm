package infra

import "gorm.io/datatypes"

type InputData struct {
	UserID      uint           `json:"user_id" form:"user_id" validate:"required"`
	TaskID      uint           `json:"task_id" form:"task_id" validate:"required"`
	Type        string         `json:"type" form:"type" validate:"required,oneof=DOC FILE LINK"`
	Title       string         `json:"title" form:"title" validate:"required"`
	Description string         `json:"description" form:"description"`
	Content     datatypes.JSON `json:"content" form:"content"`
	Link        string         `json:"link" form:"link"`
	FileName    string         `json:"file_name" form:"file_name"`
	FileType    string         `json:"file_type" form:"file_type"`
	FileSize    int            `json:"file_size" form:"file_size"`
	Order       int            `json:"order" form:"order"`
}
