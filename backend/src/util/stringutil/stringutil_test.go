package stringutil

// write unit tests for all functions in stringutil.go
import (
	"testing"
)

// Test for ToSnakeCase function
func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"HelloWorld", "hello_world"},
		{"helloWorld", "hello_world"},
		{"hello_world", "hello_world"},
		{"HELLO_WORLD", "hello_world"},
		{"hello", "hello"},
		{"", ""},
	}

	for _, test := range tests {
		result := ToSnakeCase(test.input)
		if result != test.expected {
			t.Errorf("ToSnakeCase(%s) = %s; want %s", test.input, result, test.expected)
		}
	}
}

// Test for ToCamelCase function
func TestToCamelCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello_world", "HelloWorld"},
		{"HELLO_WORLD", "HelloWorld"},
		{"hello", "Hello"},
		{"", ""},
	}

	for _, test := range tests {
		result := ToCamelCase(test.input)
		if result != test.expected {
			t.Errorf("ToCamelCase(%s) = %s; want %s", test.input, result, test.expected)
		}
	}
}

func TestRandomString(t *testing.T) {
	tests := []struct {
		length   int
		expected int
	}{
		{10, 10},
		{20, 20},
		{0, 0},
	}

	for _, test := range tests {
		result := GetRandomString(test.length)
		if len(result) != test.expected {
			t.Errorf("GetRandomString(%d) = %d; want %d", test.length, len(result), test.expected)
		}
	}
}
