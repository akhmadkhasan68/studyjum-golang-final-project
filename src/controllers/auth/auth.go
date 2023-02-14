package controllers

import (
	response "final-project/src/commons/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (c *AuthController) Register(ctx *gin.Context) {
	// var user user.InsertUser
	// if err := ctx.ShouldBind(&user); err != nil {
	// 	response.JSONErrorResponse(ctx, err)
	// 	return
	// }

	// if err := c.service.PostUser(ctx, user); err != nil {
	// 	response.JSONErrorResponse(ctx, err)
	// 	return
	// }

	response.JSONBasicResponse(ctx, http.StatusCreated, "Register controller")
}

func (c *AuthController) Login(ctx *gin.Context) {
	// var login UserLogin
	// if err := ctx.ShouldBind(&login); err != nil {
	// 	response.JSONErrorResponse(ctx, err)
	// 	return
	// }

	// token, err := c.service.UserLogin(ctx, login.Email, login.Password)
	// if err != nil {
	// 	response.JSONErrorResponse(ctx, err)
	// 	return
	// }

	response.JSONBasicData(ctx, http.StatusOK, "Login Controller", "token")
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
