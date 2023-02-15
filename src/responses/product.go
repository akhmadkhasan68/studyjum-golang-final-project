package responses

import (
	"final-project/src/commons/constants"
	"final-project/src/database/models"
)

type ProductResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Outlet      string  `json:"outlet"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

func ToProductsResponse(products *[]models.Product) []ProductResponse {
	var result []ProductResponse
	for _, value := range *products {
		result = append(result, ProductResponse{
			ID:          value.ID,
			Name:        value.Name,
			Description: value.Description,
			Price:       value.Price,
			Outlet:      value.Outlet.FirstName + " " + value.Outlet.LastName,
			CreatedAt:   value.CreatedAt.Format(constants.DateTimeLayout),
			UpdatedAt:   value.UpdatedAt.Format(constants.DateTimeLayout),
		})
	}

	return result
}

func ToProductResponse(product *models.Product) ProductResponse {
	return ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Outlet:      product.Outlet.FirstName + " " + product.Outlet.LastName,
		CreatedAt:   product.CreatedAt.Format(constants.DateTimeLayout),
		UpdatedAt:   product.UpdatedAt.Format(constants.DateTimeLayout),
	}
}
