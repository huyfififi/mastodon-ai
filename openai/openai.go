package openai

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GenerateAIPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, world!",
	})
}
