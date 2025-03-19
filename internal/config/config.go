package config

import (
	"fmt"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"strings"
)

// New creates new config
func New(configFile string) (*Config, error) {
	viper.SetConfigFile(configFile)

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("Error reading config file, %s", err)
	}

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		return nil, fmt.Errorf("Unable to decode into struct, %s", err)
	}

	return &c, nil
}
