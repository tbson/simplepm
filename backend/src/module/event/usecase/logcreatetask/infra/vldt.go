package infra

import "gorm.io/datatypes"

type InputData struct {
	TenantID     uint           `json:"tenant_id"`
	ProjectID    uint           `json:"project_id"`
	TaskID       uint           `json:"task_id"`
	UserID       uint           `json:"user_id"`
	UserFullName string         `json:"user_full_name"`
	SourceType   string         `json:"source_type"`
	SourceID     uint           `json:"source_id"`
	SourceTitle  string         `json:"source_title"`
	Action       string         `json:"action"`
	Value        datatypes.JSON `json:"value"`
}
