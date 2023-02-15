package repositories

import (
	"final-project/src/database/models"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db}
}

func (c *OrderRepository) CreateOrder(data models.Order) error {
	return c.db.Create(&data).Error
}

func (c *OrderRepository) UpdateOrder(OrderID string, data models.Order) error {
	return c.db.Where("id = ?", OrderID).Updates(&data).Error
}

func (c *OrderRepository) GetAllOrderByMember(userID string) (*[]models.Order, error) {
	var data = &[]models.Order{}

	if err := c.db.Preload("OrderDetails").Preload("OrderDetails.Product").Preload("Member").Preload("Outlet").Find(data, "member_id = ?", userID).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (c *OrderRepository) GetAllOrderByOutlet(outletID string) (*[]models.Order, error) {
	var data = &[]models.Order{}

	if err := c.db.Preload("OrderDetails").Preload("OrderDetails.Product").Preload("Member").Preload("Outlet").Find(data, "outlet_id = ?", outletID).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (c *OrderRepository) GetOrderMemberByID(UserID string, OrderID string) (*models.Order, error) {
	var data = &models.Order{}

	if err := c.db.Preload("OrderDetails").Preload("OrderDetails.Product").Preload("Member").Preload("Outlet").First(data, "id = ? AND member_id = ?", OrderID, UserID).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (c *OrderRepository) GetOrderOutletByID(OutletID string, OrderID string) (*models.Order, error) {
	var data = &models.Order{}

	if err := c.db.Preload("OrderDetails").Preload("OrderDetails.Product").Preload("Member").Preload("Outlet").First(data, "id = ? AND outlet_id = ?", OrderID, OutletID).Error; err != nil {
		return nil, err
	}

	return data, nil
}
