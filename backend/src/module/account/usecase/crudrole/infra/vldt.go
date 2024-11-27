package infra

type InputData struct {
	TenantID uint   `json:"tenant_id" form:"tenant_id" validate:"required"`
	Title    string `json:"title" form:"title" validate:"required"`
	PemIDs   []uint `json:"pem_ids" form:"pem_ids" validate:"required"`
}
