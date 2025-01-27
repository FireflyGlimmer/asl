package extractor

import (
	"ASL/utils/logger"
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
)

type UnTgz struct{}

func (t *UnTgz) Extract(inputFile, outputFolder string) {
	logger := logger.NewLogger("UnTgz")

	tgzFile, err := os.Open(inputFile)
	if err != nil {
		logger.Error("Error opening %s: %v", inputFile, err)
		return
	}
	defer tgzFile.Close()

	err = os.MkdirAll(outputFolder, 0755)
	if err != nil {
		logger.Error("Error creating %s directory: %v", outputFolder, err)
		return
	}

	gzFileContent, err := gzip.NewReader(tgzFile)
	if err != nil {
		logger.Error("Error reading content of %s: %v", inputFile, err)
		return
	}
	defer gzFileContent.Close()

	tarFileContent := tar.NewReader(gzFileContent)
	for {
		fileHeader, err := tarFileContent.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Error("Error reading header of %s: %v", inputFile, err)
			continue
		}
		targetPath := filepath.Join(outputFolder, fileHeader.Name)
		switch fileHeader.Typeflag {
		case tar.TypeDir:
			err := os.MkdirAll(targetPath, fileHeader.FileInfo().Mode())
			if err != nil {
				logger.Error("Error creating %s directory: %v", targetPath, err)
				return
			}
		case tar.TypeReg:
			newFileContent, err := os.OpenFile(targetPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, fileHeader.FileInfo().Mode())
			if err != nil {
				logger.Error("Error creating %s file: %v", targetPath, err)
				return
			}
			defer newFileContent.Close()
			_, err = io.Copy(newFileContent, tarFileContent)
			if err != nil {
				logger.Error("Error copying %s to %s: %v", fileHeader.Name, targetPath, err)
				return
			}
		default:
			logger.Error("Unsupported type: %v", fileHeader.Typeflag)
			return
		}
	}
}
