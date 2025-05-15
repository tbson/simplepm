package config

import (
	"src/common/ctype"
	"src/util/dictutil"
)

const STRING = "STRING"
const INTEGER = "INTEGER"
const FLOAT = "FLOAT"
const BOOLEAN = "BOOLEAN"
const DATE = "DATE"
const DATETIME = "DATETIME"

var VariableDataTypeDict = ctype.StrDict{
	STRING:   "String",
	INTEGER:  "Integer",
	FLOAT:    "Float",
	BOOLEAN:  "Boolean",
	DATE:     "Date",
	DATETIME: "Date time",
}

var VariableDataTypeOptions = dictutil.StrDictToSelectOptions(VariableDataTypeDict)
