package utils

import (
	"strconv"
)

func StringToInt(str string) uint64 {
	intValue, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0
	}
	return intValue
}