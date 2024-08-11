package config

import (
	"bytes"
	"fmt"
	"os"
	"reflect"

	"gopkg.in/yaml.v3"
)

// Config structure for reading the configuration file.
type Config struct {
	Port                     int    `yaml:"port" description:"Port backend service" default:"8080"`
	DisableAuth              bool   `yaml:"disable_auth" description:"Disable jwt auth" default:"false"`
	MongoUrl                 string `yaml:"mongo_url" description:"MongoDB url" default:"mongodb://mongo:27017"`
	AccessTokenExpiredHours  int    `yaml:"access_token_expired_hours" description:"Access token expired hours" default:"24"`
	RefreshTokenExpiredHours int    `yaml:"refresh_token_expired_hours" description:"Refresh token expired hours" default:"168"`
	JwtSecretKey             string `yaml:"jwt_secret_key" description:"Jwt secret key" default:""`
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
	buff := bytes.Buffer{}
	buff.WriteString("## Configuration\n")
	buff.WriteString("Configuration file is in yaml format. Default path is `config.yaml`  \n")
	buff.WriteString("Example of configuration file [here](example_config.yaml)  \n")
	buff.WriteString("### Configuration fields:\n")
	for i := 0; i < reflect.TypeOf(cfg).NumField(); i++ {
		name := reflect.TypeOf(cfg).Field(i).Tag.Get("yaml")
		description := reflect.TypeOf(cfg).Field(i).Tag.Get("description")
		defaultValue := reflect.TypeOf(cfg).Field(i).Tag.Get("default")
		buff.WriteString(fmt.Sprintf("- **%s** - %s. **Defaule value**: %s\n", name, description, defaultValue))
	}
	if err := os.WriteFile("config.md", buff.Bytes(), 0644); err != nil {
		return nil, err
	}

	return &cfg, nil
}
