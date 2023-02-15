package repositories

import (
	"final-project/src/database/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (c *ProductRepository) GetAllProduct() (*[]models.Product, error) {
	var data = &[]models.Product{}

	if err := c.db.Preload("Outlet").Find(data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (c *ProductRepository) GetAllProductByOutlet(userID string) (*[]models.Product, error) {
	var data = &[]models.Product{}

	if err := c.db.Preload("Outlet").Find(data, "outlet_id = ?", userID).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (c *ProductRepository) GetProductById(id string) (*models.Product, error) {
	var data = &models.Product{}

	if err := c.db.Preload("Outlet").First(data, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (c *ProductRepository) GetProductOutletById(id, outletID string) (*models.Product, error) {
	var data = &models.Product{}

	if err := c.db.Preload("Outlet").First(data, "id = ? ANd outlet_id = ?", id, outletID).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (c *ProductRepository) CreateProduct(data models.Product) error {
	return c.db.Create(&data).Error
}

func (c *ProductRepository) UpdateProduct(id string, data models.Product) error {
	return c.db.Where("id = ?", id).Updates(&data).Error
}

func (c *ProductRepository) DeleteProduct(id string) error {
	return c.db.Where("id = ?", id).Delete(&models.Product{}).Error
}
