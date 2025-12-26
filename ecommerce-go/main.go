package main

import (
	"github.com/gin-gonic/gin"

	"ecommerce-go/config"
	"ecommerce-go/models"
	"ecommerce-go/routes"
)

func main() {
	r := gin.Default()

	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Product{})
	config.DB.AutoMigrate(&models.User{})

	routes.RegisterRoutes(r)

	r.Run(":8080")
}
