package handlers

import (
    "net/http"

    "github.com/akmalhen/ecommerce-backend/config" 
    "github.com/akmalhen/ecommerce-backend/models" 
    "github.com/gin-gonic/gin"
)

type DashboardResponse struct {
    TotalUsers        int64           `json:"total_users"`
    ActiveUsers       int64           `json:"active_users"`
    TotalProducts     int64           `json:"total_products"`
    AvailableProducts int64           `json:"available_products"`
    LatestProducts    []models.Product `json:"latest_products"`
}

func GetDashboard(c *gin.Context) {
    var totalUsers, activeUsers, totalProducts, availableProducts int64
    var latestProducts []models.Product

    config.DB.Model(&models.User{}).Count(&totalUsers)
    config.DB.Model(&models.User{}).Where("is_active = ?", true).Count(&activeUsers)

    config.DB.Model(&models.Product{}).Count(&totalProducts)
    config.DB.Model(&models.Product{}).Where("stock > 0").Count(&availableProducts)

    if err := config.DB.Order("created_at desc").Limit(5).Find(&latestProducts).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve latest products"})
        return
    }

    response := DashboardResponse{
        TotalUsers:        totalUsers,
        ActiveUsers:       activeUsers,
        TotalProducts:     totalProducts,
        AvailableProducts: availableProducts,
        LatestProducts:    latestProducts,
    }

    c.JSON(http.StatusOK, response)
}