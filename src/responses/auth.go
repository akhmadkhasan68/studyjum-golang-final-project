package responses

import (
	"final-project/src/commons/constants"
	"final-project/src/database/models"
)

type UserResponse struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Address     string `json:"address"`
	AreaID      uint64 `json:"area_id"`
	Longitude   string `json:"longitude"`
	Latitude    string `json:"latitude"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type LoginResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}

func ToLoginResponse(user *models.User, token string) LoginResponse {
	return LoginResponse{
		User: UserResponse{
			Username:    user.Username,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Address:     user.Address,
			AreaID:      user.AreaID,
			Longitude:   user.Longitude,
			Latitude:    user.Latitude,
			CreatedAt:   user.CreatedAt.Format(constants.DateTimeLayout),
			UpdatedAt:   user.UpdatedAt.Format(constants.DateTimeLayout),
		},
		Token: token,
	}
}

func ToProfileResponse(user *models.User) UserResponse {
	return UserResponse{
		Username:    user.Username,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Address:     user.Address,
		AreaID:      user.AreaID,
		Longitude:   user.Longitude,
		Latitude:    user.Latitude,
		CreatedAt:   user.CreatedAt.Format(constants.DateTimeLayout),
		UpdatedAt:   user.UpdatedAt.Format(constants.DateTimeLayout),
	}
}
