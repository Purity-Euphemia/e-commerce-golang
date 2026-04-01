package repositories

import (
	"ecommerce-go/database"
	"ecommerce-go/models"
)

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	result := database.DB.Preload("Category").Preload("Reviews").Find(&products)
	return products, result.Error
}

func GetProductsPaginated(page, pageSize int) ([]models.Product, int64, error) {
	var products []models.Product
	var total int64

	offset := (page - 1) * pageSize

	err := database.DB.Preload("Category").Preload("Reviews").
		Offset(offset).
		Limit(pageSize).
		Order("created_at DESC").
		Find(&products).Error

	database.DB.Model(&models.Product{}).Count(&total)

	return products, total, err
}

func SearchProducts(query string, page, pageSize int) ([]models.Product, int64, error) {
	var products []models.Product
	var total int64

	offset := (page - 1) * pageSize

	err := database.DB.Preload("Category").Preload("Reviews").
		Where("name ILIKE ? OR description ILIKE ? OR sku ILIKE ?", "%"+query+"%", "%"+query+"%", "%"+query+"%").
		Offset(offset).
		Limit(pageSize).
		Order("created_at DESC").
		Find(&products).Error

	database.DB.Model(&models.Product{}).
		Where("name ILIKE ? OR description ILIKE ? OR sku ILIKE ?", "%"+query+"%", "%"+query+"%", "%"+query+"%").
		Count(&total)

	return products, total, err
}

func GetProductsByCategory(categoryID uint, page, pageSize int) ([]models.Product, int64, error) {
	var products []models.Product
	var total int64

	offset := (page - 1) * pageSize

	err := database.DB.Preload("Category").Preload("Reviews").
		Where("category_id = ?", categoryID).
		Offset(offset).
		Limit(pageSize).
		Order("created_at DESC").
		Find(&products).Error

	database.DB.Model(&models.Product{}).Where("category_id = ?", categoryID).Count(&total)

	return products, total, err
}

func GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	err := database.DB.Preload("Category").Preload("Reviews").First(&product, id).Error
	return &product, err
}

func CreateProduct(product *models.Product) error {
	result := database.DB.Create(product)
	return result.Error
}

func UpdateProduct(id uint, product *models.Product) error {
	return database.DB.Model(&models.Product{}).Where("id = ?", id).Updates(product).Error
}

func DeleteProduct(id uint) error {
	return database.DB.Delete(&models.Product{}, id).Error
}

func UpdateProductStock(productID uint, quantity int) error {
	return database.DB.Model(&models.Product{}).Where("id = ?", productID).
		Update("stock", database.DB.Raw("stock - ?", quantity)).Error
}
