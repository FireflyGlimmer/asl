package utils

import (
	"ASL/utils/config"
	"ASL/utils/logger"
	"bufio"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type Task struct {
	name string
	url  string
	out  string
}

func (t *Task) Download() {
	logger := logger.NewLogger("Download")

	client := &http.Client{}

	req, err := http.NewRequest("GET", t.url, nil)
	if err != nil {
		logger.Error("Error creating request: %v", err)
	}
	req.Header.Set("User-Agent", config.USER_AGENT)

	response, err := client.Do(req)
	if err != nil {
		logger.Error("Error senting request: %v", err)
		return
	}
	defer response.Body.Close()

	// 创建文件
	filePath := filepath.Join(t.out, t.name)
	fileContent, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		logger.Error("Error creating file: %v", err)
		return
	}
	defer fileContent.Close()

	// 写入文件
	buffer := make([]byte, 32*1024) // 32KB
	progressBar := NewProgressBar(0, int(response.ContentLength))
	reader := bufio.NewReader(response.Body)
	writer := bufio.NewWriter(fileContent)
	logger.Info("Downloading %s", t.name)
	for {
		n, err := reader.Read(buffer)
		if err != nil && err != io.EOF {
			logger.Error("Error reading response: %v", err)
			return
		}
		if n == 0 {
			break
		}
		if _, err := writer.Write(buffer[:n]); err != nil {
			logger.Error("Error writing to file: %v", err)
			return
		}
		progressBar.SetValue(progressBar.value + n)
		progressBar.Print()
	}
	writer.Flush()
	logger.Info("Download complete: %s", t.name)
}

func NewTask(inputName, inputUrl, outputFolder string) *Task {
	return &Task{name: inputName, url: inputUrl, out: outputFolder}
}
