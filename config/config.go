package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DataLoadFilePath string `json:"dataLoadFilePath"`
}

func NewConfig() (Config, error) {
	config, err := os.ReadFile(`config\config.json`)
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	if err := json.Unmarshal(config, &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
