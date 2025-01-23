package infra

type InputData struct {
	Channel   string `json:"channel" form:"channel" validate:"required"`
	TaskID    uint   `json:"task_id" form:"task_id" validate:"required"`
	ProjectID uint   `json:"project_id" form:"project_id" validate:"required"`
	Content   string `json:"content" form:"content" validate:"required"`
}
