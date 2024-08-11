package config

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"strconv"

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
	buff := bytes.Buffer{}
	buff.WriteString("## Configuration\n")
	buff.WriteString("Configuration file is in yaml format. Default path is `config.yaml`  \n")
	buff.WriteString("Example of configuration file [here](example_config.yaml)  \n")
	buff.WriteString("### Configuration fields:\n")
	for i := 0; i < reflect.TypeOf(cfg).NumField(); i++ {
		tag := reflect.TypeOf(cfg).Field(i).Tag
		name := tag.Get("yaml")
		description := tag.Get("description")
		defaultValue := tag.Get("default")
		buff.WriteString(fmt.Sprintf("- **%s** - %s. **Defaule value**: %s\n", name, description, defaultValue))
		cfgValue := reflect.ValueOf(&cfg).Elem()
		field := cfgValue.Field(i)
		switch field.Type().Kind() {
		case reflect.Int:
			if field.Int() == 0 {
				value, ok := strconv.Atoi(defaultValue)
				if ok != nil {
					return nil, fmt.Errorf("can't convert %s to int", defaultValue)
				}
				field.SetInt(int64(value))
			}
		case reflect.Bool:
			if !field.Bool() {
				value, ok := strconv.ParseBool(defaultValue)
				if ok != nil {
					return nil, fmt.Errorf("can't convert %s to bool", defaultValue)
				}
				field.SetBool(value)
			}
		case reflect.String:
			if field.String() == "" {
				field.SetString(defaultValue)
			}
		default:
			return nil, fmt.Errorf("unknown type %s", field.Type().Kind())
		}
	}
	if err := os.WriteFile("config.md", buff.Bytes(), 0644); err != nil {
		return nil, err
	}

	return &cfg, nil
}
