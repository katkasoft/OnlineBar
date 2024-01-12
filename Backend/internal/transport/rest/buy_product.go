package rest

import (
	"OnlineBar/Backend/internal/database/postgresql"
	"OnlineBar/Backend/internal/models"
	"OnlineBar/Backend/internal/services"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func BuyHandler(c *gin.Context) {
	var productList models.ProductList
	var claims models.Claims

	if err := services.CheckJWT(c.GetHeader("Authorization"), &claims); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		log.Println("Invalid token")
		return
	}

	if err := c.ShouldBindJSON(&productList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println("Invalid format")
		return
	}

	for _, product := range productList.Products {
		if product.Name == "" || product.Cost == 0 || product.Quantity == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid required fields"})
			log.Println("Error required fields")
			return
		}

		postgresql.PostBuyList(claims.ID, product.Name, product.Cost, product.Quantity, time.Now())
	}

}
