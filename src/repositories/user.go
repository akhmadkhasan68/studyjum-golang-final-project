package repositories

import (
	"final-project/src/database/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (c *UserRepository) Create(data models.User) error {
	return c.db.Create(&data).Error
}
