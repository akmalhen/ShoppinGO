package handlers

import (
	"net/http"
	"fmt"
    "github.com/xuri/excelize/v2"
	"github.com/akmalhen/ecommerce-backend/config"
	"github.com/akmalhen/ecommerce-backend/models"
	"github.com/gin-gonic/gin"
)

type ProductInput struct {
	Name        string `json:"name" binding:"required"`
	Price       uint   `json:"price" binding:"gte=0"`
	Stock       uint   `json:"stock" binding:"gte=0"`
	Description string `json:"description" binding:"required"`
}

func GetProductsAdmin(c *gin.Context) {
	var products []models.Product
	if err := config.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

func CreateProduct(c *gin.Context) {
	var input ProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := models.Product{
		Name:        input.Name,
		Price:       input.Price,
		Stock:       input.Stock,
		Description: input.Description,
	}

	if err := config.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
    productID := c.Param("id")

    var product models.Product
    if err := config.DB.First(&product, productID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    var input ProductInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    product.Name = input.Name
    product.Price = input.Price
    product.Stock = input.Stock
    product.Description = input.Description

    if err := config.DB.Save(&product).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
        return
    }

    c.JSON(http.StatusOK, product)
}

func ExportProducts(c *gin.Context) {
	var products []models.Product
	if err := config.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}

	f := excelize.NewFile()
	sheetName := "Sheet1"
	index, _ := f.NewSheet(sheetName)

	headers := []string{"ID", "Name", "Price", "Stock", "Description", "Created At"}
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
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), product.CreatedAt.Format("2006-01-02 15:04:05"))
	}

	f.SetActiveSheet(index)

	buffer, err := f.WriteToBuffer()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write excel file"})
		return
	}

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename=products.xlsx")
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", buffer.Bytes())
}