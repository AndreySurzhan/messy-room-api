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
	watchConfig(v)

	setENV(v)

	return &Config{v}, nil
}

func watchConfig(v *viper.Viper) {
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	v.WatchConfig()
}

func setENV(v *viper.Viper) {
	v.AutomaticEnv()

	v.SetEnvPrefix("ENV.")
	v.MustBindEnv(ServiceName, "SERVICE_NAME")
	v.MustBindEnv(Environment, "ENVIRONMENT")
	v.MustBindEnv(Port, "PORT")
}
