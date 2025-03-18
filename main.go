package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"mastodon-ai/config"
	"mastodon-ai/mastodon"
	"mastodon-ai/openai"
)

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func startScheduler() {
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	for range ticker.C {
		mastodon.PublishAIPost()
	}
}

func main() {
	cfg := config.LoadConfig()
	fmt.Println(cfg) // temporary

	go startScheduler()

	r := gin.Default()
	r.GET("/health", healthCheck)
	r.POST("/ai-post", openai.GenerateAIPost)
	r.Run(":8000")
}
