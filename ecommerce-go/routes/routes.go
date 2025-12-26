package routes

import (
	"github.com/gin-gonic/gin"

	"ecommerce-go/controllers"
	"ecommerce-go/middleware"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/products", controllers.GetProducts)
	r.POST("/products", controllers.CreateProduct)
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/cart", controllers.AddToCart)
	}
}
