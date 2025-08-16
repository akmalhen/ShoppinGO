package config

import (
    "log"

    "github.com/akmalhen/ecommerce-backend/models"
    "github.com/akmalhen/ecommerce-backend/utils"
)

func SeedData() {
    var userCount int64
    DB.Model(&models.User{}).Count(&userCount)

    if userCount == 0 {
        log.Println("Seeding data...")

        hashedPassword, err := utils.HashPassword("admin123")
        if err != nil {
            log.Fatalf("Could not hash password: %v", err)
        }

        adminUser := models.User{
            Name:     "Admin BNCC",
            Email:    "admin@bncc.net",
            Password: hashedPassword,
            IsActive: true,
        }
        DB.Create(&adminUser)

        products := []models.Product{
            {Name: "Laptop Gaming Pro", Price: 15000000, Stock: 10, Description: "Laptop spek dewa untuk gaming"},
            {Name: "Mouse Wireless", Price: 250000, Stock: 50, Description: "Mouse wireless nyaman dan presisi"},
            {Name: "Keyboard Mechanical", Price: 750000, Stock: 25, Description: "Keyboard mechanical dengan RGB"},
        }
        DB.Create(&products)

        log.Println("Seeding completed.")
    } else {
        log.Println("Data already seeded.")
    }
}
