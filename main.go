package main

import (
	"ASL/utils"
	"ASL/utils/config"
	"os"
)

func main() {
	// 初始化配置
	config.InitializeAslConfig()
	utils.CmdArgs(os.Args)
}
