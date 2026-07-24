package config

import (
    "encoding/json"
    "fmt"
    "os"
)

type Config struct {
    Protocol string `json:"protocol"`
    Port     string `json:"port"`
}

func Load(path string) (*Config, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, fmt.Errorf("cannot read config file: %w", err)
    }

    var cfg Config
    if err := json.Unmarshal(data, &cfg); err != nil {
        return nil, fmt.Errorf("invalid config JSON: %w", err)
    }

    // Minimal validation
    if cfg.Protocol == "" {
        return nil, fmt.Errorf("missing required field: protocol")
    }
    if cfg.Port == "" {
        return nil, fmt.Errorf("missing required field: port")
    }

    return &cfg, nil
}
