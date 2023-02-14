package models

type OrderDetail struct {
	Base
	OrderID    string  `gorm:"column:order_id;type:varchar(100);not null;"`
	ProductID  string  `gorm:"column:product_id;type:varchar(100);not null;"`
	Price      float64 `gorm:"column:price;type:decimal(16, 2);not null;"`
	Quantity   int64   `gorm:"column:quantity;type:int(20);not null;"`
	TotalPrice float64 `gorm:"column:total_price;type:decimal(16, 2);not null;"`
}
