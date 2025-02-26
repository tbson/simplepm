package schema

import (
	"encoding/json"
	"src/common/ctype"
	"src/util/dictutil"
	"time"

	"gorm.io/datatypes"
)

type GitCommit struct {
	ID            uint      `json:"id"`
	CommitID      string    `json:"commit_id"`
	CommitURL     string    `json:"commit_url"`
	CommitMessage string    `json:"commit_message"`
	CreatedAt     time.Time `json:"created_at"`
}

type GitPush struct {
	ID         uint        `json:"id"`
	GitBranch  string      `json:"git_branch"`
	GitCommits []GitCommit `json:"git_commits"`
}

type Message struct {
	ID         string                 `json:"id"`
	UserID     uint                   `json:"user_id"`
	TaskID     uint                   `json:"task_id"`
	ProjectID  uint                   `json:"project_id"`
	Content    string                 `json:"content"`
	Type       string                 `json:"type"`
	GitPush    map[string]interface{} `json:"git_push"`
	UserName   string                 `json:"user_name"`
	UserAvatar string                 `json:"user_avatar"`
	UserColor  string                 `json:"user_color"`
	CreatedAt  string                 `json:"created_at"`
	UpdatedAt  string                 `json:"update_at"`
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
	ProjectID    uint           `gorm:"not null" json:"project_id"`
	TaskID       uint           `gorm:"not null" json:"task_id"`
	UserID       uint           `gorm:"not null" json:"user_id"`
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
