package repositories

import (
	"ecommerce-go/database"
	"ecommerce-go/models"
)

func CreateCategory(category *models.Category) error {
	return database.DB.Create(category).Error
}

func GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	err := database.DB.Find(&categories).Error
	return categories, err
}

func GetCategoryBySlug(slug string) (*models.Category, error) {
	var category models.Category
	err := database.DB.Where("slug = ?", slug).First(&category).Error
	return &category, err
}

func GetCategoryByID(id uint) (*models.Category, error) {
	var category models.Category
	err := database.DB.Where("id = ?", id).First(&category).Error
	return &category, err
}

func UpdateCategory(id uint, category *models.Category) error {
	return database.DB.Model(&models.Category{}).Where("id = ?", id).Updates(category).Error
}

func DeleteCategory(id uint) error {
	return database.DB.Delete(&models.Category{}, id).Error
}
