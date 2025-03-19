package config

import (
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"log"
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
func New() (*Config, error) {
	c := &Config{
		viper.New(),
	}

	c.SetConfigFile(".env")

	// Read the .env file
	if err := c.ReadInConfig(); err != nil {
		log.Fatalf("Error reading .env file: %v", err)
	}

	c.AutomaticEnv()

	return c, nil
}
