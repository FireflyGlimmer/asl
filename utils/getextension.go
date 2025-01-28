package utils

import (
	"path/filepath"
	"strings"
)

func GetExt(filePath string) string {
	compressExts := map[string]bool{
		".tar.xz": true,
		".tar.gz": true,
	}
	fileName := filepath.Base(filePath)
	for ext := range compressExts {
		if strings.HasSuffix(fileName, ext) {
			return ext
		}
	}
	return filepath.Ext(fileName)
}
