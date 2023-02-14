package main

import (
	authBussines "final-project/src/bussiness"
	productBussiness "final-project/src/bussiness"
	"final-project/src/config"
	authController "final-project/src/controllers/auth"
	controllers "final-project/src/controllers/products"
	"final-project/src/middlewares"
	"final-project/src/repositories"
	"final-project/src/routes"

	"gorm.io/gorm"
)

func prepareModules(handler *routes.Router, db *gorm.DB) {
	jwtMid := middlewares.NewAuthenticator(config.GetEnvVariable("JWT_SECRET_KEY"))

	//init repository
	userRepository := repositories.NewUserRepository(db)
	productRepository := repositories.NewProductRepository(db)

	//init service / bussiness
	userBussiness := authBussines.NewAuthService(userRepository)
	productBussiness := productBussiness.NewProductService(productRepository)

	// Controller
	handler.User = authController.NewAuthController(userBussiness, jwtMid)
	handler.Product = controllers.NewProductsController(productBussiness, jwtMid)
}
