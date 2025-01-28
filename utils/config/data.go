package config

import (
	inlay "ASL/config"
	"os"
	"path/filepath"
	"runtime"
)

// Config
var ExecPath string
var YAMLPath string
var DeviceArch string

const MirrorsUrl string = "https://images.linuxcontainers.org/meta/simplestreams/v1/images.json"
const UserAgent string = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0"

// App
var ASLVersion string
var IsDebug bool
var IsFileLogging bool
var DefaultMode string
var WorkDir []string
var MountPartitions []string

// Deploy
var InitialSequence []string
var LaterSequence []string
var IsCreateUser bool
var UserName string
var UserPassword string
var PrivilegedUsers []string

func IsFileExists(filePath string) bool {
	_, err := os.Lstat(filePath)
	return !os.IsNotExist(err)
}

func InitializeAslConfig() {
	logger := NewTinyLogger("InitializeAslConfig")
	// Config
	ExecPath, _ = os.Executable()
	YAMLPath = filepath.Join(filepath.Dir(ExecPath), "config", "asl.yaml")
	DeviceArch = runtime.GOARCH
	if IsFileExists(YAMLPath) {
		DEFAULT_CONFIG := GetEmbededAslConfig()
		LOCAL_CONFIG := GetLocalAslConfig()
		// App
		ASLVersion = DEFAULT_CONFIG.App.ASLVersion
		IsDebug = LOCAL_CONFIG.App.IsDebug
		IsFileLogging = LOCAL_CONFIG.App.IsFileLogging
		DefaultMode = LOCAL_CONFIG.App.DefaultMode
		WorkDir = DEFAULT_CONFIG.App.WorkDir
		MountPartitions = DEFAULT_CONFIG.App.MountPartitions

		// Deploy
		InitialSequence = LOCAL_CONFIG.Deploy.InitialSequence
		LaterSequence = LOCAL_CONFIG.Deploy.LaterSequence
		IsCreateUser = LOCAL_CONFIG.Deploy.IsCreateUser
		UserName = LOCAL_CONFIG.Deploy.UserName
		UserPassword = LOCAL_CONFIG.Deploy.UserPassword
		PrivilegedUsers = LOCAL_CONFIG.Deploy.PrivilegedUsers
	} else {
		err := os.Mkdir(filepath.Dir(YAMLPath), 0755)
		if err != nil {
			logger.Error("Error creating %s: %v", filepath.Dir(YAMLPath), err)
			return
		} else {
			err := os.WriteFile(YAMLPath, []byte(inlay.AslYAML), 0644)
			if err != nil {
				logger.Error("Error writing %s: %v", YAMLPath, err)
				return
			}
			DEFAULT_CONFIG := GetEmbededAslConfig()
			LOCAL_CONFIG := GetLocalAslConfig()
			// App
			ASLVersion = DEFAULT_CONFIG.App.ASLVersion
			IsDebug = LOCAL_CONFIG.App.IsDebug
			IsFileLogging = LOCAL_CONFIG.App.IsFileLogging
			DefaultMode = LOCAL_CONFIG.App.DefaultMode
			WorkDir = DEFAULT_CONFIG.App.WorkDir
			MountPartitions = DEFAULT_CONFIG.App.MountPartitions

			// Deploy
			InitialSequence = LOCAL_CONFIG.Deploy.InitialSequence
			LaterSequence = LOCAL_CONFIG.Deploy.LaterSequence
			IsCreateUser = LOCAL_CONFIG.Deploy.IsCreateUser
			UserName = LOCAL_CONFIG.Deploy.UserName
			UserPassword = LOCAL_CONFIG.Deploy.UserPassword
			PrivilegedUsers = LOCAL_CONFIG.Deploy.PrivilegedUsers
		}
	}
}
