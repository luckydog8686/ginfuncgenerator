package generator

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestGenerateGinFunc(t *testing.T) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("0.0.0.0:8090")
}