package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

// Config structure for reading the configuration file.
type Config struct {
	Port        int    `yaml:"port"`
	DisableAuth bool   `yaml:"disable_auth"`
	MongoUrl    string `yaml:"mongo_url"`
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
		cfg.MongoUrl = "mongodb://localhost:27017"
	}
	return &cfg, nil
}
