package infra

type InputData struct {
	ProjectID   uint   `json:"project_id" form:"project_id" validate:"required"`
	Title       string `json:"title" form:"title" validate:"required"`
	Description string `json:"description" form:"description"`
	Type        string `json:"type" form:"type" validate:"required"`
	Order       int    `json:"order" form:"order"`
}
