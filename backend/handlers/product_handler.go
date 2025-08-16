package handlers

import (
	"net/http"

	"github.com/akmalhen/ecommerce-backend/config"
	"github.com/akmalhen/ecommerce-backend/models"
	"github.com/gin-gonic/gin"
)

func GetAvailableProducts(c *gin.Context) {
	var products []models.Product

	if err := config.DB.Where("stock > 0").Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}

	c.JSON(http.StatusOK, products)
}

func GetLatestProducts(c *gin.Context) {
    var products []models.Product

    if err := config.DB.Order("created_at desc").Limit(5).Find(&products).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
        return
    }

    c.JSON(http.StatusOK, products)
}