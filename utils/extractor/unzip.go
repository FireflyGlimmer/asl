package extractor

import (
	"ASL/utils/logger"
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

type UnZip struct{}

func (z *UnZip) Extract(inputFilePath, outputDir string) {
	logger := logger.NewLogger("UnZip")

	inputFile, err := zip.OpenReader(inputFilePath)
	if err != nil {
		logger.Error("Error opening %s: %v", inputFilePath, err)
		return
	}
	defer inputFile.Close()

	err = os.MkdirAll(outputDir, 0755)
	if err != nil {
		logger.Error("Error creating %s directory: %v", outputDir, err)
		return
	}

	for _, zipEntry := range inputFile.File {
		logger.Debug("Processing %s", zipEntry.Name)
		zipContent, err := zipEntry.Open()
		if err != nil {
			logger.Error("Error reading %s: %v", zipEntry.Name, err)
			return
		}
		defer zipContent.Close()

		targetPath := filepath.Join(outputDir, zipEntry.Name)
		if zipEntry.FileInfo().IsDir() {
			err := os.MkdirAll(targetPath, zipEntry.Mode()) // 还原文件夹Mode
			if err != nil {
				logger.Error("Error creating %s directory: %v", targetPath, err)
				return
			}
		} else {
			file, err := os.OpenFile(targetPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, zipEntry.Mode()) // 以压缩包的Mode创建文件
			if err != nil {
				logger.Error("Error creating %s file: %v", targetPath, err)
				return
			}
			defer file.Close()
			_, err = io.Copy(file, zipContent) // 将内容复制到创建的文件
			if err != nil {
				logger.Error("Error copying %s to %s: %v", zipEntry.Name, targetPath, err)
				return
			}
		}
	}
}
