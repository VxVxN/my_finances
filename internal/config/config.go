package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Config structure for reading the configuration file.
type Config struct {
	Port                     int    `yaml:"port"`
	DisableAuth              bool   `yaml:"disable_auth"`
	MongoUrl                 string `yaml:"mongo_url"`
	AccessTokenExpiredHours  int    `yaml:"access_token_expired_hours"`
	RefreshTokenExpiredHours int    `yaml:"refresh_token_expired_hours"`
}

func Init(configPath string) (*Config, error) {
	var cfg Config
	if file, err := os.ReadFile(configPath); err == nil {
		if err = yaml.Unmarshal(file, &cfg); err != nil {
			return nil, err
		}
	}
	if cfg.Port == 0 {
		cfg.Port = 8080
	}
	if cfg.MongoUrl == "" {
		cfg.MongoUrl = "mongodb://mongo:27017"
	}
	if cfg.AccessTokenExpiredHours == 0 {
		cfg.AccessTokenExpiredHours = 24
	}
	if cfg.RefreshTokenExpiredHours == 0 {
		cfg.RefreshTokenExpiredHours = 24 * 7
	}
	return &cfg, nil
}
