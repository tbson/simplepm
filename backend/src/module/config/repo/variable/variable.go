package variable

import (
	"src/common/ctype"
	"src/module/config/schema"
	"src/util/dictutil"
	"src/util/errutil"
	"src/util/i18nmsg"

	"src/module/config/vltd"

	"gorm.io/gorm"
)

type Schema = schema.Variable

var newSchema = schema.NewVariable
var schemaName = "variable"

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
		return &item, errutil.NewWithArgs(
			i18nmsg.NoRecordFoundDetail,
			ctype.Dict{"Value": schemaName},
		)
	}
	if count > 1 {
		return &item, errutil.NewWithArgs(
			i18nmsg.MultipleRecordsFoundDetail,
			ctype.Dict{"Value": schemaName},
		)
	}

	result := query.First(&item)
	err := result.Error
	if err != nil {
		return &item, errutil.NewGormError(err)
	}
	return &item, err
}

func (r *repo) Create(structData vltd.CreateVariableInput) (*Schema, error) {
	db := r.client
	data := dictutil.StructToDict(structData)
	item := newSchema(data)
	result := db.Create(item)
	err := result.Error
	if err != nil {
		return item, errutil.NewGormError(err)
	}
	return item, err
}

func (r *repo) GetOrCreate(
	opts ctype.QueryOpts,
	structData vltd.CreateVariableInput,
) (*Schema, error) {
	existItem, err := r.Retrieve(opts)
	if err != nil {
		return r.Create(structData)
	}
	return existItem, nil
}

func (r *repo) Update(
	opts ctype.QueryOpts,
	structData vltd.UpdateVariableInput,
	fields []string,
) (*Schema, error) {
	db := r.client
	data := dictutil.ParseStructWithFields(structData, fields)
	item, err := r.Retrieve(opts)
	if err != nil {
		return nil, err
	}
	result := db.Model(&item).Omit("ID").Updates(map[string]interface{}(data))
	err = result.Error
	if err != nil {
		return nil, errutil.NewGormError(err)
	}
	return item, err
}

func (r *repo) UpdateOrCreate(
	opts ctype.QueryOpts,
	structData vltd.CreateVariableInput,
	fields []string,
) (*Schema, error) {
	existItem, err := r.Retrieve(opts)
	if err != nil {
		return r.Create(structData)
	}
	updateOpts := ctype.QueryOpts{Filters: ctype.Dict{"ID": existItem.ID}}
	return r.Update(updateOpts, structData.ToUpdate(), fields)
}

func (r *repo) Delete(id string) ([]string, error) {
	db := r.client
	ids := []string{id}
	_, err := r.Retrieve(ctype.QueryOpts{Filters: ctype.Dict{"id": id}})
	if err != nil {
		return ids, err
	}
	result := db.Where("id = ?", id).Delete(&Schema{})
	err = result.Error
	if err != nil {
		return ids, errutil.NewGormError(err)
	}
	return ids, err
}

func (r *repo) DeleteList(ids []string) ([]string, error) {
	db := r.client
	result := db.Where("id IN (?)", ids).Delete(&Schema{})
	err := result.Error
	if err != nil {
		return ids, errutil.NewGormError(err)
	}
	return ids, err
}
