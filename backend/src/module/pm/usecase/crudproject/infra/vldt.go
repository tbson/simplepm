package infra

import "time"

type InputData struct {
	TenantID    uint       `json:"tenant_id" form:"tenant_id" validate:"required"`
	Title       string     `json:"title" form:"title" validate:"required"`
	Description string     `json:"description" form:"description"`
	Layout      string     `json:"layout" form:"layout" validate:"required"`
	WorkspaceID *uint      `json:"workspace_id" form:"workspace_id"`
	StartDate   *time.Time `json:"start_date" form:"start_date"`
	TargetDate  *time.Time `json:"target_date" form:"target_date"`
	Order       int        `json:"order" form:"order"`
}
