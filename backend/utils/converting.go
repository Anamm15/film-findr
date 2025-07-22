package utils

import (
	"time"
)

func FormatDate(t time.Time) string {
	return t.Format("02 Jan 2006")
}
