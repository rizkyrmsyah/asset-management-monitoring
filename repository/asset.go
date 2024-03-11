package repository

import (
	"asset-tracker/model"
	"database/sql"
)

func AddAsset(db *sql.DB, asset model.Asset) (err error) {
	sql := "INSERT INTO assets (name, code, in_date, source, created_at, updated_at) VALUES ($1 ,$2, $3, $4, NOW(), NOW())"
	errs := db.QueryRow(sql, asset.Name, asset.Code, asset.InDate, asset.Source)

	return errs.Err()
}
