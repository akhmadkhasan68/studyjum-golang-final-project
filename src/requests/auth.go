package requests

import "final-project/src/database/models"

type RegisterRequest struct {
	Username        string `json:"username" binding:"required"`
	Email           string `json:"email" binding:"email,max=100"`
	PhoneNumber     string `json:"phone_number" binding:"max=14"`
	Password        string `json:"password" binding:"required,max=128,min=8"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
	FirstName       string `json:"first_name" binding:"required,max=100"`
	LastName        string `json:"last_name" binding:"required,max=100"`
	Address         string `json:"address" binding:"required,max=200"`
	AreaID          uint64 `json:"area_id" binding:"required"`
	Latitude        string `json:"lat" binding:"required,max=50"`
	Longitude       string `json:"long" binding:"required,max=50"`
	Role            string `json:"role" binding:"required,oneof='MEMBER' 'OUTLET'"`
}

func (request *RegisterRequest) ToModel() models.User {
	return models.User{
		Username:    request.Username,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		Address:     request.Address,
		AreaID:      request.AreaID,
		Latitude:    request.Latitude,
		Longitude:   request.Longitude,
		Role:        models.RolesType(request.Role),
	}
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,max=128,min=8"`
}

type ChangePasswordRequest struct {
	Password        string `json:"password" binding:"required,max=128,min=8"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

type UpdateProfileRequest struct {
	Username    string `json:"username" binding:"required"`
	Email       string `json:"email" binding:"email,max=100"`
	PhoneNumber string `json:"phone_number" binding:"max=14"`
	FirstName   string `json:"first_name" binding:"required,max=100"`
	LastName    string `json:"last_name" binding:"required,max=100"`
	Address     string `json:"address" binding:"required,max=200"`
	AreaID      uint64 `json:"area_id" binding:"required"`
	Latitude    string `json:"lat" binding:"required,max=50"`
	Longitude   string `json:"long" binding:"required,max=50"`
}

func (request *UpdateProfileRequest) ToModel() models.User {
	return models.User{
		Username:    request.Username,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		Address:     request.Address,
		AreaID:      request.AreaID,
		Latitude:    request.Latitude,
		Longitude:   request.Longitude,
	}
}
