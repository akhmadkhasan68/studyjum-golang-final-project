package main

import (
	"final-project/src/config"
	"final-project/src/config/database"
	"final-project/src/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	_, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}

	// shipperClientAggregator := httpclient.NewShipperAggregatorClient(config.GetEnvVariable("SHIPPER_BASE_URL"), config.GetEnvVariable("SHIPPER_API_KEY"))

	router := gin.Default()

	app := routes.NewRouter(router)
	port := fmt.Sprintf(":%s", config.GetEnvVariable("APP_PORT"))

	app.Start(port)
}
