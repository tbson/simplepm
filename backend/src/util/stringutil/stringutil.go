package stringutil

import (
	"regexp"
	"strings"

	"math/rand/v2"

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

func ToSnakeCaseEnd(s string) string {
	split := strings.Split(s, ".")
	if len(split) == 1 {
		return ToSnakeCase(s)
	}
	return strings.Join([]string{split[0], ToSnakeCase(split[1])}, ".")
}

func ToCamelCase(s string) string {
	result := ""
	// Split the input string into words separated by underscores
	words := strings.Split(s, "_")

	// Initialize a caser for title casing with the default locale
	titleCaser := cases.Title(language.Und)

	// Capitalize each word
	for i, word := range words {
		words[i] = titleCaser.String(word)
	}

	// Join the words without separators to form CamelCase
	result = strings.Join(words, "")

	// ensure all Url convert to URL
	if strings.Contains(result, "Url") {
		return strings.ReplaceAll(result, "Url", "URL")
	}
	return result
}

func GetRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const charsetLength = len(charset)
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.IntN(charsetLength)]
	}
	return string(b)
}
