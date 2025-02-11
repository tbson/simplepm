package infra

type InputData struct {
	UserID uint   `json:"user_id" form:"user_id" validate:"required"`
	TaskID uint   `json:"task_id" form:"task_id" validate:"required"`
	Type   string `json:"type" form:"type" validate:"required,oneof=DOC FILE LINK"`
	Link   string `json:"link" validate:"required"`
}
