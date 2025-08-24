package handlers

import (
	"net/http"
	"strconv"

	"github.com/akmalhen/ecommerce-backend/config"
	"github.com/akmalhen/ecommerce-backend/models"
	"github.com/gin-gonic/gin"
)

func GetAvailableProducts(c *gin.Context) {
	var products []models.Product

	if err := config.DB.Preload("ResponsibleUser").Where("stock > 0").Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}

	c.JSON(http.StatusOK, products)
}

func GetLatestProducts(c *gin.Context) {
    var products []models.Product

    if err := config.DB.Preload("ResponsibleUser").Order("created_at desc").Limit(5).Find(&products).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
        return
    }

    c.JSON(http.StatusOK, products)
}

func GetProductByID(c *gin.Context) {
    productIDStr := c.Param("id")
    productID, err := strconv.Atoi(productIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
        return
    }

    var product models.Product
    if err := config.DB.Preload("ResponsibleUser").First(&product, productID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    c.JSON(http.StatusOK, product)
}