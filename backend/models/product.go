package models

import "gorm.io/gorm"

type Product struct {
    gorm.Model  
    Name        string `json:"name"`
    Price       uint   `json:"price"` 
    Stock       uint   `json:"stock"`
    Description string `json:"description"`
}