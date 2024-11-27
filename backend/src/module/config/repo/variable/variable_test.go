package variable

import (
	"fmt"
	"src/common/ctype"
	"src/util/dbutil"
	"src/util/localeutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

var repo Repo

var idMap = make(map[int]uint)

func TestMain(m *testing.M) {
	localeutil.Init("en")
	dbutil.InitDb()
	repo = New(dbutil.Db())

	seedData()
	m.Run()
}

func getData(index int) ctype.Dict {
	return ctype.Dict{
		"Key":         fmt.Sprintf("key%d", index),
		"Value":       fmt.Sprintf("value%d", index),
		"Description": fmt.Sprintf("description%d", index),
		"DataType":    "STRING",
	}
}

func getID(index int) uint {
	if id, ok := idMap[index]; ok {
		return id
	}
	return uint(index)
}

func setID(index int, id uint) {
	idMap[index] = id
}

func seedData() {
	for i := 0; i < 10; i++ {
		data := getData(i)
		result, _ := repo.Create(data)
		idMap[i] = result.ID
	}
}

func TestList(t *testing.T) {
	queryOptions := ctype.QueryOptions{}
	result, err := repo.List(queryOptions)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	expectedLength := 10
	if len(result) != expectedLength {
		t.Errorf("Expected %d items, got %d", expectedLength, len(result))
	}
}

func TestRetrieve(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		index := 1
		queryOptions := ctype.QueryOptions{
			Filters: ctype.Dict{"ID": getID(index)},
		}
		_, err := repo.Retrieve(queryOptions)
		if err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}
	})
	t.Run("Not found", func(t *testing.T) {
		index := 99
		queryOptions := ctype.QueryOptions{
			Filters: ctype.Dict{"ID": getID(index)},
		}
		_, err := repo.Retrieve(queryOptions)
		assert.EqualError(t, err, "record not found")
	})
}

func TestCreate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		index := 11
		data := getData(index)
		result, err := repo.Create(data)
		setID(index, result.ID)
		if err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}
	})

	t.Run("Duplicate key", func(t *testing.T) {
		index := 11
		data := getData(index)
		_, err := repo.Create(data)
		assert.EqualError(t, err, "value already exists")
	})
}

func TestUpdate(t *testing.T) {
	index := 1
	data := getData(index)
	item, _ := repo.Retrieve(
		ctype.QueryOptions{Filters: ctype.Dict{"ID": getID(index)}},
	)
	result, err := repo.Update(item.ID, data)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if result == nil {
		t.Errorf("Expected non-nil result, got nil")
	}
}

func TestGetOrCreate(t *testing.T) {
	t.Run("Get", func(t *testing.T) {
		searchIndex := 11
		index := 12
		data := getData(index)
		queryOptions := ctype.QueryOptions{
			Filters: ctype.Dict{"ID": getID(searchIndex)},
		}
		result, err := repo.GetOrCreate(queryOptions, data)
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if result == nil {
			t.Errorf("Expected non-nil result, got nil")
		}

		list, _ := repo.List(ctype.QueryOptions{})
		expectedLength := 11
		if len(list) != expectedLength {
			t.Errorf("Expected %d items, got %d", expectedLength, len(list))
		}
	})
	t.Run("Create", func(t *testing.T) {
		index := 14
		data := getData(index)
		queryOptions := ctype.QueryOptions{
			Filters: ctype.Dict{"ID": getID(index)},
		}
		result, err := repo.GetOrCreate(queryOptions, data)
		setID(index, result.ID)
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if result == nil {
			t.Errorf("Expected non-nil result, got nil")
		}

		list, _ := repo.List(ctype.QueryOptions{})
		expectedLength := 12
		if len(list) != expectedLength {
			t.Errorf("Expected %d items, got %d", expectedLength, len(list))
		}
	})
}

func TestUpdateOrCreate(t *testing.T) {
	t.Run("Update", func(t *testing.T) {
		searchIndex := 14
		index := 15
		data := getData(index)
		queryOptions := ctype.QueryOptions{
			Filters: ctype.Dict{"ID": getID(searchIndex)},
		}
		result, err := repo.UpdateOrCreate(queryOptions, data)
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if result == nil {
			t.Errorf("Expected non-nil result, got nil")
		}

		list, _ := repo.List(ctype.QueryOptions{})
		expectedLength := 12
		if len(list) != expectedLength {
			t.Errorf("Expected %d items, got %d", expectedLength, len(list))
		}

	})
	t.Run("Create", func(t *testing.T) {
		index := 16
		data := getData(index)
		queryOptions := ctype.QueryOptions{
			Filters: ctype.Dict{"ID": getID(index)},
		}
		result, err := repo.UpdateOrCreate(queryOptions, data)
		setID(index, result.ID)
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if result == nil {
			t.Errorf("Expected non-nil result, got nil")
		}

		list, _ := repo.List(ctype.QueryOptions{})
		expectedLength := 13
		if len(list) != expectedLength {
			t.Errorf("Expected %d items, got %d", expectedLength, len(list))
		}
	})
}

func TestDelete(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		index := 1
		item, _ := repo.Retrieve(
			ctype.QueryOptions{Filters: ctype.Dict{"ID": getID(index)}},
		)
		_, err := repo.Delete(item.ID)
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		list, _ := repo.List(ctype.QueryOptions{})
		if len(list) != 12 {
			t.Errorf("Expected 11 items, got %d", len(list))
		}
	})
	t.Run("Fail", func(t *testing.T) {
		_, err := repo.Delete(9999)
		assert.EqualError(t, err, "record not found")

		list, _ := repo.List(ctype.QueryOptions{})
		expectedLength := 12
		if len(list) != expectedLength {
			t.Errorf("Expected %d items, got %d", expectedLength, len(list))
		}
	})
}

func TestDeleteList(t *testing.T) {
	list, _ := repo.List(ctype.QueryOptions{})
	ids := make([]uint, 0, len(list))
	for _, item := range list {
		ids = append(ids, item.ID)
	}
	_, err := repo.DeleteList(ids)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	list, _ = repo.List(ctype.QueryOptions{})
	expectedLength := 0
	if len(list) != expectedLength {
		t.Errorf("Expected %d items, got %d", expectedLength, len(list))
	}
}
