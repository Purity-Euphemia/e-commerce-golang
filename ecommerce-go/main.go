package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"ecommerce-go/database"
	"ecommerce-go/models"
	"ecommerce-go/routes"
)

func main() {
	r := gin.Default()

	// Add CORS middleware
	r.Use(cors.Default())

	database.ConnectDatabase()
	database.DB.AutoMigrate(&models.Product{})
	database.DB.AutoMigrate(&models.User{})
	database.DB.AutoMigrate(&models.Cart{})
	database.DB.AutoMigrate(&models.CartItem{})
	database.DB.AutoMigrate(&models.Order{})
	database.DB.AutoMigrate(&models.OrderItem{})

	routes.RegisterRoutes(r)

	r.Run(":8080")
}
