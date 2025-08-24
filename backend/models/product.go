// File: backend/models/product.go
package models

import "gorm.io/gorm"

type Product struct {
    gorm.Model  
    Name              string `json:"name"`
    Price             uint   `json:"price"` 
    Stock             uint   `json:"stock"`
    Description       string `json:"description"`
    ImageURL          string `json:"image_url"`
    ResponsibleUserID uint   `json:"responsible_user_id"`
    ResponsibleUser   User   `json:"responsible_user" gorm:"foreignKey:ResponsibleUserID;references:ID"`
}