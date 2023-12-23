package rest

import (
	"github.com/gin-gonic/gin"
)

func TestFunc(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
