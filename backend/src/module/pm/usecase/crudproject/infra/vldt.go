package infra

type InputData struct {
	TenantID    uint   `json:"tenant_id" form:"tenant_id" validate:"required"`
	Title       string `json:"title" form:"title" validate:"required"`
	Description string `json:"description" form:"description"`
	Layout      string `json:"layout" form:"layout" validate:"required"`
	Status      string `json:"status" form:"status" validate:"required"`
	WorkspaceID *uint  `json:"workspace_id" form:"workspace_id"`
	Order       int    `json:"order" form:"order"`
}
