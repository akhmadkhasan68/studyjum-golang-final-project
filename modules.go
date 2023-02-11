package main

import (
	controllers "final-project/src/controllers/auth"
	"final-project/src/routes"

	"gorm.io/gorm"
)

func prepareModules(handler *routes.Router, db *gorm.DB) {
	// Persiapan proses authenticator
	// jwtMid := middlewares.NewAuthenticator(config.GetEnvVariable("JWT_SECRET_KEY"))
	// jwtDuration := config.GetEnvVariable("JWT_EXPIRED_TIME")
	// jwtExpiredTime, error := strconv.Atoi(jwtDuration)
	// if error != nil {
	// 	panic(error)
	// }

	// Persiapan repository, business dan handler
	// userRepo := usermodule.NewRepository(db)
	// userUseCase := userusecase.NewService(userRepo)
	// userUseCase.SetJWTConfig(
	// 	config.GetEnvVariable("JWT_SECRET_KEY"),
	// 	time.Duration(jwtExpiredTime)*time.Minute,
	// )

	// Controller
	handler.User = controllers.NewAuthController()
}
