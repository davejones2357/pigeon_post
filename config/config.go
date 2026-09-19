package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

const configFile = "config.json"

func Load() (*Config, error) {
	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("cannot read %s: %w", configFile, err)
	}

	var cfg Config

	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("invalid config JSON: %w", err)
	}

	return &cfg, nil
}
