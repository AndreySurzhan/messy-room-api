package config

type Config struct {
	OpenAIAPIKey string `mapstructure:"openia_api_key"`
	Port         int    `mapstructure:"port"`
}
