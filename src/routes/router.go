package routes

import (
	"final-project/src/commons/enums"
	"final-project/src/config"
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

	//Auth Route
	auth := v1.Group("/auth")
	auth.POST("/register", h.User.Register)
	auth.POST("/login", h.User.Login)

	//Auth Route With JWT Middleware
	authLoggedIn := auth.Use(middlewares.JWTMiddlewareAuth(config.GetEnvVariable("JWT_SECRET_KEY")))
	authLoggedIn.GET("/profile", h.User.Profile)
	authLoggedIn.PUT("/profile", h.User.UpdateProfile)
	authLoggedIn.PUT("/change-password", h.User.ChangePassword)

	//Products Route With JWT Middleware
	products := v1.Group("/products")
	products.Use(middlewares.JWTMiddlewareAuth(config.GetEnvVariable("JWT_SECRET_KEY")))

	//Product With All Role
	productAllRole := products.Use(middlewares.RoleMiddleware([]string{enums.MEMBER, enums.OUTLET}))
	productAllRole.GET("/", h.Product.GetAllProducts)
	productAllRole.GET("/:id", h.Product.GetDetailProduct)

	//Product With Outlet Role
	productOutletRole := products.Use(middlewares.RoleMiddleware([]string{enums.OUTLET}))
	productOutletRole.POST("/", h.Product.CreateProduct)
	productOutletRole.PUT("/:id", h.Product.UpdateProduct)
	productOutletRole.DELETE("/:id", h.Product.DeleteProduct)

	//Orders Route With JWT Middleware
	orders := v1.Group("/orders")
	orders.Use(middlewares.JWTMiddlewareAuth(config.GetEnvVariable("JWT_SECRET_KEY")))
	orders.GET("/", h.Order.GetAllOrders)
	orders.GET("/:id", h.Order.GetDetailOrder)
}
