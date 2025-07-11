package schema

import (
	"src/common/ctype"
	"src/util/dictutil"
	"src/util/iterutil"
	"time"

	"github.com/google/uuid"
)

var TypeDict = iterutil.FieldEnum{
	"STRING",
	"INTEGER",
	"FLOAT",
	"BOOLEAN",
	"DATE",
	"DATETIME",
}

var TypeOptions = iterutil.GetFieldOptions(TypeDict)

type Variable struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryUid;default:uuid_generate_v1mc()" json:"id"`
	Uid         string    `gorm:"type:text;not null;unique" json:"uid"`
	Value       string    `gorm:"type:text;not null;default:''" json:"value"`
	Description string    `gorm:"type:text;not null;default:''" json:"description"`
	DataType    string    `gorm:"type:text;not null;default:'STRING';check:data_type IN ('STRING', 'INTEGER', 'FLOAT', 'BOOLEAN', 'DATE', 'DATETIME')" json:"data_type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewVariable(data ctype.Dict) *Variable {
	return &Variable{
		Uid:         dictutil.GetValue[string](data, "Uid"),
		Value:       dictutil.GetValue[string](data, "Value"),
		Description: dictutil.GetValue[string](data, "Description"),
		DataType:    dictutil.GetValue[string](data, "DataType"),
	}
}
