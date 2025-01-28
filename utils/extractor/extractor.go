package extractor

import (
	"ASL/utils/logger"
)

type Extractor interface {
	Extract(inputFile, outputDir string)
}

// 通用解包函数
func NewExtractor(fileType string) Extractor {
	logger := logger.NewLogger("NewExtractor")

	logger.Debug("InputFile extention: %s", fileType)
	switch fileType {
	case ".zip":
		return &UnZip{}
	case ".tar":
		return &UnTar{}
	case ".tar.gz", ".tgz":
		return &UnTgz{}
	case ".tar.xz", ".txz":
		return &UnTxz{}
	default:
		logger.Warn("Unsupported type: %s", fileType)
		return nil
	}
}
