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

func (t *UnTgz) Extract(inputFilePath, outputDir string) {
	logger := logger.NewLogger("UnTgz")

	inputFile, err := os.Open(inputFilePath)
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

	gzReader, err := gzip.NewReader(inputFile)
	if err != nil {
		logger.Error("Error reading content of %s: %v", inputFile, err)
		return
	}
	defer gzReader.Close()

	tarReader := tar.NewReader(gzReader)
	hardLinks := make(map[string]string)
	for {
		fileHeader, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Error("Error reading header of %s: %v", inputFile, err)
			continue
		}
		targetPath := filepath.Join(outputDir, fileHeader.Name)
		switch fileHeader.Typeflag {
		// 文件夹
		case tar.TypeDir:
			err := os.MkdirAll(targetPath, fileHeader.FileInfo().Mode())
			if err != nil {
				logger.Error("Error creating %s directory: %v", targetPath, err)
				return
			}
		// 常规文件
		case tar.TypeReg:
			file, err := os.OpenFile(targetPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, fileHeader.FileInfo().Mode())
			if err != nil {
				logger.Error("Error creating %s file: %v", targetPath, err)
				return
			}
			defer file.Close()
			_, err = io.Copy(file, tarReader)
			if err != nil {
				logger.Error("Error copying %s to %s: %v", fileHeader.Name, targetPath, err)
				return
			}
		// 软链接
		case tar.TypeSymlink:
			err := os.Symlink(fileHeader.Linkname, targetPath)
			if err != nil {
				logger.Error("Error creating symlink %s -> %s: %v", targetPath, fileHeader.Linkname, err)
				return
			}
		// 硬链接
		case tar.TypeLink:
			hardLinks[targetPath] = fileHeader.Linkname
		default:
			logger.Warn("Skipping unsupported type: %v (%s)", fileHeader.Typeflag, fileHeader.Name)
			continue
		}
	}
	// 处理硬链接
	for target, origin := range hardLinks {
		err := os.Link(filepath.Join(outputDir, origin), target)
		if err != nil {
			logger.Error("Error creating hardlink %s -> %s: %v", target, origin, err)
			return
		}
	}
}
