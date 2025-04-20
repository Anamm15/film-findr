package utils

import (
	"strconv"
)

func StringToInt(str string) int {
	intValue, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return intValue
}
