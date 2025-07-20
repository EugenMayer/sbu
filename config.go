package main

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type UserConfig struct {
	GlobalConfig GlobalConfig `yaml:"global,omitempty"`
}

type GlobalConfig struct {
	DefaultCredentials DefaultCredentials `yaml:"credentials,omitempty"`
}

type DefaultCredentials struct {
	Username string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`
}

func UserConfigPath() (string, error) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/.shelly.yml", userHome), nil
}

func UserConfigExistsInHome() bool {
	path, _ := UserConfigPath()
	return UserConfigExists(path)
}

func UserConfigExists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		// no config present
		return false
	}
	return true
}

func LoadUserConfigFromHome() (*UserConfig, error) {
	path, _ := UserConfigPath()
	return LoadUserConfig(path)
}

func LoadUserConfig(path string) (*UserConfig, error) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		// no config present
		return nil, nil
	}

	data, readErr := os.ReadFile(path)
	if readErr != nil {
		return nil, readErr
	}

	config := UserConfig{}
	unmarshalErr := yaml.Unmarshal(data, &config)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	return &config, nil
}
