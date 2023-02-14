package models

import (
	"database/sql/driver"
	"final-project/src/commons/enums"
)

type OrderStatus string

const (
	PENDING   OrderStatus = OrderStatus(enums.PENDING)
	PAID      OrderStatus = OrderStatus(enums.PAID)
	CANCELED  OrderStatus = OrderStatus(enums.CANCELED)
	SHIPPED   OrderStatus = OrderStatus(enums.SHIPPED)
	COMPLETED OrderStatus = OrderStatus(enums.COMPLETED)
)

func (ct *OrderStatus) Scan(value interface{}) error {
	*ct = OrderStatus(value.([]byte))
	return nil
}

func (ct OrderStatus) Value() (driver.Value, error) {
	return string(ct), nil
}

type Order struct {
	Base
	MemberID      string      `gorm:"column:member_id;type:varchar(100);not null;"`
	OutletID      string      `gorm:"column:outlet_id;type:varchar(100);not null;"`
	OrderNumber   string      `gorm:"column:order_number;type:varchar(100);not null;"`
	SubTotalPrice float64     `gorm:"column:sub_total_price;type:decimal(16, 2);not null;"`
	AdminFee      float64     `gorm:"column:admin_fee;type:decimal(16, 2);not null;"`
	TotalPrice    float64     `gorm:"column:total_price;type:decimal(16, 2);not null;"`
	Status        OrderStatus `gorm:"column:status;type:enum('PENDING', 'PAID', 'CANCELED', 'SHIPPED', 'COMPLETED');not null;default PENDING;"`
}
