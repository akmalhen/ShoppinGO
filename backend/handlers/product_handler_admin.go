package handlers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/akmalhen/ecommerce-backend/config"
	"github.com/akmalhen/ecommerce-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

type AdminProductInput struct {
	Name        string `json:"name" binding:"required"`
	Price       uint   `json:"price" binding:"gte=0"`
	Stock       uint   `json:"stock" binding:"gte=0"`
	Description string `json:"description"`
}

func GetProductsAdmin(c *gin.Context) {
    var products []models.Product
    err := config.DB.Preload("ResponsibleUser").Order("id asc").Find(&products).Error
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
        return
    }
    c.JSON(http.StatusOK, products)
}

func CreateProduct(c *gin.Context) {
	var input AdminProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIDClaim, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	userID := uint(userIDClaim.(float64))

	product := models.Product{
		Name:              input.Name,
		Price:             input.Price,
		Stock:             input.Stock,
		Description:       input.Description,
		ResponsibleUserID: userID,
	}

	if err := config.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}
	
	config.DB.Preload("ResponsibleUser").First(&product, product.ID)

	c.JSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
	productID := c.Param("id")

	var input AdminProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	userIDClaim, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	userID := uint(userIDClaim.(float64))

	updateData := map[string]interface{}{
		"name":                input.Name,
		"price":               input.Price,
		"stock":               input.Stock,
		"description":         input.Description,
		"responsible_user_id": userID,
	}

	dbResult := config.DB.Model(&models.Product{}).Where("id = ?", productID).Updates(updateData)

	if dbResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}
	if dbResult.RowsAffected == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

	var updatedProduct models.Product
	config.DB.Preload("ResponsibleUser").First(&updatedProduct, productID)

	c.JSON(http.StatusOK, updatedProduct)
}

func ExportProducts(c *gin.Context) {
	var products []models.Product
	if err := config.DB.Preload("ResponsibleUser").Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}

	f := excelize.NewFile()
	sheetName := "Products"
	index, _ := f.NewSheet(sheetName)
	f.SetActiveSheet(index)

	headers := []string{"ID", "Name", "Price", "Stock", "Description", "Responsible User", "Created At"}
	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheetName, cell, header)
	}

	for i, product := range products {
		row := i + 2
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), product.ID)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), product.Name)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), product.Price)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), product.Stock)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), product.Description)
        f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), product.ResponsibleUser.Name)
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", row), product.CreatedAt.Format("2006-01-02 15:04:05"))
	}

	buffer, err := f.WriteToBuffer()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write excel file"})
		return
	}

	c.Header("Content-Disposition", "attachment; filename=products.xlsx")
	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", buffer.Bytes())
}

func UploadProductImage(c *gin.Context) {
    productID := c.Param("id")
    var product models.Product
    if err := config.DB.First(&product, productID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    file, err := c.FormFile("image")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Image file is required"})
        return
    }

    filename := fmt.Sprintf("product_%s_%d%s", productID, time.Now().Unix(), filepath.Ext(file.Filename))
    savePath := filepath.Join("public", "images", filename)

    if err := c.SaveUploadedFile(file, savePath); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
        return
    }

    imageURL := "http://localhost:8080/static/images/" + filename

    updateData := map[string]interface{}{"image_url": imageURL}
    if err := config.DB.Model(&product).Updates(updateData).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product image URL"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully", "url": imageURL})
}