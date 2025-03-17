package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func main() {
	r := gin.Default()
	r.GET("/health", healthCheck)
	r.Run(":8000")
}
