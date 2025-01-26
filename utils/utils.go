package utils

import (
	"ASL/utils/extractor"
	"ASL/utils/logger"
	"ASL/utils/lxc"
	"fmt"
)

func Extractor(inputFile, outputFolder string) {
	logger := logger.NewLogger("Extractor")

	fileExt := GetExt(inputFile)
	extractor := extractor.NewExtractor(fileExt)
	if extractor != nil {
		extractor.Extract(inputFile, outputFolder)
	} else {
		logger.Error("Error extracting %s: %v", inputFile, extractor)
	}
}

func GetImage(linuxType, linuxVersion string) {
	imageUrl, imageSha256 := lxc.ParserMirrors(linuxType, linuxVersion)
	task := NewTask("root.tar.xz", imageUrl, "config")
	task.Download()
	fmt.Println(imageUrl, imageSha256)
}
