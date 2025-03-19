package config

type Config struct {
	OpenAIAPIKey string `mapstructure:"openai_api_key"`
	Port         int    `mapstructure:"port"`
	LogLevel     string `mapstructure:"log_level"`
}
