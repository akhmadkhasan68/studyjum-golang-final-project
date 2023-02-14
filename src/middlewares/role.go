package middlewares

import (
	"bytes"
	"encoding/json"
	response "final-project/src/commons/responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type IRoleMiddleware interface {
	ExtractJWTUser(ctx *gin.Context) (*token, error)
}

type roleMiddleware struct {
	JWTMiddleware IAuthenticator
}

func NewRoleMiddleware(JWTMiddleware IAuthenticator) *roleMiddleware {
	return &roleMiddleware{JWTMiddleware}
}

func RoleMiddleware(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := c.Get("user")
		claims := user.(*jwt.Token).Claims.(jwt.MapClaims)
		res := new(token)
		buff := new(bytes.Buffer)
		json.NewEncoder(buff).Encode(&claims)
		json.NewDecoder(buff).Decode(res)

		isExist := false

		for _, role := range roles {
			if role == res.Role {
				isExist = true
				break
			}
		}

		if !isExist {
			c.AbortWithStatusJSON(
				http.StatusForbidden,
				response.ErrResponse{
					Message:     "Request Forbidden",
					Status:      http.StatusForbidden,
					Description: "",
				},
			)
			return
		}

		c.Next()
	}
}
