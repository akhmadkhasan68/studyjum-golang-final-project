package models

import (
	"database/sql/driver"
	"final-project/src/commons/enums"
)

type RolesType string

const (
	MEMBER RolesType = RolesType(enums.MEMBER)
	OUTLET RolesType = RolesType(enums.OUTLET)
)

func (ct *RolesType) Scan(value interface{}) error {
	*ct = RolesType(value.([]byte))
	return nil
}

func (ct RolesType) Value() (driver.Value, error) {
	return string(ct), nil
}

type User struct {
	Base
	Username    string    `gorm:"column:username;type:varchar(50);not null;"`
	Email       string    `gorm:"column:email;type:varchar(50);not null;"`
	PhoneNumber string    `gorm:"column:phone_number;type:varchar(50);not null;"`
	Password    string    `gorm:"column:password;type:varchar(200);not null;"`
	Role        RolesType `gorm:"type:enum('MEMBER', 'OUTLET');column:role"`
	FirstName   string    `gorm:"column:first_name;type:varchar(50);not null"`
	LastName    string    `gorm:"column:last_name;type:varchar(50);not null"`
	Address     string    `gorm:"column:address;type:text;null"`
	AreaID      uint64    `gorm:"column:area_id;type:int unsigned;not null"`
	Latitude    string    `gorm:"column:latitude;type:varchar(100);not null"`
	Longitude   string    `gorm:"column:longitude;type:varchar(100);not null"`

	Products []Product `gorm:"foreignKey:outlet_id;"`
}
