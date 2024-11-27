package iterutil

import (
	"strings"
	"unicode"
)

type FieldEnum []string
type fieldOption struct {
	Value string
	Label string
}
type FieldOptions []fieldOption

func getLabel(s string) string {
	if len(s) == 0 {
		return ""
	}

	// Step 1: Replace underscores with spaces
	s = strings.ReplaceAll(s, "_", " ")

	// Step 2: Convert the string to lowercase
	s = strings.ToLower(s)

	// Step 3: Capitalize the first character
	// Convert the string to a slice of runes to handle Unicode characters properly
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// Label is the Value that replace _ with " " capitalizing the first letter
func GetFieldOptions(enum FieldEnum) FieldOptions {
	var options FieldOptions
	for _, v := range enum {
		options = append(options, fieldOption{v, getLabel(v)})
	}
	return options
}
