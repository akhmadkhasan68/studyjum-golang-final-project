package responses

import (
	"final-project/src/commons/constants"
	"final-project/src/database/models"
)

type ProductResponse struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Price       float64      `json:"price"`
	User        UserResponse `json:"user"`
	CreatedAt   string       `json:"created_at"`
	UpdatedAt   string       `json:"updated_at"`
}

func ToProductsResponse(products *[]models.Product) []ProductResponse {
	var result []ProductResponse
	for _, value := range *products {
		result = append(result, ProductResponse{
			ID:          value.ID,
			Name:        value.Name,
			Description: value.Description,
			Price:       value.Price,
			User: UserResponse{
				Username:    value.Outlet.Username,
				Email:       value.Outlet.Email,
				PhoneNumber: value.Outlet.PhoneNumber,
				FirstName:   value.Outlet.FirstName,
				LastName:    value.Outlet.LastName,
				Address:     value.Outlet.Address,
				AreaID:      value.Outlet.AreaID,
				Longitude:   value.Outlet.Longitude,
				Latitude:    value.Outlet.Latitude,
				CreatedAt:   value.Outlet.CreatedAt.Format(constants.DateTimeLayout),
				UpdatedAt:   value.Outlet.UpdatedAt.Format(constants.DateTimeLayout),
			},
			CreatedAt: value.CreatedAt.Format(constants.DateTimeLayout),
			UpdatedAt: value.UpdatedAt.Format(constants.DateTimeLayout),
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
		User: UserResponse{
			Username:    product.Outlet.Username,
			Email:       product.Outlet.Email,
			PhoneNumber: product.Outlet.PhoneNumber,
			FirstName:   product.Outlet.FirstName,
			LastName:    product.Outlet.LastName,
			Address:     product.Outlet.Address,
			AreaID:      product.Outlet.AreaID,
			Longitude:   product.Outlet.Longitude,
			Latitude:    product.Outlet.Latitude,
			CreatedAt:   product.Outlet.CreatedAt.Format(constants.DateTimeLayout),
			UpdatedAt:   product.Outlet.UpdatedAt.Format(constants.DateTimeLayout),
		},
		CreatedAt: product.CreatedAt.Format(constants.DateTimeLayout),
		UpdatedAt: product.UpdatedAt.Format(constants.DateTimeLayout),
	}
}
