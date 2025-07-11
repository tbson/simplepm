package variable

/*
import (
	"fmt"
	"src/common/ctype"
	"src/util/dbutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var r *repo

var idMap = make(map[int]uint)

func TestMain(m *testing.M) {
	dbutil.InitDb()
	dbClient := dbutil.Db(nil)
	r = New(dbClient)

	seedData()

	m.Run()

	cleanup(dbClient)
}

func cleanup(dbClient *gorm.DB) {
	dbClient.Exec("TRUNCATE TABLE variables")
}

func seedData() {
	for i := 0; i < 10; i++ {
		data := getData(i)
		result, _ := r.Create(data)
		idMap[i] = result.ID
	}
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

func TestList(t *testing.T) {
	opts := ctype.QueryOpts{}
	result, err := r.List(opts)
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
		opts := ctype.QueryOpts{
			Filters: ctype.Dict{"ID": getID(index)},
		}
		_, err := r.Retrieve(opts)
		if err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}
	})
	t.Run("Not found", func(t *testing.T) {
		index := 99
		opts := ctype.QueryOpts{
			Filters: ctype.Dict{"ID": getID(index)},
		}
		_, err := r.Retrieve(opts)
		assert.EqualError(t, err, "no record found")
	})
}

func TestCreate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		index := 11
		data := getData(index)
		result, err := r.Create(data)
		setID(index, result.ID)
		if err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}
	})

	t.Run("Duplicate key", func(t *testing.T) {
		index := 11
		data := getData(index)
		_, err := r.Create(data)
		assert.EqualError(t, err, "value already exists")
	})
}

func TestUpdate(t *testing.T) {
	index := 1
	data := getData(index)
	item, _ := r.Retrieve(
		ctype.QueryOpts{Filters: ctype.Dict{"ID": getID(index)}},
	)
	updateOpts := ctype.QueryOpts{Filters: ctype.Dict{"ID": item.ID}}
	result, err := r.Update(updateOpts, data)
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
		opts := ctype.QueryOpts{
			Filters: ctype.Dict{"ID": getID(searchIndex)},
		}
		result, err := r.GetOrCreate(opts, data)
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if result == nil {
			t.Errorf("Expected non-nil result, got nil")
		}

		list, _ := r.List(ctype.QueryOpts{})
		expectedLength := 11
		if len(list) != expectedLength {
			t.Errorf("Expected %d items, got %d", expectedLength, len(list))
		}
	})
	t.Run("Create", func(t *testing.T) {
		index := 14
		data := getData(index)
		opts := ctype.QueryOpts{
			Filters: ctype.Dict{"ID": getID(index)},
		}
		result, err := r.GetOrCreate(opts, data)
		setID(index, result.ID)
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if result == nil {
			t.Errorf("Expected non-nil result, got nil")
		}

		list, _ := r.List(ctype.QueryOpts{})
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
		opts := ctype.QueryOpts{
			Filters: ctype.Dict{"ID": getID(searchIndex)},
		}
		result, err := r.UpdateOrCreate(opts, data)
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if result == nil {
			t.Errorf("Expected non-nil result, got nil")
		}

		list, _ := r.List(ctype.QueryOpts{})
		expectedLength := 12
		if len(list) != expectedLength {
			t.Errorf("Expected %d items, got %d", expectedLength, len(list))
		}

	})
	t.Run("Create", func(t *testing.T) {
		index := 16
		data := getData(index)
		opts := ctype.QueryOpts{
			Filters: ctype.Dict{"ID": getID(index)},
		}
		result, err := r.UpdateOrCreate(opts, data)
		setID(index, result.ID)
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if result == nil {
			t.Errorf("Expected non-nil result, got nil")
		}

		list, _ := r.List(ctype.QueryOpts{})
		expectedLength := 13
		if len(list) != expectedLength {
			t.Errorf("Expected %d items, got %d", expectedLength, len(list))
		}
	})
}

func TestDelete(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		index := 1
		item, _ := r.Retrieve(
			ctype.QueryOpts{Filters: ctype.Dict{"ID": getID(index)}},
		)
		_, err := r.Delete(item.ID)
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		list, _ := r.List(ctype.QueryOpts{})
		if len(list) != 12 {
			t.Errorf("Expected 11 items, got %d", len(list))
		}
	})
	t.Run("Fail", func(t *testing.T) {
		_, err := r.Delete(9999)
		assert.EqualError(t, err, "no record found")

		list, _ := r.List(ctype.QueryOpts{})
		expectedLength := 12
		if len(list) != expectedLength {
			t.Errorf("Expected %d items, got %d", expectedLength, len(list))
		}
	})
}

func TestDeleteList(t *testing.T) {
	list, _ := r.List(ctype.QueryOpts{})
	ids := make([]uint, 0, len(list))
	for _, item := range list {
		ids = append(ids, item.ID)
	}
	_, err := r.DeleteList(ids)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	list, _ = r.List(ctype.QueryOpts{})
	expectedLength := 0
	if len(list) != expectedLength {
		t.Errorf("Expected %d items, got %d", expectedLength, len(list))
	}
}
*/
