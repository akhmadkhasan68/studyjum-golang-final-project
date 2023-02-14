package controllers

import (
	bussiness "final-project/src/bussiness/auth"
	response "final-project/src/commons/responses"
	"final-project/src/requests"
	"final-project/src/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *bussiness.AuthService
}

func NewAuthController(authService *bussiness.AuthService) *AuthController {
	return &AuthController{authService}
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
	// token, err := c.service.UserLogin(ctx, login.Email, login.Password)
	// if err != nil {
	// 	response.JSONErrorResponse(ctx, err)
	// 	return
	// }

	response.JSONBasicData(ctx, http.StatusOK, "Success login to your account!", responses.ToLoginResponse(user, token))
}

func (c *AuthController) Profile(ctx *gin.Context) {
	// claims, err := c.mid.ExtractJWTUser(ctx)
	// if err != nil {
	// 	response.JSONErrorResponse(ctx, err)
	// 	return
	// }

	// user, err := c.service.GetUserWithID(ctx, claims.ID)
	// if err != nil {
	// 	response.JSONErrorResponse(ctx, err)
	// 	return
	// }

	response.JSONBasicData(ctx, http.StatusOK, "Profile Controller", "")
}

func (c *AuthController) UpdateProfile(ctx *gin.Context) {
	// claims, err := c.mid.ExtractJWTUser(ctx)
	// if err != nil {
	// 	response.JSONErrorResponse(ctx, err)
	// 	return
	// }

	// var editProfile user.EditProfile
	// if err := ctx.ShouldBind(&editProfile); err != nil {
	// 	response.JSONErrorResponse(ctx, err)
	// 	return
	// }

	// if err := c.service.PutUserWithID(ctx, claims.ID, editProfile); err != nil {
	// 	response.JSONErrorResponse(ctx, err)
	// 	return
	// }

	response.JSONBasicResponse(ctx, http.StatusOK, "Update Profile Controller")
}
