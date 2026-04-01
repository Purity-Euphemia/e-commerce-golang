package repositories

import (
	"ecommerce-go/database"
	"ecommerce-go/models"

	"gorm.io/gorm"
)

func CreateCoupon(coupon *models.Coupon) error {
	return database.DB.Create(coupon).Error
}

func GetCouponByCode(code string) (*models.Coupon, error) {
	var coupon models.Coupon
	err := database.DB.Where("code = ? AND is_active = ?", code, true).First(&coupon).Error
	return &coupon, err
}

func UpdateCouponUsage(code string) error {
	return database.DB.Model(&models.Coupon{}).
		Where("code = ?", code).
		Update("usage_count", gorm.Expr("usage_count + ?", 1)).Error
}

func GetAllCoupons() ([]models.Coupon, error) {
	var coupons []models.Coupon
	err := database.DB.Find(&coupons).Error
	return coupons, err
}

func UpdateCoupon(code string, coupon *models.Coupon) error {
	return database.DB.Model(&models.Coupon{}).Where("code = ?", code).Updates(coupon).Error
}

func DeleteCoupon(code string) error {
	return database.DB.Where("code = ?", code).Delete(&models.Coupon{}).Error
}
