package routes

import (
	"github.com/akmalhen/ecommerce-backend/handlers"
	"github.com/gin-gonic/gin"
	"github.com/akmalhen/ecommerce-backend/middlewares" 
)



func SetupRoutes(router *gin.Engine) {
	public := router.Group("/products")
	{
		public.GET("/available", handlers.GetAvailableProducts)
		public.GET("/latest", handlers.GetLatestProducts)
		public.GET("/:id", handlers.GetProductByID)
		
	}

	admin := router.Group("/admin")
	{
		admin.POST("/login", handlers.Login)
		admin.POST("/products/:id/upload", handlers.UploadProductImage)
		
		protected := admin.Group("/")
		protected.Use(middlewares.AuthMiddleware())
		{
			protected.GET("/dashboard", handlers.GetDashboard)
			
			userRoutes := protected.Group("/users")
			{
				userRoutes.POST("", handlers.CreateUser)
				userRoutes.GET("", handlers.GetUsers)
				userRoutes.PUT("/:id", handlers.UpdateUser)
			}

			productRoutes := protected.Group("/products")
			{
				productRoutes.GET("", handlers.GetProductsAdmin)
				productRoutes.POST("", handlers.CreateProduct)
				productRoutes.PUT("/:id", handlers.UpdateProduct)
				productRoutes.GET("/export", handlers.ExportProducts)
			}
		}
	}
}
