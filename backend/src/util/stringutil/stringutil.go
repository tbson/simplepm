package stringutil

import (
	"regexp"
	"strings"

	"golang.org/x/exp/rand"
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
	// Define the character set: lowercase letters and digits
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"

	// Create a string builder for efficient string concatenation
	var builder strings.Builder
	builder.Grow(length)

	// Generate the random string
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(charset))
		builder.WriteByte(charset[randomIndex])
	}

	return builder.String()
}
