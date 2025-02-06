package infra

type InputData struct {
	UserID      uint   `json:"user_id" form:"user_id" validate:"required"`
	TaskID      uint   `json:"task_id" form:"task_id" validate:"required"`
	Type        string `json:"type" form:"type" validate:"required,oneof=DOC FILE LINK"`
	Title       string `json:"title" form:"title" validate:"required"`
	Description string `json:"description" form:"description"`
	Content     string `json:"content" form:"content"`
	URL         string `json:"url" form:"url"`
	Order       int    `json:"order" form:"order"`
}
