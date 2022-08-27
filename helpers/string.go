package helpers

import "strings"

func RemoveUnusedStringAndToUpper(input string) string {
	input = strings.TrimSpace(input)
	input = strings.ToUpper(input)
	input = strings.ReplaceAll(input, ".", "")
	input = strings.ReplaceAll(input, ",", "")
	input = strings.ReplaceAll(input, "RT", "")
	input = strings.ReplaceAll(input, "RW", "")
	input = strings.ReplaceAll(input, "NOMOR", "")
	input = strings.ReplaceAll(input, "NOMER", "")
	input = strings.ReplaceAll(input, "NO", "")
	input = strings.ReplaceAll(input, "N0 ", "")
	input = strings.ReplaceAll(input, "KELURAHAN", "")
	input = strings.ReplaceAll(input, "KECAMATAN", "")
	input = strings.ReplaceAll(input, "KEL", "")
	input = strings.ReplaceAll(input, "BLOC", "")
	input = strings.ReplaceAll(input, "BLOK", "")
	input = strings.ReplaceAll(input, "RUKO", "")
	input = strings.TrimSpace(input)
	return input
}
