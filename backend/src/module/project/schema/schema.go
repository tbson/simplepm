package schema

import (
	"src/common/ctype"
	"src/module/account/schema"
	"src/util/dictutil"
	"time"
)

type Workspace struct {
	ID          uint          `gorm:"primaryKey" json:"id"`
	TenantID    uint          `gorm:"not null" json:"tenant_id"`
	Tenant      schema.Tenant `gorm:"foreignKey:TenantID;constraint:OnDelete:CASCADE;" json:"tenant"`
	Title       string        `gorm:"not null" json:"title"`
	Description string        `gorm:"not null;default:''" json:"description"`
	Avatar      string        `gorm:"not null;default:''" json:"avatar"`
	Order       int           `gorm:"not null;default:0" json:"order"`
	Users       []schema.User `gorm:"many2many:workspaces_users;" json:"users"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}

func NewWorkspace(data ctype.Dict) *Workspace {
	return &Workspace{
		TenantID:    dictutil.GetValue[uint](data, "TenantID"),
		Title:       dictutil.GetValue[string](data, "Title"),
		Description: dictutil.GetValue[string](data, "Description"),
		Avatar:      dictutil.GetValue[string](data, "Avatar"),
		Order:       dictutil.GetValue[int](data, "Order"),
	}
}

type WorkspaceUser struct {
	ID          uint         `gorm:"primaryKey" json:"id"`
	WorkspaceID uint         `gorm:"not null;uniqueIndex:idx_workspace_user" json:"workspace_id"`
	Workspace   Workspace    `gorm:"foreignKey:WorkspaceID;constraint:OnDelete:CASCADE;" json:"workspace"`
	UserID      uint         `gorm:"not null;uniqueIndex:idx_workspace_user" json:"user_id"`
	User        schema.User  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"user"`
	CreatorID   uint         `gorm:"not null" json:"creator_id"`
	Creator     *schema.User `gorm:"foreignKey:CreatorID;constraint:OnDelete:SET NULL;" json:"creator"`
	CreatedAt   time.Time    `json:"created_at"`
}

func (WorkspaceUser) TableName() string {
	return "workspaces_users"
}

func NewWorkspaceUser(data ctype.Dict) *WorkspaceUser {
	return &WorkspaceUser{
		WorkspaceID: dictutil.GetValue[uint](data, "WorkspaceID"),
		UserID:      dictutil.GetValue[uint](data, "UserID"),
		CreatorID:   dictutil.GetValue[uint](data, "CreatorID"),
	}
}
