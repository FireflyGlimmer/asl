package utils

import (
	"ASL/utils/config"
	"fmt"
	"os"
)

func ShowUsage() {
	fmt.Println("Usage:\n\t", os.Args[0], "<Option1> <Option2> <Option3> <Option4>")
	fmt.Println("Options:")
	fmt.Println("\t -i | --install: Install a Linux container from Tsinghua Open Source Mirror")
	fmt.Println("\t\t	Option2: Linux Type		Option3: Linux Codename")
	fmt.Println("\t\tBut when used Option4: --local, asl will run in local-mode. You must redefine the Option2 and Option3.")
	fmt.Println("\t\t	Option2: Rootfs Path		Option3: Container Name")
	fmt.Println("\t -c | --execute: Execute commands in container")
	fmt.Println("\t\t	Option2: Your Command		Option3: Container Name")
	fmt.Println("\t -l | --login: Login container")
	fmt.Println("\t\t	Option2: Container Name")
	fmt.Println("\t -d | --delete: Delete container")
	fmt.Println("\t\t	Option2: Container Name")
	fmt.Println("\t -v | --version: Show ASLGO version")
	fmt.Println("\t -h | --help: Show this usage")
}

func ShowVersion() {
	fmt.Printf("ASL Version: %s\n", config.ASLVersion)
	fmt.Println("Author: FireflyGlimmer")
	fmt.Println("Repo Url: https://github.com/FireflyGlimmer/asl")
}
