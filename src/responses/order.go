package responses

import (
	"final-project/src/commons/constants"
	"final-project/src/database/models"
)

type OrderResponse struct {
	ID            string                `json:"id"`
	Member        string                `json:"member"`
	Outlet        string                `json:"outlet"`
	OrderNumber   string                `json:"order_number"`
	SubTotalPrice float64               `json:"sub_total_price"`
	AdminFee      float64               `json:"admin_fee"`
	TotalPrice    float64               `json:"total_price"`
	Status        string                `json:"status"`
	OrderDetails  []OrderDetailResponse `json:"order_details"`
	CreatedAt     string                `json:"created_at"`
	UpdatedAt     string                `json:"updated_at"`
}

type OrderDetailResponse struct {
	Product    string  `json:"product"`
	Price      float64 `json:"price"`
	Quantity   int64   `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}

func ToOrdersResponse(orders *[]models.Order) []OrderResponse {
	var result []OrderResponse
	for _, value := range *orders {
		var orderDetails []OrderDetailResponse
		for _, detail := range value.OrderDetails {
			orderDetails = append(orderDetails, OrderDetailResponse{
				Product:    detail.Product.Name,
				Price:      detail.Price,
				Quantity:   detail.Quantity,
				TotalPrice: detail.TotalPrice,
			})
		}

		result = append(result, OrderResponse{
			ID:            value.ID,
			Member:        value.Member.FirstName + " " + value.Member.LastName,
			Outlet:        value.Outlet.FirstName + " " + value.Outlet.LastName,
			OrderNumber:   value.OrderNumber,
			SubTotalPrice: value.SubTotalPrice,
			AdminFee:      value.AdminFee,
			TotalPrice:    value.TotalPrice,
			Status:        string(value.Status),
			OrderDetails:  orderDetails,
			CreatedAt:     value.CreatedAt.Format(constants.DateTimeLayout),
			UpdatedAt:     value.UpdatedAt.Format(constants.DateTimeLayout),
		})
	}

	return result
}

func ToOrderResponse(order *models.Order) OrderResponse {
	var orderDetails []OrderDetailResponse
	for _, detail := range order.OrderDetails {
		orderDetails = append(orderDetails, OrderDetailResponse{
			Product:    detail.Product.Name,
			Price:      detail.Price,
			Quantity:   detail.Quantity,
			TotalPrice: detail.TotalPrice,
		})
	}

	return OrderResponse{
		ID:            order.ID,
		Member:        order.Member.FirstName + " " + order.Member.LastName,
		Outlet:        order.Outlet.FirstName + " " + order.Outlet.LastName,
		OrderNumber:   order.OrderNumber,
		SubTotalPrice: order.SubTotalPrice,
		AdminFee:      order.AdminFee,
		TotalPrice:    order.TotalPrice,
		Status:        string(order.Status),
		OrderDetails:  orderDetails,
		CreatedAt:     order.CreatedAt.Format(constants.DateTimeLayout),
		UpdatedAt:     order.UpdatedAt.Format(constants.DateTimeLayout),
	}
}
