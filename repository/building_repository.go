package repository

import (
	"database/sql"
	"strings"

	"TestTask/models"
)

type BuildingRepository struct {
	db *sql.DB
}

func NewBuildingRepository(db *sql.DB) *BuildingRepository {
	return &BuildingRepository{db: db}
}

func (r *BuildingRepository) CreateBuilding(building *models.Building) error {
	_, err := r.db.Exec(`
		INSERT INTO buildings (name, city, year_built, floors)
		VALUES ($1, $2, $3, $4)
	`, building.Name, building.City, building.YearBuilt, building.Floors)
	return err
}

func (r *BuildingRepository) GetBuildings(city string, yearBuilt int, floors int) ([]models.Building, error) {
	var buildings []models.Building
	var query string

	query = `
		SELECT id, name, city, year_built, floors, created_at
		FROM buildings
	`

	params := make([]interface{}, 0)
	conditions := make([]string, 0)

	if city != "" {
		conditions = append(conditions, "city = $1")
		params = append(params, city)
	}

	if yearBuilt != 0 {
		conditions = append(conditions, "year_built = $2")
		params = append(params, yearBuilt)
	}

	if floors != 0 {
		conditions = append(conditions, "floors = $3")
		params = append(params, floors)
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	rows, err := r.db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var building models.Building
		err := rows.Scan(&building.ID, &building.Name, &building.City, &building.YearBuilt, &building.Floors, &building.CreateTime)
		if err != nil {
			return nil, err
		}
		buildings = append(buildings, building)
	}

	return buildings, nil
}