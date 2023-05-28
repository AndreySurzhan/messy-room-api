package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// Config ...
type Config struct {
	App struct {
		Env         string
		ServiceName string
	}
	Server struct {
		Port string
	}
	Logger struct {
		Level     string
		SentryDsn string
	}
}

// New creates new config
func New() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./internal/config")

	var config Config

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Errorf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, errors.Errorf("unable to decode into struct, %v", err)
	}
	return &config, nil
}

// Watch config file for changes
func (c *Config) Watch() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}
