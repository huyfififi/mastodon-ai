package config

import (
	"os"

	"github.com/openai/openai-go"
)

type Config struct {
	OpenAIClient   *openai.Client
	MastodonAPIKey string
	MastodonURL    string
}

func LoadConfig() *Config {
	openaiClient := openai.NewClient() // os.LookupEnv("OPENAI_API_KEY")
	return &Config{
		OpenAIClient:   openaiClient,
		MastodonAPIKey: getEnv("MASTODON_API_KEY"),
		MastodonURL:    getEnv("MASTODON_URL"),
	}
}

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	panic("missing required environment variable: " + key)
}
