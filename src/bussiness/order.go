package bussiness

import (
	"final-project/src/commons/constants"
	"final-project/src/commons/enums"
	"final-project/src/database/models"
	"final-project/src/httpclient"
	"final-project/src/repositories"
	"final-project/src/requests"
	"time"
)

type OrderService struct {
	shipperClient     *httpclient.ShipperClient
	orderRepository   *repositories.OrderRepository
	productRepository *repositories.ProductRepository
}

func NewOrderService(shipperClient *httpclient.ShipperClient, orderRepository *repositories.OrderRepository, productRepository *repositories.ProductRepository) *OrderService {
	return &OrderService{shipperClient, orderRepository, productRepository}
}

func (c *OrderService) GetAllOrdersMember(MemberID string) (*[]models.Order, error) {
	data, err := c.orderRepository.GetAllOrderByMember(MemberID)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *OrderService) GetAllOrdersOutlet(OutletID string) (*[]models.Order, error) {
	data, err := c.orderRepository.GetAllOrderByOutlet(OutletID)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *OrderService) GetDetailOrderMember(MemberID, OrderID string) (*models.Order, error) {
	data, err := c.orderRepository.GetOrderMemberByID(MemberID, OrderID)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *OrderService) GetDetailOrderOutlet(MemberID, OrderID string) (*models.Order, error) {
	data, err := c.orderRepository.GetOrderOutletByID(MemberID, OrderID)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *OrderService) CreateOrder(userID string, request requests.CreateOrderRequest) error {
	data := &models.Order{
		OutletID: request.OutletID,
		MemberID: userID,
	}

	var subtotalPrice float64 = 0
	var adminFee float64 = 0
	var orderDetails []models.OrderDetail
	for _, product := range request.Products {
		productDetail, err := c.productRepository.GetProductById(product.ProductID)
		if err != nil {
			return err
		}
		productPrice := productDetail.Price
		subtotalPrice += productPrice

		orderDetails = append(orderDetails, models.OrderDetail{
			ProductID:  product.ProductID,
			Quantity:   int64(product.Quantity),
			Price:      productPrice,
			TotalPrice: productPrice * float64(product.Quantity),
		})
	}

	data.OrderDetails = orderDetails

	data.OrderNumber = c.generateOrderNumber()
	data.SubTotalPrice = subtotalPrice
	data.AdminFee = adminFee
	data.TotalPrice = adminFee + subtotalPrice
	data.Status = models.OrderStatus(enums.PENDING)

	return c.orderRepository.CreateOrder(*data)
}

func (c *OrderService) CancelOrder(userID string, OrderID string) error {
	data, err := c.orderRepository.GetOrderMemberByID(userID, OrderID)
	if err != nil {
		return err
	}

	data.Status = models.OrderStatus(enums.CANCELED)

	updateOrderErr := c.orderRepository.UpdateOrder(OrderID, *data)
	if updateOrderErr != nil {
		return updateOrderErr
	}

	return nil
}

func (c *OrderService) generateOrderNumber() string {
	return time.Now().Format(constants.TimestampLayout)
}
