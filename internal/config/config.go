package config

import (
	"github.com/spf13/viper"
	"strings"
)

type dbConfig struct {
	User     string
	Password string
	Database string
	Host     string
	Port     uint16
}

type serviceConfig struct {
	Debug bool
	Port  uint16
}

// Config contains configuration data that is injected to application to avoid hardcode values.
type Config struct {
	Database dbConfig
	Service  serviceConfig
}

// NewConfig gets configuration from environment. Variable names are listed in ".env" file.
func NewConfig() *Config {
	confer := viper.New()
	confer.AutomaticEnv()
	confer.SetEnvPrefix("chimpanzee")
	confer.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	return &Config{
		Database: dbConfig{
			User:     confer.GetString("db.user"),
			Password: confer.GetString("db.password"),
			Database: confer.GetString("db.database"),
			Host:     confer.GetString("db.host"),
			Port:     uint16(confer.GetUint("db.port")),
		},
		Service: serviceConfig{
			Debug: confer.GetBool("service.debug"),
			Port:  uint16(confer.GetUint("service.port")),
		},
	}
}
