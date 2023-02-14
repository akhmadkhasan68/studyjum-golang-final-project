package bussiness

import (
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
