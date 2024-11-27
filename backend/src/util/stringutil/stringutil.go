package stringutil

import (
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(s string) string {
	snake := matchFirstCap.ReplaceAllString(s, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func ToCamelCase(s string) string {
	// Split the input string into words separated by underscores
	words := strings.Split(s, "_")

	// Initialize a caser for title casing with the default locale
	titleCaser := cases.Title(language.Und)

	// Capitalize each word
	for i, word := range words {
		words[i] = titleCaser.String(word)
	}

	// Join the words without separators to form CamelCase
	return strings.Join(words, "")
}
