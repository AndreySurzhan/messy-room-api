package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

// Config keys
const (
	Port         = "PORT"
	OpenAIAPIKey = "OPENAI_API_KEY"
	LoggerLevel  = "LOG_LEVEL"
)

// Config ...
type Config struct {
	*viper.Viper
}

// New creates new config
func New(configPath string) (*Config, error) {
	c := &Config{
		viper.New(),
	}

	_ = c.readFileConfig(configPath)

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
	c.MustBindEnv(Port, "PORT")
	c.MustBindEnv(OpenAIAPIKey, "OPENAI_API_KEY")

	c.AutomaticEnv()
}

func (c *Config) readFileConfig(configPath string) error {
	c.SetConfigFile(configPath)

	if err := c.ReadInConfig(); err != nil {
		return err
	}
	return nil
}
