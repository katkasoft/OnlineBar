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
	var hasError bool

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

	tx, err := postgresql.StartTransaction()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		log.Println("Error starting the transaction")
		return
	}

	for _, product := range productList.Products {
		if product.Name == "" || product.Cost == 0 || product.Quantity == 0 {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid required fields"})
			log.Println("Error required fields")
			return
		}

		err := postgresql.PostBuyList(tx, claims.ID, product.Name, product.Cost, product.Quantity, time.Now())
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			log.Println("Error inserting product into the database")
			hasError = true
			break
		}
	}

	if hasError {
		return
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error committing the transaction"})
		log.Println("Error committing the transaction:", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"done": "Product list added"})
	log.Printf("User with id %s added list of product", claims.ID)
}
