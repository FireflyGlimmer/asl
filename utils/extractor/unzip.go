package extractor

import (
	"ASL/utils/logger"
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

type UnZip struct{}

func (z *UnZip) Extract(inputFile, outputFolder string) {
	logger := logger.NewLogger("UnZip")

	// 全部的内存地址
	zipFile, err := zip.OpenReader(inputFile)
	if err != nil {
		logger.Error("Error opening %s: %v", inputFile, err)
		return
	}
	defer zipFile.Close()

	// 创建目标路径
	err = os.MkdirAll(outputFolder, 0755)
	if err != nil {
		logger.Error("Error creating %s directory: %v", outputFolder, err)
		return
	}

	/*
		遍历文件条目
		file是文件内存地址
	*/
	for _, file := range zipFile.File { // 转换成文件条目
		logger.Debug("Processing %s", file.Name)
		zipContent, err := file.Open()
		if err != nil {
			logger.Error("Error reading %s: %v", file.Name, err)
			return
		}
		defer zipContent.Close()

		targetPath := filepath.Join(outputFolder, file.Name)
		if file.FileInfo().IsDir() {
			err := os.MkdirAll(targetPath, file.Mode()) // 还原文件夹Mode
			if err != nil {
				logger.Error("Error creating %s directory: %v", targetPath, err)
				return
			}
		} else {
			newFileContent, err := os.OpenFile(targetPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode()) // 以压缩包的Mode创建文件
			if err != nil {
				logger.Error("Error creating %s file: %v", targetPath, err)
				return
			}
			defer newFileContent.Close()
			_, err = io.Copy(newFileContent, zipContent) // 将内容复制到创建的文件
			if err != nil {
				logger.Error("Error copying %s to %s: %v", file.Name, targetPath, err)
				return
			}
		}
	}
}
