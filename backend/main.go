package main

import (
	"github.com/akmalhen/ecommerce-backend/config"
	"github.com/akmalhen/ecommerce-backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	config.SeedData() 

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run() 
}
