package infra

type InputData struct {
	ProjectID   uint   `json:"project_id" form:"project_id" validate:"required"`
	FeatureID   uint   `json:"feature_id" form:"feature_id"`
	Title       string `json:"title" form:"title" validate:"required"`
	Description string `json:"description" form:"description"`
	Order       int    `json:"order" form:"order"`
}
