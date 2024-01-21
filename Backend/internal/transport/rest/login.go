package rest

import (
	"OnlineBar/Backend/internal/database/postgresql"
	"OnlineBar/Backend/internal/models"
	"OnlineBar/Backend/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println("Invalid format")
		return
	}

	if user.Name == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid required fields"})
		log.Println("Error required fields")
		return
	}

	// Check user existing
	if err, exist := postgresql.UserExist(user.Name, user.Email); err != nil || exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong password or name"})
		log.Println(err)
		return
	}

	storedPassword, err := postgresql.GetUserPassword(user.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong password or name"})
		log.Println(err)
		return
	}

	// Check password valid
	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong password or name"})
		log.Println(err)
		return
	}

	if user.ID, err = postgresql.GetUserID(user.Name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error to claim ID"})
		log.Println(err)
		return
	}

	// Generate JWT token
	token, err := services.GenerateJWT(user.Name, user.ID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unable to generate JWT"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
	log.Printf("User %s authorized", user.Name)
}
