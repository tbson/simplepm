package errutil

import (
	"errors"
	"fmt"
	"strings"

	"src/common/ctype"
	"src/util/localeutil"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

const DEFAULT_FIELD = "detail"

type errorItem struct {
	Field    string   `json:"field"`
	Messages []ErrMsg `json:"messages"`
}

type ErrMsg struct {
	MsgCode *i18n.Message `json:"msg_code"`
	Args    ctype.Dict    `json:"args"`
}

type CustomError struct {
	Errors []errorItem `json:"errors"`
}

func (e *CustomError) Error() string {
	if e == nil || len(e.Errors) == 0 {
		return "empty custom error" // Handle the nil or empty case safely
	}

	// This string is mainly for debugging/logging, not the final user message
	var sb strings.Builder
	sb.WriteString("structured error:")
	for i, item := range e.Errors {
		if i > 0 {
			sb.WriteString(";") // Separator between field errors
		}
		sb.WriteString(" Field '")
		sb.WriteString(item.Field)
		sb.WriteString("' errors: ")
		if len(item.Messages) > 0 {
			// Just append the first message code for brevity in the summary
			sb.WriteString(item.Messages[0].MsgCode.ID)
			if len(item.Messages) > 1 {
				sb.WriteString(fmt.Sprintf(" (+%d more)", len(item.Messages)-1))
			}
		} else {
			sb.WriteString("no messages")
		}
	}
	return sb.String()
}

type localizedErrorItem struct {
	Field    string   `json:"field"`
	Messages []string `json:"messages"`
}

type localizedError struct {
	Errors []localizedErrorItem `json:"errors"`
}

func (e *CustomError) Localize() localizedError {
	localizer := localeutil.Get()
	localizedErrors := make([]localizedErrorItem, len(e.Errors))

	for i, item := range e.Errors {
		localizedMessages := make([]string, len(item.Messages))
		for j, msg := range item.Messages {
			localizedMessages[j] = localizer.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: msg.MsgCode,
				TemplateData:   msg.Args,
			})
		}
		localizedErrors[i] = localizedErrorItem{
			Field:    item.Field,
			Messages: localizedMessages,
		}
	}

	return localizedError{Errors: localizedErrors}
}

func (e *CustomError) Update(field string, msgCode *i18n.Message) *CustomError {
	errMsg := ErrMsg{
		MsgCode: msgCode,
		Args:    ctype.Dict{},
	}
	return buildError(e, field, errMsg)
}

func (e *CustomError) UpdateWithArgs(field string, msgCode *i18n.Message, args ctype.Dict) *CustomError {
	errMsg := ErrMsg{
		MsgCode: msgCode,
		Args:    args,
	}
	return buildError(e, field, errMsg)
}

func (e *CustomError) Merge(err *CustomError) *CustomError {
	if err == nil {
		return e
	}

	for _, item := range err.Errors {
		foundIndex := -1
		for i, existingItem := range e.Errors {
			if existingItem.Field == item.Field {
				foundIndex = i
				break
			}
		}

		if foundIndex != -1 {
			e.Errors[foundIndex].Messages = append(e.Errors[foundIndex].Messages, item.Messages...)
		} else {
			e.Errors = append(e.Errors, item)
		}
	}

	return e
}

func NewStandard(field string, msgCode *i18n.Message, args ctype.Dict) *CustomError {
	err := &CustomError{}
	if msgCode == nil {
		msgCode = localeutil.UnknownError
	}
	errMsg := ErrMsg{
		MsgCode: msgCode,
		Args:    args,
	}
	return buildError(err, field, errMsg)
}

func NewEmpty() *CustomError {
	return &CustomError{}
}

func New(msgCode *i18n.Message) *CustomError {
	return NewStandard(DEFAULT_FIELD, msgCode, ctype.Dict{})
}

func NewRaw(msg string) *CustomError {
	return NewWithArgs(localeutil.RawMsg, ctype.Dict{"Msg": msg})
}

func NewWithArgs(msgCode *i18n.Message, args ctype.Dict) *CustomError {
	return NewStandard(DEFAULT_FIELD, msgCode, args)
}

func buildError(err *CustomError, field string, errMsg ErrMsg) *CustomError {
	if err == nil {
		err = &CustomError{}
	}

	if field == "" {
		field = DEFAULT_FIELD
	}

	// Try to find an existing error item for this field
	foundIndex := -1
	for i, item := range err.Errors {
		if item.Field == field {
			foundIndex = i
			break
		}
	}

	if foundIndex != -1 {
		// If found, append the new message to the existing item's messages
		err.Errors[foundIndex].Messages = append(err.Errors[foundIndex].Messages, errMsg)
	} else {
		// If not found, create a new errorItem and append it to the slice
		err.Errors = append(err.Errors, errorItem{
			Field:    field,
			Messages: []ErrMsg{errMsg},
		})
	}

	return err
}

func getGormField(message string) string {
	field := message
	field = strings.Split(field, "(")[1]
	field = strings.Split(field, ")")[0]
	return field
}

func NewGormError(err error) *CustomError {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {

		switch pgErr.Code {
		case "23505":
			msg := ErrMsg{
				MsgCode: localeutil.GormDuplicateKey,
				Args:    ctype.Dict{},
			}

			return &CustomError{
				Errors: []errorItem{
					{
						Field:    getGormField(pgErr.Detail),
						Messages: []ErrMsg{msg},
					},
				},
			}
		}
	}
	return &CustomError{
		Errors: []errorItem{
			{
				Field:    "detail",
				Messages: []ErrMsg{{localeutil.RawMsg, ctype.Dict{"Msg": err.Error()}}},
			},
		},
	}
}
