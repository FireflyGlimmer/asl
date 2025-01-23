package config

import (
	inlay "ASL/config"
	"os"
	"path/filepath"
)

// Config
var EXEC_PATH string
var YAML_PATH string

// App
var ASL_VERSION string
var IS_DEBUG bool
var IS_FILELOGGING bool
var DEFAULT_MODE string
var WORK_DIR []string
var MOUNT_PARTITIONS []string

// Linux
var PATH []string
var LD_LIBRARY_PATH []string
var PROOT_TMP_DIR string
var PROOT_LOADER string

// Deploy
var INITIAL_SEQUENCE []string
var LATER_SEQUENCE []string
var CREATE_USER bool
var USER_NAME string
var USER_PASSWORD string
var PRIVILEGED_USERS []string

func IsFileExists(filePath string) bool {
	_, err := os.Lstat(filePath)
	return !os.IsNotExist(err)
}

func InitializeAslConfig() {
	Logger := NewTinyLogger("InitializeAslConfig")
	EXEC_PATH, _ = os.Executable()
	YAML_PATH = filepath.Join(filepath.Dir(EXEC_PATH), "config", "asl.yaml")
	if IsFileExists(YAML_PATH) {
		DEFAULT_CONFIG := GetEmbededAslConfig()
		LOCAL_CONFIG := GetLocalAslConfig()
		// App
		ASL_VERSION = DEFAULT_CONFIG.App.ASL_VERSION
		IS_DEBUG = LOCAL_CONFIG.App.IS_DEBUG
		IS_FILELOGGING = LOCAL_CONFIG.App.IS_FILELOGGING
		DEFAULT_MODE = LOCAL_CONFIG.App.DEFAULT_MODE
		WORK_DIR = DEFAULT_CONFIG.App.WORK_DIR
		MOUNT_PARTITIONS = DEFAULT_CONFIG.App.MOUNT_PARTITIONS

		// Linux
		PATH = DEFAULT_CONFIG.Linux.PATH
		LD_LIBRARY_PATH = DEFAULT_CONFIG.Linux.LD_LIBRARY_PATH
		PROOT_TMP_DIR = DEFAULT_CONFIG.Linux.PROOT_TMP_DIR
		PROOT_LOADER = DEFAULT_CONFIG.Linux.PROOT_LOADER

		// Deploy
		INITIAL_SEQUENCE = DEFAULT_CONFIG.Deploy.INITIAL_SEQUENCE
		LATER_SEQUENCE = DEFAULT_CONFIG.Deploy.LATER_SEQUENCE
		CREATE_USER = DEFAULT_CONFIG.Deploy.CREATE_USER
		USER_NAME = DEFAULT_CONFIG.Deploy.USER_NAME
		USER_PASSWORD = DEFAULT_CONFIG.Deploy.USER_PASSWORD
		PRIVILEGED_USERS = DEFAULT_CONFIG.Deploy.PRIVILEGED_USERS
	} else {
		err := os.Mkdir(filepath.Dir(YAML_PATH), 0755)
		if err != nil {
			Logger.Error("Error creating %s: %v", filepath.Dir(YAML_PATH), err)
		} else {
			os.WriteFile(YAML_PATH, []byte(inlay.AslYAML), 0644)
			DEFAULT_CONFIG := GetEmbededAslConfig()
			LOCAL_CONFIG := GetLocalAslConfig()
			// App
			ASL_VERSION = DEFAULT_CONFIG.App.ASL_VERSION
			IS_DEBUG = LOCAL_CONFIG.App.IS_DEBUG
			IS_FILELOGGING = LOCAL_CONFIG.App.IS_FILELOGGING
			DEFAULT_MODE = LOCAL_CONFIG.App.DEFAULT_MODE
			WORK_DIR = DEFAULT_CONFIG.App.WORK_DIR
			MOUNT_PARTITIONS = DEFAULT_CONFIG.App.MOUNT_PARTITIONS

			// Linux
			PATH = LOCAL_CONFIG.Linux.PATH
			LD_LIBRARY_PATH = LOCAL_CONFIG.Linux.LD_LIBRARY_PATH
			PROOT_TMP_DIR = LOCAL_CONFIG.Linux.PROOT_TMP_DIR
			PROOT_LOADER = LOCAL_CONFIG.Linux.PROOT_LOADER

			// Deploy
			INITIAL_SEQUENCE = LOCAL_CONFIG.Deploy.INITIAL_SEQUENCE
			LATER_SEQUENCE = LOCAL_CONFIG.Deploy.LATER_SEQUENCE
			CREATE_USER = LOCAL_CONFIG.Deploy.CREATE_USER
			USER_NAME = LOCAL_CONFIG.Deploy.USER_NAME
			USER_PASSWORD = LOCAL_CONFIG.Deploy.USER_PASSWORD
			PRIVILEGED_USERS = LOCAL_CONFIG.Deploy.PRIVILEGED_USERS
		}
	}
}
