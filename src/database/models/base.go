package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        string         `gorm:"column:id;type:varchar(100);primary_key;"`
	CreatedAt time.Time      `gorm:"colum:created_at;type:datetime;not null"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:datetime;not null"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime"`
}

func (base *Base) BeforeCreate(scope *gorm.DB) (err error) {
	uuid := uuid.New()

	base.ID = uuid.String()

	return
}
