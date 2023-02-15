package bussiness

import (
	"final-project/src/database/models"
	"final-project/src/repositories"
	"final-project/src/requests"
)

type ProductService struct {
	productRepository *repositories.ProductRepository
}

func NewProductService(productRepository *repositories.ProductRepository) *ProductService {
	return &ProductService{productRepository}
}

func (c *ProductService) GetAllProducts() (*[]models.Product, error) {
	data, err := c.productRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *ProductService) DetailProductById(id string) (*models.Product, error) {
	data, err := c.productRepository.GetProductById(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *ProductService) CreateProduct(request requests.CreateProductRequest, userID string) error {
	data := request.ToModel()
	data.OutletID = userID

	return c.productRepository.Create(data)
}

func (c *ProductService) UpdateProduct(productID string, request requests.CreateProductRequest) error {
	data := request.ToModel()

	return c.productRepository.Update(productID, data)
}

func (c *ProductService) DeleteProduct(id string) error {
	return c.productRepository.Delete(id)
}
