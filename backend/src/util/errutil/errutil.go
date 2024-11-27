package errutil

import (
	"errors"
	"strings"

	"src/util/localeutil"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type errorItem struct {
	Field    string   `json:"field"`
	Messages []string `json:"messages"`
}

type CustomError struct {
	Errors []errorItem `json:"errors"`
}

func (e *CustomError) Error() string {
	return e.Errors[0].Messages[0]
}

func buildErrorItem(field string, messages []string) errorItem {
	if field == "" {
		field = "detail"
	}
	return errorItem{
		Field:    field,
		Messages: messages,
	}
}

func getGormField(message string) string {
	field := message
	field = strings.Split(field, "(")[1]
	field = strings.Split(field, ")")[0]
	return field
}

func NewGormError(err error) *CustomError {
	localizer := localeutil.Get()
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {

		switch pgErr.Code {
		case "23505":
			msg := localizer.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: localeutil.GormDuplicateKey,
			})

			return &CustomError{
				Errors: []errorItem{
					{
						Field:    getGormField(pgErr.Detail),
						Messages: []string{msg},
					},
				},
			}
		}
	}
	return &CustomError{
		Errors: []errorItem{
			{
				Field:    "detail",
				Messages: []string{err.Error()},
			},
		},
	}
}

func New(field string, messages []string) *CustomError {
	err := buildErrorItem(field, messages)
	return &CustomError{
		Errors: []errorItem{err},
	}
}

func (e *CustomError) Add(field string, messages []string) *CustomError {
	err := buildErrorItem(field, messages)
	e.Errors = append(e.Errors, err)
	return e
}
