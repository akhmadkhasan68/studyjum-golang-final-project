package main

import (
	"final-project/src/bussiness"
	"final-project/src/config"
	"final-project/src/controllers"
	"final-project/src/httpclient"
	"final-project/src/middlewares"
	"final-project/src/repositories"
	"final-project/src/routes"

	"gorm.io/gorm"
)

func prepareModules(handler *routes.Router, db *gorm.DB) {
	jwtMid := middlewares.NewAuthenticator(config.GetEnvVariable("JWT_SECRET_KEY"))
	shipperClientAggregator := httpclient.NewShipperAggregatorClient(config.GetEnvVariable("SHIPPER_BASE_URL"), config.GetEnvVariable("SHIPPER_API_KEY"))

	//init repository
	userRepository := repositories.NewUserRepository(db)
	productRepository := repositories.NewProductRepository(db)
	orderRepository := repositories.NewOrderRepository(db)

	//init service / bussiness
	authBussines := bussiness.NewAuthService(userRepository)
	productBussines := bussiness.NewProductService(productRepository)
	orderBussiness := bussiness.NewOrderService(shipperClientAggregator, orderRepository, productRepository)

	// Controller
	handler.User = controllers.NewAuthController(authBussines, jwtMid)
	handler.Product = controllers.NewProductsController(productBussines, jwtMid)
	handler.Order = controllers.NewOrdersController(orderBussiness, jwtMid)
}
