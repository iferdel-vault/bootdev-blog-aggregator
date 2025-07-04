package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (cfg *Config) SetUser(currentUserName string) error {

	cfg.CurrentUserName = currentUserName
	return write(*cfg)
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error getting user home directory: %w", err)
	}
	configFilePath := homeDir + "/" + configFileName
	return configFilePath, nil
}

func Read() (Config, error) {

	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("error getting config file path: %w", err)
	}

	var config Config

	dat, err := os.ReadFile(configFilePath)
	if err != nil {
		return Config{}, fmt.Errorf("error reading config file: %w", err)
	}

	err = json.Unmarshal(dat, &config)
	if err != nil {
		return Config{}, fmt.Errorf("error parsing config file: %w", err)
	}

	return config, nil
}

func write(cfg Config) error {

	configFilePath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("error getting config file path: %w", err)
	}

	dat, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("error marshalling config file after SetUser: %w", err)
	}

	err = os.WriteFile(configFilePath, dat, 0644)
	if err != nil {
		return fmt.Errorf("error writing config file: %w", err)
	}

	return nil

}
