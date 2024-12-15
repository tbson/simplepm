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
	Tenant      schema.Tenant `gorm:"foreignKey:TenantID" json:"tenant"`
	Projects    []Project     `gorm:"constraint:OnDelete:CASCADE;" json:"projects"`
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
	Workspace   Workspace    `gorm:"foreignKey:WorkspaceID" json:"workspace"`
	UserID      uint         `gorm:"not null;uniqueIndex:idx_workspace_user" json:"user_id"`
	User        schema.User  `gorm:"foreignKey:UserID" json:"user"`
	CreatorID   *uint        `gorm:"default:null" json:"creator_id"`
	Creator     *schema.User `gorm:"foreignKey:CreatorID" json:"creator"`
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
	Tenant      schema.Tenant `gorm:"foreignKey:TenantID" json:"tenant"`
	WorkspaceID *uint         `gorm:"default:null" json:"workspace_id"`
	Workspace   *Workspace    `gorm:"foreignKey:WorkspaceID" json:"workspace"`
	Tasks       []Task        `gorm:"constraint:OnDelete:CASCADE;" json:"tasks"`
	TaskFields  []TaskField   `gorm:"constraint:OnDelete:CASCADE;" json:"task_fields"`
	Title       string        `gorm:"type:text;not null" json:"title"`
	Description string        `gorm:"ntype:text;ot null;default:''" json:"description"`
	Avatar      string        `gorm:"type:text;not null;default:''" json:"avatar"`
	Layout      string        `gorm:"type:text;not null;default:'TABLE';check:layout IN ('TABLE', 'KANBAN', 'ROADMAP')" json:"layout"`
	Status      string        `gorm:"type:text;not null;default:'ACTIVE';check:status IN ('ACTIVE', 'FINISHED', 'ARCHIEVED')" json:"status"`
	Order       int           `gorm:"not null;default:0" json:"order"`
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
		Status:      dictutil.GetValue[string](data, "Status"),
		Order:       dictutil.GetValue[int](data, "Order"),
		FinishedAt:  dictutil.GetValue[*time.Time](data, "FinishedAt"),
	}
}

type ProjectUser struct {
	ID        uint         `gorm:"primaryKey" json:"id"`
	ProjectID uint         `gorm:"not null;uniqueIndex:idx_project_user" json:"project_id"`
	Project   Project      `gorm:"foreignKey:ProjectID" json:"project"`
	UserID    uint         `gorm:"not null;uniqueIndex:idx_project_user" json:"user_id"`
	User      schema.User  `gorm:"foreignKey:UserID" json:"user"`
	CreatorID *uint        `gorm:"default:null" json:"creator_id"`
	Creator   *schema.User `gorm:"foreignKey:CreatorID" json:"creator"`
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
	ID               uint              `gorm:"primaryKey" json:"id"`
	ProjectID        uint              `gorm:"not null" json:"project_id"`
	Project          Project           `gorm:"foreignKey:ProjectID" json:"project"`
	TaskFieldOptions []TaskFieldOption `gorm:"constraint:OnDelete:CASCADE;" json:"task_field_options"`
	TaskFieldValues  []TaskFieldValue  `gorm:"constraint:OnDelete:CASCADE;" json:"task_field_values"`
	Title            string            `gorm:"type:text;not null" json:"title"`
	Type             string            `gorm:"type:text;not null;default:'TEXT';check:type IN ('TEXT', 'NUMBER', 'DATE', 'SELECT', 'MULTIPLE_SELECT')" json:"type"`
	Description      string            `gorm:"type:text;not null;default:''" json:"description"`
	Order            int               `gorm:"not null;default:0" json:"order"`
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
	ID              uint             `gorm:"primaryKey" json:"id"`
	TaskFieldID     uint             `gorm:"not null" json:"task_field_id"`
	TaskField       TaskField        `gorm:"foreignKey:TaskFieldID" json:"task_field"`
	TaskFieldValues []TaskFieldValue `gorm:"constraint:OnDelete:CASCADE;" json:"task_field_values"`
	Title           string           `gorm:"type:text;not null" json:"title"`
	Description     string           `gorm:"type:text;default:''" json:"description"`
	Color           string           `gorm:"type:text;default:''" json:"color"`
	Order           int              `gorm:"not null;default:0" json:"order"`
}

func NewTaskFieldOption(data ctype.Dict) *TaskFieldOption {
	return &TaskFieldOption{
		TaskFieldID: dictutil.GetValue[uint](data, "TaskFieldID"),
		Title:       dictutil.GetValue[string](data, "Title"),
		Order:       dictutil.GetValue[int](data, "Order"),
	}
}

type Feature struct {
	ID          uint          `gorm:"primaryKey" json:"id"`
	ProjectID   uint          `gorm:"not null" json:"project_id"`
	Project     schema.Tenant `gorm:"foreignKey:ProjectID" json:"project"`
	Tasks       []Task        `gorm:"constraint:OnDelete:CASCADE;" json:"tasks"`
	Title       string        `gorm:"type:text;not null" json:"title"`
	Description string        `gorm:"ntype:text;ot null;default:''" json:"description"`
	Status      string        `gorm:"type:text;not null;default:'ACTIVE';check:status IN ('ACTIVE', 'FINISHED', 'ARCHIEVED')" json:"status"`
	Order       int           `gorm:"not null;default:0" json:"order"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}

func NewFeature(data ctype.Dict) *Feature {
	return &Feature{
		ProjectID:   dictutil.GetValue[uint](data, "ProjectID"),
		Title:       dictutil.GetValue[string](data, "Title"),
		Description: dictutil.GetValue[string](data, "Description"),
		Status:      dictutil.GetValue[string](data, "Status"),
		Order:       dictutil.GetValue[int](data, "Order"),
	}
}

type Task struct {
	ID              uint             `gorm:"primaryKey" json:"id"`
	ProjectID       uint             `gorm:"not null" json:"project_id"`
	Project         schema.Tenant    `gorm:"foreignKey:ProjectID" json:"project"`
	FeatureID       uint             `gorm:"not null" json:"feature_id"`
	Feature         schema.Tenant    `gorm:"foreignKey:FeatureID" json:"feature"`
	UserID          *uint            `gorm:"default:null" json:"user_id"`
	User            *schema.User     `gorm:"foreignKey:UserID" json:"user"`
	TaskFieldValues []TaskFieldValue `gorm:"constraint:OnDelete:CASCADE;" json:"task_field_values"`
	Title           string           `gorm:"type:text;not null" json:"title"`
	Description     string           `gorm:"ntype:text;ot null;default:''" json:"description"`
	Order           int              `gorm:"not null;default:0" json:"order"`
	CreatedAt       time.Time        `json:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at"`
}

func NewTask(data ctype.Dict) *Task {
	return &Task{
		ProjectID:   dictutil.GetValue[uint](data, "ProjectID"),
		FeatureID:   dictutil.GetValue[uint](data, "FeatureID"),
		UserID:      dictutil.GetValue[*uint](data, "UserID"),
		Title:       dictutil.GetValue[string](data, "Title"),
		Description: dictutil.GetValue[string](data, "Description"),
		Order:       dictutil.GetValue[int](data, "Order"),
	}
}

type TaskFieldValue struct {
	ID                uint             `gorm:"primaryKey" json:"id"`
	TaskID            uint             `gorm:"not null" json:"task_id"`
	Task              Task             `gorm:"foreignKey:TaskID" json:"task"`
	TaskFieldID       uint             `gorm:"not null" json:"task_field_id"`
	TaskField         TaskField        `gorm:"foreignKey:TaskFieldID" json:"task_field"`
	TaskFieldOptionID *uint            `gorm:"default:null" json:"task_field_option_id"`
	TaskFieldOption   *TaskFieldOption `gorm:"foreignKey:TaskFieldOptionID" json:"task_field_option"`
	NumberValue       *int             `gorm:"default:null" json:"number_value"`
	DateValue         *time.Time       `gorm:"default:null" json:"date_value"`
	Value             string           `gorm:"type:text;not null" json:"value"`
	CreatedAt         time.Time        `json:"created_at"`
	UpdatedAt         time.Time        `json:"updated_at"`
}

func NewTaskFieldValue(data ctype.Dict) *TaskFieldValue {
	return &TaskFieldValue{
		TaskID:            dictutil.GetValue[uint](data, "TaskID"),
		TaskFieldID:       dictutil.GetValue[uint](data, "TaskFieldID"),
		TaskFieldOptionID: dictutil.GetValue[*uint](data, "TaskFieldOptionID"),
		NumberValue:       dictutil.GetValue[*int](data, "NumberValue"),
		DateValue:         dictutil.GetValue[*time.Time](data, "DateValue"),
		Value:             dictutil.GetValue[string](data, "Value"),
	}
}
