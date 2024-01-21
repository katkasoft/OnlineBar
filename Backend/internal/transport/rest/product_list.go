package rest

import (
	"OnlineBar/Backend/internal/database/postgresql"
	"OnlineBar/Backend/internal/models"
	"OnlineBar/Backend/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProductList(c *gin.Context) {
	var claims models.Claims
	var productList models.ProductList

	if err := services.CheckJWT(c.GetHeader("Authorization"), &claims); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		log.Println("Invalid token")
		return
	}

	// Get massive of rpdoduct list
	productList, err := postgresql.GetBuyList(claims.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error get product list"})
		log.Println(err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"succes": productList})

}
