package rest

import (
	"OnlineBar/Backend/internal/database/postgresql"
	"OnlineBar/Backend/internal/models"
	"OnlineBar/Backend/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBalance(c *gin.Context) {
	var claims models.Claims
	var balance float64

	if err := services.CheckJWT(c.GetHeader("Authorization"), &claims); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid token")
		log.Println("Invalid token")
		return
	}

	balance, err := postgresql.GetBalance(claims.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Internal Error")
		log.Println("Error to get User Balance")
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, balance)
	log.Printf("User with id %s get balance", claims.ID)

}
