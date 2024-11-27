package iterutil

// write unit tests for all functions in iterutil.go

import (
	"reflect"
	"testing"
)

// Test for getLabel function
func TestGetLabel(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello_world", "Hello world"},
		{"HELLO_WORLD", "Hello world"},
		{"hello", "Hello"},
		{"", ""},
	}
	for _, test := range tests {
		result := getLabel(test.input)
		if result != test.expected {
			t.Errorf("getLabel(%s) = %s; want %s", test.input, result, test.expected)
		}
	}
}

// Test for GetFieldOptions function
func TestGetFieldOptions(t *testing.T) {
	enum := FieldEnum{"hello_world", "HELLO_WORLD", "hello", ""}
	expected := FieldOptions{
		{"hello_world", "Hello world"},
		{"HELLO_WORLD", "Hello world"},
		{"hello", "Hello"},
		{"", ""},
	}

	result := GetFieldOptions(enum)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetFieldOptions(%v) = %v; want %v", enum, result, expected)
	}
}
