package main

import (
	"final-project/src/config"
	"final-project/src/config/database"
	"final-project/src/database/models"
	"final-project/src/routes"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}

	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.User{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Product{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Order{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.OrderDetail{})

	// shipperClientAggregator := httpclient.NewShipperAggregatorClient(config.GetEnvVariable("SHIPPER_BASE_URL"), config.GetEnvVariable("SHIPPER_API_KEY"))

	r := gin.New()

	routing := &routes.Router{}

	prepareModules(routing, db)
	routing.CreateRouting(r)

	portEnv := config.GetEnvVariable("APP_PORT")
	appPort, error := strconv.Atoi(portEnv)
	if error != nil {
		panic(error)
	}

	svc := http.Server{
		Addr:    fmt.Sprintf(":%d", appPort),
		Handler: r,
	}

	fmt.Printf("Starting server at port %d\n", appPort)
	if err := svc.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
