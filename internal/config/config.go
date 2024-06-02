package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

// Config structure for reading the configuration file.
type Config struct {
	Port        int  `yaml:"port"`
	DisableAuth bool `yaml:"disable_auth"`
}

func Init(configPath string) (*Config, error) {
	var file []byte
	var err error
	if file, err = os.ReadFile(configPath); err != nil {
		return nil, err
	}
	var cfg Config
	if err = yaml.Unmarshal(file, &cfg); err != nil {
		return nil, err
	}
	if cfg.Port == 0 {
		cfg.Port = 8080
	}
	return &cfg, nil
}
