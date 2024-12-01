package infra

type InputData struct {
	TenantID    uint   `json:"tenant_id" form:"tenant_id" validate:"required"`
	Title       string `json:"title" form:"title" validate:"required"`
	Description string `json:"description" form:"description"`
	Order       int    `json:"order" form:"order"`
}
