package bussiness

import (
	"final-project/src/database/models"
	"final-project/src/repositories"
)

type ProductService struct {
	productRepository *repositories.ProductRepository
}

func NewProductService(productRepository *repositories.ProductRepository) *ProductService {
	return &ProductService{productRepository}
}

func (c *ProductService) GetAll() (*[]models.Product, error) {
	data, err := c.productRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *ProductService) DetailById(id string) (*models.Product, error) {
	data, err := c.productRepository.GetProductById(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *ProductService) Create() {

}

func (c *ProductService) Update() {

}

func (c *ProductService) Delete() {

}
