package utils

import (
	"fmt"
	"time"
)

func GenerateUniqueImageName(identifier string, originalFilename string) string {
	timestamp := time.Now().Format("20060102_150405.000")
	ext := GetFileExtension(originalFilename)
	return fmt.Sprintf("film_%s_%s%s", identifier, timestamp, ext)
}

func GetFileExtension(filename string) string {
	for i := len(filename) - 1; i >= 0; i-- {
		if filename[i] == '.' {
			return filename[i:]
		}
	}
	return ""
}
