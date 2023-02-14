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

func (c *UserRepository) GetUserWithEmail(email string) (*models.User, error) {
	var data = &models.User{}

	if err := c.db.First(data, "email = ?", email).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (c *UserRepository) GetUserWithUsername(username string) (*models.User, error) {
	var data = &models.User{}

	if err := c.db.First(data, "username = ?", username).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (c *UserRepository) GetUserWithPhone(phone string) (*models.User, error) {
	var data = &models.User{}

	if err := c.db.First(data, "phone_number = ?", phone).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (c *UserRepository) GetUserWithID(id string) (*models.User, error) {
	var data = &models.User{}

	if err := c.db.First(data, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return data, nil
}
