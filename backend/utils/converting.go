package utils

import (
	"strconv"
	"time"
)

func StringToInt(str string) int {
	intValue, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return intValue
}

func FormatDate(t time.Time) string {
	return t.Format("02 Jan 2006")
}
