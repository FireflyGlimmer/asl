package config

import (
	inlay "ASL/config"
	"os"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	App    AppConfiguration    `yaml:"app"`
	Deploy DeployConfiguration `yaml:"deploy"`
}

// app配置
type AppConfiguration struct {
	ASLVersion      string   `yaml:"asl_version"`
	IsDebug         bool     `yaml:"is_debug"`
	IsFileLogging   bool     `yaml:"is_filelogging"`
	DefaultMode     string   `yaml:"default_mode"`
	WorkDir         []string `yaml:"work_dir"`
	MountPartitions []string `yaml:"mount_partitions"`
}

// deploy配置
type DeployConfiguration struct {
	InitialSequence []string `yaml:"initial_sequence"`
	LaterSequence   []string `yaml:"later_sequence"`
	IsCreateUser    bool     `yaml:"is_create_user"`
	UserName        string   `yaml:"user_name"`
	UserPassword    string   `yaml:"user_password"`
	PrivilegedUsers []string `yaml:"privileged_users"`
}

func GetEmbededAslConfig() *Configuration {
	logger := NewTinyLogger("GetEmbededAslConfig")

	var aslConfig Configuration
	err := yaml.Unmarshal([]byte(inlay.AslYAML), &aslConfig)
	if err != nil {
		logger.Error("Error unmarshalling asl.yaml: %v", err)
		return nil
	}
	return &aslConfig
}

func GetLocalAslConfig() *Configuration {
	logger := NewTinyLogger("GetLocalAslConfig")

	var aslConfig Configuration
	aslYAML, err := os.ReadFile(YAMLPath)
	if err != nil {
		logger.Error("Error opening asl.yaml: %v", err)
		return nil
	}
	err = yaml.Unmarshal([]byte(aslYAML), &aslConfig)
	if err != nil {
		logger.Error("Error unmarshalling asl.yaml: %v", err)
		return nil
	}
	return &aslConfig
}
