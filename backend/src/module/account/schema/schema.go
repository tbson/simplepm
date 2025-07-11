package schema

import (
	"encoding/json"
	"strings"
	"time"

	"src/common/ctype"
	"src/util/dictutil"
	"src/util/pwdutil"

	"gorm.io/datatypes"
)

type Tenant struct {
	ID          uint         `gorm:"primaryKey" json:"id"`
	Roles       []Role       `gorm:"constraint:OnDelete:CASCADE;" json:"roles"`
	Users       []User       `gorm:"constraint:OnDelete:CASCADE;" json:"users"`
	GitAccounts []GitAccount `gorm:"constraint:OnDelete:CASCADE;" json:"git_accounts"`
	Uid         string       `gorm:"type:text;not null;unique" json:"uid"`
	Title       string       `gorm:"type:text;not null" json:"title"`
	Avatar      string       `gorm:"type:text;not null;default:''" json:"avatar"`
	AvatarStr   string       `gorm:"type:text;not null;default:''" json:"avatar_str"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

func NewTenant(data ctype.Dict) *Tenant {
	return &Tenant{
		Uid:       dictutil.GetValue[string](data, "Uid"),
		Title:     dictutil.GetValue[string](data, "Title"),
		Avatar:    dictutil.GetValue[string](data, "Avatar"),
		AvatarStr: dictutil.GetValue[string](data, "AvatarStr"),
	}
}

type User struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	TenantID       uint           `gorm:"not null;uniqueIndex:idx_users_tenant_external;uniqueIndex:idx_users_tenant_email" json:"tenant_id"`
	Tenant         Tenant         `gorm:"constraint:OnDelete:CASCADE;" json:"tenant"`
	TenantTmpID    *uint          `json:"tenant_tmp_id"`
	Sub            *string        `gorm:"type:text;default:null;unique" json:"sub"`
	Pwd            string         `gorm:"type:text;not null;default:''" json:"pwd"`
	PwdResetToken  string         `gorm:"type:text;not null;default:''" json:"pwd_reset_token"`
	PwdResetAt     *time.Time     `gorm:"type:timestamp;default:null" json:"pwd_reset_at"`
	Roles          []Role         `gorm:"many2many:users_roles;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"roles"`
	ExternalID     string         `gorm:"type:text;not null;uniqueIndex:idx_users_tenant_external" json:"uid"`
	Email          string         `gorm:"type:text;not null;uniqueIndex:idx_users_tenant_email" json:"email"`
	Mobile         *string        `gorm:"type:text" json:"mobile"`
	FirstName      string         `gorm:"type:text;not null;default:''" json:"first_name"`
	LastName       string         `gorm:"type:text;not null;default:''" json:"last_name"`
	Avatar         string         `gorm:"type:text;not null;default:''" json:"avatar"`
	GithubUsername string         `gorm:"type:text;not null;default:''" json:"github_username"`
	Color          string         `gorm:"type:text;not null;default:''" json:"color"`
	ExtraInfo      datatypes.JSON `gorm:"type:json;not null;default:'{}'" json:"extra_info"`
	Admin          bool           `gorm:"type:boolean;not null;default:false" json:"admin"`
	LockedAt       *time.Time     `gorm:"type:timestamp;default:null" json:"locked_at"`
	LockedReason   string         `gorm:"type:text;not null;default:''" json:"locked_reason"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
}

func (u *User) FullName() string {
	result := u.FirstName + " " + u.LastName
	return strings.TrimSpace(result)
}

func NewUser(data ctype.Dict) *User {
	extraInfoJSON, err := json.Marshal(data["ExtraInfo"])
	if err != nil {
		panic("Failed to marshal ExtraInfo")
	}

	pwd := pwdutil.MakePwd(dictutil.GetValue[string](data, "Pwd"))

	result := &User{
		TenantID:    dictutil.GetValue[uint](data, "TenantID"),
		TenantTmpID: dictutil.GetValue[*uint](data, "TenantTmpID"),
		ExternalID:  dictutil.GetValue[string](data, "ExternalID"),
		Sub:         dictutil.GetValue[*string](data, "Sub"),
		Email:       dictutil.GetValue[string](data, "Email"),
		Mobile:      dictutil.GetValue[*string](data, "Mobile"),
		FirstName:   dictutil.GetValue[string](data, "FirstName"),
		LastName:    dictutil.GetValue[string](data, "LastName"),
		Avatar:      dictutil.GetValue[string](data, "Avatar"),
		Color:       dictutil.GetValue[string](data, "Color"),
		ExtraInfo:   datatypes.JSON(extraInfoJSON),
		Roles:       dictutil.GetValue[[]Role](data, "Roles"),
	}

	if pwd != "" {
		result.Pwd = pwd
	}

	return result
}

type Role struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Users     []User    `gorm:"many2many:users_roles;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"users"`
	Pems      []Pem     `gorm:"many2many:roles_pems;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"pems"`
	TenantID  uint      `gorm:"not null;uniqueIndex:idx_roles_tenant_title" json:"tenant_id"`
	Tenant    Tenant    `gorm:"constraint:OnDelete:CASCADE;" json:"tenant"`
	Title     string    `gorm:"type:text;not null;uniqueIndex:idx_roles_tenant_title" json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewRole(data ctype.Dict) *Role {
	return &Role{
		TenantID: dictutil.GetValue[uint](data, "TenantID"),
		Title:    dictutil.GetValue[string](data, "Title"),
		Pems:     dictutil.GetValue[[]Pem](data, "Pems"),
	}
}

type Pem struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Roles  []Role `gorm:"many2many:roles_pems;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"roles"`
	Title  string `gorm:"type:text;not null" json:"title"`
	Module string `gorm:"type:text;not null;uniqueIndex:idx_pems_module_action" json:"module"`
	Action string `gorm:"type:text;not null;uniqueIndex:idx_pems_module_action" json:"action"`
	Admin  bool   `gorm:"type:boolean;not null;default:false" json:"admin"`
}

func NewPem(data ctype.Dict) *Pem {
	return &Pem{
		Title:  dictutil.GetValue[string](data, "Title"),
		Module: dictutil.GetValue[string](data, "Module"),
		Action: dictutil.GetValue[string](data, "Action"),
		Admin:  dictutil.GetValue[bool](data, "Admin"),
	}
}

type GitAccount struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	TenantID  *uint     `gorm:"default:null;" json:"tenant_id"`
	Tenant    *Tenant   `gorm:"constraint:OnDelete:CASCADE;" json:"tenant"`
	Uid       *string   `gorm:"type:text;default:null" json:"uid"`
	Title     string    `gorm:"type:text;not null;default:''" json:"title"`
	Avatar    string    `gorm:"type:text;not null;default:''" json:"avatar"`
	Type      string    `gorm:"type:text;not null;default:'GITHUB';check:type IN ('GITHUB', 'GITLAB')" json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewGitAccount(data ctype.Dict) *GitAccount {
	return &GitAccount{
		TenantID: dictutil.GetValue[*uint](data, "TenantID"),
		Uid:      dictutil.GetValue[*string](data, "Uid"),
		Title:    dictutil.GetValue[string](data, "Title"),
		Avatar:   dictutil.GetValue[string](data, "Avatar"),
		Type:     dictutil.GetValue[string](data, "Type"),
	}
}

type GitRepo struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	GitAccountID uint       `json:"git_account_id"`
	GitAccount   GitAccount `gorm:"constraint:OnDelete:CASCADE;" json:"git_account"`
	RepoID       string     `gorm:"type:text;not null" json:"repo_id"`
	Uid          string     `gorm:"type:text;not null" json:"uid"`
	Private      bool       `gorm:"type:boolean;not null" json:"private"`
}

func NewGitRepo(data ctype.Dict) *GitRepo {
	return &GitRepo{
		GitAccountID: dictutil.GetValue[uint](data, "GitAccountID"),
		RepoID:       dictutil.GetValue[string](data, "RepoID"),
		Uid:          dictutil.GetValue[string](data, "Uid"),
		Private:      dictutil.GetValue[bool](data, "Private"),
	}
}
