package routes

import "github.com/gin-gonic/gin"

type Router struct {
	router *gin.Engine
}

func NewRouter(router *gin.Engine) *Router {
	return &Router{
		router: router,
	}
}

func (r *Router) Start(port string) {
	// Auth route
	// auth := r.router.Group("/auth")
	// auth.POST("/register", r.user.Register)
	// auth.POST("/login", r.user.Login)

	// // User route
	// user := r.router.Group("/users", r.middleware.Auth)
	// user.POST("/", r.middleware.CheckRole(r.user.Create, []string{enums.Admin}))
	// user.GET("/", r.middleware.CheckRole(r.user.AllUsers, []string{enums.Admin}))
	// user.GET("/email/:email", r.middleware.CheckRole(r.user.DetailUserByEmail, []string{enums.Admin}))
	// user.GET("/profile", r.middleware.CheckRole(r.user.DetailUserById, []string{enums.User}))
	// user.PUT("/profile", r.middleware.CheckRole(r.user.UpdateUserById, []string{enums.User}))

	// // Product route
	// product := r.router.Group("/products", r.middleware.Auth)
	// product.GET("/", r.middleware.CheckRole(r.product.GetAllProducts, []string{enums.Admin, enums.User}))
	// product.GET("/id/:id", r.middleware.CheckRole(r.product.GetProductById, []string{enums.Admin, enums.User}))
	// product.POST("/", r.middleware.CheckRole(r.product.CreateProduct, []string{enums.Admin}))
	// product.PUT("/id/:id", r.middleware.CheckRole(r.product.UpdateProduct, []string{enums.Admin}))
	// product.DELETE("/id/:id", r.middleware.CheckRole(r.product.DeleteProduct, []string{enums.Admin}))

	// // Transaction route
	// transaction := r.router.Group("/transactions", r.middleware.Auth)
	// transaction.POST("/inquire", r.middleware.CheckRole(r.transaction.InquireTransaction, []string{enums.User}))
	// transaction.PUT("/id/:id", r.middleware.CheckRole(r.transaction.UpdateStatTransaction, []string{enums.Admin, enums.Cashier}))
	// transaction.POST("/confirm", r.middleware.CheckRole(r.transaction.ConfirmTransaction, []string{enums.User}))

	v1 := r.router.Group("/v1")
	v1.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	r.router.Run(port)
}
