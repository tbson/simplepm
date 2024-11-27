package schema

import (
	"encoding/json"
	"time"

	"src/common/ctype"
	"src/util/dictutil"

	"gorm.io/datatypes"
)

type AuthClient struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Tenants     []Tenant  `gorm:"constraint:OnDelete:SET NULL;" json:"tenants"`
	Uid         string    `gorm:"type:text;not null;unique" json:"uid"`
	Description string    `gorm:"type:text;not null;default:''" json:"description"`
	Secret      string    `gorm:"type:text;not null" json:"secret"`
	Partition   string    `gorm:"type:text;not null" json:"partition"`
	Default     bool      `gorm:"type:boolean;not null;default:false" json:"default"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewAuthClient(data ctype.Dict) *AuthClient {
	return &AuthClient{
		Uid:         dictutil.GetValue[string](data, "Uid"),
		Description: dictutil.GetValue[string](data, "Description"),
		Secret:      dictutil.GetValue[string](data, "Secret"),
		Partition:   dictutil.GetValue[string](data, "Partition"),
		Default:     dictutil.GetValue[bool](data, "Default"),
	}
}

type Tenant struct {
	ID           uint        `gorm:"primaryKey" json:"id"`
	AuthClientID uint        `json:"auth_client_id"`
	AuthClient   *AuthClient `json:"auth_client"`
	Roles        []Role      `gorm:"constraint:OnDelete:CASCADE;" json:"roles"`
	Users        []User      `gorm:"constraint:OnDelete:CASCADE;" json:"users"`
	Uid          string      `gorm:"type:text;not null;unique" json:"uid"`
	Title        string      `gorm:"type:text;not null" json:"title"`
	Avatar       string      `gorm:"type:text;not null;default:''" json:"avatar"`
	AvatarStr    string      `gorm:"type:text;not null;default:''" json:"avatar_str"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
}

func NewTenant(data ctype.Dict) *Tenant {
	return &Tenant{
		AuthClientID: dictutil.GetValue[uint](data, "AuthClientID"),
		Uid:          dictutil.GetValue[string](data, "Uid"),
		Title:        dictutil.GetValue[string](data, "Title"),
		Avatar:       dictutil.GetValue[string](data, "Avatar"),
		AvatarStr:    dictutil.GetValue[string](data, "AvatarStr"),
	}
}

type User struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	TenantID     uint           `gorm:"not null;uniqueIndex:idx_users_tenant_external;uniqueIndex:idx_users_tenant_email" json:"tenant_id"`
	Tenant       Tenant         `gorm:"constraint:OnDelete:CASCADE;" json:"tenant"`
	TenantTmpID  *uint          `json:"tenant_tmp_id"`
	Sub          *string        `gorm:"type:text;default:null;unique" json:"sub"`
	Roles        []Role         `gorm:"many2many:users_roles;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"roles"`
	ExternalID   string         `gorm:"type:text;not null;uniqueIndex:idx_users_tenant_external" json:"uid"`
	Email        string         `gorm:"type:text;not null;uniqueIndex:idx_users_tenant_email" json:"email"`
	Mobile       *string        `gorm:"type:text" json:"mobile"`
	FirstName    string         `gorm:"type:text;not null;default:''" json:"first_name"`
	LastName     string         `gorm:"type:text;not null;default:''" json:"last_name"`
	Avatar       string         `gorm:"type:text;not null;default:''" json:"avatar"`
	AvatarStr    string         `gorm:"type:text;not null;default:''" json:"avatar_str"`
	ExtraInfo    datatypes.JSON `gorm:"type:json;not null;default:'{}'" json:"extra_info"`
	Admin        bool           `gorm:"type:boolean;not null;default:false" json:"admin"`
	LockedAt     *time.Time     `gorm:"type:timestamp;default:null" json:"locked_at"`
	LockedReason string         `gorm:"type:text;not null;default:''" json:"locked_reason"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

func NewUser(data ctype.Dict) *User {
	extraInfoJSON, err := json.Marshal(data["ExtraInfo"])
	if err != nil {
		panic("Failed to marshal ExtraInfo")
	}
	return &User{
		TenantID:    dictutil.GetValue[uint](data, "TenantID"),
		TenantTmpID: dictutil.GetValue[*uint](data, "TenantTmpID"),
		ExternalID:  dictutil.GetValue[string](data, "ExternalID"),
		Email:       dictutil.GetValue[string](data, "Email"),
		Mobile:      dictutil.GetValue[*string](data, "Mobile"),
		FirstName:   dictutil.GetValue[string](data, "FirstName"),
		LastName:    dictutil.GetValue[string](data, "LastName"),
		Avatar:      dictutil.GetValue[string](data, "Avatar"),
		AvatarStr:   dictutil.GetValue[string](data, "AvatarStr"),
		ExtraInfo:   datatypes.JSON(extraInfoJSON),
		Roles:       dictutil.GetValue[[]Role](data, "Roles"),
	}
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
