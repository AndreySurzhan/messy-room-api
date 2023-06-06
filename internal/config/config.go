package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Config keys
const (
	ServiceName = "ENV.SERVICE_NAME"
	Environment = "ENV.ENVIRONMENT"
	Port        = "ENV.PORT"

	LoggerLevel     = "Runtime.Logger.Level"
	LoggerSentryDNS = "Runtime.Logger.SentryDns"
)

const (
	configPath     = "."
	configFileName = "config"
	configFileType = "yaml"
)

// Config ...
type Config struct {
	*viper.Viper
}

// New creates new config
func New() (*Config, error) {
	v := viper.New()
	v.SetConfigName(configFileName)
	v.AddConfigPath(configPath)
	v.SetConfigType(configFileType)
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	v.WatchConfig()

	return &Config{
		v,
	}, nil
}
