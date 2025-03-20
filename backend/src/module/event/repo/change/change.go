package change

import (
	"src/common/ctype"
	"src/module/event/schema"
	"src/util/dictutil"
	"src/util/errutil"
	"src/util/localeutil"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"gorm.io/gorm"
)

type Schema = schema.Change

var newSchema = schema.NewChange

type repo struct {
	client *gorm.DB
}

func New(client *gorm.DB) *repo {
	return &repo{client: client}
}

func (r *repo) WithTx(tx *gorm.DB) {
	r.client = tx
}

func (r *repo) List(opts ctype.QueryOpts) ([]Schema, error) {
	db := r.client
	if opts.Order == "" {
		db = db.Order("id DESC")
	} else {
		db = db.Order(opts.Order)
	}
	filters := dictutil.DictCamelToSnake(opts.Filters)
	preloads := opts.Preloads
	if len(preloads) > 0 {
		for _, preload := range preloads {
			db = db.Preload(preload)
		}
	}

	var items []Schema

	if len(filters) > 0 {
		db = db.Where(map[string]interface{}(filters))
	}
	result := db.Find(&items)
	err := result.Error
	if err != nil {
		return items, errutil.NewGormError(err)
	}
	return items, err
}

func (r *repo) Retrieve(opts ctype.QueryOpts) (*Schema, error) {
	db := r.client
	localizer := localeutil.Get()
	filters := dictutil.DictCamelToSnake(opts.Filters)
	preloads := opts.Preloads
	if len(preloads) > 0 {
		for _, preload := range preloads {
			db = db.Preload(preload)
		}
	}

	var item Schema
	var count int64
	query := db.Where(map[string]interface{}(filters))
	query.Model(&Schema{}).Count(&count)
	if count == 0 {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.NoRecordFound,
		})
		return &item, errutil.New("", []string{msg})
	}
	if count > 1 {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.MultipleRecordsFound,
		})
		return &item, errutil.New("", []string{msg})
	}

	result := query.First(&item)
	err := result.Error
	if err != nil {
		return &item, errutil.NewGormError(err)
	}
	return &item, err
}

func (r *repo) Create(data ctype.Dict) (*Schema, error) {
	item := newSchema(data)
	result := r.client.Create(item)
	err := result.Error
	if err != nil {
		return item, errutil.NewGormError(err)
	}
	return item, err
}
