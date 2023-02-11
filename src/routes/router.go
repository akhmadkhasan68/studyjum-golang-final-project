package routes

import (
	"final-project/src/config"
	controllers "final-project/src/controllers/auth"
	"final-project/src/middlewares"

	"github.com/gin-gonic/gin"
)

type Router struct {
	User *controllers.AuthController
}

func (h *Router) CreateRouting(r *gin.Engine) {
	r.Use(middlewares.Headers())
	r.Use(middlewares.CustomLogger())
	r.Use(gin.CustomRecovery(middlewares.ErrorHandler))

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello Gin Gonic App!",
		})
	})

	// Group routing /v1
	v1 := r.Group("/v1")

	auth := v1.Group("/auth")
	auth.POST("/register", h.User.Register)
	auth.POST("/login", h.User.Login)

	// Group routing /v1 dengan auth JWT
	authLoggedIn := auth.Use(middlewares.JWTMiddlewareAuth(config.GetEnvVariable("JWT_SECRET_KEY")))
	authLoggedIn.GET("/profile", h.User.Profile)
	authLoggedIn.PUT("/profile", h.User.UpdateProfile)
}
