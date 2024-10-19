package main

import (
	"TestTask/config"
	"TestTask/controllers"
	"TestTask/repository"

	//"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func main() {
	cfg, err := config.LoadConfig("config/config.json")
	if err != nil {
		log.Fatal(err)
	}

	db, err := config.NewDBConnection(cfg)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewBuildingRepository(db)
	controller := controllers.NewBuildingController(repo)

	r := gin.Default()

	r.POST("/api/v1/buildings", controller.CreateBuilding)
	r.GET("/api/v1/buildings", controller.GetBuildings)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
