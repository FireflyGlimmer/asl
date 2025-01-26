package config

import (
	inlay "ASL/config"
	"os"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	App    AppConfiguration    `yaml:"app"`
	Linux  LinuxConfiguration  `yaml:"linux"`
	Deploy DeployConfiguration `yaml:"deploy"`
}

// app配置
type AppConfiguration struct {
	ASL_VERSION      string   `yaml:"ASL_VERSION"`
	WORK_DIR         []string `yaml:"WORK_DIR"`
	DEFAULT_MODE     string   `yaml:"DEFAULT_MODE"`
	MOUNT_PARTITIONS []string `yaml:"MOUNT_PARTITIONS"`
	IS_DEBUG         bool     `yaml:"IS_DEBUG"`
	IS_FILELOGGING   bool     `yaml:"IS_FILELOGGING"`
}

// linux配置
type LinuxConfiguration struct {
	PATH            []string `yaml:"PATH"`
	LD_LIBRARY_PATH []string `yaml:"LD_LIBRARY_PATH"`
	PROOT_TMP_DIR   string   `yaml:"PROOT_TMP_DIR"`
	PROOT_LOADER    string   `yaml:"PROOT_LOADER"`
}

// deploy配置
type DeployConfiguration struct {
	INITIAL_SEQUENCE []string `yaml:"INITIAL_SEQUENCE"`
	LATER_SEQUENCE   []string `yaml:"LATER_SEQUENCE"`
	CREATE_USER      bool     `yaml:"CREATE_USER"`
	USER_NAME        string   `yaml:"USER_NAME"`
	USER_PASSWORD    string   `yaml:"USER_PASSWORD"`
	PRIVILEGED_USERS []string `yaml:"PRIVILEGED_USERS"`
}

func GetEmbededAslConfig() *Configuration {
	logger := NewTinyLogger("GetEmbededAslConfig")

	var aslConfig Configuration
	err := yaml.Unmarshal([]byte(inlay.AslYAML), &aslConfig)
	if err != nil {
		logger.Error("Error unmarshalling asl.yaml: %v", err)
	}
	return &aslConfig
}

func GetLocalAslConfig() *Configuration {
	logger := NewTinyLogger("GetLocalAslConfig")

	var aslConfig Configuration
	aslYAML, err := os.ReadFile(YAML_PATH)
	if err != nil {
		logger.Error("Error opening asl.yaml: %v", err)
	}
	err = yaml.Unmarshal([]byte(aslYAML), &aslConfig)
	if err != nil {
		logger.Error("Error unmarshalling asl.yaml: %v", err)
	}
	return &aslConfig
}
