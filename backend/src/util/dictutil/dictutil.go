package dictutil

import (
	"fmt"
	"reflect"
	"src/common/ctype"
	"src/util/stringutil"
)

func StructToDict(obj interface{}) ctype.Dict {
	result := make(ctype.Dict)
	if obj == nil {
		return result
	}
	val := reflect.ValueOf(obj)
	// If obj is a nil pointer, return an empty dict.
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return result
		}
		val = val.Elem()
	}
	// Return empty dict if not a struct
	if val.Kind() != reflect.Struct {
		return result
	}

	typ := val.Type()
	// Iterate through the struct fields
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if field.PkgPath != "" {
			continue
		}
		result[field.Name] = val.Field(i).Interface()
	}

	return result
}

func DictCamelToSnake(data ctype.Dict) ctype.Dict {
	result := make(ctype.Dict)
	for k, v := range data {
		result[stringutil.ToSnakeCaseEnd(k)] = v
	}
	return result
}

func GetValue[T any](data ctype.Dict, key string) T {
	if val, ok := data[key]; ok {
		if converted, ok := val.(T); ok {
			return converted
		}
	}
	var defaultValue T
	return defaultValue
}

func DiffDict(m1, m2 ctype.Dict) ctype.Dict {
	differences := make(ctype.Dict)
	for key, v1 := range m1 {
		if v2, ok := m2[key]; !ok || !reflect.DeepEqual(v1, v2) {
			differences[key] = v1
		}
	}
	return differences
}

func StrDictToSelectOptions(data ctype.StrDict) []ctype.SelectOption[string] {
	var result []ctype.SelectOption[string]
	for k, v := range data {
		value := fmt.Sprintf("%v", k)
		label := fmt.Sprintf("%v", v)
		result = append(result, ctype.SelectOption[string]{
			Value: value,
			Label: label,
		})
	}
	return result
}
