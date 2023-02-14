package models

type Product struct {
	Base
	OutletID    string  `gorm:"column:outlet_id;type:varchar(100);not null;"`
	Name        string  `gorm:"column:name;type:varchar(50);not null"`
	Description string  `gorm:"column:description;type:text;null"`
	Price       float64 `gorm:"column:price;type:decimal(16, 2);not null"`

	Outlet User `gorm:"foreignKey:outlet_id;"`
}
