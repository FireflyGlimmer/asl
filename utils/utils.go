package utils

import (
	"ASL/utils/extractor"
	"ASL/utils/logger"
	"ASL/utils/lxc"
	"fmt"
)

func Extractor(inputFile, outputDir string) {
	logger := logger.NewLogger("Extractor")

	fileExt := GetExt(inputFile)
	extractor := extractor.NewExtractor(fileExt)
	if extractor != nil {
		extractor.Extract(inputFile, outputDir)
	} else {
		logger.Error("Error extracting %s: %v", inputFile, extractor)
		return
	}
}

func GetImage(linuxType, linuxVersion string) {
	imageUrl, imageSha256 := lxc.ParserMirrors(linuxType, linuxVersion)
	task := NewTask("root.tar.xz", imageUrl, "out")
	task.Download()
	fmt.Println(imageUrl, imageSha256)
}
