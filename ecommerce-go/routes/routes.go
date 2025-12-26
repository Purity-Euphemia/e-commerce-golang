package routes

import (
	"github.com/gin-gonic/gin"
	"ecommerce-go/controllers"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/products", controllers.GetProducts)
	r.POST("/products", controllers.CreateProduct)
	r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)
    r.POST("/cart", controllers.AddToCart)


}
