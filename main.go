package main

import (
	"ASL/utils"
	"ASL/utils/config"
	"os"
)

func main() {
	config.InitializeAslConfig()
	if len(os.Args) > 1 {
		utils.CmdArgs(os.Args)
	}
}
