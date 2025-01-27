package extractor

import (
	"ASL/utils/logger"
	"archive/tar"
	"io"
	"os"
	"path/filepath"
)

type UnTar struct{}

func (t *UnTar) Extract(inputFile, outputFolder string) {
	logger := logger.NewLogger("UnTar")

	tarFile, err := os.Open(inputFile)
	if err != nil {
		logger.Error("Error opening %s: %v", inputFile, err)
		return
	}
	defer tarFile.Close()

	err = os.MkdirAll(outputFolder, 0755)
	if err != nil {
		logger.Error("Error creating %s directory: %v", outputFolder, err)
		return
	}

	tarFileContent := tar.NewReader(tarFile)
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
