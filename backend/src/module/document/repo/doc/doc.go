package doc

import (
	"src/common/ctype"
	"src/module/document/schema"
	"src/util/dictutil"
	"src/util/errutil"
	"src/util/localeutil"

	"gorm.io/gorm"
)

type Schema = schema.Doc

var newSchema = schema.NewDoc

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
		return &item, errutil.New(localeutil.NoRecordFound)
	}
	if count > 1 {
		return &item, errutil.New(localeutil.MultipleRecordsFound)
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

func (r *repo) GetOrCreate(opts ctype.QueryOpts, data ctype.Dict) (*Schema, error) {
	existItem, err := r.Retrieve(opts)
	if err != nil {
		return r.Create(data)
	}
	return existItem, nil
}

func (r *repo) Update(opts ctype.QueryOpts, data ctype.Dict) (*Schema, error) {
	item, err := r.Retrieve(opts)
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

func (r *repo) UpdateOrCreate(
	opts ctype.QueryOpts,
	data ctype.Dict,
) (*Schema, error) {
	existItem, err := r.Retrieve(opts)
	if err != nil {
		return r.Create(data)
	}
	updateOpts := ctype.QueryOpts{Filters: ctype.Dict{"ID": existItem.ID}}
	return r.Update(updateOpts, data)
}

func (r *repo) Delete(id uint) ([]uint, error) {
	ids := []uint{id}
	_, err := r.Retrieve(ctype.QueryOpts{Filters: ctype.Dict{"ID": id}})
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

func (r *repo) DeleteList(ids []uint) ([]uint, error) {
	result := r.client.Where("id IN (?)", ids).Delete(&Schema{})
	err := result.Error
	if err != nil {
		return ids, errutil.NewGormError(err)
	}
	return ids, err
}
