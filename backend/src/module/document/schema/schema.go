package schema

import (
	"src/common/ctype"
	account "src/module/account/schema"
	pm "src/module/pm/schema"
	"src/util/dictutil"
	"time"
)

type Doc struct {
	ID          uint         `gorm:"primaryKey" json:"id"`
	UserID      uint         `gorm:"not null" json:"user_id"`
	User        account.User `gorm:"foreignKey:UserID" json:"user"`
	TaskID      uint         `gorm:"not null" json:"task_id"`
	Task        pm.Task      `gorm:"foreignKey:TaskID" json:"task"`
	Type        string       `gorm:"type:text;not null;default:'DOC';check:type IN ('DOC', 'FILE','LINK')" json:"type"`
	Title       string       `gorm:"type:text;not null" json:"title"`
	Description string       `gorm:"type:text;not null;default:''" json:"description"`
	Content     string       `gorm:"type:text;not null;default:''" json:"content"`
	URL         string       `gorm:"type:text;not null;default:''" json:"url"`
	Order       int          `gorm:"not null;default:0" json:"order"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

func NewDoc(data ctype.Dict) *Doc {
	return &Doc{
		UserID:      dictutil.GetValue[uint](data, "UserID"),
		TaskID:      dictutil.GetValue[uint](data, "TaskID"),
		Type:        dictutil.GetValue[string](data, "Type"),
		Title:       dictutil.GetValue[string](data, "Title"),
		Description: dictutil.GetValue[string](data, "Description"),
		Content:     dictutil.GetValue[string](data, "Content"),
		URL:         dictutil.GetValue[string](data, "URL"),
		Order:       dictutil.GetValue[int](data, "Order"),
	}
}
