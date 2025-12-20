package routes

import (
	"github.com/gin-gonic/gin"
	"ecommerce-go/controllers"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/products", controllers.GetProducts)
	r.POST("/products", controllers.CreateProduct)
}
