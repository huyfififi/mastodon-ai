package openaiutil

import (
	"mastodon-ai/config"
)

func GenerateAIPost(cfg *config.Config) string {
	post := "Hello, world!" + cfg.MastodonURL
	return post
}
