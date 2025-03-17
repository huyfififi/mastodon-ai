package config

import (
	"os"
)

type Config struct {
	OpenAIAPIKey	string
	MastodonAPIKey	string
	MastodonURL	string
}

func LoadConfig() *Config {
	return &Config{
		OpenAIAPIKey: getEnv("OPENAI_API_KEY"),
		MastodonAPIKey: getEnv("MASTODON_API_KEY"),
		MastodonURL: getEnv("MASTODON_URL"),
	}
}

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	panic("missing required environment variable: " + key)
}
