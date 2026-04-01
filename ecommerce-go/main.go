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

	// Run migrations
	database.DB.AutoMigrate(&models.User{})
	database.DB.AutoMigrate(&models.Category{})
	database.DB.AutoMigrate(&models.Product{})
	database.DB.AutoMigrate(&models.Review{})
	database.DB.AutoMigrate(&models.Cart{})
	database.DB.AutoMigrate(&models.CartItem{})
	database.DB.AutoMigrate(&models.Order{})
	database.DB.AutoMigrate(&models.OrderItem{})
	database.DB.AutoMigrate(&models.Wishlist{})
	database.DB.AutoMigrate(&models.Coupon{})

	routes.RegisterRoutes(r)

	r.Run(":8080")
}
