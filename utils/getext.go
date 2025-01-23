package utils

import (
	"path/filepath"
	"strings"
)

func GetExt(filePath string) string {
	fileBase := filepath.Base(filePath)
	fileExt := filepath.Ext(filePath)
	fileExtIndex := strings.LastIndexByte(filepath.Base(filePath), '.')
	if fileExtIndex == -1 {
		return fileExt
	} else {
		return filepath.Ext(fileBase[:fileExtIndex]) + fileExt
	}
}
