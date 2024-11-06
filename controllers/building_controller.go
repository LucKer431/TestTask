package controllers

import (
	"net/http"
	"strconv"

	"TestTask/models"
	"TestTask/repository"

	"github.com/gin-gonic/gin"
)

type BuildingController struct {
	repo *repository.BuildingRepository
}

func NewBuildingController(repo *repository.BuildingRepository) *BuildingController {
	return &BuildingController{repo: repo}
}

// @Summary Get all buildings
// @Description Get all buildings with optional filtering by city, year built, and number of floors
// @Accept json
// @Produce json
// @Param city query string false "City filter"
// @Param year_built query int false "Year built filter"
// @Param floors query int false "Number of floors filter"
// @Success 200 {array} models.Building
// @Failure 400 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /buildings [get]
func (c *BuildingController) GetBuildings(ctx *gin.Context) {
	city := ctx.Query("city")
	yearBuiltStr := ctx.Query("year_built")
	floorsStr := ctx.Query("floors")

	var yearBuilt, floors int
	var err error

	if yearBuiltStr != "" {
		yearBuilt, err = strconv.Atoi(yearBuiltStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный год"})
			return
		}
	}

	if floorsStr != "" {
		floors, err = strconv.Atoi(floorsStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверное кол-во этажей"})
			return
		}
	}

	buildings, err := c.repo.GetBuildings(city, yearBuilt, floors)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка чтения"})
		return
	}

	ctx.JSON(http.StatusOK, buildings)
}

// @Summary Create a new building
// @Description Create a new building and save it to the database
// @Accept json
// @Produce json
// @Param building body models.Building true "Building to create"
// @Success 201 {object} models.Building
// @Failure 400 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /buildings [post]
func (c *BuildingController) CreateBuilding(ctx *gin.Context) {
	var building models.Building
	err := ctx.BindJSON(&building)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверное тело записи"})
		return
	}

	err = c.repo.CreateBuilding(&building)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Создание прошло успешно"})
}