package vldtutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"reflect"
	"slices"
	"src/client/s3client"
	"src/common/ctype"
	"src/module/aws/repo/s3"
	"src/util/dictutil"
	"src/util/errutil"
	"src/util/localeutil"
	"src/util/stringutil"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func BytesToStruct[T any](data []byte, target T) (T, error) {
	localizer := localeutil.Get()
	if err := json.Unmarshal(data, &target); err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.CannotReadRequestBody,
		})
		return target, errutil.New("", []string{msg})
	}
	return target, nil
}

func ValidatePayload[T any](c echo.Context, target T) (T, error) {
	// result := ctype.Dict{}
	result := target
	localizer := localeutil.Get()
	// bind the payload to the target struct
	if err := c.Bind(&target); err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.CannotReadRequestBody,
		})
		return result, errutil.New("", []string{msg})
	}
	// Validate the struct
	if err := c.Validate(target); err != nil {
		// Map to collect messages per field
		error := errutil.CustomError{}
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, fe := range ve {
				// Map struct field name to JSON field name
				fieldName := fe.Field()
				structField, _ := reflect.TypeOf(target).FieldByName(fe.StructField())
				jsonTag := structField.Tag.Get("json")
				if jsonTag != "" {
					fieldName = strings.Split(jsonTag, ",")[0]
				}

				// Customize the error message based on the validation tag
				var msg string
				switch fe.Tag() {
				case "required":
					msg = localizer.MustLocalize(&i18n.LocalizeConfig{
						DefaultMessage: localeutil.FieldRequired,
					})
				case "oneof":
					msg = localizer.MustLocalize(&i18n.LocalizeConfig{
						DefaultMessage: localeutil.MustBeOneOf,
						TemplateData: ctype.Dict{
							"Values": fe.Param(),
						},
					})
				default:
					msg = localizer.MustLocalize(&i18n.LocalizeConfig{
						DefaultMessage: localeutil.InvalidValue,
					})
				}

				// Append the error message to the field's error list
				error.Add(fieldName, []string{msg})
			}
		} else {
			// For other errors, return a general message
			return result, errutil.New("", []string{err.Error()})
		}
		return result, &error
	}

	// return dictutil.StructToDict(target), nil
	return target, nil
}

func ValidateUpdatePayload[T any](c echo.Context, target T) (T, []string, error) {
	defaultFields := []string{}
	structResult := target
	localizer := localeutil.Get()

	fields, err := getFields(c)
	if err != nil {
		return structResult, defaultFields, err
	}

	if err := c.Bind(&target); err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.CannotReadRequestBody,
		})
		return structResult, defaultFields, errutil.New("", []string{msg})
	}

	return target, fields, nil
}

func GetDictByFields[T any](target T, fields []string, fieldModifier []string) ctype.Dict {
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

	data := dictutil.StructToDict(target)

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

func ValidateId(id string) uint {
	if id == "" {
		return 0
	}
	if id, err := strconv.Atoi(id); err == nil {
		return uint(id)
	}
	return 0
}

func ValidateIds(ids string) []uint {
	var idList []uint
	if ids == "" {
		return idList
	}
	for _, id := range strings.Split(ids, ",") {
		if id, err := strconv.Atoi(id); err == nil {
			idList = append(idList, uint(id))
		}
	}
	return idList
}

func getFiles(c echo.Context) (map[string][]*multipart.FileHeader, error) {
	if !strings.Contains(c.Request().Header.Get("Content-Type"), "multipart/form-data") {
		return map[string][]*multipart.FileHeader{}, nil
	}

	result := map[string][]*multipart.FileHeader{}
	localizer := localeutil.Get()
	form, err := c.MultipartForm()
	if err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.CannotReadRequestBody,
		})
		return result, errutil.New("", []string{msg})
	}

	// Add form data to keys map
	for k, v := range form.File {
		result[stringutil.ToCamelCase(k)] = v
	}

	return result, nil
}

func UploadAndUPdatePayload(
	c echo.Context,
	folder string,
	result ctype.Dict,
) (ctype.Dict, error) {
	if c.Request().Header.Get("Content-Type") == "application/json" {
		return result, nil
	}
	s3Repo := s3.New(s3client.NewClient())
	files, err := getFiles(c)
	if err != nil {
		return result, err
	}

	s3Result, err := s3Repo.Uploads(c.Request().Context(), folder, files)
	for k, v := range s3Result {
		result[k] = v.FileURL
		fmt.Println(k, v.FileURL)
	}
	return result, nil
}

func UploadAndGetMetadata(
	c echo.Context,
	folder string,
) ([]s3.FileInfo, error) {
	resultFiles := []s3.FileInfo{}
	if c.Request().Header.Get("Content-Type") == "application/json" {
		return resultFiles, nil
	}
	s3Repo := s3.New(s3client.NewClient())
	files, err := getFiles(c)
	if err != nil {
		return resultFiles, err
	}

	s3Result, err := s3Repo.Uploads(c.Request().Context(), folder, files)
	for _, v := range s3Result {
		resultFiles = append(resultFiles, v)
	}
	return resultFiles, nil
}

func getFields(c echo.Context) ([]string, error) {
	if c.Request().Header.Get("Content-Type") == "application/json" {
		return getJsonFields(c)
	}
	return getFormFields(c)
}

func getJsonFields(c echo.Context) ([]string, error) {
	result := []string{}
	localizer := localeutil.Get()
	bodyBytes, err := io.ReadAll(c.Request().Body)
	if err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.CannotReadRequestBody,
		})
		return result, errutil.New("", []string{msg})
	}

	// Reset the body so it can be read again if needed
	c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// Unmarshal into a map to get the keys present in the payload
	var payloadMap ctype.Dict
	if err := json.Unmarshal(bodyBytes, &payloadMap); err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.InvalidJSONPayload,
		})

		return result, errutil.New("", []string{msg})
	}

	// Extract only the keys
	var keyList []string
	for k := range payloadMap {
		keyList = append(keyList, k)
	}

	return keyList, nil
}

func getFormFields(c echo.Context) ([]string, error) {
	result := []string{}
	localizer := localeutil.Get()
	var keyList []string
	keys := make(map[string]interface{})
	form, err := c.FormParams()
	if err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.CannotReadRequestBody,
		})
		return result, errutil.New("", []string{msg})
	}

	// Add form data to keys map
	for k, v := range form {
		keys[k] = v
	}

	// Extract only the keys
	for k := range keys {
		keyList = append(keyList, k)
	}

	return keyList, nil
}

func CheckRequiredFilter(c echo.Context, param string) error {
	localizer := localeutil.Get()
	if c.QueryParam(param) == "" {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.MissingQueryParam,
			TemplateData: ctype.Dict{
				"Value": param,
			},
		})
		return errutil.New("", []string{msg})
	}
	return nil
}
