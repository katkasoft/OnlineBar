package rest

import (
	"OnlineBar/Backend/internal/models"
	"OnlineBar/Backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestFunc(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	var claims models.Claims

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		return
	}

	if err := services.CheckJWT(authHeader, &claims); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JWT"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
