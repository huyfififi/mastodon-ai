package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"mastodon-ai/config"
)

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func main() {
	cfg := config.LoadConfig()
	fmt.Println(cfg) // temporary

	r := gin.Default()
	r.GET("/health", healthCheck)
	r.Run(":8000")
}
