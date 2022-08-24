package config

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func New(basePath string, environment string) (*Config, error) {
	mode := Development
	if environment == "production" {
		mode = Production
	}

	configPath := fmt.Sprintf("%s/%s_config.yaml", basePath, mode.ConfigFileName())

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, errors.New(configPath + " path doesn't exists")
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	config := Config{
		Environment: mode,
		BasePath:    basePath,
	}
	err = yaml.Unmarshal(data, &config)

	return &config, err

}
