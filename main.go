package main

import (
	"TestTask/config"
	"TestTask/controllers"
	"TestTask/repository"
	"TestTask/db"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"

	_ "TestTask/docs" // Import generated docs
)

// @title Building API
// @version 1.0
// @description This is a sample server for managing buildings.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api/v1

func main() {
	cfg, err := config.LoadConfig("config/config.json")
	if err != nil {
		log.Fatal(err)
	}

	dbConn, err := db.InitDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	repo := repository.NewBuildingRepository(dbConn)
	controller := controllers.NewBuildingController(repo)

	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		buildings := v1.Group("/buildings")
		{
			buildings.POST("", controller.CreateBuilding)
			buildings.GET("", controller.GetBuildings)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}