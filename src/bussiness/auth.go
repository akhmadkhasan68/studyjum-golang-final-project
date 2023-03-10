package bussiness

import (
	"errors"
	response "final-project/src/commons/responses"
	"final-project/src/config"
	"final-project/src/database/models"
	"final-project/src/repositories"
	"final-project/src/requests"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepository *repositories.UserRepository
	jwtSecretKey   string
	jwtExpired     time.Duration
}

func NewAuthService(userRepository *repositories.UserRepository) *AuthService {
	JWTExpiredTime, _ := strconv.Atoi(config.GetEnvVariable("JWT_EXPIRED_TIME"))

	return &AuthService{
		userRepository: userRepository,
		jwtSecretKey:   config.GetEnvVariable("JWT_SECRET_KEY"),
		jwtExpired:     time.Duration(JWTExpiredTime) * time.Minute,
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

func (c *AuthService) Login(loginRequest requests.LoginRequest) (string, *models.User, error) {
	user, err := c.userRepository.GetUserWithUsername(loginRequest.Username)
	if err != nil {
		if errors.Is(err, &response.ErrNotFound{}) {
			return "", nil, response.NewErrUnauthorized("Incorrect username entered")
		}
		return "", nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		return "", nil, response.NewErrUnauthorized("Incorrect password entered")
	}

	token, err := c.generateToken(user)
	if err != nil {
		return "", nil, fmt.Errorf("generate token failed: %w", err)
	}

	return token, user, nil
}

func (c *AuthService) GetByUserID(id string) (*models.User, error) {
	user, err := c.userRepository.GetUserWithID(id)
	if err != nil {
		if errors.Is(err, &response.ErrNotFound{}) {
			return nil, response.NewErrUnauthorized("Incorrect username entered")
		}
		return nil, err
	}

	return user, nil
}

func (c *AuthService) ChangePassword(userID string, changePasswordRequest requests.ChangePasswordRequest) error {
	user, err := c.userRepository.GetUserWithID(userID)
	if err != nil {
		if errors.Is(err, &response.ErrNotFound{}) {
			return response.NewErrUnauthorized("Incorrect username entered")
		}
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(changePasswordRequest.Password), bcrypt.MinCost)
	if err != nil {
		return fmt.Errorf("generate from password failed: %w", err)
	}
	user.Password = string(hash)

	return c.userRepository.Update(userID, *user)
}

func (c *AuthService) UpdateProfile(userID string, updateProfileRequest requests.UpdateProfileRequest) error {
	requestData := updateProfileRequest.ToModel()
	user, err := c.userRepository.GetUserWithID(userID)
	if err != nil {
		if errors.Is(err, &response.ErrNotFound{}) {
			return response.NewErrUnauthorized("Incorrect username entered")
		}
		return err
	}

	if user.Email != requestData.Email {
		getUserWithEmail, _ := c.userRepository.GetUserWithEmail(requestData.Email)
		if getUserWithEmail != nil {
			return response.NewErrDuplicateUniqueColumn("Email")
		}
	}

	if user.Username != requestData.Username {
		getUserWithUsername, _ := c.userRepository.GetUserWithUsername(requestData.Username)
		if getUserWithUsername != nil {
			return response.NewErrDuplicateUniqueColumn("Username")
		}
	}

	if user.PhoneNumber != requestData.PhoneNumber {
		getUserWithPhone, _ := c.userRepository.GetUserWithPhone(requestData.PhoneNumber)
		if getUserWithPhone != nil {
			return response.NewErrDuplicateUniqueColumn("Phone number")
		}
	}

	return c.userRepository.Update(userID, requestData)
}

func (c *AuthService) generateToken(user *models.User) (token string, err error) {
	eJWT := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":   user.ID,
			"role": user.Role,
			"exp":  time.Now().Add(c.jwtExpired).Unix(),
		},
	)

	return eJWT.SignedString([]byte(c.jwtSecretKey))
}
