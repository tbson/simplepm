package taskfieldvalue

import (
	"src/common/ctype"
	"src/module/pm/schema"
	"src/util/dictutil"
	"src/util/errutil"
	"src/util/localeutil"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"gorm.io/gorm"
)

type Schema = schema.TaskFieldValue

var newSchema = schema.NewTaskFieldValue

type Repo struct {
	client *gorm.DB
}

func New(client *gorm.DB) Repo {
	return Repo{client: client}
}

func (r Repo) List(queryOptions ctype.QueryOptions) ([]Schema, error) {
	db := r.client
	if queryOptions.Order == "" {
		db = db.Order("id DESC")
	} else {
		db = db.Order(queryOptions.Order)
	}
	filters := dictutil.DictCamelToSnake(queryOptions.Filters)
	preloads := queryOptions.Preloads
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

func (r Repo) Retrieve(queryOptions ctype.QueryOptions) (*Schema, error) {
	db := r.client
	localizer := localeutil.Get()
	filters := dictutil.DictCamelToSnake(queryOptions.Filters)
	preloads := queryOptions.Preloads
	if len(preloads) > 0 {
		for _, preload := range preloads {
			db = db.Preload(preload)
		}
	}

	joins := queryOptions.Joins
	if len(joins) > 0 {
		for _, join := range joins {
			db = db.Joins(join)
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

func (r Repo) Create(data ctype.Dict) (*Schema, error) {
	item := newSchema(data)
	result := r.client.Create(item)
	err := result.Error
	if err != nil {
		return item, errutil.NewGormError(err)
	}
	return item, err
}

func (r Repo) GetOrCreate(queryOptions ctype.QueryOptions, data ctype.Dict) (*Schema, error) {
	existItem, err := r.Retrieve(queryOptions)
	if err != nil {
		return r.Create(data)
	}
	return existItem, nil
}

func (r Repo) Update(queryOptions ctype.QueryOptions, data ctype.Dict) (*Schema, error) {
	item, err := r.Retrieve(queryOptions)
	if err != nil {
		return nil, err
	}
	result := r.client.Model(&item).Omit("ID").Updates(map[string]interface{}(data))
	err = result.Error
	if err != nil {
		return nil, errutil.NewGormError(err)
	}
	return item, err
}

func (r Repo) UpdateOrCreate(
	queryOptions ctype.QueryOptions,
	data ctype.Dict,
) (*Schema, error) {
	existItem, err := r.Retrieve(queryOptions)
	if err != nil {
		return r.Create(data)
	}
	updateOptions := ctype.QueryOptions{Filters: ctype.Dict{"ID": existItem.ID}}
	return r.Update(updateOptions, data)
}

func (r Repo) Delete(id uint) ([]uint, error) {
	ids := []uint{id}
	_, err := r.Retrieve(ctype.QueryOptions{Filters: ctype.Dict{"ID": id}})
	if err != nil {
		return ids, err
	}
	result := r.client.Where("id = ?", id).Delete(&Schema{})
	err = result.Error
	if err != nil {
		return ids, errutil.NewGormError(err)
	}
	return ids, err
}

func (r Repo) DeleteList(ids []uint) ([]uint, error) {
	result := r.client.Where("id IN (?)", ids).Delete(&Schema{})
	err := result.Error
	if err != nil {
		return ids, errutil.NewGormError(err)
	}
	return ids, err
}
