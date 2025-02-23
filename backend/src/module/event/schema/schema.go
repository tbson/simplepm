package schema

import (
	"encoding/json"
	"src/common/ctype"
	"src/util/dictutil"
	"time"

	"gorm.io/datatypes"

	account "src/module/account/schema"
	pm "src/module/pm/schema"
)

type Message struct {
	ID         string `json:"id"`
	UserID     uint   `json:"user_id"`
	TaskID     uint   `json:"task_id"`
	ProjectID  uint   `json:"project_id"`
	Content    string `json:"content"`
	UserName   string `json:"user_name"`
	UserAvatar string `json:"user_avatar"`
	UserColor  string `json:"user_color"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"update_at"`
}

type Attachment struct {
	ID        string `json:"id"`
	MessageID string `json:"message_id"`
	FileName  string `json:"file_name"`
	FileType  string `json:"file_type"`
	FileURL   string `json:"file_url"`
	FileSize  int    `json:"file_size"`
	CreatedAt string `json:"created_at"`
}

type Change struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	TenantID     uint           `gorm:"not null" json:"tenant_id"`
	Tenant       account.Tenant `gorm:"foreignKey:TenantID" json:"tenant"`
	ProjectID    uint           `gorm:"not null" json:"project_id"`
	Project      pm.Project     `gorm:"foreignKey:ProjectID" json:"project"`
	TaskID       uint           `gorm:"not null" json:"task_id"`
	Task         pm.Task        `gorm:"foreignKey:TaskID" json:"task"`
	UserID       uint           `gorm:"not null" json:"user_id"`
	User         account.User   `gorm:"foreignKey:UserID" json:"user"`
	UserFullName string         `gorm:"type:text;not null;default:''" json:"user_full_name"`
	SourceType   string         `gorm:"type:text;not null" json:"source_type"`
	SourceID     uint           `gorm:"not null" json:"source_id"`
	SourceTitle  string         `gorm:"type:text;not null;default:''" json:"source_title"`
	Action       string         `gorm:"type:text;not null" json:"action"`
	Value        datatypes.JSON `gorm:"type:jsonb;not null;default:'{}'" json:"value"`
	CreatedAt    time.Time      `json:"created_at"`
}

func NewChange(data ctype.Dict) *Change {
	valueJSON, err := json.Marshal(data["Value"])
	if err != nil {
		panic("Failed to marshal Value")
	}
	return &Change{
		TenantID:     dictutil.GetValue[uint](data, "TenantID"),
		ProjectID:    dictutil.GetValue[uint](data, "ProjectID"),
		TaskID:       dictutil.GetValue[uint](data, "TaskID"),
		UserID:       dictutil.GetValue[uint](data, "UserID"),
		UserFullName: dictutil.GetValue[string](data, "UserFullName"),
		SourceType:   dictutil.GetValue[string](data, "SourceType"),
		SourceID:     dictutil.GetValue[uint](data, "SourceID"),
		SourceTitle:  dictutil.GetValue[string](data, "SourceTitle"),
		Action:       dictutil.GetValue[string](data, "Action"),
		Value:        datatypes.JSON(valueJSON),
	}
}
