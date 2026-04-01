package services

import (
	"errors"
	"strings"

	"ecommerce-go/database"
	"ecommerce-go/models"
	"ecommerce-go/repositories"
)

func GetProducts() ([]models.Product, error) {
	return repositories.GetAllProducts()
}

func GetProductsPaginated(page, pageSize int) ([]models.Product, int64, error) {
	return repositories.GetProductsPaginated(page, pageSize)
}

func SearchProducts(query string, page, pageSize int) ([]models.Product, int64, error) {
	return repositories.SearchProducts(query, page, pageSize)
}

func GetProductsByCategory(categoryID uint, page, pageSize int) ([]models.Product, int64, error) {
	return repositories.GetProductsByCategory(categoryID, page, pageSize)
}

func GetProductByID(id uint) (*models.Product, error) {
	return repositories.GetProductByID(id)
}

func CreateProduct(name, description string, price, discountPrice float64, stock int, categoryID uint, slug, image, sku string) (*models.Product, error) {
	if slug == "" {
		slug = strings.ToLower(strings.ReplaceAll(name, " ", "-"))
	}

	product := &models.Product{
		Name:          name,
		Description:   description,
		Price:         price,
		DiscountPrice: discountPrice,
		Stock:         stock,
		CategoryID:    categoryID,
		Slug:          slug,
		Image:         image,
		SKU:           sku,
	}

	err := repositories.CreateProduct(product)
	return product, err
}

func UpdateProduct(id uint, name, description string, price, discountPrice float64, stock int, categoryID uint, slug, image, sku string) (*models.Product, error) {
	if slug == "" {
		slug = strings.ToLower(strings.ReplaceAll(name, " ", "-"))
	}

	product := &models.Product{
		Name:          name,
		Description:   description,
		Price:         price,
		DiscountPrice: discountPrice,
		Stock:         stock,
		CategoryID:    categoryID,
		Slug:          slug,
		Image:         image,
		SKU:           sku,
	}

	err := repositories.UpdateProduct(id, product)
	return product, err
}

func DeleteProduct(id uint) error {
	return repositories.DeleteProduct(id)
}

func ReduceStock(productID uint, qty int) error {
	var product models.Product
	err := database.DB.First(&product, productID).Error
	if err != nil {
		return err
	}

	if product.Stock < qty {
		return errors.New("not enough stock")
	}

	product.Stock -= qty
	return database.DB.Save(&product).Error
}
