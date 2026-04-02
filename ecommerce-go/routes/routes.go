package routes

import (
	"github.com/gin-gonic/gin"

	"ecommerce-go/controllers"
	"ecommerce-go/middleware"
)

func RegisterRoutes(r *gin.Engine) {
	// Public routes
	r.GET("/products", controllers.GetProducts)
	r.GET("/products/search", controllers.SearchProducts)
	r.GET("/products/:id", controllers.GetProductByID)
	r.GET("/categories", controllers.GetAllCategories)
	r.GET("/categories/:slug", controllers.GetCategoryBySlug)
	r.GET("/products/category/:category_id", controllers.GetProductsByCategory)
	r.GET("/reviews/product/:product_id", controllers.GetProductReviews)

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Admin routes
	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminOnly())
	{
		// Categories
		admin.POST("/categories", controllers.CreateCategory)

		// Products
		admin.POST("/products", controllers.CreateProduct)
		admin.PUT("/products/:id", controllers.UpdateProduct)
		admin.DELETE("/products/:id", controllers.DeleteProduct)

		// Orders
		admin.GET("/orders", controllers.GetAllOrders)
		admin.PUT("/orders/:order_id/status", controllers.UpdateOrderStatus)

		// Dashboard
		admin.GET("/dashboard/stats", controllers.GetDashboardStats)

		// Coupons
		admin.POST("/coupons", controllers.CreateCoupon)
		admin.GET("/coupons", controllers.GetAllCoupons)
		admin.DELETE("/coupons/:code", controllers.DeleteCoupon)
	}

	// Authenticated routes
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		// Cart
		auth.POST("/cart", controllers.AddToCart)
		auth.POST("/checkout", controllers.Checkout)
		auth.GET("/my-orders", controllers.GetMyOrders)
		auth.PUT("/cart", controllers.UpdateCartItem)
		auth.DELETE("/cart/:product_id", controllers.DeleteCartItem)

		// Wishlist
		auth.POST("/wishlist", controllers.AddToWishlist)
		auth.DELETE("/wishlist/:product_id", controllers.RemoveFromWishlist)
		auth.GET("/my-wishlist", controllers.GetMyWishlist)

		// Reviews
		auth.POST("/reviews/product/:product_id", controllers.AddReview)
		auth.DELETE("/reviews/:review_id", controllers.DeleteReview)

		// User profile
		auth.GET("/profile", controllers.GetProfile)
		auth.PUT("/profile", controllers.UpdateProfile)
	}
}
