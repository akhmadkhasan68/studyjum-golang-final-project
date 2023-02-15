package requests

import "final-project/src/database/models"

type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
}

func (request *CreateProductRequest) ToModel() models.Product {
	return models.Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
	}
}
