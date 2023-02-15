package bussiness

import "final-project/src/httpclient"

type OrderService struct {
	shipperClient *httpclient.ShipperClient
}

func NewOrderService(shipperClient *httpclient.ShipperClient) *OrderService {
	return &OrderService{shipperClient}
}
