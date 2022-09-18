package helpers

import "strings"

var ReplacerStrings []string

func RemoveUnusedStringAndToUpper(input string) string {
	input = strings.TrimSpace(input)
	input = strings.ToUpper(input)
	for _, replacer := range ReplacerStrings {
		input = strings.ReplaceAll(input, replacer, "")
	}
	input = strings.TrimSpace(input)
	return input
}
