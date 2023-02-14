package bussiness

import (
	response "final-project/src/commons/responses"
	"final-project/src/repositories"
	"final-project/src/requests"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepository *repositories.UserRepository
}

func NewAuthService(userRepository *repositories.UserRepository) *AuthService {
	return &AuthService{
		userRepository: userRepository,
	}
}

func (c *AuthService) Register(registerRequest requests.RegisterRequest) error {
	data := registerRequest.ToModel()

	getUserWithEmail, _ := c.userRepository.GetUserWithEmail(data.Email)
	if getUserWithEmail != nil {
		return response.NewErrDuplicateUniqueColumn("Email")
	}

	getUserWithUsername, _ := c.userRepository.GetUserWithUsername(data.Username)
	if getUserWithUsername != nil {
		return response.NewErrDuplicateUniqueColumn("Username")
	}

	getUserWithPhone, _ := c.userRepository.GetUserWithPhone(data.PhoneNumber)
	if getUserWithPhone != nil {
		return response.NewErrDuplicateUniqueColumn("Phone number")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.MinCost)
	if err != nil {
		return fmt.Errorf("generate from password failed: %w", err)
	}
	data.Password = string(hash)

	return c.userRepository.Create(data)
}

func (c *AuthService) Login() {

}

func (c *AuthService) GetByUserID() {

}

func (c *AuthService) generateToken() {

}
