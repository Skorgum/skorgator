package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	workingDir, err := os.Getwd()
	if err != nil {
		return Config{}, err
	}

	configPath := filepath.Join(workingDir, ".gatorconfig.json")

	data, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username

	workingDir, err := os.Getwd()
	if err != nil {
		return err
	}

	configPath := filepath.Join(workingDir, ".gatorconfig.json")

	data, err := json.Marshal(c)
	if err != nil {
		return err
	}

	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
