package repositories

import (
	"ecommerce-go/database"
	"ecommerce-go/models"
)

func CreateReview(review *models.Review) error {
	return database.DB.Create(review).Error
}

func GetProductReviews(productID uint, page, pageSize int) ([]models.Review, int64, error) {
	var reviews []models.Review
	var total int64

	offset := (page - 1) * pageSize

	err := database.DB.Where("product_id = ?", productID).
		Offset(offset).
		Limit(pageSize).
		Order("created_at DESC").
		Find(&reviews).Error

	database.DB.Model(&models.Review{}).Where("product_id = ?", productID).Count(&total)

	return reviews, total, err
}

func GetUserReviews(userID uint) ([]models.Review, error) {
	var reviews []models.Review
	err := database.DB.Where("user_id = ?", userID).Find(&reviews).Error
	return reviews, err
}

func DeleteReview(id uint, userID uint) error {
	return database.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Review{}).Error
}

func UpdateReview(id uint, review *models.Review) error {
	return database.DB.Model(&models.Review{}).Where("id = ?", id).Updates(review).Error
}
