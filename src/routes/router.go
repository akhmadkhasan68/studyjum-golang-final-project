package routes

import (
	authController "final-project/src/controllers/auth"
	ordersController "final-project/src/controllers/orders"
	productsController "final-project/src/controllers/products"
	"final-project/src/middlewares"

	"github.com/gin-gonic/gin"
)

type Router struct {
	User    *authController.AuthController
	Order   *ordersController.OrdersController
	Product *productsController.ProductsController
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
	// authLoggedIn := auth.Use(middlewares.JWTMiddlewareAuth(config.GetEnvVariable("JWT_SECRET_KEY")))
	auth.GET("/profile", h.User.Profile)
	auth.PUT("/profile", h.User.UpdateProfile)

	products := v1.Group("products")
	products.GET("/", h.Product.GetAllProducts)
	products.GET("/:id", h.Product.GetDetailProduct)

	orders := v1.Group("/orders")
	orders.GET("/", h.Order.GetAllOrders)
	orders.GET("/:id", h.Order.GetDetailOrder)
}
