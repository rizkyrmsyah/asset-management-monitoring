package repository

import (
	"asset-tracker/model"
	"database/sql"
)

func CreateLocation(db *sql.DB, location model.Location) (err error) {
	sql := "INSERT INTO locations (name, created_at, updated_at) VALUES ($1, NOW(), NOW())"
	errs := db.QueryRow(sql, location.Name)

	return errs.Err()
}

func GetAllLocation(db *sql.DB) (locations []model.Location, err error) {
	sql := "SELECT * FROM locations WHERE deleted_at IS NULL ORDER BY name ASC"
	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var locationData = model.Location{}
		err = rows.Scan(&locationData.ID, &locationData.Name, &locationData.CreatedAt, &locationData.UpdatedAt, &locationData.DeletedAt)
		if err != nil {
			return
		}

		locations = append(locations, locationData)
	}

	return
}

func GetLocationById(db *sql.DB, id int) (location *model.Location, err error) {
	var loc model.Location
	sql := "SELECT * FROM locations WHERE id = $1 AND deleted_at IS NULL"
	err = db.QueryRow(sql, id).Scan(&loc.ID, &loc.Name, &loc.CreatedAt, &loc.UpdatedAt, &loc.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &loc, nil
}

func UpdateLocation(db *sql.DB, location model.LocationUpdateRequest) (err error) {
	sql := "UPDATE locations SET name = $2, updated_at = NOW() WHERE id = $1"
	errs := db.QueryRow(sql, location.ID, location.Name)

	return errs.Err()
}

func DeleteLocation(db *sql.DB, locationId int) (err error) {
	sql := "UPDATE locations SET deleted_at = NOW() WHERE id = $1"
	errs := db.QueryRow(sql, locationId)

	return errs.Err()
}
