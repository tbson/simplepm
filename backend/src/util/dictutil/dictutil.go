package dictutil

import (
	"fmt"
	"reflect"
	"slices"
	"src/common/ctype"
	"src/util/stringutil"
	"strings"
)

func StructToDict(obj any) ctype.Dict {
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
	for i := range typ.NumField() {
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

func ParseStructWithFilters[T any](target T, fields []string, fieldModifier []string) ctype.Dict {
	// fieldModifer can be included fields or excluded fields, the excluded fields are the fields that prefix with "-"
	// if the fieldModifier item is prefixed with "-", the field will be excluded
	// if the fieldModifier item not prefixed with "-", the field will be included if it is not present in the fields variable
	// if the fieldModifier is empty, fields variable will be used

	newFields := []string{}

	includeFields := []string{}
	excludeFields := []string{}

	for _, field := range fieldModifier {
		if strings.HasPrefix(field, "-") {
			excludeFields = append(excludeFields, strings.TrimPrefix(field, "-"))
		} else {
			includeFields = append(includeFields, field)
		}
	}

	for _, field := range fields {
		if !slices.Contains(excludeFields, field) {
			newFields = append(newFields, field)
		}
	}

	for _, field := range includeFields {
		if !slices.Contains(newFields, field) {
			newFields = append(newFields, field)
		}
	}

	data := StructToDict(target)

	// remove the fields that are not present in the payload, check json tags
	for k := range data {
		structField, _ := reflect.TypeOf(target).FieldByName(k)
		jsonTag := structField.Tag.Get("json")
		if jsonTag != "" {
			fieldName := strings.Split(jsonTag, ",")[0]
			if !slices.Contains(newFields, fieldName) {
				delete(data, k)
			}
		}
	}

	return data
}

func ParseStructWithFields[T any](target T, fields []string) ctype.Dict {
	fieldModifier := []string{}
	return ParseStructWithFilters(target, fields, fieldModifier)
}
