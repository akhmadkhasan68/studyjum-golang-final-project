package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        uuid.UUID    `gorm:"type:varchar(100);primary_key;"`
	CreatedAt time.Time    `gorm:"type:datetime;not null"`
	UpdatedAt time.Time    `gorm:"type:datetime;not null"`
	DeletedAt sql.NullTime `gorm:"type:datetime"`
}

func (base *Base) BeforeCreate(scope *gorm.DB) (err error) {
	uuid := uuid.New()

	base.ID = uuid

	return
}
