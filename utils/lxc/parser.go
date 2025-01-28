package lxc

import (
	"ASL/utils/config"
	"ASL/utils/logger"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ParserMirrors(linuxType, linuxVersion string) (string, string) {
	logger := logger.NewLogger("ParserMirrors")

	productName := linuxType + ":" + linuxVersion + ":" + config.DeviceArch + ":" + "default"

	// 创建Http Client
	client := &http.Client{}

	// 创建Http Request
	req, err := http.NewRequest("GET", config.MirrorsUrl, nil)
	if err != nil {
		logger.Error("Error creating request: %v", err)
		return "", ""
	}
	req.Header.Set("User-Agent", config.UserAgent)

	// 发送Http Request
	response, err := client.Do(req)
	if err != nil {
		logger.Error("Error senting request: %v", err)
		return "", ""
	}
	defer response.Body.Close()

	// 读取Response Body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		logger.Error("Error reading response body: %v", err)
		return "", ""
	}
	imagesJson := make(map[string]interface{})
	err = json.Unmarshal(body, &imagesJson)
	if err != nil {
		logger.Error("Error unmarshaling json: %v", err)
		return "", ""
	}

	// 读取products的内容
	products, exists := imagesJson["products"].(map[string]interface{})
	if !exists {
		logger.Error("Error getting products from json: %v", exists)
		return "", ""
	}

	// 读取productName的内容
	productInfo, exists := products[productName].(map[string]interface{})
	if !exists {
		logger.Error("Error getting %s from products: %v", productName, exists)
		return "", ""
	}

	// 读取versions的内容
	productVersions, exists := productInfo["versions"].(map[string]interface{})
	if !exists {
		logger.Error("Error getting versions from productInfo: %v", exists)
		return "", ""
	}

	// 比较得到最新版本
	productLatestVersion := ""
	for productVersion := range productVersions {
		if productVersion > productLatestVersion {
			productLatestVersion = productVersion
		}
	}

	// 选定最新版本
	product, exists := productVersions[productLatestVersion].(map[string]interface{})
	if !exists {
		logger.Error("Error getting latest version items: %s", productLatestVersion)
		return "", ""
	}

	// 确保从 items 字段中获取数据
	productItems, exists := product["items"].(map[string]interface{})
	if !exists {
		logger.Error("Error getting items from product")
		return "", ""
	}

	// 选定root.tar.xz
	productFile, exists := productItems["root.tar.xz"].(map[string]interface{})
	if !exists {
		logger.Error("Error getting root.tar.xz from items")
		return "", ""
	}

	// root.tar.xz的path
	productFilePath, exists := productFile["path"].(string)
	if !exists {
		logger.Error("Error getting path from root.tar.xz")
		return "", ""
	}

	// root.tar.xz的sha256
	productFileSha256, exists := productFile["sha256"].(string)
	if !exists {
		logger.Error("Error getting sha256 from root.tar.xz")
		return "", ""
	}

	productFileUrl := fmt.Sprintf("https://images.linuxcontainers.org/%s", productFilePath)
	return productFileUrl, productFileSha256
}

func DownloadLXCImage() {

}
