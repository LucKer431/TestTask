package main

import (
    "github.com/gin-gonic/gin"
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/gin-swagger/swaggerFiles"
    "github.com/swaggo/swag/example/celler/httputil"
    "log"
    "your-project-name/db"
    "your-project-name/models"
)

func main() {
    db.InitDB()
    r := gin.Default()

    // Swagger UI
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    r.POST("/api/v1/buildings", createBuilding)
    r.GET("/api/v1/buildings", getBuildings)

    log.Fatal(r.Run(":8080"))
}

func createBuilding(c *gin.Context) {
    var building models.Building
    if err := c.BindJSON(&building); err != nil {
        httputil.NewError(c, http.StatusBadRequest, err)
        return
    }
    result := db.DB.Create(&building)
    if result.Error != nil {
        httputil.NewError(c, http.StatusInternalServerError, result.Error)
        return
    }
    c.JSON(http.StatusCreated, building)
}

func getBuildings(c *gin.Context) {
    var buildings []models.Building
    city := c.Query("city")
    yearBuilt := c.Query("yearBuilt")
    floors := c.Query("floors")

    query := db.DB.Model(&models.Building{})
    if city != "" {
        query = query.Where("city = ?", city)
    }
    if yearBuilt != "" {
        query = query.Where("year_built = ?", yearBuilt)
    }
    if floors != "" {
        query = query.Where("floors = ?", floors)
    }

    result := query.Find(&buildings)
    if result.Error != nil {
        httputil.NewError(c, http.StatusInternalServerError, result.Error)
        return
    }
    c.JSON(http.StatusOK, buildings)
}