package rest

import (
	"OnlineBar/internal/database/postgresql"
	"OnlineBar/internal/models"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invaild required fields"})
		log.Println("Error reqired fields")
		return
	}

	if err, exist := postgresql.UserExist(user.Name, user.Email); err != nil || exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error wrong password or name"})
		log.Println(err)
		return
	}

	storedPassword, err := postgresql.GetUserPassword(user.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error retrieving user data"})
		log.Println(err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong password"})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"succesful": "User autorizated"})
	log.Printf("User %s autorizated", user.Name)

}
