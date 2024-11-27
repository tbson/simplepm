package variable

import (
	"src/common/ctype"
	"src/module/config/schema"
	"src/util/dictutil"
	"src/util/errutil"

	"gorm.io/gorm"
)

type Schema = schema.Variable

var newSchema = schema.NewVariable

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
	filters := dictutil.DictCamelToSnake(queryOptions.Filters)
	preloads := queryOptions.Preloads
	if len(preloads) > 0 {
		for _, preload := range preloads {
			db = db.Preload(preload)
		}
	}

	var item Schema
	result := db.Where(map[string]interface{}(filters)).First(&item)
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

func (r Repo) Update(id uint, data ctype.Dict) (*Schema, error) {
	queryOptions := ctype.QueryOptions{
		Filters: ctype.Dict{"id": id},
	}
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
	return r.Update(existItem.ID, data)
}

func (r Repo) Delete(id uint) ([]uint, error) {
	ids := []uint{id}
	_, err := r.Retrieve(ctype.QueryOptions{Filters: ctype.Dict{"id": id}})
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
