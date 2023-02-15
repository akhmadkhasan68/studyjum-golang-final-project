package controllers

import (
	authBussines "final-project/src/bussiness"
	response "final-project/src/commons/responses"
	"final-project/src/middlewares"
	"final-project/src/requests"
	"final-project/src/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService   *authBussines.AuthService
	JWTMiddleware middlewares.IAuthenticator
}

func NewAuthController(authService *authBussines.AuthService, JWTMiddleware middlewares.IAuthenticator) *AuthController {
	return &AuthController{
		authService:   authService,
		JWTMiddleware: JWTMiddleware,
	}
}

func (c *AuthController) Register(ctx *gin.Context) {
	var registerRequest requests.RegisterRequest
	if err := ctx.ShouldBind(&registerRequest); err != nil {
		response.JSONErrorResponse(ctx, err)
		return
	}

	if err := c.authService.Register(registerRequest); err != nil {
		response.JSONErrorResponse(ctx, err)
		return
	}

	response.JSONBasicResponse(ctx, http.StatusCreated, "Success to register account!")
}

func (c *AuthController) Login(ctx *gin.Context) {
	var loginRequest requests.LoginRequest
	if err := ctx.ShouldBind(&loginRequest); err != nil {
		response.JSONErrorResponse(ctx, err)
		return
	}

	token, user, err := c.authService.Login(loginRequest)
	if err != nil {
		response.JSONErrorResponse(ctx, err)
		return
	}

	response.JSONBasicData(ctx, http.StatusOK, "Success login to your account!", responses.ToLoginResponse(user, token))
}

func (c *AuthController) Profile(ctx *gin.Context) {
	data, err := c.JWTMiddleware.ExtractJWTUser(ctx)
	if err != nil {
		response.JSONErrorResponse(ctx, err)
		return
	}

	user, err := c.authService.GetByUserID(data.ID)
	if err != nil {
		response.JSONErrorResponse(ctx, err)
		return
	}

	response.JSONBasicData(ctx, http.StatusOK, "Profile Controller", responses.ToProfileResponse(user))
}

func (c *AuthController) UpdateProfile(ctx *gin.Context) {
	var updateProfileRequest requests.UpdateProfileRequest
	if err := ctx.ShouldBind(&updateProfileRequest); err != nil {
		response.JSONErrorResponse(ctx, err)
		return
	}

	data, err := c.JWTMiddleware.ExtractJWTUser(ctx)
	if err != nil {
		response.JSONErrorResponse(ctx, err)
		return
	}

	if err := c.authService.UpdateProfile(data.ID, updateProfileRequest); err != nil {
		response.JSONErrorResponse(ctx, err)
		return
	}

	response.JSONBasicResponse(ctx, http.StatusOK, "Success update your profile!")
}

func (c *AuthController) ChangePassword(ctx *gin.Context) {
	var changePasswordRequest requests.ChangePasswordRequest
	if err := ctx.ShouldBind(&changePasswordRequest); err != nil {
		response.JSONErrorResponse(ctx, err)
		return
	}

	data, err := c.JWTMiddleware.ExtractJWTUser(ctx)
	if err != nil {
		response.JSONErrorResponse(ctx, err)
		return
	}

	if err := c.authService.ChangePassword(data.ID, changePasswordRequest); err != nil {
		response.JSONErrorResponse(ctx, err)
		return
	}

	response.JSONBasicResponse(ctx, http.StatusOK, "Success change your password!")
}
