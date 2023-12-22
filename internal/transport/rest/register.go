package rest

import (
	"OnlineBar/internal/database/postgresql"
	"OnlineBar/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if user.Name == "" || user.Password == "" || user.Email == "" || user.OS == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invaild required fields"})
		log.Println("Error reqired fields")
		return
	}

	if err, exist := postgresql.UserExist(user.Name, user.Email); err != nil || exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User exist"})
		log.Println(err)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error hashing password"})
		log.Println(err)
		return
	}

	if err := postgresql.AddUser(user.Name, user.Email, string(hashedPassword), user.OS); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Database error"})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"succesful": "User registred"})
	log.Println("User registred")
}
