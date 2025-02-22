package schema

import (
	"src/common/ctype"
	account "src/module/account/schema"
	pm "src/module/pm/schema"
	"src/util/dictutil"
	"time"

	"gorm.io/datatypes"
)

type Doc struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	UserID      uint           `gorm:"not null" json:"user_id"`
	User        account.User   `gorm:"foreignKey:UserID" json:"user"`
	TaskID      uint           `gorm:"not null" json:"task_id"`
	Task        pm.Task        `gorm:"foreignKey:TaskID" json:"task"`
	Type        string         `gorm:"type:text;not null;default:'DOC';check:type IN ('DOC', 'FILE','LINK')" json:"type"`
	Title       string         `gorm:"type:text;not null" json:"title"`
	Description string         `gorm:"type:text;not null;default:''" json:"description"`
	Content     datatypes.JSON `gorm:"type:jsonb;not null;default:'{}'" json:"content"`
	Link        string         `gorm:"type:text;not null;default:''" json:"link"`
	FileName    string         `gorm:"type:text;not null;default:''" json:"file_name"`
	FileType    string         `gorm:"type:text;not null;default:''" json:"file_type"`
	FileSize    int            `gorm:"not null;default:0" json:"file_size"`
	FileURL     string         `gorm:"type:text;not null;default:''" json:"file_url"`
	Order       int            `gorm:"not null;default:0" json:"order"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

func NewDoc(data ctype.Dict) *Doc {
	return &Doc{
		UserID:      dictutil.GetValue[uint](data, "UserID"),
		TaskID:      dictutil.GetValue[uint](data, "TaskID"),
		Type:        dictutil.GetValue[string](data, "Type"),
		Title:       dictutil.GetValue[string](data, "Title"),
		Description: dictutil.GetValue[string](data, "Description"),
		Content:     dictutil.GetValue[datatypes.JSON](data, "Content"),
		Link:        dictutil.GetValue[string](data, "Link"),
		FileName:    dictutil.GetValue[string](data, "FileName"),
		FileType:    dictutil.GetValue[string](data, "FileType"),
		FileSize:    dictutil.GetValue[int](data, "FileSize"),
		FileURL:     dictutil.GetValue[string](data, "FileURL"),
		Order:       dictutil.GetValue[int](data, "Order"),
	}
}

type DocAttachment struct {
	ID        uint         `gorm:"primaryKey" json:"id"`
	UserID    uint         `gorm:"not null" json:"user_id"`
	User      account.User `gorm:"foreignKey:UserID" json:"user"`
	TaskID    uint         `gorm:"not null" json:"task_id"`
	Task      pm.Task      `gorm:"foreignKey:TaskID" json:"task"`
	FileName  string       `gorm:"type:text;not null" json:"file_name"`
	FileType  string       `gorm:"type:text;not null" json:"file_type"`
	FileSize  int          `gorm:"not null" json:"file_size"`
	FileURL   string       `gorm:"type:text;not null" json:"file_url"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

func NewDocAttachment(data ctype.Dict) *DocAttachment {
	return &DocAttachment{
		UserID:   dictutil.GetValue[uint](data, "UserID"),
		TaskID:   dictutil.GetValue[uint](data, "TaskID"),
		FileName: dictutil.GetValue[string](data, "FileName"),
		FileType: dictutil.GetValue[string](data, "FileType"),
		FileSize: dictutil.GetValue[int](data, "FileSize"),
		FileURL:  dictutil.GetValue[string](data, "FileURL"),
	}
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
		Value:        dictutil.GetValue[datatypes.JSON](data, "Value"),
	}
}
