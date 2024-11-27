package dictutil

import (
	"reflect"
	"src/common/ctype"
	"src/util/stringutil"
)

func StructToDict(obj interface{}) ctype.Dict {
	result := make(ctype.Dict)
	val := reflect.ValueOf(obj)

	// Iterate through the struct fields
	for i := 0; i < val.NumField(); i++ {
		// Get the struct field name and value
		fieldName := reflect.TypeOf(obj).Field(i).Name
		fieldValue := val.Field(i).Interface()
		result[fieldName] = fieldValue
	}

	return result
}

func DictCamelToSnake(data ctype.Dict) ctype.Dict {
	result := make(ctype.Dict)
	for k, v := range data {
		result[stringutil.ToSnakeCase(k)] = v
	}
	return result
}

func GetValue[T any](data ctype.Dict, key string) T {
	var defaultValue T
	if val, ok := data[key]; ok {
		return val.(T)
	}
	return defaultValue
}
