package rest

import (
	"OnlineBar/internal/models"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

}
