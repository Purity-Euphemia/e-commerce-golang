package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ecommerce-go/services"
	"ecommerce-go/utils"
)

type OrderStatusInput struct {
	Status string `json:"status" binding:"required"`
}

type CouponInput struct {
	Code          string  `json:"code" binding:"required"`
	DiscountType  string  `json:"discount_type" binding:"required"` // percentage or fixed
	DiscountValue float64 `json:"discount_value" binding:"required,gt=0"`
	MinAmount     float64 `json:"min_amount"`
	MaxUsage      int     `json:"max_usage"`
}

// GetAllOrders - Admin endpoint
func GetAllOrders(c *gin.Context) {
	page, pageSize := utils.GetPaginationFromQuery(c)

	orders, total, err := services.GetAllOrders(page, pageSize)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    orders,
		"pagination": gin.H{
			"page":      page,
			"page_size": pageSize,
			"total":     total,
		},
	})
}

// UpdateOrderStatus - Admin endpoint
func UpdateOrderStatus(c *gin.Context) {
	orderIDStr := c.Param("order_id")
	orderID, err := strconv.ParseUint(orderIDStr, 10, 32)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "invalid order id")
		return
	}

	var input OrderStatusInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	order, err := services.UpdateOrderStatus(uint(orderID), input.Status)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, order)
}

// GetDashboardStats - Admin endpoint
func GetDashboardStats(c *gin.Context) {
	stats, err := services.GetDashboardStats()
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, stats)
}

// CreateCoupon - Admin endpoint
func CreateCoupon(c *gin.Context) {
	var input CouponInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	coupon, err := services.CreateCoupon(input.Code, input.DiscountType, input.DiscountValue, input.MinAmount, input.MaxUsage)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    coupon,
	})
}

// GetAllCoupons - Admin endpoint
func GetAllCoupons(c *gin.Context) {
	coupons, err := services.GetAllCoupons()
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, coupons)
}

// DeleteCoupon - Admin endpoint
func DeleteCoupon(c *gin.Context) {
	code := c.Param("code")
	err := services.DeleteCoupon(code)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, gin.H{"message": "coupon deleted"})
}
