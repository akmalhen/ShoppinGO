package main

import (
	"github.com/akmalhen/ecommerce-backend/config"
	"github.com/akmalhen/ecommerce-backend/routes"
	"github.com/gin-contrib/cors" // Pastikan import ini ada
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	config.SeedData() 

	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	r.Use(cors.New(corsConfig))

	r.Static("/static", "./public")

	routes.SetupRoutes(r)

	r.Run()
}