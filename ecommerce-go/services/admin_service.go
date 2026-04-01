package services

import (
	"errors"
	"time"

	"ecommerce-go/database"
	"ecommerce-go/models"
	"ecommerce-go/repositories"
)

func GetAllOrders(page, pageSize int) ([]models.Order, int64, error) {
	var orders []models.Order
	var total int64

	offset := (page - 1) * pageSize

	err := database.DB.Preload("User").Preload("Items").
		Offset(offset).
		Limit(pageSize).
		Order("created_at DESC").
		Find(&orders).Error

	database.DB.Model(&models.Order{}).Count(&total)

	return orders, total, err
}

func UpdateOrderStatus(orderID uint, status string) (*models.Order, error) {
	validStatuses := map[string]bool{
		"pending":   true,
		"confirmed": true,
		"shipped":   true,
		"delivered": true,
		"cancelled": true,
	}

	if !validStatuses[status] {
		return nil, errors.New("invalid order status")
	}

	order := &models.Order{Status: status}
	err := database.DB.Model(&models.Order{}).Where("id = ?", orderID).Updates(order).First(order).Error

	return order, err
}

func GetDashboardStats() (map[string]interface{}, error) {
	var totalUsers, totalOrders, totalProducts int64
	var totalRevenue float64

	database.DB.Model(&models.User{}).Count(&totalUsers)
	database.DB.Model(&models.Order{}).Count(&totalOrders)
	database.DB.Model(&models.Product{}).Count(&totalProducts)
	database.DB.Model(&models.Order{}).
		Where("status = ?", "delivered").
		Select("SUM(total_amount)").
		Row().
		Scan(&totalRevenue)

	// Get today's revenue
	today := time.Now().Format("2006-01-02")
	var todayRevenue float64
	database.DB.Model(&models.Order{}).
		Where("DATE(created_at) = ? AND status = ?", today, "delivered").
		Select("SUM(total_amount)").
		Row().
		Scan(&todayRevenue)

	// Get pending orders
	var pendingOrders int64
	database.DB.Model(&models.Order{}).Where("status = ?", "pending").Count(&pendingOrders)

	return map[string]interface{}{
		"total_users":    totalUsers,
		"total_orders":   totalOrders,
		"total_products": totalProducts,
		"total_revenue":  totalRevenue,
		"today_revenue":  todayRevenue,
		"pending_orders": pendingOrders,
	}, nil
}

func CreateCoupon(code, discountType string, discountValue, minAmount float64, maxUsage int) (*models.Coupon, error) {
	coupon := &models.Coupon{
		Code:          code,
		DiscountType:  discountType,
		DiscountValue: discountValue,
		MinAmount:     minAmount,
		MaxUsage:      maxUsage,
		StartDate:     time.Now(),
		EndDate:       time.Now().AddDate(0, 1, 0),
		IsActive:      true,
	}

	err := repositories.CreateCoupon(coupon)
	return coupon, err
}

func GetAllCoupons() ([]models.Coupon, error) {
	return repositories.GetAllCoupons()
}

func DeleteCoupon(code string) error {
	return repositories.DeleteCoupon(code)
}
