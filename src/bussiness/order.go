package bussiness

import (
	"final-project/src/commons/constants"
	"final-project/src/commons/enums"
	"final-project/src/database/models"
	"final-project/src/httpclient"
	"final-project/src/repositories"
	"final-project/src/requests"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OrderService struct {
	shipperClient     *httpclient.ShipperClient
	orderRepository   *repositories.OrderRepository
	productRepository *repositories.ProductRepository
	userRepository    *repositories.UserRepository
}

func NewOrderService(
	shipperClient *httpclient.ShipperClient,
	orderRepository *repositories.OrderRepository,
	productRepository *repositories.ProductRepository,
	userRepository *repositories.UserRepository,
) *OrderService {
	return &OrderService{shipperClient, orderRepository, productRepository, userRepository}
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

func (c *OrderService) CreateOrder(ctx *gin.Context, userID string, request requests.CreateOrderRequest) error {
	data := &models.Order{
		OutletID: request.OutletID,
		MemberID: userID,
	}

	outletDetail, errOutletDetail := c.userRepository.GetUserWithID(request.OutletID)
	if errOutletDetail != nil {
		return errOutletDetail
	}

	memberDetail, errMemberDetail := c.userRepository.GetUserWithID(userID)
	if errMemberDetail != nil {
		return errMemberDetail
	}

	var subtotalPrice float64 = 0
	var adminFee float64 = 0
	var orderNumber = c.generateOrderNumber()
	var orderID = uuid.New().String()
	var orderDetails []models.OrderDetail
	var itemPackageRequest []httpclient.ItemPackageRequest

	for _, product := range request.Products {
		productDetail, err := c.productRepository.GetProductById(product.ProductID)
		if err != nil {
			return err
		}
		productName := productDetail.Name
		productPrice := productDetail.Price
		subtotalPrice += productPrice

		orderDetails = append(orderDetails, models.OrderDetail{
			ProductID:  product.ProductID,
			Quantity:   int64(product.Quantity),
			Price:      productPrice,
			TotalPrice: productPrice * float64(product.Quantity),
		})

		itemPackageRequest = append(itemPackageRequest, httpclient.ItemPackageRequest{
			Name:  productName,
			Price: int64(productPrice),
			Qty:   int64(product.Quantity),
		})
	}

	data.OrderDetails = orderDetails
	data.ID = orderID
	data.OrderNumber = orderNumber
	data.SubTotalPrice = subtotalPrice
	data.AdminFee = adminFee
	totalPrice := adminFee + subtotalPrice
	data.TotalPrice = totalPrice
	data.Status = models.OrderStatus(enums.PENDING)

	_, errShipper := c.shipperClient.CreateOrder(ctx, httpclient.CreateOrderRequest{
		Consignee: httpclient.ConsigneRequest{
			Name:        memberDetail.FirstName + " " + memberDetail.LastName,
			PhoneNumber: memberDetail.PhoneNumber,
		},
		Consigner: httpclient.ConsigneRequest{
			Name:        outletDetail.FirstName + " " + outletDetail.LastName,
			PhoneNumber: outletDetail.PhoneNumber,
		},
		Courier: httpclient.CourierRequest{
			Cod:          false,
			RateID:       58,
			UseInsurance: true,
		},
		Coverage: "domestic",
		Destination: httpclient.DestinationRequest{
			Address: memberDetail.Address,
			AreaID:  int64(memberDetail.AreaID),
			Lat:     memberDetail.Latitude,
			Lng:     memberDetail.Longitude,
		},
		ExternalID: orderID,
		Origin: httpclient.DestinationRequest{
			Address: outletDetail.Address,
			AreaID:  int64(outletDetail.AreaID),
			Lat:     outletDetail.Latitude,
			Lng:     outletDetail.Longitude,
		},
		Package: httpclient.PackageRequest{
			Height:      60,
			Items:       itemPackageRequest,
			Length:      30,
			PackageType: 2,
			Price:       int64(totalPrice),
			Weight:      1.0,
			Width:       30,
		},
		PaymentType: "postpay",
	})

	if errShipper != nil {
		return errShipper
	}

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
