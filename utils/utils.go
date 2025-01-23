package utils

import (
	"ASL/utils/extractor"
	"ASL/utils/logger"
)

func Extractor(inputFile, outputFolder string) {
	logger := logger.NewLogger("Extractor")
	fileExt := GetExt(inputFile)
	extractor := extractor.NewExtractor(fileExt)
	if extractor != nil {
		extractor.Extract(inputFile, outputFolder)
	} else {
		logger.Error("Error extracting %s: ", inputFile, extractor)
	}
}
