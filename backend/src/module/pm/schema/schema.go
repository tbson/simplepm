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
	Title       string        `gorm:"type:text;not null" json:"title"`
	Description string        `gorm:"type:text;not null;default:''" json:"description"`
	Avatar      string        `gorm:"type:text;not null;default:''" json:"avatar"`
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
	CreatorID   *uint        `gorm:"default:null" json:"creator_id"`
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
		CreatorID:   dictutil.GetValue[*uint](data, "CreatorID"),
	}
}

type Project struct {
	ID          uint          `gorm:"primaryKey" json:"id"`
	TenantID    uint          `gorm:"not null" json:"tenant_id"`
	Tenant      schema.Tenant `gorm:"foreignKey:TenantID;constraint:OnDelete:CASCADE;" json:"tenant"`
	WorkspaceID *uint         `gorm:"default:null" json:"workspace_id"`
	Workspace   *Workspace    `gorm:"foreignKey:WorkspaceID;constraint:OnDelete:CASCADE;" json:"workspace"`
	Title       string        `gorm:"type:text;not null" json:"title"`
	Description string        `gorm:"ntype:text;ot null;default:''" json:"description"`
	Avatar      string        `gorm:"type:text;not null;default:''" json:"avatar"`
	Layout      string        `gorm:"type:text;not null;default:'TABLE';check:layout IN ('TABLE', 'KANBAN', 'ROADMAP')" json:"layout"`
	Order       int           `gorm:"not null;default:0" json:"order"`
	StartDate   *time.Time    `json:"start_date"`
	TargetDate  *time.Time    `json:"target_date"`
	FinishedAt  *time.Time    `json:"finished_at"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}

func NewProject(data ctype.Dict) *Project {
	return &Project{
		TenantID:    dictutil.GetValue[uint](data, "TenantID"),
		WorkspaceID: dictutil.GetValue[*uint](data, "WorkspaceID"),
		Title:       dictutil.GetValue[string](data, "Title"),
		Description: dictutil.GetValue[string](data, "Description"),
		Avatar:      dictutil.GetValue[string](data, "Avatar"),
		Layout:      dictutil.GetValue[string](data, "Layout"),
		Order:       dictutil.GetValue[int](data, "Order"),
		StartDate:   dictutil.GetValue[*time.Time](data, "StartDate"),
		TargetDate:  dictutil.GetValue[*time.Time](data, "TargetDate"),
		FinishedAt:  dictutil.GetValue[*time.Time](data, "FinishedAt"),
	}
}

type ProjectUser struct {
	ID        uint         `gorm:"primaryKey" json:"id"`
	ProjectID uint         `gorm:"not null;uniqueIndex:idx_project_user" json:"project_id"`
	Project   Project      `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE;" json:"project"`
	UserID    uint         `gorm:"not null;uniqueIndex:idx_project_user" json:"user_id"`
	User      schema.User  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"user"`
	CreatorID *uint        `gorm:"default:null" json:"creator_id"`
	Creator   *schema.User `gorm:"foreignKey:CreatorID;constraint:OnDelete:SET NULL;" json:"creator"`
	CreatedAt time.Time    `json:"created_at"`
}

func (ProjectUser) TableName() string {
	return "projects_users"
}

func NewProjectUser(data ctype.Dict) *ProjectUser {
	return &ProjectUser{
		ProjectID: dictutil.GetValue[uint](data, "ProjectID"),
		UserID:    dictutil.GetValue[uint](data, "UserID"),
		CreatorID: dictutil.GetValue[*uint](data, "CreatorID"),
	}
}

type TaskField struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	ProjectID   uint    `gorm:"not null" json:"project_id"`
	Project     Project `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE;" json:"project"`
	Title       string  `gorm:"type:text;not null" json:"title"`
	Type        string  `gorm:"type:text;not null" json:"type"`
	Description string  `gorm:"type:text;not null;default:''" json:"description"`
	Order       int     `gorm:"not null;default:0" json:"order"`
}

func NewTaskField(data ctype.Dict) *TaskField {
	return &TaskField{
		ProjectID:   dictutil.GetValue[uint](data, "ProjectID"),
		Title:       dictutil.GetValue[string](data, "Title"),
		Type:        dictutil.GetValue[string](data, "Type"),
		Description: dictutil.GetValue[string](data, "Description"),
		Order:       dictutil.GetValue[int](data, "Order"),
	}
}

type TaskFieldOption struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	TaskFieldID uint      `gorm:"not null" json:"task_field_id"`
	TaskField   TaskField `gorm:"foreignKey:TaskFieldID;constraint:OnDelete:CASCADE;" json:"task_field"`
	Title       string    `gorm:"type:text;not null" json:"title"`
	Order       int       `gorm:"not null;default:0" json:"order"`
}

func NewTaskFieldOption(data ctype.Dict) *TaskFieldOption {
	return &TaskFieldOption{
		TaskFieldID: dictutil.GetValue[uint](data, "TaskFieldID"),
		Title:       dictutil.GetValue[string](data, "Title"),
		Order:       dictutil.GetValue[int](data, "Order"),
	}
}
