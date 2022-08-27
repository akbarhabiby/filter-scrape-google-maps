package helpers

import (
	"strconv"
	"strings"
)

func GetInt(input string) (result int, err error) {
	input = strings.ToLower(strings.TrimSpace(input))
	result, err = strconv.Atoi(input)
	if err != nil {
		input = strings.ReplaceAll(input, ".", "")
		input = strings.ReplaceAll(input, "no", "")
		input = strings.TrimSpace(input)

		result, err = strconv.Atoi(input)
	}

	return
}
