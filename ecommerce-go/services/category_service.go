package services

import (
	"ecommerce-go/models"
	"ecommerce-go/repositories"
	"strings"
)

func CreateCategory(name, description string) (*models.Category, error) {
	slug := strings.ToLower(strings.ReplaceAll(name, " ", "-"))

	category := &models.Category{
		Name:        name,
		Description: description,
		Slug:        slug,
	}

	err := repositories.CreateCategory(category)
	return category, err
}

func GetAllCategories() ([]models.Category, error) {
	return repositories.GetAllCategories()
}

func GetCategoryBySlug(slug string) (*models.Category, error) {
	return repositories.GetCategoryBySlug(slug)
}

func UpdateCategory(id uint, name, description string) (*models.Category, error) {
	slug := strings.ToLower(strings.ReplaceAll(name, " ", "-"))
	category := &models.Category{
		Name:        name,
		Description: description,
		Slug:        slug,
	}

	err := repositories.UpdateCategory(id, category)
	return category, err
}

func DeleteCategory(id uint) error {
	return repositories.DeleteCategory(id)
}
