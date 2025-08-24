package models

import "gorm.io/gorm"

type User struct {
    gorm.Model 
    Name       string `json:"name"`
    Email      string `json:"email" gorm:"unique"`
    Password   string `json:"-"` 
    IsActive   bool   `json:"is_active" gorm:"default:true"`
    Products   []Product `json:"-" gorm:"foreignKey:ResponsibleUserID"`
}