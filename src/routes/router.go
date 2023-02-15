package routes

import (
	"final-project/src/commons/enums"
	"final-project/src/config"
	"final-project/src/controllers"
	"final-project/src/middlewares"

	"github.com/gin-gonic/gin"
)

type Router struct {
	User    *controllers.AuthController
	Order   *controllers.OrdersController
	Product *controllers.ProductsController
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
	products.GET("/", h.Product.GetAllProducts)
	products.GET("/:id", h.Product.GetDetailProduct)

	//Product With Outlet Role
	productOutletRole := products.Use(middlewares.RoleMiddleware([]string{enums.OUTLET}))
	productOutletRole.POST("/", h.Product.CreateProduct)
	productOutletRole.PUT("/:id", h.Product.UpdateProduct)
	productOutletRole.DELETE("/:id", h.Product.DeleteProduct)

	//Orders Route With JWT Middleware
	orders := v1.Group("/orders")
	orders.Use(middlewares.JWTMiddlewareAuth(config.GetEnvVariable("JWT_SECRET_KEY")))

	//Orders Route All Role
	orders.GET("/", h.Order.GetAllOrders)
	orders.GET("/:id", h.Order.GetDetailOrder)

	//Orders Member Role
	ordersMemberRole := orders.Use(middlewares.RoleMiddleware([]string{enums.MEMBER}))
	ordersMemberRole.POST("/", h.Order.CreateOrder)
	ordersMemberRole.DELETE("/:id", h.Order.CancelOrder)

	//Order Outlet Role
	ordersOutletRole := v1.Group("/orders").Group("/outlet")
	ordersOutletRole.Use(middlewares.JWTMiddlewareAuth(config.GetEnvVariable("JWT_SECRET_KEY")))
	ordersOutletRole.Use(middlewares.RoleMiddleware([]string{enums.OUTLET}))
	ordersOutletRole.POST("/ship", h.Order.ShipOrder)
	ordersOutletRole.DELETE("/reject/:id", h.Order.RejectOrder)
}
