package config

import "src/common/ctype"

const VDTString = "STRING"
const VDTInteger = "INTEGER"
const VDTFloat = "FLOAT"
const VDTBoolean = "BOOLEAN"
const VDTDate = "DATE"
const VDTDateTime = "DATETIME"

type option = ctype.SelectOption[string]

var VariableDataTypeOptions = []option{
	{Value: VDTString, Label: "String"},
	{Value: VDTInteger, Label: "Integer"},
	{Value: VDTFloat, Label: "Float"},
	{Value: VDTBoolean, Label: "Boolean"},
	{Value: VDTDate, Label: "Date"},
	{Value: VDTDateTime, Label: "Date time"},
}
