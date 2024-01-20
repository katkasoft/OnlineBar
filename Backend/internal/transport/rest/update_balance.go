package rest

import (
	"OnlineBar/Backend/internal/database/postgresql"
	"OnlineBar/Backend/internal/models"
	"OnlineBar/Backend/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateBalance(c *gin.Context) {
	var claims models.Claims
	var user models.User

	if err := services.CheckJWT(c.GetHeader("Authorization"), &claims); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		log.Println("Invalid token")
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println("Invalid format")
		return
	}

	err := postgresql.UpdateBalance(user.Balance, claims.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Internal Error"})
		log.Println("Error to update User Balance")
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"User balance updated: ": user.Balance})
	log.Printf("User with id %s get balance", claims.ID)

}
