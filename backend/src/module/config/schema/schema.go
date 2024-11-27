package schema

import (
	"src/common/ctype"
	"src/util/dictutil"
	"src/util/iterutil"
	"time"
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
	ID          uint      `gorm:"primaryKey" json:"id"`
	Key         string    `gorm:"type:text;not null;unique" json:"key"`
	Value       string    `gorm:"type:text;not null;default:''" json:"value"`
	Description string    `gorm:"type:text;not null;default:''" json:"description"`
	DataType    string    `gorm:"type:text;not null;default:'STRING';check:data_type IN ('STRING', 'INTEGER', 'FLOAT', 'BOOLEAN', 'DATE', 'DATETIME')" json:"data_type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewVariable(data ctype.Dict) *Variable {
	return &Variable{
		Key:         dictutil.GetValue[string](data, "Key"),
		Value:       dictutil.GetValue[string](data, "Value"),
		Description: dictutil.GetValue[string](data, "Description"),
		DataType:    dictutil.GetValue[string](data, "DataType"),
	}
}
