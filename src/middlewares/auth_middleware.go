package middleware

import "github.com/gin-gonic/gin"

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (c *AuthMiddleware) Auth(ctx *gin.Context) {

}

func (c *AuthMiddleware) CheckRole(ctx *gin.Context) {

}
