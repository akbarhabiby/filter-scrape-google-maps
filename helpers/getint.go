package helpers

import (
	"strconv"
)

func GetInt(input string) (result int, err error) {
	input = RemoveUnusedStringAndToUpper(input)
	result, err = strconv.Atoi(input)
	return
}
