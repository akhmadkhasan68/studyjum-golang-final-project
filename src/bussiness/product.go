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

func (c *ProductService) GetAllProductsByOutlet(outletID string) (*[]models.Product, error) {
	data, err := c.productRepository.GetAllProductByOutlet(outletID)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *ProductService) GetAllProducts() (*[]models.Product, error) {
	data, err := c.productRepository.GetAllProduct()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *ProductService) DetailProductOutletById(id, outletID string) (*models.Product, error) {
	data, err := c.productRepository.GetProductOutletById(id, outletID)
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

	return c.productRepository.CreateProduct(data)
}

func (c *ProductService) UpdateProduct(productID string, request requests.CreateProductRequest) error {
	data := request.ToModel()

	return c.productRepository.UpdateProduct(productID, data)
}

func (c *ProductService) DeleteProduct(id string) error {
	return c.productRepository.DeleteProduct(id)
}
