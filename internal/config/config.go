package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
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
	c := &Config{
		viper.New(),
	}

	_ = c.readFileConfig()

	c.setENV()
	c.watchConfig()
	return c, nil
}

func (c *Config) watchConfig() {
	c.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	c.WatchConfig()
}

func (c *Config) setENV() {
	c.SetEnvPrefix("ENV.")
	c.MustBindEnv(ServiceName, "SERVICE_NAME")
	c.MustBindEnv(Environment, "ENVIRONMENT")
	c.MustBindEnv(Port, "PORT")

	c.AutomaticEnv()
}

func (c *Config) readFileConfig() error {
	c.SetConfigType(configFileType)
	c.SetConfigName(configFileName)
	c.AddConfigPath(configPath)

	if err := c.ReadInConfig(); err != nil {
		return err
	}
	return nil
}
