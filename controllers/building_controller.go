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
