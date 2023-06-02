package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	*viper.Viper
}

func New() (*Config, error) {
	pwd, _ := os.Getwd()

	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(pwd + "/internal/config")
	v.SetConfigType("yaml")
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
