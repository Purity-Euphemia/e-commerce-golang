package routes

import (
	"github.com/gin-gonic/gin"

	"ecommerce-go/controllers"
	"ecommerce-go/middleware"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/products", controllers.GetProducts)
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	admin := r.Group("/")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminOnly())
	{
		admin.POST("/products", controllers.CreateProduct)
	}

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/cart", controllers.AddToCart)
		auth.POST("/checkout", controllers.Checkout)
		auth.GET("/my-orders", controllers.GetMyOrders)
	}
}
